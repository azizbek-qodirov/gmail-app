package service

import (
	"context"
	pb "gmail-service/internal/pkg/genproto"

	"gmail-service/internal/storage"
)

type DraftService struct {
	stg storage.StorageI
	pb.UnimplementedDraftServiceServer
}

func NewDraftService(stg storage.StorageI) *DraftService {
	return &DraftService{stg: stg}
}

func (s *DraftService) Create(ctx context.Context, req *pb.DraftCreateUpdateReq) (*pb.Void, error) {
	return s.stg.Draft().Create(ctx, req)
}

func (s *DraftService) Delete(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	return s.stg.Draft().Delete(ctx, req)
}

func (s *DraftService) SendDraft(ctx context.Context, req *pb.ByID) (*pb.MessageSentRes, error) {
	return s.stg.Draft().SendDraft(ctx, req)
}

func (s *DraftService) Update(ctx context.Context, req *pb.DraftCreateUpdateReq) (*pb.Void, error) {
	return s.stg.Draft().Update(ctx, req)
}
