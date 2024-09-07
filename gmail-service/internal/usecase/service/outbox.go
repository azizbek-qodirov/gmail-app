package service

import (
	"context"
	pb "gmail-service/internal/pkg/genproto"

	"gmail-service/internal/storage"
)

type OutboxService struct {
	stg storage.StorageI
	pb.UnimplementedOutboxServiceServer
}

func NewOutboxService(stg storage.StorageI) *OutboxService {
	return &OutboxService{stg: stg}
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
	return s.stg.Outbox().Send(ctx, req)
}

func (s *OutboxService) StarMessage(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	return s.stg.Outbox().StarMessage(ctx, req)
}
