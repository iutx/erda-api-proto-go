// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: search.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/dop/search/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// SearchService search.proto
	SearchService() pb.SearchServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		searchService: pb.NewSearchServiceClient(cc),
	}
}

type serviceClients struct {
	searchService pb.SearchServiceClient
}

func (c *serviceClients) SearchService() pb.SearchServiceClient {
	return c.searchService
}

type searchServiceWrapper struct {
	client pb.SearchServiceClient
	opts   []grpc1.CallOption
}

func (s *searchServiceWrapper) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	return s.client.Search(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
