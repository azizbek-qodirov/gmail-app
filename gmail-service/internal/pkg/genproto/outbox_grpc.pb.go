// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.1
// source: gmailapp-submodule/outbox.proto

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
	OutboxService_Send_FullMethodName           = "/protos.OutboxService/Send"
	OutboxService_Get_FullMethodName            = "/protos.OutboxService/Get"
	OutboxService_GetAll_FullMethodName         = "/protos.OutboxService/GetAll"
	OutboxService_MoveToTrash_FullMethodName    = "/protos.OutboxService/MoveToTrash"
	OutboxService_Delete_FullMethodName         = "/protos.OutboxService/Delete"
	OutboxService_StarMessage_FullMethodName    = "/protos.OutboxService/StarMessage"
	OutboxService_ArchiveMessage_FullMethodName = "/protos.OutboxService/ArchiveMessage"
)

// OutboxServiceClient is the client API for OutboxService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OutboxServiceClient interface {
	Send(ctx context.Context, in *OutboxMessageSentReq, opts ...grpc.CallOption) (*MessageSentRes, error)
	Get(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*OutboxMessageGetRes, error)
	GetAll(ctx context.Context, in *OutboxMessagesGetAllReq, opts ...grpc.CallOption) (*OutboxMessagesGetAllRes, error)
	MoveToTrash(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error)
	Delete(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error)
	StarMessage(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error)
	ArchiveMessage(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error)
}

type outboxServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOutboxServiceClient(cc grpc.ClientConnInterface) OutboxServiceClient {
	return &outboxServiceClient{cc}
}

func (c *outboxServiceClient) Send(ctx context.Context, in *OutboxMessageSentReq, opts ...grpc.CallOption) (*MessageSentRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MessageSentRes)
	err := c.cc.Invoke(ctx, OutboxService_Send_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outboxServiceClient) Get(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*OutboxMessageGetRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OutboxMessageGetRes)
	err := c.cc.Invoke(ctx, OutboxService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outboxServiceClient) GetAll(ctx context.Context, in *OutboxMessagesGetAllReq, opts ...grpc.CallOption) (*OutboxMessagesGetAllRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OutboxMessagesGetAllRes)
	err := c.cc.Invoke(ctx, OutboxService_GetAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outboxServiceClient) MoveToTrash(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, OutboxService_MoveToTrash_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outboxServiceClient) Delete(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, OutboxService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outboxServiceClient) StarMessage(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, OutboxService_StarMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outboxServiceClient) ArchiveMessage(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, OutboxService_ArchiveMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OutboxServiceServer is the server API for OutboxService service.
// All implementations must embed UnimplementedOutboxServiceServer
// for forward compatibility
type OutboxServiceServer interface {
	Send(context.Context, *OutboxMessageSentReq) (*MessageSentRes, error)
	Get(context.Context, *ByID) (*OutboxMessageGetRes, error)
	GetAll(context.Context, *OutboxMessagesGetAllReq) (*OutboxMessagesGetAllRes, error)
	MoveToTrash(context.Context, *ByID) (*Void, error)
	Delete(context.Context, *ByID) (*Void, error)
	StarMessage(context.Context, *ByID) (*Void, error)
	ArchiveMessage(context.Context, *ByID) (*Void, error)
	mustEmbedUnimplementedOutboxServiceServer()
}

// UnimplementedOutboxServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOutboxServiceServer struct {
}

func (UnimplementedOutboxServiceServer) Send(context.Context, *OutboxMessageSentReq) (*MessageSentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedOutboxServiceServer) Get(context.Context, *ByID) (*OutboxMessageGetRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedOutboxServiceServer) GetAll(context.Context, *OutboxMessagesGetAllReq) (*OutboxMessagesGetAllRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedOutboxServiceServer) MoveToTrash(context.Context, *ByID) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveToTrash not implemented")
}
func (UnimplementedOutboxServiceServer) Delete(context.Context, *ByID) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedOutboxServiceServer) StarMessage(context.Context, *ByID) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StarMessage not implemented")
}
func (UnimplementedOutboxServiceServer) ArchiveMessage(context.Context, *ByID) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArchiveMessage not implemented")
}
func (UnimplementedOutboxServiceServer) mustEmbedUnimplementedOutboxServiceServer() {}

// UnsafeOutboxServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OutboxServiceServer will
// result in compilation errors.
type UnsafeOutboxServiceServer interface {
	mustEmbedUnimplementedOutboxServiceServer()
}

func RegisterOutboxServiceServer(s grpc.ServiceRegistrar, srv OutboxServiceServer) {
	s.RegisterService(&OutboxService_ServiceDesc, srv)
}

func _OutboxService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OutboxMessageSentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboxServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OutboxService_Send_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboxServiceServer).Send(ctx, req.(*OutboxMessageSentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OutboxService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboxServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OutboxService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboxServiceServer).Get(ctx, req.(*ByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _OutboxService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OutboxMessagesGetAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboxServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OutboxService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboxServiceServer).GetAll(ctx, req.(*OutboxMessagesGetAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OutboxService_MoveToTrash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboxServiceServer).MoveToTrash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OutboxService_MoveToTrash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboxServiceServer).MoveToTrash(ctx, req.(*ByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _OutboxService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboxServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OutboxService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboxServiceServer).Delete(ctx, req.(*ByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _OutboxService_StarMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboxServiceServer).StarMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OutboxService_StarMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboxServiceServer).StarMessage(ctx, req.(*ByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _OutboxService_ArchiveMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboxServiceServer).ArchiveMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OutboxService_ArchiveMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboxServiceServer).ArchiveMessage(ctx, req.(*ByID))
	}
	return interceptor(ctx, in, info, handler)
}

// OutboxService_ServiceDesc is the grpc.ServiceDesc for OutboxService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OutboxService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.OutboxService",
	HandlerType: (*OutboxServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _OutboxService_Send_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _OutboxService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _OutboxService_GetAll_Handler,
		},
		{
			MethodName: "MoveToTrash",
			Handler:    _OutboxService_MoveToTrash_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _OutboxService_Delete_Handler,
		},
		{
			MethodName: "StarMessage",
			Handler:    _OutboxService_StarMessage_Handler,
		},
		{
			MethodName: "ArchiveMessage",
			Handler:    _OutboxService_ArchiveMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gmailapp-submodule/outbox.proto",
}