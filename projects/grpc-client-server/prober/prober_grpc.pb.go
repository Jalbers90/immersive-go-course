// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: prober/prober.proto

package prober

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

// ProberClient is the client API for Prober service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProberClient interface {
	DoProbes(ctx context.Context, in *ProbeRequest, opts ...grpc.CallOption) (*ProbeReply, error)
}

type proberClient struct {
	cc grpc.ClientConnInterface
}

func NewProberClient(cc grpc.ClientConnInterface) ProberClient {
	return &proberClient{cc}
}

func (c *proberClient) DoProbes(ctx context.Context, in *ProbeRequest, opts ...grpc.CallOption) (*ProbeReply, error) {
	out := new(ProbeReply)
	err := c.cc.Invoke(ctx, "/prober.Prober/DoProbes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProberServer is the server API for Prober service.
// All implementations must embed UnimplementedProberServer
// for forward compatibility
type ProberServer interface {
	DoProbes(context.Context, *ProbeRequest) (*ProbeReply, error)
	mustEmbedUnimplementedProberServer()
}

// UnimplementedProberServer must be embedded to have forward compatible implementations.
type UnimplementedProberServer struct {
}

func (UnimplementedProberServer) DoProbes(context.Context, *ProbeRequest) (*ProbeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoProbes not implemented")
}
func (UnimplementedProberServer) mustEmbedUnimplementedProberServer() {}

// UnsafeProberServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProberServer will
// result in compilation errors.
type UnsafeProberServer interface {
	mustEmbedUnimplementedProberServer()
}

func RegisterProberServer(s grpc.ServiceRegistrar, srv ProberServer) {
	s.RegisterService(&Prober_ServiceDesc, srv)
}

func _Prober_DoProbes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProbeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProberServer).DoProbes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prober.Prober/DoProbes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProberServer).DoProbes(ctx, req.(*ProbeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Prober_ServiceDesc is the grpc.ServiceDesc for Prober service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Prober_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "prober.Prober",
	HandlerType: (*ProberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoProbes",
			Handler:    _Prober_DoProbes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prober/prober.proto",
}
