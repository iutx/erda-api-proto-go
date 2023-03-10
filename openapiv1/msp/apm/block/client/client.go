// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: msp_apm_block.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/openapiv1/msp/apm/block/pb"
	grpc1 "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// Client provide all service clients.
type Client interface {
	// MspApmBlock msp_apm_block.proto
	MspApmBlock() pb.MspApmBlockClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		mspApmBlock: pb.NewMspApmBlockClient(cc),
	}
}

type serviceClients struct {
	mspApmBlock pb.MspApmBlockClient
}

func (c *serviceClients) MspApmBlock() pb.MspApmBlockClient {
	return c.mspApmBlock
}

type mspApmBlockWrapper struct {
	client pb.MspApmBlockClient
	opts   []grpc1.CallOption
}

func (s *mspApmBlockWrapper) CREATE_BLOCK(ctx context.Context, req *pb.CREATE_BLOCK_Request) (*emptypb.Empty, error) {
	return s.client.CREATE_BLOCK(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *mspApmBlockWrapper) DELETE_BLOCK(ctx context.Context, req *pb.DELETE_BLOCK_Request) (*emptypb.Empty, error) {
	return s.client.DELETE_BLOCK(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *mspApmBlockWrapper) GET_BLOCK(ctx context.Context, req *pb.GET_BLOCK_Request) (*emptypb.Empty, error) {
	return s.client.GET_BLOCK(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *mspApmBlockWrapper) LIST_BLOCKS(ctx context.Context, req *pb.LIST_BLOCKS_Request) (*emptypb.Empty, error) {
	return s.client.LIST_BLOCKS(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *mspApmBlockWrapper) TMC_METRIC_DASHBOARD_UPDATE(ctx context.Context, req *pb.TMC_METRIC_DASHBOARD_UPDATE_Request) (*emptypb.Empty, error) {
	return s.client.TMC_METRIC_DASHBOARD_UPDATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
