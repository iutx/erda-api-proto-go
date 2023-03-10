// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: dop_filetree.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/openapiv1/dop/filetree/pb"
	grpc1 "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// Client provide all service clients.
type Client interface {
	// DopFiletree dop_filetree.proto
	DopFiletree() pb.DopFiletreeClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		dopFiletree: pb.NewDopFiletreeClient(cc),
	}
}

type serviceClients struct {
	dopFiletree pb.DopFiletreeClient
}

func (c *serviceClients) DopFiletree() pb.DopFiletreeClient {
	return c.dopFiletree
}

type dopFiletreeWrapper struct {
	client pb.DopFiletreeClient
	opts   []grpc1.CallOption
}

func (s *dopFiletreeWrapper) CMDB_PROJECT_FILETREE_CREATE(ctx context.Context, req *pb.CMDB_PROJECT_FILETREE_CREATE_Request) (*emptypb.Empty, error) {
	return s.client.CMDB_PROJECT_FILETREE_CREATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *dopFiletreeWrapper) CMDB_PROJECT_FILETREE_DELETE(ctx context.Context, req *pb.CMDB_PROJECT_FILETREE_DELETE_Request) (*emptypb.Empty, error) {
	return s.client.CMDB_PROJECT_FILETREE_DELETE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *dopFiletreeWrapper) CMDB_PROJECT_FILETREE_FIND_ANCESTORS(ctx context.Context, req *pb.CMDB_PROJECT_FILETREE_FIND_ANCESTORS_Request) (*emptypb.Empty, error) {
	return s.client.CMDB_PROJECT_FILETREE_FIND_ANCESTORS(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *dopFiletreeWrapper) CMDB_PROJECT_FILETREE_FUZZY_SEARCH(ctx context.Context, req *pb.CMDB_PROJECT_FILETREE_FUZZY_SEARCH_Request) (*emptypb.Empty, error) {
	return s.client.CMDB_PROJECT_FILETREE_FUZZY_SEARCH(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *dopFiletreeWrapper) CMDB_PROJECT_FILETREE_GET(ctx context.Context, req *pb.CMDB_PROJECT_FILETREE_GET_Request) (*emptypb.Empty, error) {
	return s.client.CMDB_PROJECT_FILETREE_GET(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *dopFiletreeWrapper) CMDB_PROJECT_FILETREE_LIST(ctx context.Context, req *pb.CMDB_PROJECT_FILETREE_LIST_Request) (*emptypb.Empty, error) {
	return s.client.CMDB_PROJECT_FILETREE_LIST(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
