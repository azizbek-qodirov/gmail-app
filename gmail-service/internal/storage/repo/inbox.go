package repo

import (
	"context"
	"database/sql"
	"errors"
	pb "gmail-service/internal/pkg/genproto"
	"time"
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
	res, err := r.db.ExecContext(ctx, query, time.Now(), req.Id)
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
		WHERE i.receiver_id = $1 AND i.deleted_at = 0
	`

	rows, err := r.db.QueryContext(ctx, query, req.Body.SenderId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var message pb.InboxMessageGetRes
		message.Outbox = &pb.OutboxMessageGetRes{}
		message.Outbox.Sender = &pb.UserGetRes{}

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
			&message.Outbox.AttachmentIds,
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
	message.Outbox = &pb.OutboxMessageGetRes{}
	message.Outbox.Sender = &pb.UserGetRes{}

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
		&message.Outbox.AttachmentIds,
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

	res, err := r.db.ExecContext(ctx, query, time.Now(), req.Id)
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
		SET deleted_at = $1
		WHERE id = $2
	`
	res, err := r.db.ExecContext(ctx, query, time.Now(), req.Id)
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
