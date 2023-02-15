// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: org.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/org/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// OrgService org.proto
	OrgService() pb.OrgServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		orgService: pb.NewOrgServiceClient(cc),
	}
}

type serviceClients struct {
	orgService pb.OrgServiceClient
}

func (c *serviceClients) OrgService() pb.OrgServiceClient {
	return c.orgService
}

type orgServiceWrapper struct {
	client pb.OrgServiceClient
	opts   []grpc1.CallOption
}

func (s *orgServiceWrapper) CreateOrg(ctx context.Context, req *pb.CreateOrgRequest) (*pb.CreateOrgResponse, error) {
	return s.client.CreateOrg(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *orgServiceWrapper) UpdateOrg(ctx context.Context, req *pb.UpdateOrgRequest) (*pb.UpdateOrgResponse, error) {
	return s.client.UpdateOrg(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *orgServiceWrapper) GetOrg(ctx context.Context, req *pb.GetOrgRequest) (*pb.GetOrgResponse, error) {
	return s.client.GetOrg(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *orgServiceWrapper) DeleteOrg(ctx context.Context, req *pb.DeleteOrgRequest) (*pb.DeleteOrgResponse, error) {
	return s.client.DeleteOrg(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *orgServiceWrapper) ListOrg(ctx context.Context, req *pb.ListOrgRequest) (*pb.ListOrgResponse, error) {
	return s.client.ListOrg(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *orgServiceWrapper) ListPublicOrg(ctx context.Context, req *pb.ListOrgRequest) (*pb.ListOrgResponse, error) {
	return s.client.ListPublicOrg(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *orgServiceWrapper) GetOrgByDomain(ctx context.Context, req *pb.GetOrgByDomainRequest) (*pb.GetOrgByDomainResponse, error) {
	return s.client.GetOrgByDomain(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *orgServiceWrapper) ChangeCurrentOrg(ctx context.Context, req *pb.ChangeCurrentOrgRequest) (*pb.ChangeCurrentOrgResponse, error) {
	return s.client.ChangeCurrentOrg(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *orgServiceWrapper) CreateOrgClusterRelation(ctx context.Context, req *pb.OrgClusterRelationCreateRequest) (*pb.OrgClusterRelationCreateResponse, error) {
	return s.client.CreateOrgClusterRelation(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *orgServiceWrapper) ListOrgClusterRelation(ctx context.Context, req *pb.ListOrgClusterRelationRequest) (*pb.ListOrgClusterRelationResponse, error) {
	return s.client.ListOrgClusterRelation(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *orgServiceWrapper) SetReleaseCrossCluster(ctx context.Context, req *pb.SetReleaseCrossClusterRequest) (*pb.SetReleaseCrossClusterResponse, error) {
	return s.client.SetReleaseCrossCluster(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *orgServiceWrapper) GenVerifyCode(ctx context.Context, req *pb.GenVerifyCodeRequest) (*pb.GenVerifyCodeResponse, error) {
	return s.client.GenVerifyCode(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *orgServiceWrapper) SetNotifyConfig(ctx context.Context, req *pb.SetNotifyConfigRequest) (*pb.SetNotifyConfigResponse, error) {
	return s.client.SetNotifyConfig(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *orgServiceWrapper) GetNotifyConfig(ctx context.Context, req *pb.GetNotifyConfigRequest) (*pb.GetNotifyConfigResponse, error) {
	return s.client.GetNotifyConfig(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *orgServiceWrapper) GetOrgClusterRelationsByOrg(ctx context.Context, req *pb.GetOrgClusterRelationsByOrgRequest) (*pb.GetOrgClusterRelationsByOrgResponse, error) {
	return s.client.GetOrgClusterRelationsByOrg(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *orgServiceWrapper) DereferenceCluster(ctx context.Context, req *pb.DereferenceClusterRequest) (*pb.DereferenceClusterResponse, error) {
	return s.client.DereferenceCluster(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
