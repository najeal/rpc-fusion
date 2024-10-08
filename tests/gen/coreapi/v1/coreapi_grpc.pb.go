// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: coreapi/v1/coreapi.proto

package coreapiv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CoreApiService_Ping_FullMethodName   = "/coreapi.v1.CoreApiService/Ping"
	CoreApiService_Order_FullMethodName  = "/coreapi.v1.CoreApiService/Order"
	CoreApiService_Cancel_FullMethodName = "/coreapi.v1.CoreApiService/Cancel"
)

// CoreApiServiceClient is the client API for CoreApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoreApiServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	Order(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error)
	Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error)
}

type coreApiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCoreApiServiceClient(cc grpc.ClientConnInterface) CoreApiServiceClient {
	return &coreApiServiceClient{cc}
}

func (c *coreApiServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, CoreApiService_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreApiServiceClient) Order(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, CoreApiService_Order_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreApiServiceClient) Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelResponse)
	err := c.cc.Invoke(ctx, CoreApiService_Cancel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoreApiServiceServer is the server API for CoreApiService service.
// All implementations must embed UnimplementedCoreApiServiceServer
// for forward compatibility.
type CoreApiServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	Order(context.Context, *OrderRequest) (*OrderResponse, error)
	Cancel(context.Context, *CancelRequest) (*CancelResponse, error)
	mustEmbedUnimplementedCoreApiServiceServer()
}

// UnimplementedCoreApiServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCoreApiServiceServer struct{}

func (UnimplementedCoreApiServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedCoreApiServiceServer) Order(context.Context, *OrderRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Order not implemented")
}
func (UnimplementedCoreApiServiceServer) Cancel(context.Context, *CancelRequest) (*CancelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (UnimplementedCoreApiServiceServer) mustEmbedUnimplementedCoreApiServiceServer() {}
func (UnimplementedCoreApiServiceServer) testEmbeddedByValue()                        {}

// UnsafeCoreApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoreApiServiceServer will
// result in compilation errors.
type UnsafeCoreApiServiceServer interface {
	mustEmbedUnimplementedCoreApiServiceServer()
}

func RegisterCoreApiServiceServer(s grpc.ServiceRegistrar, srv CoreApiServiceServer) {
	// If the following call pancis, it indicates UnimplementedCoreApiServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CoreApiService_ServiceDesc, srv)
}

func _CoreApiService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreApiServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoreApiService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreApiServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreApiService_Order_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreApiServiceServer).Order(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoreApiService_Order_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreApiServiceServer).Order(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreApiService_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreApiServiceServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoreApiService_Cancel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreApiServiceServer).Cancel(ctx, req.(*CancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CoreApiService_ServiceDesc is the grpc.ServiceDesc for CoreApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoreApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coreapi.v1.CoreApiService",
	HandlerType: (*CoreApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _CoreApiService_Ping_Handler,
		},
		{
			MethodName: "Order",
			Handler:    _CoreApiService_Order_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _CoreApiService_Cancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coreapi/v1/coreapi.proto",
}
