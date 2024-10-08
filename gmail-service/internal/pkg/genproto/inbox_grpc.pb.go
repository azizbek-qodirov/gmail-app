// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.1
// source: gmailapp-submodule/inbox.proto

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
	InboxService_GetByID_FullMethodName        = "/protos.InboxService/GetByID"
	InboxService_GetAll_FullMethodName         = "/protos.InboxService/GetAll"
	InboxService_MoveToTrash_FullMethodName    = "/protos.InboxService/MoveToTrash"
	InboxService_Delete_FullMethodName         = "/protos.InboxService/Delete"
	InboxService_MarkAsRead_FullMethodName     = "/protos.InboxService/MarkAsRead"
	InboxService_MarkAsSpam_FullMethodName     = "/protos.InboxService/MarkAsSpam"
	InboxService_StarMessage_FullMethodName    = "/protos.InboxService/StarMessage"
	InboxService_ArchiveMessage_FullMethodName = "/protos.InboxService/ArchiveMessage"
)

// InboxServiceClient is the client API for InboxService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InboxServiceClient interface {
	GetByID(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*InboxMessageGetRes, error)
	GetAll(ctx context.Context, in *InboxMessageGetAllReq, opts ...grpc.CallOption) (*InboxMessagesGetAllRes, error)
	MoveToTrash(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error)
	Delete(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error)
	MarkAsRead(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error)
	MarkAsSpam(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error)
	StarMessage(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error)
	ArchiveMessage(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error)
}

type inboxServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInboxServiceClient(cc grpc.ClientConnInterface) InboxServiceClient {
	return &inboxServiceClient{cc}
}

func (c *inboxServiceClient) GetByID(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*InboxMessageGetRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InboxMessageGetRes)
	err := c.cc.Invoke(ctx, InboxService_GetByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inboxServiceClient) GetAll(ctx context.Context, in *InboxMessageGetAllReq, opts ...grpc.CallOption) (*InboxMessagesGetAllRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InboxMessagesGetAllRes)
	err := c.cc.Invoke(ctx, InboxService_GetAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inboxServiceClient) MoveToTrash(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, InboxService_MoveToTrash_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inboxServiceClient) Delete(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, InboxService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inboxServiceClient) MarkAsRead(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, InboxService_MarkAsRead_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inboxServiceClient) MarkAsSpam(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, InboxService_MarkAsSpam_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inboxServiceClient) StarMessage(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, InboxService_StarMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inboxServiceClient) ArchiveMessage(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, InboxService_ArchiveMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InboxServiceServer is the server API for InboxService service.
// All implementations must embed UnimplementedInboxServiceServer
// for forward compatibility
type InboxServiceServer interface {
	GetByID(context.Context, *ByID) (*InboxMessageGetRes, error)
	GetAll(context.Context, *InboxMessageGetAllReq) (*InboxMessagesGetAllRes, error)
	MoveToTrash(context.Context, *ByID) (*Void, error)
	Delete(context.Context, *ByID) (*Void, error)
	MarkAsRead(context.Context, *ByID) (*Void, error)
	MarkAsSpam(context.Context, *ByID) (*Void, error)
	StarMessage(context.Context, *ByID) (*Void, error)
	ArchiveMessage(context.Context, *ByID) (*Void, error)
	mustEmbedUnimplementedInboxServiceServer()
}

// UnimplementedInboxServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInboxServiceServer struct {
}

func (UnimplementedInboxServiceServer) GetByID(context.Context, *ByID) (*InboxMessageGetRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedInboxServiceServer) GetAll(context.Context, *InboxMessageGetAllReq) (*InboxMessagesGetAllRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedInboxServiceServer) MoveToTrash(context.Context, *ByID) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveToTrash not implemented")
}
func (UnimplementedInboxServiceServer) Delete(context.Context, *ByID) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedInboxServiceServer) MarkAsRead(context.Context, *ByID) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkAsRead not implemented")
}
func (UnimplementedInboxServiceServer) MarkAsSpam(context.Context, *ByID) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkAsSpam not implemented")
}
func (UnimplementedInboxServiceServer) StarMessage(context.Context, *ByID) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StarMessage not implemented")
}
func (UnimplementedInboxServiceServer) ArchiveMessage(context.Context, *ByID) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArchiveMessage not implemented")
}
func (UnimplementedInboxServiceServer) mustEmbedUnimplementedInboxServiceServer() {}

// UnsafeInboxServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InboxServiceServer will
// result in compilation errors.
type UnsafeInboxServiceServer interface {
	mustEmbedUnimplementedInboxServiceServer()
}

func RegisterInboxServiceServer(s grpc.ServiceRegistrar, srv InboxServiceServer) {
	s.RegisterService(&InboxService_ServiceDesc, srv)
}

func _InboxService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InboxServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InboxService_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InboxServiceServer).GetByID(ctx, req.(*ByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _InboxService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InboxMessageGetAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InboxServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InboxService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InboxServiceServer).GetAll(ctx, req.(*InboxMessageGetAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _InboxService_MoveToTrash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InboxServiceServer).MoveToTrash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InboxService_MoveToTrash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InboxServiceServer).MoveToTrash(ctx, req.(*ByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _InboxService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InboxServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InboxService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InboxServiceServer).Delete(ctx, req.(*ByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _InboxService_MarkAsRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InboxServiceServer).MarkAsRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InboxService_MarkAsRead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InboxServiceServer).MarkAsRead(ctx, req.(*ByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _InboxService_MarkAsSpam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InboxServiceServer).MarkAsSpam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InboxService_MarkAsSpam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InboxServiceServer).MarkAsSpam(ctx, req.(*ByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _InboxService_StarMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InboxServiceServer).StarMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InboxService_StarMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InboxServiceServer).StarMessage(ctx, req.(*ByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _InboxService_ArchiveMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InboxServiceServer).ArchiveMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InboxService_ArchiveMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InboxServiceServer).ArchiveMessage(ctx, req.(*ByID))
	}
	return interceptor(ctx, in, info, handler)
}

// InboxService_ServiceDesc is the grpc.ServiceDesc for InboxService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InboxService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.InboxService",
	HandlerType: (*InboxServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByID",
			Handler:    _InboxService_GetByID_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _InboxService_GetAll_Handler,
		},
		{
			MethodName: "MoveToTrash",
			Handler:    _InboxService_MoveToTrash_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _InboxService_Delete_Handler,
		},
		{
			MethodName: "MarkAsRead",
			Handler:    _InboxService_MarkAsRead_Handler,
		},
		{
			MethodName: "MarkAsSpam",
			Handler:    _InboxService_MarkAsSpam_Handler,
		},
		{
			MethodName: "StarMessage",
			Handler:    _InboxService_StarMessage_Handler,
		},
		{
			MethodName: "ArchiveMessage",
			Handler:    _InboxService_ArchiveMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gmailapp-submodule/inbox.proto",
}
