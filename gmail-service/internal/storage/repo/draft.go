package repo

import (
	"context"
	"database/sql"
	"errors"
	pb "gmail-service/internal/pkg/genproto"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type DraftRepo struct {
	db             *sql.DB
	attachmentRepo *AttachmentRepo
}

func NewDraftRepo(db *sql.DB) *DraftRepo {
	return &DraftRepo{db: db, attachmentRepo: NewAttachmentRepo(db)}
}

func (r *DraftRepo) Create(ctx context.Context, req *pb.DraftCreateUpdateReq) (*pb.Void, error) {
	query := `
		INSERT INTO outbox 
			(
				subject, 
				body, 
				attachment_ids, 
				sender_id, 
				receiver_to_emails, 
				receiver_cc_emails, 
				receiver_bcc_emails, 
				is_draft
			)
		VALUES ($1, $2, $3, $4, true)
	`
	res, err := r.db.ExecContext(
		ctx,
		query,
		req.Body.Subject,
		req.Body.Body,
		pq.Array(req.Body.AttachmentIds),
		req.SenderId,
		pq.Array(req.Body.Receivers.To.Emails),
		pq.Array(req.Body.Receivers.Cc.Emails),
		pq.Array(req.Body.Receivers.Bcc.Emails),
	)
	if err != nil {
		return nil, err
	}

	if aff, err := res.RowsAffected(); aff == 0 || err != nil {
		return nil, errors.New(err.Error())
	}

	return nil, nil
}

func (r *DraftRepo) Update(ctx context.Context, req *pb.DraftCreateUpdateReq) (*pb.Void, error) {
	query := `
		UPDATE outbox
		SET subject = $1,
			body = $2,
			attachment_ids = $3
			receiver_to_emails = $4,
			receiver_cc_emails = $5,
			receiver_bcc_emails = $6
		WHERE id = $7 AND is_draft = true
	`
	res, err := r.db.ExecContext(
		ctx,
		query,
		req.Body.Subject,
		req.Body.Body,
		pq.Array(req.Body.AttachmentIds),
		pq.Array(req.Body.Receivers.To.Emails),
		pq.Array(req.Body.Receivers.Cc.Emails),
		pq.Array(req.Body.Receivers.Bcc.Emails),
		req.SenderId, // this actually is draft id
	)

	if err != nil {
		return nil, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		return nil, errors.New("draft not found")
	}

	return &pb.Void{}, nil
}

func (r *DraftRepo) Delete(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	query := `
		DELETE FROM outbox
		WHERE id = $1 AND is_draft = true
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

func (r *DraftRepo) SendDraft(ctx context.Context, req *pb.ByID) (*pb.MessageSentRes, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var draft pb.OutboxMessageSentBody
	var draftId uuid.UUID

	query := `
		SELECT id, subject, body, attachment_ids 
		FROM outbox
		WHERE id = $1 AND is_draft = true
	`
	row := tx.QueryRowContext(ctx, query, req.Id)
	err = row.Scan(
		&draftId,
		&draft.Subject,
		&draft.Body,
		&draft.AttachmentIds,
	)
	if err != nil {
		return nil, err
	}

	query = `
		UPDATE outbox
		SET is_draft = false,
			sent_at = $1
		WHERE id = $2
	`
	_, err = tx.ExecContext(ctx, query, time.Now(), draftId)
	if err != nil {
		return nil, err
	}

	var sent int64 = 0
	var failed int64 = 0
	var failedEmails []string

	for _, receiver := range draft.Receivers.To.Emails {
		query = `
			INSERT INTO inbox (outbox_id, receiver_id, type)
			SELECT $1, id, 'to'
			FROM users
			WHERE email = $2
		`
		res, err := tx.ExecContext(ctx, query, draftId, receiver)
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

	for _, cc := range draft.Receivers.Cc.Emails {
		query = `
			INSERT INTO inbox (outbox_id, receiver_id, type)
			SELECT $1, id, 'cc'
			FROM users
			WHERE email = $2
		`
		res, err := tx.ExecContext(ctx, query, draftId, cc)
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

	for _, bcc := range draft.Receivers.Bcc.Emails {
		query = `
			INSERT INTO inbox (outbox_id, receiver_id, type)
			SELECT $1, id, 'bcc'
			FROM users
			WHERE email = $2
		`
		res, err := tx.ExecContext(ctx, query, draftId, bcc)
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
		TotalSent:    sent,
		TotalFailed:  failed,
		FailedEmails: failedEmails,
	}, nil
}
