// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.1
// source: gmailapp-submodule/user.proto

package genproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	UserService_CreateUser_FullMethodName             = "/protos.UserService/CreateUser"
	UserService_ConfirmUser_FullMethodName            = "/protos.UserService/ConfirmUser"
	UserService_GetUserByID_FullMethodName            = "/protos.UserService/GetUserByID"
	UserService_GetUserByEmail_FullMethodName         = "/protos.UserService/GetUserByEmail"
	UserService_GetUserSecurityByEmail_FullMethodName = "/protos.UserService/GetUserSecurityByEmail"
	UserService_UpdateUser_FullMethodName             = "/protos.UserService/UpdateUser"
	UserService_ChangeUserPFP_FullMethodName          = "/protos.UserService/ChangeUserPFP"
	UserService_ChangeUserPassword_FullMethodName     = "/protos.UserService/ChangeUserPassword"
	UserService_DeleteUser_FullMethodName             = "/protos.UserService/DeleteUser"
	UserService_IsEmailExists_FullMethodName          = "/protos.UserService/IsEmailExists"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *UserCreateReq, opts ...grpc.CallOption) (*Void, error)
	ConfirmUser(ctx context.Context, in *ByEmail, opts ...grpc.CallOption) (*Void, error)
	GetUserByID(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*UserGetRes, error)
	GetUserByEmail(ctx context.Context, in *ByEmail, opts ...grpc.CallOption) (*UserGetRes, error)
	GetUserSecurityByEmail(ctx context.Context, in *ByEmail, opts ...grpc.CallOption) (*UserGetSecurityRes, error)
	UpdateUser(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*Void, error)
	ChangeUserPFP(ctx context.Context, in *UserChangePFPReq, opts ...grpc.CallOption) (*Void, error)
	ChangeUserPassword(ctx context.Context, in *UserRecoverPasswordReq, opts ...grpc.CallOption) (*Void, error)
	DeleteUser(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error)
	IsEmailExists(ctx context.Context, in *ByEmail, opts ...grpc.CallOption) (*UserEmailCheckRes, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *UserCreateReq, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, UserService_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ConfirmUser(ctx context.Context, in *ByEmail, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, UserService_ConfirmUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByID(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*UserGetRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserGetRes)
	err := c.cc.Invoke(ctx, UserService_GetUserByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByEmail(ctx context.Context, in *ByEmail, opts ...grpc.CallOption) (*UserGetRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserGetRes)
	err := c.cc.Invoke(ctx, UserService_GetUserByEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserSecurityByEmail(ctx context.Context, in *ByEmail, opts ...grpc.CallOption) (*UserGetSecurityRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserGetSecurityRes)
	err := c.cc.Invoke(ctx, UserService_GetUserSecurityByEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, UserService_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangeUserPFP(ctx context.Context, in *UserChangePFPReq, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, UserService_ChangeUserPFP_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangeUserPassword(ctx context.Context, in *UserRecoverPasswordReq, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, UserService_ChangeUserPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, UserService_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) IsEmailExists(ctx context.Context, in *ByEmail, opts ...grpc.CallOption) (*UserEmailCheckRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserEmailCheckRes)
	err := c.cc.Invoke(ctx, UserService_IsEmailExists_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	CreateUser(context.Context, *UserCreateReq) (*Void, error)
	ConfirmUser(context.Context, *ByEmail) (*Void, error)
	GetUserByID(context.Context, *ByID) (*UserGetRes, error)
	GetUserByEmail(context.Context, *ByEmail) (*UserGetRes, error)
	GetUserSecurityByEmail(context.Context, *ByEmail) (*UserGetSecurityRes, error)
	UpdateUser(context.Context, *UserUpdateReq) (*Void, error)
	ChangeUserPFP(context.Context, *UserChangePFPReq) (*Void, error)
	ChangeUserPassword(context.Context, *UserRecoverPasswordReq) (*Void, error)
	DeleteUser(context.Context, *ByID) (*Void, error)
	IsEmailExists(context.Context, *ByEmail) (*UserEmailCheckRes, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *UserCreateReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) ConfirmUser(context.Context, *ByEmail) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmUser not implemented")
}
func (UnimplementedUserServiceServer) GetUserByID(context.Context, *ByID) (*UserGetRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByID not implemented")
}
func (UnimplementedUserServiceServer) GetUserByEmail(context.Context, *ByEmail) (*UserGetRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByEmail not implemented")
}
func (UnimplementedUserServiceServer) GetUserSecurityByEmail(context.Context, *ByEmail) (*UserGetSecurityRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserSecurityByEmail not implemented")
}
func (UnimplementedUserServiceServer) UpdateUser(context.Context, *UserUpdateReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServiceServer) ChangeUserPFP(context.Context, *UserChangePFPReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUserPFP not implemented")
}
func (UnimplementedUserServiceServer) ChangeUserPassword(context.Context, *UserRecoverPasswordReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUserPassword not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *ByID) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) IsEmailExists(context.Context, *ByEmail) (*UserEmailCheckRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsEmailExists not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*UserCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ConfirmUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByEmail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ConfirmUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ConfirmUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ConfirmUser(ctx, req.(*ByEmail))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByID(ctx, req.(*ByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByEmail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByEmail(ctx, req.(*ByEmail))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserSecurityByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByEmail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserSecurityByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserSecurityByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserSecurityByEmail(ctx, req.(*ByEmail))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UserUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangeUserPFP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserChangePFPReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangeUserPFP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ChangeUserPFP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangeUserPFP(ctx, req.(*UserChangePFPReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangeUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRecoverPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangeUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ChangeUserPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangeUserPassword(ctx, req.(*UserRecoverPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*ByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_IsEmailExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByEmail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).IsEmailExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_IsEmailExists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).IsEmailExists(ctx, req.(*ByEmail))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "ConfirmUser",
			Handler:    _UserService_ConfirmUser_Handler,
		},
		{
			MethodName: "GetUserByID",
			Handler:    _UserService_GetUserByID_Handler,
		},
		{
			MethodName: "GetUserByEmail",
			Handler:    _UserService_GetUserByEmail_Handler,
		},
		{
			MethodName: "GetUserSecurityByEmail",
			Handler:    _UserService_GetUserSecurityByEmail_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "ChangeUserPFP",
			Handler:    _UserService_ChangeUserPFP_Handler,
		},
		{
			MethodName: "ChangeUserPassword",
			Handler:    _UserService_ChangeUserPassword_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
		{
			MethodName: "IsEmailExists",
			Handler:    _UserService_IsEmailExists_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gmailapp-submodule/user.proto",
}
