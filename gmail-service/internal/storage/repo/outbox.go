package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	pb "gmail-service/internal/pkg/genproto"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type OutboxRepo struct {
	db *sql.DB
}

func NewOutboxRepo(db *sql.DB) *OutboxRepo {
	return &OutboxRepo{db: db}
}

func (r *OutboxRepo) ArchiveMessage(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	query := `
		UPDATE outbox
		SET is_archived = NOT is_archived
		WHERE id = $1 AND deleted_at = 0 AND is_draft = false
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

func (r *OutboxRepo) Delete(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	query := `
		UPDATE outbox
		SET deleted_at = $1
		WHERE id = $2 AND is_draft = false AND deleted_at = 0
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

func (r *OutboxRepo) Get(ctx context.Context, req *pb.ByID) (*pb.OutboxMessageGetRes, error) {
	message := pb.OutboxMessageGetRes{
		Sender:        &pb.UserGetRes{},
		AttachmentIds: &pb.AttachmentIdsWrapper{},
		Receivers: &pb.Receivers{
			To:  &pb.MessageSendTo{},
			Cc:  &pb.MessageSendCC{},
			Bcc: &pb.MessageSendBCC{},
		},
	}

	query := `
		SELECT 
		    o.id,
		    o.subject,
		    o.body,
		    o.attachment_ids,
			o.receiver_to_emails,
			o.receiver_cc_emails,
			o.receiver_bcc_emails,
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
		FROM outbox AS o
		JOIN users AS u ON o.sender_id = u.id
		WHERE o.id = $1 AND o.deleted_at = 0 OR o.deleted_at = 1
	`
	row := r.db.QueryRowContext(ctx, query, req.Id)

	err := row.Scan(
		&message.Id,
		&message.Subject,
		&message.Body,
		pq.Array(&message.AttachmentIds.AttachmentIds),
		pq.Array(&message.Receivers.To.Emails),
		pq.Array(&message.Receivers.Cc.Emails),
		pq.Array(&message.Receivers.Bcc.Emails),
		&message.IsDraft,
		&message.IsArchived,
		&message.IsStarred,
		&message.SentAt,
		&message.DeletedAt,
		&message.Sender.Id,
		&message.Sender.FirstName,
		&message.Sender.LastName,
		&message.Sender.Dob,
		&message.Sender.Email,
		&message.Sender.Gender,
		&message.Sender.PfpUrl,
	)

	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (r *OutboxRepo) GetAll(ctx context.Context, req *pb.OutboxMessagesGetAllReq) (*pb.OutboxMessagesGetAllRes, error) {
	var messages []*pb.OutboxMessageGetRes

	query := `
		SELECT 
		    o.id,
		    o.subject,
		    o.body,
		    o.attachment_ids,
			o.receiver_to_emails,
			o.receiver_cc_emails,
			o.receiver_bcc_emails,
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
		FROM outbox AS o
		JOIN users AS u ON o.sender_id = u.id
		WHERE o.sender_id = $1
	`

	var args []interface{}
	argCount := 2
	args = append(args, req.SenderId)

	if req.Body.Query != "" {
		query += fmt.Sprintf(" AND (o.subject ILIKE $%d OR o.body ILIKE $%d)", argCount, argCount)
		args = append(args, "%"+req.Body.Query+"%")
		argCount++
	}
	if req.Body.IsTrashed {
		query += " AND o.deleted_at = 1"
	} else {
		query += " AND o.deleted_at = 0"
	}
	if req.Body.IsArchived {
		query += fmt.Sprintf(" AND o.is_archived = $%d", argCount)
		args = append(args, req.Body.IsArchived)
		argCount++
	}
	if req.Body.IsStarred {
		query += fmt.Sprintf(" AND o.is_starred = $%d", argCount)
		args = append(args, req.Body.IsStarred)
		argCount++
	}
	if req.Body.IsDraft {
		query += fmt.Sprintf(" AND o.is_draft = $%d", argCount)
		args = append(args, req.Body.IsDraft)
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

	query += fmt.Sprintf(" OFFSET $%d LIMIT $%d", argCount, argCount+1)
	args = append(args, req.Pagination.Skip, req.Pagination.Limit)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		message := pb.OutboxMessageGetRes{
			Sender:        &pb.UserGetRes{},
			AttachmentIds: &pb.AttachmentIdsWrapper{},
			Receivers: &pb.Receivers{
				To:  &pb.MessageSendTo{},
				Cc:  &pb.MessageSendCC{},
				Bcc: &pb.MessageSendBCC{},
			},
		}

		err = rows.Scan(
			&message.Id,
			&message.Subject,
			&message.Body,
			pq.Array(&message.AttachmentIds.AttachmentIds),
			pq.Array(&message.Receivers.To.Emails),
			pq.Array(&message.Receivers.Cc.Emails),
			pq.Array(&message.Receivers.Bcc.Emails),
			&message.IsDraft,
			&message.IsArchived,
			&message.IsStarred,
			&message.SentAt,
			&message.DeletedAt,
			&message.Sender.Id,
			&message.Sender.FirstName,
			&message.Sender.LastName,
			&message.Sender.Dob,
			&message.Sender.Email,
			&message.Sender.Gender,
			&message.Sender.PfpUrl,
		)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &pb.OutboxMessagesGetAllRes{
		Messages: messages,
	}, nil
}

func (r *OutboxRepo) MoveToTrash(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	query := `
		UPDATE outbox
		SET deleted_at = 
			CASE 
				WHEN deleted_at = 0 THEN 1 
				WHEN deleted_at = 1 THEN 0
				ELSE deleted_at
			END
		WHERE id = $1 AND is_draft = false AND deleted_at = 0 OR deleted_at = 1
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

func (r *OutboxRepo) Send(ctx context.Context, req *pb.OutboxMessageSentReq) (*pb.MessageSentRes, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := `
		INSERT INTO outbox 
			(subject, body, attachment_ids, receiver_to_emails, receiver_cc_emails, receiver_bcc_emails, sender_id, is_draft, sent_at)
		VALUES 
			($1, $2, $3, $4, $5, $6, $7, false, NOW()) 
		RETURNING id
	`

	var messageId uuid.UUID
	err = tx.QueryRowContext(
		ctx,
		query,
		req.Body.Subject,
		req.Body.Body,
		pq.Array(req.Body.AttachmentIds),
		pq.Array(req.Body.Receivers.To.Emails),
		pq.Array(req.Body.Receivers.Cc.Emails),
		pq.Array(req.Body.Receivers.Bcc.Emails),
		req.SenderId,
	).Scan(&messageId)

	if err != nil {
		if err2 := tx.Rollback(); err2 != nil {
			return nil, err
		}
		return nil, err
	}

	var sent int64 = 0
	var failed int64 = 0
	var failedEmails []string

	for _, receiver := range req.Body.Receivers.To.Emails {
		query = `
			INSERT INTO inbox (outbox_id, receiver_id, type)
			SELECT $1, id, 'to'
			FROM users
			WHERE email = $2
		`

		res, err := tx.ExecContext(ctx, query, messageId, receiver)
		if err != nil {
			failed++
			failedEmails = append(failedEmails, receiver)
			continue
		}

		rowsAffected, _ := res.RowsAffected()
		if rowsAffected == 0 {
			failed++
			failedEmails = append(failedEmails, receiver)
		} else {
			sent++
		}
	}
	for _, cc := range req.Body.Receivers.Cc.Emails {
		query = `
			INSERT INTO inbox (outbox_id, receiver_id, type)
			SELECT $1, id, 'cc'
			FROM users
			WHERE email = $2
		`

		res, err := tx.ExecContext(ctx, query, messageId, cc)
		if err != nil {
			failed++
			failedEmails = append(failedEmails, cc)
			continue
		}

		rowsAffected, _ := res.RowsAffected()
		if rowsAffected == 0 {
			failed++
			failedEmails = append(failedEmails, cc)
		} else {
			sent++
		}
	}
	for _, bcc := range req.Body.Receivers.Bcc.Emails {
		query = `
			INSERT INTO inbox (outbox_id, receiver_id, type)
			SELECT $1, id, 'bcc'
			FROM users
			WHERE email = $2
		`

		res, err := tx.ExecContext(ctx, query, messageId, bcc)
		if err != nil {
			failed++
			failedEmails = append(failedEmails, bcc)
			continue
		}

		rowsAffected, _ := res.RowsAffected()
		if rowsAffected == 0 {
			failed++
			failedEmails = append(failedEmails, bcc)
		} else {
			sent++
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &pb.MessageSentRes{
		TotalSent:    &sent,
		TotalFailed:  &failed,
		FailedEmails: failedEmails,
	}, nil
}

func (r *OutboxRepo) StarMessage(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	query := `
		UPDATE outbox
		SET is_starred = NOT is_starred
		WHERE id = $1 AND deleted_at = 0 AND is_draft = false
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
