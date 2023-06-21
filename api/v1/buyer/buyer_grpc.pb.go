// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: v1/buyer/buyer.proto

package buyer

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

const (
	Buyer_Self_FullMethodName = "/api.v1.auth.buyer.Buyer/Self"
)

// BuyerClient is the client API for Buyer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BuyerClient interface {
	Self(ctx context.Context, in *SelfRequest, opts ...grpc.CallOption) (*SelfReply, error)
}

type buyerClient struct {
	cc grpc.ClientConnInterface
}

func NewBuyerClient(cc grpc.ClientConnInterface) BuyerClient {
	return &buyerClient{cc}
}

func (c *buyerClient) Self(ctx context.Context, in *SelfRequest, opts ...grpc.CallOption) (*SelfReply, error) {
	out := new(SelfReply)
	err := c.cc.Invoke(ctx, Buyer_Self_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuyerServer is the server API for Buyer service.
// All implementations must embed UnimplementedBuyerServer
// for forward compatibility
type BuyerServer interface {
	Self(context.Context, *SelfRequest) (*SelfReply, error)
	mustEmbedUnimplementedBuyerServer()
}

// UnimplementedBuyerServer must be embedded to have forward compatible implementations.
type UnimplementedBuyerServer struct {
}

func (UnimplementedBuyerServer) Self(context.Context, *SelfRequest) (*SelfReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Self not implemented")
}
func (UnimplementedBuyerServer) mustEmbedUnimplementedBuyerServer() {}

// UnsafeBuyerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BuyerServer will
// result in compilation errors.
type UnsafeBuyerServer interface {
	mustEmbedUnimplementedBuyerServer()
}

func RegisterBuyerServer(s grpc.ServiceRegistrar, srv BuyerServer) {
	s.RegisterService(&Buyer_ServiceDesc, srv)
}

func _Buyer_Self_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuyerServer).Self(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Buyer_Self_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuyerServer).Self(ctx, req.(*SelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Buyer_ServiceDesc is the grpc.ServiceDesc for Buyer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Buyer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.auth.buyer.Buyer",
	HandlerType: (*BuyerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Self",
			Handler:    _Buyer_Self_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/buyer/buyer.proto",
}
