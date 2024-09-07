package repo

import (
	"context"
	"database/sql"
	"errors"
	pb "gmail-service/internal/pkg/genproto"
	"time"

	"github.com/google/uuid"
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
		WHERE id = $2 AND is_draft = false
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

func (r *OutboxRepo) Get(ctx context.Context, req *pb.ByID) (*pb.OutboxMessageGetRes, error) {
	var message pb.OutboxMessageGetRes
	message.Sender = &pb.UserGetRes{}
	query := `
		SELECT 
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
		FROM outbox AS o
		JOIN users AS u ON o.sender_id = u.id
		WHERE o.id = $1 AND o.deleted_at = 0
	`
	row := r.db.QueryRowContext(ctx, query, req.Id)

	err := row.Scan(
		&message.Id,
		&message.Subject,
		&message.Body,
		&message.AttachmentIds,
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
		WHERE o.deleted_at = 0
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var message pb.OutboxMessageGetRes
		message.Sender = &pb.UserGetRes{}
		err = rows.Scan(
			&message.Id,
			&message.Subject,
			&message.Body,
			&message.AttachmentIds,
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
		SET deleted_at = $1
		WHERE id = $2 AND is_draft = false
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

func (r *OutboxRepo) Send(ctx context.Context, req *pb.OutboxMessageSentReq) (*pb.MessageSentRes, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := `
		INSERT INTO outbox (subject, body, attachment_ids, sender_id, is_draft, sent_at)
		VALUES ($1, $2, $3, $4, false, $5)
		RETURNING id
	`

	var messageId uuid.UUID
	err = tx.QueryRowContext(
		ctx,
		query,
		req.Body.Subject,
		req.Body.Body,
		req.Body.AttachmentIds,
		req.SenderId,
		time.Now(),
	).Scan(&messageId)

	if err != nil {
		return nil, err
	}

	var sent int64 = 0
	var failed int64 = 0
	var failedEmails []string

	for _, receiver := range req.Body.Receivers.To.Email {
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

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &pb.MessageSentRes{
		TotalSent:    sent,
		TotalFailed:  failed,
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
