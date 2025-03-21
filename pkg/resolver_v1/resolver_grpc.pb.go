// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: resolver.proto

package resolver_v1

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
	ResolverV1_Resolve_FullMethodName = "/resolver_v1.ResolverV1/Resolve"
)

// ResolverV1Client is the client API for ResolverV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResolverV1Client interface {
	Resolve(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*ResolveResponse, error)
}

type resolverV1Client struct {
	cc grpc.ClientConnInterface
}

func NewResolverV1Client(cc grpc.ClientConnInterface) ResolverV1Client {
	return &resolverV1Client{cc}
}

func (c *resolverV1Client) Resolve(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*ResolveResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResolveResponse)
	err := c.cc.Invoke(ctx, ResolverV1_Resolve_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResolverV1Server is the server API for ResolverV1 service.
// All implementations must embed UnimplementedResolverV1Server
// for forward compatibility.
type ResolverV1Server interface {
	Resolve(context.Context, *ResolveRequest) (*ResolveResponse, error)
	mustEmbedUnimplementedResolverV1Server()
}

// UnimplementedResolverV1Server must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedResolverV1Server struct{}

func (UnimplementedResolverV1Server) Resolve(context.Context, *ResolveRequest) (*ResolveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resolve not implemented")
}
func (UnimplementedResolverV1Server) mustEmbedUnimplementedResolverV1Server() {}
func (UnimplementedResolverV1Server) testEmbeddedByValue()                    {}

// UnsafeResolverV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResolverV1Server will
// result in compilation errors.
type UnsafeResolverV1Server interface {
	mustEmbedUnimplementedResolverV1Server()
}

func RegisterResolverV1Server(s grpc.ServiceRegistrar, srv ResolverV1Server) {
	// If the following call pancis, it indicates UnimplementedResolverV1Server was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ResolverV1_ServiceDesc, srv)
}

func _ResolverV1_Resolve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResolverV1Server).Resolve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResolverV1_Resolve_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResolverV1Server).Resolve(ctx, req.(*ResolveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ResolverV1_ServiceDesc is the grpc.ServiceDesc for ResolverV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResolverV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "resolver_v1.ResolverV1",
	HandlerType: (*ResolverV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Resolve",
			Handler:    _ResolverV1_Resolve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resolver.proto",
}
