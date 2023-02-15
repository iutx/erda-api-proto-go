// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: user.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/user/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// UserService user.proto
	UserService() pb.UserServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		userService: pb.NewUserServiceClient(cc),
	}
}

type serviceClients struct {
	userService pb.UserServiceClient
}

func (c *serviceClients) UserService() pb.UserServiceClient {
	return c.userService
}

type userServiceWrapper struct {
	client pb.UserServiceClient
	opts   []grpc1.CallOption
}

func (s *userServiceWrapper) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return s.client.GetUser(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *userServiceWrapper) FindUsers(ctx context.Context, req *pb.FindUsersRequest) (*pb.FindUsersResponse, error) {
	return s.client.FindUsers(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *userServiceWrapper) FindUsersByKey(ctx context.Context, req *pb.FindUsersByKeyRequest) (*pb.FindUsersByKeyResponse, error) {
	return s.client.FindUsersByKey(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
