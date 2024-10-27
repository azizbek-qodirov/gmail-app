package service

import (
	"context"
	"database/sql"
	"errors"
	pb "gmail-service/internal/pkg/genproto"
	"gmail-service/internal/storage/repo"

	"gmail-service/internal/storage"

	"github.com/google/uuid"
)

var ctx = context.Background()

type DraftService struct {
	stg            storage.StorageI
	attachmentRepo *repo.AttachmentRepo

	pb.UnimplementedDraftServiceServer
}

func NewDraftService(stg storage.StorageI, db *sql.DB) *DraftService {
	return &DraftService{
		stg:            stg,
		attachmentRepo: repo.NewAttachmentRepo(db),
	}
}

func (s *DraftService) Create(ctx context.Context, req *pb.DraftCreateUpdateReq) (*pb.Void, error) {
	err := s.Validate(req)
	if err != nil {
		return nil, err
	}

	return s.stg.Draft().Create(ctx, req)
}

func (s *DraftService) Update(ctx context.Context, req *pb.DraftCreateUpdateReq) (*pb.Void, error) {
	err := s.Validate(req)
	if err != nil {
		return nil, err
	}

	return s.stg.Draft().Update(ctx, req)
}

func (s *DraftService) Delete(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	return s.stg.Draft().Delete(ctx, req)
}

func (s *DraftService) SendDraft(ctx context.Context, req *pb.ByID) (*pb.MessageSentRes, error) {
	return s.stg.Draft().SendDraft(ctx, req)
}

func (s *DraftService) Validate(req *pb.DraftCreateUpdateReq) error {
	// ignore whitespaces from attachment_ids
	no_ws := []string{}
	for _, v := range req.Body.AttachmentIds {
		if v != "" {
			no_ws = append(no_ws, v)
		}
	}
	req.Body.AttachmentIds = no_ws

	// check every incoming element if they are valid uuid
	for _, v := range req.Body.AttachmentIds {
		err := uuid.Validate(v)
		if err != nil {
			return errors.New("invalid attachment uuid")
		}
	}

	// check if all attachments exists in database
	for _, v := range req.Body.AttachmentIds {
		exists, err := s.attachmentRepo.IsExists(ctx, &pb.ByID{Id: v})
		if err != nil {
			return err
		}
		if !exists {
			return errors.New("attachment with id " + v + " not found")
		}
	}

	// ignore whitespaces from emails
	no_ws = []string{}
	for _, v := range req.Body.Receivers.To.Emails {
		if v != "" {
			no_ws = append(no_ws, v)
		}
	}
	req.Body.Receivers.To.Emails = no_ws

	no_ws = []string{}
	for _, v := range req.Body.Receivers.Cc.Emails {
		if v != "" {
			no_ws = append(no_ws, v)
		}
	}
	req.Body.Receivers.Cc.Emails = no_ws

	no_ws = []string{}
	for _, v := range req.Body.Receivers.Bcc.Emails {
		if v != "" {
			no_ws = append(no_ws, v)
		}
	}
	req.Body.Receivers.Bcc.Emails = no_ws

	return nil
}
