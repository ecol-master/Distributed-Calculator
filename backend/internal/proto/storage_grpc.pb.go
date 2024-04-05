// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: proto/storage.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StorageServiceClient is the client API for StorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageServiceClient interface {
	CreateExpression(ctx context.Context, in *CreateExpressionRequest, opts ...grpc.CallOption) (*CreateExpressionResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	UpdateExpression(ctx context.Context, in *UpdateExpressionRequest, opts ...grpc.CallOption) (*Empty, error)
	SelectUserExpressions(ctx context.Context, in *SelectUserExpressionsRequest, opts ...grpc.CallOption) (*SelectUserExpressionsResponse, error)
	SelectExpression(ctx context.Context, in *SelectExpressionRequest, opts ...grpc.CallOption) (*SelectExpressionResponse, error)
}

type storageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageServiceClient(cc grpc.ClientConnInterface) StorageServiceClient {
	return &storageServiceClient{cc}
}

func (c *storageServiceClient) CreateExpression(ctx context.Context, in *CreateExpressionRequest, opts ...grpc.CallOption) (*CreateExpressionResponse, error) {
	out := new(CreateExpressionResponse)
	err := c.cc.Invoke(ctx, "/storage.StorageService/CreateExpression", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/storage.StorageService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) UpdateExpression(ctx context.Context, in *UpdateExpressionRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/storage.StorageService/UpdateExpression", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) SelectUserExpressions(ctx context.Context, in *SelectUserExpressionsRequest, opts ...grpc.CallOption) (*SelectUserExpressionsResponse, error) {
	out := new(SelectUserExpressionsResponse)
	err := c.cc.Invoke(ctx, "/storage.StorageService/SelectUserExpressions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) SelectExpression(ctx context.Context, in *SelectExpressionRequest, opts ...grpc.CallOption) (*SelectExpressionResponse, error) {
	out := new(SelectExpressionResponse)
	err := c.cc.Invoke(ctx, "/storage.StorageService/SelectExpression", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServiceServer is the server API for StorageService service.
// All implementations must embed UnimplementedStorageServiceServer
// for forward compatibility
type StorageServiceServer interface {
	CreateExpression(context.Context, *CreateExpressionRequest) (*CreateExpressionResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	UpdateExpression(context.Context, *UpdateExpressionRequest) (*Empty, error)
	SelectUserExpressions(context.Context, *SelectUserExpressionsRequest) (*SelectUserExpressionsResponse, error)
	SelectExpression(context.Context, *SelectExpressionRequest) (*SelectExpressionResponse, error)
	mustEmbedUnimplementedStorageServiceServer()
}

// UnimplementedStorageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStorageServiceServer struct {
}

func (UnimplementedStorageServiceServer) CreateExpression(context.Context, *CreateExpressionRequest) (*CreateExpressionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExpression not implemented")
}
func (UnimplementedStorageServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedStorageServiceServer) UpdateExpression(context.Context, *UpdateExpressionRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExpression not implemented")
}
func (UnimplementedStorageServiceServer) SelectUserExpressions(context.Context, *SelectUserExpressionsRequest) (*SelectUserExpressionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectUserExpressions not implemented")
}
func (UnimplementedStorageServiceServer) SelectExpression(context.Context, *SelectExpressionRequest) (*SelectExpressionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectExpression not implemented")
}
func (UnimplementedStorageServiceServer) mustEmbedUnimplementedStorageServiceServer() {}

// UnsafeStorageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServiceServer will
// result in compilation errors.
type UnsafeStorageServiceServer interface {
	mustEmbedUnimplementedStorageServiceServer()
}

func RegisterStorageServiceServer(s grpc.ServiceRegistrar, srv StorageServiceServer) {
	s.RegisterService(&StorageService_ServiceDesc, srv)
}

func _StorageService_CreateExpression_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExpressionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).CreateExpression(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/CreateExpression",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).CreateExpression(ctx, req.(*CreateExpressionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_UpdateExpression_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateExpressionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).UpdateExpression(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/UpdateExpression",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).UpdateExpression(ctx, req.(*UpdateExpressionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_SelectUserExpressions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectUserExpressionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).SelectUserExpressions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/SelectUserExpressions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).SelectUserExpressions(ctx, req.(*SelectUserExpressionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_SelectExpression_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectExpressionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).SelectExpression(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/SelectExpression",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).SelectExpression(ctx, req.(*SelectExpressionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StorageService_ServiceDesc is the grpc.ServiceDesc for StorageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StorageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storage.StorageService",
	HandlerType: (*StorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExpression",
			Handler:    _StorageService_CreateExpression_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _StorageService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateExpression",
			Handler:    _StorageService_UpdateExpression_Handler,
		},
		{
			MethodName: "SelectUserExpressions",
			Handler:    _StorageService_SelectUserExpressions_Handler,
		},
		{
			MethodName: "SelectExpression",
			Handler:    _StorageService_SelectExpression_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/storage.proto",
}