package service

import (
	"context"

	pb "gmail-service/internal/pkg/genproto"

	"gmail-service/internal/storage"
)

type UserService struct {
	stg storage.StorageI
	pb.UnimplementedUserServiceServer
}

func NewUserService(stg storage.StorageI) *UserService {
	return &UserService{stg: stg}
}

func (s *UserService) ChangeUserPFP(ctx context.Context, req *pb.UserChangePFPReq) (*pb.Void, error) {
	return s.stg.User().ChangeUserPFP(ctx, req)
}

func (s *UserService) ChangeUserPassword(ctx context.Context, req *pb.UserRecoverPasswordReq) (*pb.Void, error) {
	return s.stg.User().ChangeUserPassword(ctx, req)
}

func (s *UserService) ConfirmUser(ctx context.Context, req *pb.ByEmail) (*pb.Void, error) {
	return s.stg.User().ConfirmUser(ctx, req)
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.UserCreateReq) (*pb.Void, error) {
	return s.stg.User().CreateUser(ctx, req)
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	return s.stg.User().DeleteUser(ctx, req)
}

func (s *UserService) GetUserByEmail(ctx context.Context, req *pb.ByEmail) (*pb.UserGetRes, error) {
	return s.stg.User().GetUserByEmail(ctx, req)
}

func (s *UserService) GetUserByID(ctx context.Context, req *pb.ByID) (*pb.UserGetRes, error) {
	return s.stg.User().GetUserByID(ctx, req)
}

func (s *UserService) GetUserSecurityByEmail(ctx context.Context, req *pb.ByEmail) (*pb.UserGetSecurityRes, error) {
	return s.stg.User().GetUserSecurityByEmail(ctx, req)
}

func (s *UserService) IsEmailExists(ctx context.Context, req *pb.ByEmail) (*pb.UserEmailCheckRes, error) {
	return s.stg.User().IsEmailExists(ctx, req)
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UserUpdateReq) (*pb.Void, error) {
	return s.stg.User().UpdateUser(ctx, req)
}
