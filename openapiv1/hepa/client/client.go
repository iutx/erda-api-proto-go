// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: hepa.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/openapiv1/hepa/pb"
	grpc1 "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// Client provide all service clients.
type Client interface {
	// Hepa hepa.proto
	Hepa() pb.HepaClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		hepa: pb.NewHepaClient(cc),
	}
}

type serviceClients struct {
	hepa pb.HepaClient
}

func (c *serviceClients) Hepa() pb.HepaClient {
	return c.hepa
}

type hepaWrapper struct {
	client pb.HepaClient
	opts   []grpc1.CallOption
}

func (s *hepaWrapper) HEPA_DOMAINS_GET(ctx context.Context, req *pb.HEPA_DOMAINS_GET_Request) (*emptypb.Empty, error) {
	return s.client.HEPA_DOMAINS_GET(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *hepaWrapper) HEPA_RUNTIME_DOMAIN_GET(ctx context.Context, req *pb.DomainListRequest) (*pb.DomainListResponse, error) {
	return s.client.HEPA_RUNTIME_DOMAIN_GET(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *hepaWrapper) HEPA_RUNTIME_SERVICE_DOMAIN_UPDATE(ctx context.Context, req *pb.HEPA_RUNTIME_SERVICE_DOMAIN_UPDATE_Request) (*emptypb.Empty, error) {
	return s.client.HEPA_RUNTIME_SERVICE_DOMAIN_UPDATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
