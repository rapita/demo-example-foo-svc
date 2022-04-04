// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package foov1

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

// FooClient is the client API for Foo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FooClient interface {
	Say(ctx context.Context, in *SayRequest, opts ...grpc.CallOption) (*SayResponse, error)
}

type fooClient struct {
	cc grpc.ClientConnInterface
}

func NewFooClient(cc grpc.ClientConnInterface) FooClient {
	return &fooClient{cc}
}

func (c *fooClient) Say(ctx context.Context, in *SayRequest, opts ...grpc.CallOption) (*SayResponse, error) {
	out := new(SayResponse)
	err := c.cc.Invoke(ctx, "/demo.example.foo.v1.Foo/Say", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FooServer is the server API for Foo service.
// All implementations should embed UnimplementedFooServer
// for forward compatibility
type FooServer interface {
	Say(context.Context, *SayRequest) (*SayResponse, error)
}

// UnimplementedFooServer should be embedded to have forward compatible implementations.
type UnimplementedFooServer struct {
}

func (UnimplementedFooServer) Say(context.Context, *SayRequest) (*SayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Say not implemented")
}

// UnsafeFooServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FooServer will
// result in compilation errors.
type UnsafeFooServer interface {
	mustEmbedUnimplementedFooServer()
}

func RegisterFooServer(s grpc.ServiceRegistrar, srv FooServer) {
	s.RegisterService(&Foo_ServiceDesc, srv)
}

func _Foo_Say_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FooServer).Say(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.example.foo.v1.Foo/Say",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FooServer).Say(ctx, req.(*SayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Foo_ServiceDesc is the grpc.ServiceDesc for Foo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Foo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "demo.example.foo.v1.Foo",
	HandlerType: (*FooServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Say",
			Handler:    _Foo_Say_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/foo/v1/foo.proto",
}