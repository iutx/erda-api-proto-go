// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: dicehub.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/openapiv1/dicehub/pb"
	grpc1 "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// Client provide all service clients.
type Client interface {
	// Dicehub dicehub.proto
	Dicehub() pb.DicehubClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		dicehub: pb.NewDicehubClient(cc),
	}
}

type serviceClients struct {
	dicehub pb.DicehubClient
}

func (c *serviceClients) Dicehub() pb.DicehubClient {
	return c.dicehub
}

type dicehubWrapper struct {
	client pb.DicehubClient
	opts   []grpc1.CallOption
}

func (s *dicehubWrapper) DICEHUB_RELEASES_DOWNLOAD(ctx context.Context, req *pb.DICEHUB_RELEASES_DOWNLOAD_Request) (*emptypb.Empty, error) {
	return s.client.DICEHUB_RELEASES_DOWNLOAD(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *dicehubWrapper) DICEHUB_RELEASES_YAML_GET(ctx context.Context, req *pb.ReleaseGetDiceYmlRequest) (*emptypb.Empty, error) {
	return s.client.DICEHUB_RELEASES_YAML_GET(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *dicehubWrapper) IMAGEHUB_IMAGE_LIST(ctx context.Context, req *pb.ReleaseListRequest) (*pb.ReleaseListResponse, error) {
	return s.client.IMAGEHUB_IMAGE_LIST(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *dicehubWrapper) RELEASE_RULE_CREATE(ctx context.Context, req *pb.CreateUpdateDeleteReleaseRuleRequest) (*pb.CreateBranchRuleResponse, error) {
	return s.client.RELEASE_RULE_CREATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *dicehubWrapper) RELEASE_RULE_LIST(ctx context.Context, req *pb.RELEASE_RULE_LIST_Request) (*emptypb.Empty, error) {
	return s.client.RELEASE_RULE_LIST(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *dicehubWrapper) RELEASE_RULE_UPDATE(ctx context.Context, req *pb.RELEASE_RULE_UPDATE_Request) (*emptypb.Empty, error) {
	return s.client.RELEASE_RULE_UPDATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *dicehubWrapper) RELEASE_RULEDelete(ctx context.Context, req *pb.RELEASE_RULEDelete_Request) (*emptypb.Empty, error) {
	return s.client.RELEASE_RULEDelete(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}