// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: fdp.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/openapiv1/fdp/pb"
	grpc1 "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// Client provide all service clients.
type Client interface {
	// Fdp fdp.proto
	Fdp() pb.FdpClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		fdp: pb.NewFdpClient(cc),
	}
}

type serviceClients struct {
	fdp pb.FdpClient
}

func (c *serviceClients) Fdp() pb.FdpClient {
	return c.fdp
}

type fdpWrapper struct {
	client pb.FdpClient
	opts   []grpc1.CallOption
}

func (s *fdpWrapper) CDP_DELETE(ctx context.Context, req *pb.CDP_DELETE_Request) (*emptypb.Empty, error) {
	return s.client.CDP_DELETE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *fdpWrapper) CDP_GET(ctx context.Context, req *pb.CDP_GET_Request) (*emptypb.Empty, error) {
	return s.client.CDP_GET(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *fdpWrapper) CDP_PATCH(ctx context.Context, req *pb.CDP_PATCH_Request) (*emptypb.Empty, error) {
	return s.client.CDP_PATCH(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *fdpWrapper) CDP_POST(ctx context.Context, req *pb.CDP_POST_Request) (*emptypb.Empty, error) {
	return s.client.CDP_POST(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *fdpWrapper) CDP_PUT(ctx context.Context, req *pb.CDP_PUT_Request) (*emptypb.Empty, error) {
	return s.client.CDP_PUT(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *fdpWrapper) FDP_WEBSOCKET(ctx context.Context, req *pb.FDP_WEBSOCKET_Request) (*emptypb.Empty, error) {
	return s.client.FDP_WEBSOCKET(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
