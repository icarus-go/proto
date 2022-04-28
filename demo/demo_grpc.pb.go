// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: demo/demo.proto

package demo

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

// DemoClient is the client API for Demo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DemoClient interface {
	CreateDemo(ctx context.Context, in *CreateDemoRequest, opts ...grpc.CallOption) (*CreateDemoReply, error)
	UpdateDemo(ctx context.Context, in *UpdateDemoRequest, opts ...grpc.CallOption) (*UpdateDemoReply, error)
	DeleteDemo(ctx context.Context, in *DeleteDemoRequest, opts ...grpc.CallOption) (*DeleteDemoReply, error)
	GetDemo(ctx context.Context, in *GetDemoRequest, opts ...grpc.CallOption) (*GetDemoReply, error)
	ListDemo(ctx context.Context, in *ListDemoRequest, opts ...grpc.CallOption) (*ListDemoReply, error)
}

type demoClient struct {
	cc grpc.ClientConnInterface
}

func NewDemoClient(cc grpc.ClientConnInterface) DemoClient {
	return &demoClient{cc}
}

func (c *demoClient) CreateDemo(ctx context.Context, in *CreateDemoRequest, opts ...grpc.CallOption) (*CreateDemoReply, error) {
	out := new(CreateDemoReply)
	err := c.cc.Invoke(ctx, "/demo.Demo/CreateDemo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoClient) UpdateDemo(ctx context.Context, in *UpdateDemoRequest, opts ...grpc.CallOption) (*UpdateDemoReply, error) {
	out := new(UpdateDemoReply)
	err := c.cc.Invoke(ctx, "/demo.Demo/UpdateDemo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoClient) DeleteDemo(ctx context.Context, in *DeleteDemoRequest, opts ...grpc.CallOption) (*DeleteDemoReply, error) {
	out := new(DeleteDemoReply)
	err := c.cc.Invoke(ctx, "/demo.Demo/DeleteDemo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoClient) GetDemo(ctx context.Context, in *GetDemoRequest, opts ...grpc.CallOption) (*GetDemoReply, error) {
	out := new(GetDemoReply)
	err := c.cc.Invoke(ctx, "/demo.Demo/GetDemo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoClient) ListDemo(ctx context.Context, in *ListDemoRequest, opts ...grpc.CallOption) (*ListDemoReply, error) {
	out := new(ListDemoReply)
	err := c.cc.Invoke(ctx, "/demo.Demo/ListDemo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DemoServer is the server API for Demo service.
// All implementations must embed UnimplementedDemoServer
// for forward compatibility
type DemoServer interface {
	CreateDemo(context.Context, *CreateDemoRequest) (*CreateDemoReply, error)
	UpdateDemo(context.Context, *UpdateDemoRequest) (*UpdateDemoReply, error)
	DeleteDemo(context.Context, *DeleteDemoRequest) (*DeleteDemoReply, error)
	GetDemo(context.Context, *GetDemoRequest) (*GetDemoReply, error)
	ListDemo(context.Context, *ListDemoRequest) (*ListDemoReply, error)
	mustEmbedUnimplementedDemoServer()
}

// UnimplementedDemoServer must be embedded to have forward compatible implementations.
type UnimplementedDemoServer struct {
}

func (UnimplementedDemoServer) CreateDemo(context.Context, *CreateDemoRequest) (*CreateDemoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDemo not implemented")
}
func (UnimplementedDemoServer) UpdateDemo(context.Context, *UpdateDemoRequest) (*UpdateDemoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDemo not implemented")
}
func (UnimplementedDemoServer) DeleteDemo(context.Context, *DeleteDemoRequest) (*DeleteDemoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDemo not implemented")
}
func (UnimplementedDemoServer) GetDemo(context.Context, *GetDemoRequest) (*GetDemoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDemo not implemented")
}
func (UnimplementedDemoServer) ListDemo(context.Context, *ListDemoRequest) (*ListDemoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDemo not implemented")
}
func (UnimplementedDemoServer) mustEmbedUnimplementedDemoServer() {}

// UnsafeDemoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DemoServer will
// result in compilation errors.
type UnsafeDemoServer interface {
	mustEmbedUnimplementedDemoServer()
}

func RegisterDemoServer(s grpc.ServiceRegistrar, srv DemoServer) {
	s.RegisterService(&Demo_ServiceDesc, srv)
}

func _Demo_CreateDemo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDemoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServer).CreateDemo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.Demo/CreateDemo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServer).CreateDemo(ctx, req.(*CreateDemoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Demo_UpdateDemo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDemoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServer).UpdateDemo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.Demo/UpdateDemo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServer).UpdateDemo(ctx, req.(*UpdateDemoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Demo_DeleteDemo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDemoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServer).DeleteDemo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.Demo/DeleteDemo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServer).DeleteDemo(ctx, req.(*DeleteDemoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Demo_GetDemo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDemoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServer).GetDemo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.Demo/GetDemo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServer).GetDemo(ctx, req.(*GetDemoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Demo_ListDemo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDemoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServer).ListDemo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.Demo/ListDemo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServer).ListDemo(ctx, req.(*ListDemoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Demo_ServiceDesc is the grpc.ServiceDesc for Demo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Demo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "demo.Demo",
	HandlerType: (*DemoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDemo",
			Handler:    _Demo_CreateDemo_Handler,
		},
		{
			MethodName: "UpdateDemo",
			Handler:    _Demo_UpdateDemo_Handler,
		},
		{
			MethodName: "DeleteDemo",
			Handler:    _Demo_DeleteDemo_Handler,
		},
		{
			MethodName: "GetDemo",
			Handler:    _Demo_GetDemo_Handler,
		},
		{
			MethodName: "ListDemo",
			Handler:    _Demo_ListDemo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo/demo.proto",
}
