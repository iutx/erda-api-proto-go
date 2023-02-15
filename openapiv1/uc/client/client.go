// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: uc.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/openapiv1/uc/pb"
	grpc1 "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// Client provide all service clients.
type Client interface {
	// Uc uc.proto
	Uc() pb.UcClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		uc: pb.NewUcClient(cc),
	}
}

type serviceClients struct {
	uc pb.UcClient
}

func (c *serviceClients) Uc() pb.UcClient {
	return c.uc
}

type ucWrapper struct {
	client pb.UcClient
	opts   []grpc1.CallOption
}

func (s *ucWrapper) UC_PWD_SECURITY_CONFIG_GET(ctx context.Context, req *pb.UC_PWD_SECURITY_CONFIG_GET_Request) (*pb.PwdSecurityConfigGetResponse, error) {
	return s.client.UC_PWD_SECURITY_CONFIG_GET(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *ucWrapper) UC_PWD_SECURITY_CONFIG_UPDATE(ctx context.Context, req *pb.PwdSecurityConfigUpdateRequest) (*pb.PwdSecurityConfigUpdateResponse, error) {
	return s.client.UC_PWD_SECURITY_CONFIG_UPDATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *ucWrapper) UC_USER_BATCH_FREEZE(ctx context.Context, req *pb.UserBatchFreezeRequest) (*pb.UserBatchFreezeResponse, error) {
	return s.client.UC_USER_BATCH_FREEZE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *ucWrapper) UC_USER_BATCH_UNFREEZE(ctx context.Context, req *pb.UserBatchUnFreezeRequest) (*pb.UserBatchUnFreezeResponse, error) {
	return s.client.UC_USER_BATCH_UNFREEZE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *ucWrapper) UC_USER_BATCH_UPDATE_LOGIN_METHOD(ctx context.Context, req *pb.UserBatchUpdateLoginMethodRequest) (*pb.UserBatchUpdateLoginMethodResponse, error) {
	return s.client.UC_USER_BATCH_UPDATE_LOGIN_METHOD(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *ucWrapper) UC_USER_CREATE(ctx context.Context, req *pb.UserCreateRequest) (*pb.UserCreateResponse, error) {
	return s.client.UC_USER_CREATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *ucWrapper) UC_USER_EXPORT(ctx context.Context, req *pb.UserPagingRequest) (*emptypb.Empty, error) {
	return s.client.UC_USER_EXPORT(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *ucWrapper) UC_USER_FREEZE(ctx context.Context, req *pb.UserFreezeRequest) (*pb.UserFreezeResponse, error) {
	return s.client.UC_USER_FREEZE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *ucWrapper) UC_USER_LIST_LOGIN_METHOD(ctx context.Context, req *pb.UC_USER_LIST_LOGIN_METHOD_Request) (*pb.UserListLoginMethodResponse, error) {
	return s.client.UC_USER_LIST_LOGIN_METHOD(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *ucWrapper) UC_USER_PAGING(ctx context.Context, req *pb.UserPagingRequest) (*pb.UserPagingResponse, error) {
	return s.client.UC_USER_PAGING(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *ucWrapper) UC_USER_UNFREEZE(ctx context.Context, req *pb.UserUnfreezeRequest) (*pb.UserUnfreezeResponse, error) {
	return s.client.UC_USER_UNFREEZE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *ucWrapper) UC_USER_UPDATE_LOGIN_METHOD(ctx context.Context, req *pb.UserUpdateLoginMethodRequest) (*pb.UserUpdateLoginMethodResponse, error) {
	return s.client.UC_USER_UPDATE_LOGIN_METHOD(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *ucWrapper) UC_USER_UPDATE_USERINFO(ctx context.Context, req *pb.UserUpdateInfoRequset) (*pb.UserUpdateInfoResponse, error) {
	return s.client.UC_USER_UPDATE_USERINFO(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
