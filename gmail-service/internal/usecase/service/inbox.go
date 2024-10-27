package service

import (
	"context"
	pb "gmail-service/internal/pkg/genproto"

	"gmail-service/internal/storage"
)

type InboxService struct {
	stg storage.StorageI
	pb.UnimplementedInboxServiceServer
}

func NewInboxService(stg storage.StorageI) *InboxService {
	return &InboxService{stg: stg}
}

func (s *InboxService) ArchiveMessage(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	return s.stg.Inbox().ArchiveMessage(ctx, req)
}

func (s *InboxService) Delete(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	return s.stg.Inbox().Delete(ctx, req)
}

func (s *InboxService) GetAll(ctx context.Context, req *pb.InboxMessageGetAllReq) (*pb.InboxMessagesGetAllRes, error) {
	return s.stg.Inbox().GetAll(ctx, req)
}

func (s *InboxService) GetByID(ctx context.Context, req *pb.ByID) (*pb.InboxMessageGetRes, error) {
	return s.stg.Inbox().GetByID(ctx, req)
}

func (s *InboxService) MarkAsRead(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	return s.stg.Inbox().MarkAsRead(ctx, req)
}

func (s *InboxService) MarkAsSpam(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	return s.stg.Inbox().MarkAsSpam(ctx, req)
}

func (s *InboxService) MoveToTrash(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	return s.stg.Inbox().MoveToTrash(ctx, req)
}

func (s *InboxService) StarMessage(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	return s.stg.Inbox().StarMessage(ctx, req)
}
