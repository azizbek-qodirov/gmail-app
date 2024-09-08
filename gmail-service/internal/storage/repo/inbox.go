package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	pb "gmail-service/internal/pkg/genproto"
	"time"

	"github.com/lib/pq"
)

type InboxRepo struct {
	db *sql.DB
}

func NewInboxRepo(db *sql.DB) *InboxRepo {
	return &InboxRepo{db: db}
}

func (r *InboxRepo) ArchiveMessage(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	query := `
		UPDATE inbox
		SET is_archived = NOT is_archived
		WHERE id = $1 AND deleted_at = 0
	`

	res, err := r.db.ExecContext(ctx, query, req.Id)
	if err != nil {
		return nil, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		return nil, sql.ErrNoRows
	}

	return &pb.Void{}, nil
}

func (r *InboxRepo) Delete(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	query := `
		UPDATE inbox
		SET deleted_at = $1
		WHERE id = $2
	`
	res, err := r.db.ExecContext(ctx, query, time.Now().Unix(), req.Id)
	if err != nil {
		return nil, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		return nil, sql.ErrNoRows
	}

	return &pb.Void{}, nil
}

func (r *InboxRepo) GetAll(ctx context.Context, req *pb.InboxMessageGetAllReq) (*pb.InboxMessagesGetAllRes, error) {
	var messages []*pb.InboxMessageGetRes

	query := `
		SELECT 
			i.id,
			i.receiver_id,
			i.type,
			i.is_spam,
			i.is_archived,
			i.is_starred,
			i.read_at,
			i.deleted_at,
			o.id,
			o.subject,
			o.body,
			o.attachment_ids,
			o.receiver_to_emails,
			o.receiver_cc_emails,
			CASE WHEN i.type = 'bcc' THEN o.receiver_bcc_emails ELSE NULL END AS receiver_bcc_emails,
			o.is_draft,
			o.is_archived,
			o.is_starred,
			o.sent_at,
			o.deleted_at,
			u.id,
			u.first_name,
			u.last_name,
			u.dob,
			u.email,
			u.gender,
			u.pfp_url
		FROM inbox AS i
		JOIN outbox AS o ON i.outbox_id = o.id
		JOIN users AS u ON o.sender_id = u.id
		WHERE i.receiver_id = $1 
	`

	var args []interface{}
	args = append(args, req.ReceiverId)
	argCount := 2

	if req.Body.Query != "" {
		query += fmt.Sprintf(" AND (o.subject ILIKE $%d OR o.body ILIKE $%d)", argCount, argCount)
		args = append(args, "%"+req.Body.Query+"%", "%"+req.Body.Query+"%")
		argCount++
	}

	if req.Body.SenderId != "" {
		query += fmt.Sprintf(" AND o.sender_id = $%d", argCount)
		args = append(args, req.Body.SenderId)
		argCount++
	}

	if req.Body.Type != "" {
		query += fmt.Sprintf(" AND i.type = $%d", argCount)
		args = append(args, req.Body.Type)
		argCount++
	}

	if req.Body.IsSpam {
		query += fmt.Sprintf(" AND i.is_spam = $%d", argCount)
		args = append(args, req.Body.IsSpam)
		argCount++
	}

	if req.Body.IsArchived {
		query += fmt.Sprintf(" AND i.is_archived = $%d", argCount)
		args = append(args, req.Body.IsArchived)
		argCount++
	}

	if req.Body.IsStarred {
		query += fmt.Sprintf(" AND i.is_starred = $%d", argCount)
		args = append(args, req.Body.IsStarred)
		argCount++
	}

	if req.Body.IsTrashed {
		query += fmt.Sprintf(" AND i.deleted_at = $%d", argCount)
		args = append(args, 1)
		argCount++
	} else {
		query += fmt.Sprintf(" AND i.deleted_at = $%d", argCount)
		args = append(args, 0)
		argCount++
	}

	if req.Body.SentFrom != "" {
		query += fmt.Sprintf(" AND o.sent_at >= $%d", argCount)
		args = append(args, req.Body.SentFrom)
		argCount++
	}

	if req.Body.SentTo != "" {
		query += fmt.Sprintf(" AND o.sent_at <= $%d", argCount)
		args = append(args, req.Body.SentTo)
		argCount++
	}

	if req.Body.UnreadOnly {
		query += fmt.Sprintf(" AND i.read_at = $%d", argCount)
		args = append(args, 0)
		argCount++
	}

	query += fmt.Sprintf(" OFFSET $%d LIMIT $%d", argCount, argCount+1)
	args = append(args, req.Pagination.Skip, req.Pagination.Limit)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		message := pb.InboxMessageGetRes{
			Outbox: &pb.OutboxMessageGetRes{
				AttachmentIds: &pb.AttachmentIdsWrapper{},
				Sender:        &pb.UserGetRes{},
				Receivers: &pb.Receivers{
					To:  &pb.MessageSendTo{},
					Cc:  &pb.MessageSendCC{},
					Bcc: &pb.MessageSendBCC{},
				},
			},
		}
		err = rows.Scan(
			&message.Id,
			&message.ReceiverId,
			&message.Type,
			&message.IsSpam,
			&message.IsArchived,
			&message.IsStarred,
			&message.ReadAt,
			&message.DeletedAt,
			&message.Outbox.Id,
			&message.Outbox.Subject,
			&message.Outbox.Body,
			pq.Array(&message.Outbox.AttachmentIds.AttachmentIds),
			pq.Array(&message.Outbox.Receivers.To.Emails),
			pq.Array(&message.Outbox.Receivers.Cc.Emails),
			pq.Array(&message.Outbox.Receivers.Bcc.Emails),
			&message.Outbox.IsDraft,
			&message.Outbox.IsArchived,
			&message.Outbox.IsStarred,
			&message.Outbox.SentAt,
			&message.Outbox.DeletedAt,
			&message.Outbox.Sender.Id,
			&message.Outbox.Sender.FirstName,
			&message.Outbox.Sender.LastName,
			&message.Outbox.Sender.Dob,
			&message.Outbox.Sender.Email,
			&message.Outbox.Sender.Gender,
			&message.Outbox.Sender.PfpUrl,
		)
		if err != nil {
			return nil, err
		}

		messages = append(messages, &message)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &pb.InboxMessagesGetAllRes{
		Messages: messages,
	}, nil
}

