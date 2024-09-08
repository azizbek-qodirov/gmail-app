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

type OutboxService struct {
	stg            storage.StorageI
	attachmentRepo *repo.AttachmentRepo

	pb.UnimplementedOutboxServiceServer
}

func NewOutboxService(stg storage.StorageI, db *sql.DB) *OutboxService {
	return &OutboxService{stg: stg, attachmentRepo: repo.NewAttachmentRepo(db)}
}

func (s *OutboxService) ArchiveMessage(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	return s.stg.Outbox().ArchiveMessage(ctx, req)
}

func (s *OutboxService) Delete(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	return s.stg.Outbox().Delete(ctx, req)
}

func (s *OutboxService) Get(ctx context.Context, req *pb.ByID) (*pb.OutboxMessageGetRes, error) {
	return s.stg.Outbox().Get(ctx, req)
}

func (s *OutboxService) GetAll(ctx context.Context, req *pb.OutboxMessagesGetAllReq) (*pb.OutboxMessagesGetAllRes, error) {
	return s.stg.Outbox().GetAll(ctx, req)
}

func (s *OutboxService) MoveToTrash(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	return s.stg.Outbox().MoveToTrash(ctx, req)
}

func (s *OutboxService) Send(ctx context.Context, req *pb.OutboxMessageSentReq) (*pb.MessageSentRes, error) {
	err := s.Validate(req)
	if err != nil {
		return nil, err
	}

	return s.stg.Outbox().Send(ctx, req)
}

func (s *OutboxService) StarMessage(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	return s.stg.Outbox().StarMessage(ctx, req)
}

func (s *OutboxService) Validate(req *pb.OutboxMessageSentReq) error {
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
