// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: payment/api/subscription/v1/subscription_internal.proto

package v1

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
	SubscriptionInternalService_GetInternalSubscription_FullMethodName = "/payment.api.subscription.v1.SubscriptionInternalService/GetInternalSubscription"
)

// SubscriptionInternalServiceClient is the client API for SubscriptionInternalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubscriptionInternalServiceClient interface {
	GetInternalSubscription(ctx context.Context, in *GetInternalSubscriptionRequest, opts ...grpc.CallOption) (*Subscription, error)
}

type subscriptionInternalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscriptionInternalServiceClient(cc grpc.ClientConnInterface) SubscriptionInternalServiceClient {
	return &subscriptionInternalServiceClient{cc}
}

func (c *subscriptionInternalServiceClient) GetInternalSubscription(ctx context.Context, in *GetInternalSubscriptionRequest, opts ...grpc.CallOption) (*Subscription, error) {
	out := new(Subscription)
	err := c.cc.Invoke(ctx, SubscriptionInternalService_GetInternalSubscription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriptionInternalServiceServer is the server API for SubscriptionInternalService service.
// All implementations should embed UnimplementedSubscriptionInternalServiceServer
// for forward compatibility
type SubscriptionInternalServiceServer interface {
	GetInternalSubscription(context.Context, *GetInternalSubscriptionRequest) (*Subscription, error)
}

// UnimplementedSubscriptionInternalServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSubscriptionInternalServiceServer struct {
}

func (UnimplementedSubscriptionInternalServiceServer) GetInternalSubscription(context.Context, *GetInternalSubscriptionRequest) (*Subscription, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInternalSubscription not implemented")
}

// UnsafeSubscriptionInternalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscriptionInternalServiceServer will
// result in compilation errors.
type UnsafeSubscriptionInternalServiceServer interface {
	mustEmbedUnimplementedSubscriptionInternalServiceServer()
}

func RegisterSubscriptionInternalServiceServer(s grpc.ServiceRegistrar, srv SubscriptionInternalServiceServer) {
	s.RegisterService(&SubscriptionInternalService_ServiceDesc, srv)
}

func _SubscriptionInternalService_GetInternalSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInternalSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionInternalServiceServer).GetInternalSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionInternalService_GetInternalSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionInternalServiceServer).GetInternalSubscription(ctx, req.(*GetInternalSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SubscriptionInternalService_ServiceDesc is the grpc.ServiceDesc for SubscriptionInternalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubscriptionInternalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payment.api.subscription.v1.SubscriptionInternalService",
	HandlerType: (*SubscriptionInternalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInternalSubscription",
			Handler:    _SubscriptionInternalService_GetInternalSubscription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment/api/subscription/v1/subscription_internal.proto",
}