func (r *InboxRepo) GetByID(ctx context.Context, req *pb.ByID) (*pb.InboxMessageGetRes, error) {
	var message pb.InboxMessageGetRes
	message.Outbox = &pb.OutboxMessageGetRes{
		AttachmentIds: &pb.AttachmentIdsWrapper{},
		Sender:        &pb.UserGetRes{},
		Receivers: &pb.Receivers{
			To:  &pb.MessageSendTo{},
			Cc:  &pb.MessageSendCC{},
			Bcc: &pb.MessageSendBCC{},
		},
	}

	query := `
		SELECT 
			i.id,
			i.receiver_id,
			i.type,
			i.is_spam,
			i.is_archived,
			i.is_starred,
			i.read_at,
			i.deleted_at,
			o.id,
			o.subject,
			o.body,
			o.attachment_ids,
			o.receiver_to_emails,
			o.receiver_cc_emails,
			CASE WHEN i.type = 'bcc' THEN o.receiver_bcc_emails ELSE NULL END AS receiver_bcc_emails, -- Conditional BCC
			o.is_draft,
			o.is_archived,
			o.is_starred,
			o.sent_at,
			o.deleted_at,
			u.id,
			u.first_name,
			u.last_name,
			u.dob,
			u.email,
			u.gender,
			u.pfp_url
		FROM inbox AS i
		JOIN outbox AS o ON i.outbox_id = o.id
		JOIN users AS u ON o.sender_id = u.id
		WHERE i.id = $1 AND i.deleted_at = 0
	`

	row := r.db.QueryRowContext(ctx, query, req.Id)
	err := row.Scan(
		&message.Id,
		&message.ReceiverId,
		&message.Type,
		&message.IsSpam,
		&message.IsArchived,
		&message.IsStarred,
		&message.ReadAt,
		&message.DeletedAt,
		&message.Outbox.Id,
		&message.Outbox.Subject,
		&message.Outbox.Body,
		pq.Array(&message.Outbox.AttachmentIds.AttachmentIds),
		pq.Array(&message.Outbox.Receivers.To.Emails),
		pq.Array(&message.Outbox.Receivers.Cc.Emails),
		pq.Array(&message.Outbox.Receivers.Bcc.Emails),
		&message.Outbox.IsDraft,
		&message.Outbox.IsArchived,
		&message.Outbox.IsStarred,
		&message.Outbox.SentAt,
		&message.Outbox.DeletedAt,
		&message.Outbox.Sender.Id,
		&message.Outbox.Sender.FirstName,
		&message.Outbox.Sender.LastName,
		&message.Outbox.Sender.Dob,
		&message.Outbox.Sender.Email,
		&message.Outbox.Sender.Gender,
		&message.Outbox.Sender.PfpUrl,
	)

	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (r *InboxRepo) MarkAsRead(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	query := `
		UPDATE inbox
		SET read_at = $1
		WHERE id = $2 AND deleted_at = 0
	`

	res, err := r.db.ExecContext(ctx, query, time.Now().Unix(), req.Id)
	if err != nil {
		return nil, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		return nil, sql.ErrNoRows
	}

	return &pb.Void{}, nil
}

func (r *InboxRepo) MarkAsSpam(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	query := `
		UPDATE inbox
		SET is_spam = NOT is_spam
		WHERE id = $1 AND deleted_at = 0
	`

	res, err := r.db.ExecContext(ctx, query, req.Id)
	if err != nil {
		return nil, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		return nil, sql.ErrNoRows
	}

	return &pb.Void{}, nil
}

func (r *InboxRepo) MoveToTrash(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	query := `
		UPDATE inbox
		SET deleted_at = 
			CASE 
				WHEN deleted_at = 0 THEN 1 
				WHEN deleted_at = 1 THEN 0
				ELSE deleted_at
			END
		WHERE id = $1 AND deleted_at = 0 OR deleted_at = 1
	`
	res, err := r.db.ExecContext(ctx, query, req.Id)
	if err != nil {
		return nil, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		return nil, sql.ErrNoRows
	}

	return &pb.Void{}, nil
}

func (r *InboxRepo) StarMessage(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	query := `
		UPDATE inbox
		SET is_starred = NOT is_starred
		WHERE id = $1 AND deleted_at = 0
	`

	res, err := r.db.ExecContext(ctx, query, req.Id)
	if err != nil {
		return nil, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		return nil, errors.New("message not found")
	}

	return &pb.Void{}, nil
}
