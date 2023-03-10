// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: msp.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/openapiv1/msp/pb"
	grpc1 "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// Client provide all service clients.
type Client interface {
	// Msp msp.proto
	Msp() pb.MspClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		msp: pb.NewMspClient(cc),
	}
}

type serviceClients struct {
	msp pb.MspClient
}

func (c *serviceClients) Msp() pb.MspClient {
	return c.msp
}

type mspWrapper struct {
	client pb.MspClient
	opts   []grpc1.CallOption
}

func (s *mspWrapper) MONITOR_INSTANCES(ctx context.Context, req *pb.MONITOR_INSTANCES_Request) (*emptypb.Empty, error) {
	return s.client.MONITOR_INSTANCES(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *mspWrapper) MONITOR_RUNTIME(ctx context.Context, req *pb.MONITOR_RUNTIME_Request) (*emptypb.Empty, error) {
	return s.client.MONITOR_RUNTIME(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
