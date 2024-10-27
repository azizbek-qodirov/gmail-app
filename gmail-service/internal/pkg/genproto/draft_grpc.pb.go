// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.1
// source: gmailapp-submodule/draft.proto

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
	DraftService_Create_FullMethodName    = "/protos.DraftService/Create"
	DraftService_Update_FullMethodName    = "/protos.DraftService/Update"
	DraftService_Delete_FullMethodName    = "/protos.DraftService/Delete"
	DraftService_SendDraft_FullMethodName = "/protos.DraftService/SendDraft"
)

// DraftServiceClient is the client API for DraftService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DraftServiceClient interface {
	Create(ctx context.Context, in *DraftCreateUpdateReq, opts ...grpc.CallOption) (*Void, error)
	Update(ctx context.Context, in *DraftCreateUpdateReq, opts ...grpc.CallOption) (*Void, error)
	Delete(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error)
	SendDraft(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*MessageSentRes, error)
}

type draftServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDraftServiceClient(cc grpc.ClientConnInterface) DraftServiceClient {
	return &draftServiceClient{cc}
}

func (c *draftServiceClient) Create(ctx context.Context, in *DraftCreateUpdateReq, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, DraftService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *draftServiceClient) Update(ctx context.Context, in *DraftCreateUpdateReq, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, DraftService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *draftServiceClient) Delete(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, DraftService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *draftServiceClient) SendDraft(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*MessageSentRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MessageSentRes)
	err := c.cc.Invoke(ctx, DraftService_SendDraft_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DraftServiceServer is the server API for DraftService service.
// All implementations must embed UnimplementedDraftServiceServer
// for forward compatibility
type DraftServiceServer interface {
	Create(context.Context, *DraftCreateUpdateReq) (*Void, error)
	Update(context.Context, *DraftCreateUpdateReq) (*Void, error)
	Delete(context.Context, *ByID) (*Void, error)
	SendDraft(context.Context, *ByID) (*MessageSentRes, error)
	mustEmbedUnimplementedDraftServiceServer()
}

// UnimplementedDraftServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDraftServiceServer struct {
}

func (UnimplementedDraftServiceServer) Create(context.Context, *DraftCreateUpdateReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDraftServiceServer) Update(context.Context, *DraftCreateUpdateReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDraftServiceServer) Delete(context.Context, *ByID) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDraftServiceServer) SendDraft(context.Context, *ByID) (*MessageSentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendDraft not implemented")
}
func (UnimplementedDraftServiceServer) mustEmbedUnimplementedDraftServiceServer() {}

// UnsafeDraftServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DraftServiceServer will
// result in compilation errors.
type UnsafeDraftServiceServer interface {
	mustEmbedUnimplementedDraftServiceServer()
}

func RegisterDraftServiceServer(s grpc.ServiceRegistrar, srv DraftServiceServer) {
	s.RegisterService(&DraftService_ServiceDesc, srv)
}

func _DraftService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DraftCreateUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DraftServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DraftService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DraftServiceServer).Create(ctx, req.(*DraftCreateUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DraftService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DraftCreateUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DraftServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DraftService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DraftServiceServer).Update(ctx, req.(*DraftCreateUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DraftService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DraftServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DraftService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DraftServiceServer).Delete(ctx, req.(*ByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DraftService_SendDraft_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DraftServiceServer).SendDraft(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DraftService_SendDraft_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DraftServiceServer).SendDraft(ctx, req.(*ByID))
	}
	return interceptor(ctx, in, info, handler)
}

// DraftService_ServiceDesc is the grpc.ServiceDesc for DraftService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DraftService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.DraftService",
	HandlerType: (*DraftServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _DraftService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DraftService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DraftService_Delete_Handler,
		},
		{
			MethodName: "SendDraft",
			Handler:    _DraftService_SendDraft_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gmailapp-submodule/draft.proto",
}