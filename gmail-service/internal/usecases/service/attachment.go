package service

import (
	"context"
	pb "gmail-service/internal/pkg/genproto"

	"gmail-service/internal/storage"
)

type AttachmentService struct {
	stg storage.StorageI
	pb.UnimplementedAttachmentServiceServer
}

func NewAttachmentService(stg storage.StorageI) *AttachmentService {
	return &AttachmentService{stg: stg}
}

func (s *AttachmentService) Create(ctx context.Context, req *pb.AttachmentCreateReq) (*pb.AttachmentCreateRes, error) {
	return s.stg.Attachment().Create(ctx, req)
}

func (s *AttachmentService) GetByID(ctx context.Context, req *pb.ByID) (*pb.AttachmentGetRes, error) {
	return s.stg.Attachment().GetByID(ctx, req)
}

func (s *AttachmentService) Delete(ctx context.Context, req *pb.ByID) (*pb.AttachmentDeleteRes, error) {
	return s.stg.Attachment().Delete(ctx, req)
}

func (s *AttachmentService) GetAll(ctx context.Context, req *pb.AttachmentGetAllReq) (*pb.AttachmentGetAllRes, error) {
	return s.stg.Attachment().GetAll(ctx, req)
}

func (s AttachmentService) GetMyUploads(ctx context.Context, req *pb.ByID) (*pb.AttachmentGetAllRes, error) {
	return s.stg.Attachment().GetMyUploads(ctx, req)
}
