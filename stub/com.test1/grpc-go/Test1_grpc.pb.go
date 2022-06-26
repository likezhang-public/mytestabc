// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: Test1.proto

package grpc_go

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

// Test1Client is the client API for Test1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Test1Client interface {
	Heathcheck(ctx context.Context, in *HealthcheckRequest, opts ...grpc.CallOption) (*HealthcheckResponse, error)
	Helloworld(ctx context.Context, in *HelloworldRequest, opts ...grpc.CallOption) (*HelloworldResponse, error)
}

type test1Client struct {
	cc grpc.ClientConnInterface
}

func NewTest1Client(cc grpc.ClientConnInterface) Test1Client {
	return &test1Client{cc}
}

func (c *test1Client) Heathcheck(ctx context.Context, in *HealthcheckRequest, opts ...grpc.CallOption) (*HealthcheckResponse, error) {
	out := new(HealthcheckResponse)
	err := c.cc.Invoke(ctx, "/com.test1.Test1/Heathcheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *test1Client) Helloworld(ctx context.Context, in *HelloworldRequest, opts ...grpc.CallOption) (*HelloworldResponse, error) {
	out := new(HelloworldResponse)
	err := c.cc.Invoke(ctx, "/com.test1.Test1/Helloworld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Test1Server is the server API for Test1 service.
// All implementations must embed UnimplementedTest1Server
// for forward compatibility
type Test1Server interface {
	Heathcheck(context.Context, *HealthcheckRequest) (*HealthcheckResponse, error)
	Helloworld(context.Context, *HelloworldRequest) (*HelloworldResponse, error)
	mustEmbedUnimplementedTest1Server()
}

// UnimplementedTest1Server must be embedded to have forward compatible implementations.
type UnimplementedTest1Server struct {
}

func (UnimplementedTest1Server) Heathcheck(context.Context, *HealthcheckRequest) (*HealthcheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heathcheck not implemented")
}
func (UnimplementedTest1Server) Helloworld(context.Context, *HelloworldRequest) (*HelloworldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Helloworld not implemented")
}
func (UnimplementedTest1Server) mustEmbedUnimplementedTest1Server() {}

// UnsafeTest1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Test1Server will
// result in compilation errors.
type UnsafeTest1Server interface {
	mustEmbedUnimplementedTest1Server()
}

func RegisterTest1Server(s grpc.ServiceRegistrar, srv Test1Server) {
	s.RegisterService(&Test1_ServiceDesc, srv)
}

func _Test1_Heathcheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthcheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Test1Server).Heathcheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.test1.Test1/Heathcheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Test1Server).Heathcheck(ctx, req.(*HealthcheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test1_Helloworld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloworldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Test1Server).Helloworld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.test1.Test1/Helloworld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Test1Server).Helloworld(ctx, req.(*HelloworldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Test1_ServiceDesc is the grpc.ServiceDesc for Test1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Test1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.test1.Test1",
	HandlerType: (*Test1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Heathcheck",
			Handler:    _Test1_Heathcheck_Handler,
		},
		{
			MethodName: "Helloworld",
			Handler:    _Test1_Helloworld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Test1.proto",
}
