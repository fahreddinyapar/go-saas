// Code generated by protoc-gen-go-grpc-proxy. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc-proxy v1.2.0
// - protoc             (unknown)
// source: event/api/v1/event.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

var _ EventServiceServer = (*eventServiceClientProxy)(nil)

const GrpcOperationEventServiceEvent = "/event.api.v1.EventService/Event"

// eventServiceClientProxy is the proxy to turn EventService client to server interface.
type eventServiceClientProxy struct {
	cc EventServiceClient
}

func NewEventServiceClientProxy(cc EventServiceClient) EventServiceServer {
	return &eventServiceClientProxy{cc}
}

func (c *eventServiceClientProxy) Event(ctx context.Context, in *EventRequest) (*emptypb.Empty, error) {
	return c.cc.Event(ctx, in)
}
