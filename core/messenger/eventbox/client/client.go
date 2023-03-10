// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: eventbox.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/messenger/eventbox/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// EventBoxService eventbox.proto
	EventBoxService() pb.EventBoxServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		eventBoxService: pb.NewEventBoxServiceClient(cc),
	}
}

type serviceClients struct {
	eventBoxService pb.EventBoxServiceClient
}

func (c *serviceClients) EventBoxService() pb.EventBoxServiceClient {
	return c.eventBoxService
}

type eventBoxServiceWrapper struct {
	client pb.EventBoxServiceClient
	opts   []grpc1.CallOption
}

func (s *eventBoxServiceWrapper) CreateMessage(ctx context.Context, req *pb.CreateMessageRequest) (*pb.CreateMessageResponse, error) {
	return s.client.CreateMessage(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *eventBoxServiceWrapper) PrefixGet(ctx context.Context, req *pb.PrefixGetRequest) (*pb.PrefixGetResponse, error) {
	return s.client.PrefixGet(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *eventBoxServiceWrapper) Put(ctx context.Context, req *pb.PutRequest) (*pb.PutResponse, error) {
	return s.client.Put(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *eventBoxServiceWrapper) Del(ctx context.Context, req *pb.DelRequest) (*pb.DelResponse, error) {
	return s.client.Del(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *eventBoxServiceWrapper) GetVersion(ctx context.Context, req *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	return s.client.GetVersion(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *eventBoxServiceWrapper) GetSMTPInfo(ctx context.Context, req *pb.GetSMTPInfoRequest) (*pb.GetSMTPInfoResponse, error) {
	return s.client.GetSMTPInfo(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *eventBoxServiceWrapper) ListHooks(ctx context.Context, req *pb.ListHooksRequest) (*pb.ListHooksResponse, error) {
	return s.client.ListHooks(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *eventBoxServiceWrapper) InspectHook(ctx context.Context, req *pb.InspectHookRequest) (*pb.InspectHookResponse, error) {
	return s.client.InspectHook(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *eventBoxServiceWrapper) CreateHook(ctx context.Context, req *pb.CreateHookRequest) (*pb.CreateHookResponse, error) {
	return s.client.CreateHook(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *eventBoxServiceWrapper) EditHook(ctx context.Context, req *pb.EditHookRequest) (*pb.EditHookResponse, error) {
	return s.client.EditHook(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *eventBoxServiceWrapper) PingHook(ctx context.Context, req *pb.PingHookRequest) (*pb.PingHookResponse, error) {
	return s.client.PingHook(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *eventBoxServiceWrapper) DeleteHook(ctx context.Context, req *pb.DeleteHookRequest) (*pb.DeleteHookResponse, error) {
	return s.client.DeleteHook(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *eventBoxServiceWrapper) ListHookEvents(ctx context.Context, req *pb.ListHookEventsRequest) (*pb.ListHookEventsResponse, error) {
	return s.client.ListHookEvents(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *eventBoxServiceWrapper) Stat(ctx context.Context, req *pb.StatRequest) (*pb.StatResponse, error) {
	return s.client.Stat(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
