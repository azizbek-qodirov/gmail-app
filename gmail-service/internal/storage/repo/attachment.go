package repo

import (
	"context"
	"database/sql"
	"errors"
	pb "gmail-service/internal/pkg/genproto"
	"time"

	"github.com/google/uuid"
)

type AttachmentRepo struct {
	db *sql.DB
}

func NewAttachmentRepo(db *sql.DB) *AttachmentRepo {
	return &AttachmentRepo{db: db}
}

func (r *AttachmentRepo) Create(ctx context.Context, req *pb.AttachmentCreateReq) (*pb.AttachmentCreateRes, error) {
	query := `
		INSERT INTO attachments (file_url, file_name, file_size, mime_type, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING file_url
	`

	var fileUrl string
	err := r.db.QueryRowContext(
		ctx,
		query,
		req.FileUrl,
		req.FileName,
		req.FileSize,
		req.MimeType,
		time.Now(),
	).Scan(&fileUrl)

	if err != nil {
		return nil, err
	}

	return &pb.AttachmentCreateRes{
		FileUrl: fileUrl,
	}, nil
}

func (r *AttachmentRepo) GetByID(ctx context.Context, req *pb.ByID) (*pb.AttachmentGetRes, error) {
	var attachment pb.AttachmentGetRes

	query := `
		SELECT 
			id,
			file_url,
			file_name,
			file_size,
			mime_type,
		FROM attachments
		WHERE id = $1 AND deleted_at = 0
	`

	row := r.db.QueryRowContext(ctx, query, req.Id)
	err := row.Scan(
		&attachment.Id,
		&attachment.FileUrl,
		&attachment.FileName,
		&attachment.FileSize,
		&attachment.MimeType,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("attachment not found")
		}
		return nil, err
	}

	return &attachment, nil
}

func (r *AttachmentRepo) IsExists(ctx context.Context, req *pb.ByID) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM attachments WHERE id = $1)`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, req.Id).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *AttachmentRepo) Delete(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	query := `
		DELETE FROM attachments
		WHERE id = $1
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

	return nil, nil
}

func (r *AttachmentRepo) GetAll(ctx context.Context, req *pb.AttachmentGetAllReq) (*pb.AttachmentGetAllRes, error) {
	var attachments []*pb.AttachmentGetRes

	var attachmentIDs []uuid.UUID
	query := `SELECT attachment_ids FROM outbox WHERE id = $1 AND deleted_at = 0`
	err := r.db.QueryRowContext(ctx, query, req.OutboxId).Scan(&attachmentIDs)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("outbox not found")
		}
		return nil, err
	}

	if len(attachmentIDs) == 0 {
		return &pb.AttachmentGetAllRes{Attachments: attachments}, nil // Return empty result
	}

	query = `
		SELECT 
			id,
			file_url,
			file_name,
			file_size,
			mime_type
		FROM attachments
		WHERE id = ANY($1) AND deleted_at = 0
	`
	rows, err := r.db.QueryContext(ctx, query, attachmentIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var attachment pb.AttachmentGetRes
		err = rows.Scan(
			&attachment.Id,
			&attachment.FileUrl,
			&attachment.FileName,
			&attachment.FileSize,
			&attachment.MimeType,
		)
		if err != nil {
			return nil, err
		}
		attachments = append(attachments, &attachment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &pb.AttachmentGetAllRes{
		Attachments: attachments,
	}, nil
}
