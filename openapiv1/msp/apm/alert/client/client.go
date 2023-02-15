// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: msp_apm_alert.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/openapiv1/msp/apm/alert/pb"
	grpc1 "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// Client provide all service clients.
type Client interface {
	// MspApmAlert msp_apm_alert.proto
	MspApmAlert() pb.MspApmAlertClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		mspApmAlert: pb.NewMspApmAlertClient(cc),
	}
}

type serviceClients struct {
	mspApmAlert pb.MspApmAlertClient
}

func (c *serviceClients) MspApmAlert() pb.MspApmAlertClient {
	return c.mspApmAlert
}

type mspApmAlertWrapper struct {
	client pb.MspApmAlertClient
	opts   []grpc1.CallOption
}

func (s *mspApmAlertWrapper) APM_ALERT(ctx context.Context, req *pb.APM_ALERT_Request) (*emptypb.Empty, error) {
	return s.client.APM_ALERT(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *mspApmAlertWrapper) APM_ALERTS(ctx context.Context, req *pb.APM_ALERTS_Request) (*emptypb.Empty, error) {
	return s.client.APM_ALERTS(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *mspApmAlertWrapper) APM_ALERTS_RULES(ctx context.Context, req *pb.APM_ALERTS_RULES_Request) (*emptypb.Empty, error) {
	return s.client.APM_ALERTS_RULES(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *mspApmAlertWrapper) APM_ALERT_CREATE(ctx context.Context, req *pb.APM_ALERT_CREATE_Request) (*emptypb.Empty, error) {
	return s.client.APM_ALERT_CREATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *mspApmAlertWrapper) APM_ALERT_DELETE(ctx context.Context, req *pb.APM_ALERT_DELETE_Request) (*emptypb.Empty, error) {
	return s.client.APM_ALERT_DELETE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *mspApmAlertWrapper) APM_ALERT_RECORD(ctx context.Context, req *pb.APM_ALERT_RECORD_Request) (*emptypb.Empty, error) {
	return s.client.APM_ALERT_RECORD(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *mspApmAlertWrapper) APM_ALERT_RECORDS(ctx context.Context, req *pb.APM_ALERT_RECORDS_Request) (*emptypb.Empty, error) {
	return s.client.APM_ALERT_RECORDS(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *mspApmAlertWrapper) APM_ALERT_RECORD_ATTRS(ctx context.Context, req *pb.APM_ALERT_RECORD_ATTRS_Request) (*emptypb.Empty, error) {
	return s.client.APM_ALERT_RECORD_ATTRS(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *mspApmAlertWrapper) APM_ALERT_RECORD_HISTORIES(ctx context.Context, req *pb.APM_ALERT_RECORD_HISTORIES_Request) (*emptypb.Empty, error) {
	return s.client.APM_ALERT_RECORD_HISTORIES(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *mspApmAlertWrapper) APM_ALERT_RECORD_ISSUE_CREATE(ctx context.Context, req *pb.APM_ALERT_RECORD_ISSUE_CREATE_Request) (*emptypb.Empty, error) {
	return s.client.APM_ALERT_RECORD_ISSUE_CREATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *mspApmAlertWrapper) APM_ALERT_RECORD_ISSUE_UPDATE(ctx context.Context, req *pb.APM_ALERT_RECORD_ISSUE_UPDATE_Request) (*emptypb.Empty, error) {
	return s.client.APM_ALERT_RECORD_ISSUE_UPDATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *mspApmAlertWrapper) APM_ALERT_SWITCH(ctx context.Context, req *pb.APM_ALERT_SWITCH_Request) (*emptypb.Empty, error) {
	return s.client.APM_ALERT_SWITCH(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *mspApmAlertWrapper) APM_ALERT_UPDATE(ctx context.Context, req *pb.APM_ALERT_UPDATE_Request) (*emptypb.Empty, error) {
	return s.client.APM_ALERT_UPDATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *mspApmAlertWrapper) APM_CUSTOMIZE_ALERT(ctx context.Context, req *pb.APM_CUSTOMIZE_ALERT_Request) (*emptypb.Empty, error) {
	return s.client.APM_CUSTOMIZE_ALERT(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *mspApmAlertWrapper) APM_CUSTOMIZE_ALERTS(ctx context.Context, req *pb.APM_CUSTOMIZE_ALERTS_Request) (*emptypb.Empty, error) {
	return s.client.APM_CUSTOMIZE_ALERTS(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *mspApmAlertWrapper) APM_CUSTOMIZE_ALERT_CREATE(ctx context.Context, req *pb.APM_CUSTOMIZE_ALERT_CREATE_Request) (*emptypb.Empty, error) {
	return s.client.APM_CUSTOMIZE_ALERT_CREATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *mspApmAlertWrapper) APM_CUSTOMIZE_ALERT_DASH_PREVIEW(ctx context.Context, req *pb.APM_CUSTOMIZE_ALERT_DASH_PREVIEW_Request) (*emptypb.Empty, error) {
	return s.client.APM_CUSTOMIZE_ALERT_DASH_PREVIEW(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *mspApmAlertWrapper) APM_CUSTOMIZE_ALERT_DELETE(ctx context.Context, req *pb.APM_CUSTOMIZE_ALERT_DELETE_Request) (*emptypb.Empty, error) {
	return s.client.APM_CUSTOMIZE_ALERT_DELETE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *mspApmAlertWrapper) APM_CUSTOMIZE_ALERT_METRICS(ctx context.Context, req *pb.APM_CUSTOMIZE_ALERT_METRICS_Request) (*emptypb.Empty, error) {
	return s.client.APM_CUSTOMIZE_ALERT_METRICS(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *mspApmAlertWrapper) APM_CUSTOMIZE_ALERT_NOTIFY_TARGET(ctx context.Context, req *pb.APM_CUSTOMIZE_ALERT_NOTIFY_TARGET_Request) (*emptypb.Empty, error) {
	return s.client.APM_CUSTOMIZE_ALERT_NOTIFY_TARGET(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *mspApmAlertWrapper) APM_CUSTOMIZE_ALERT_SWITCH(ctx context.Context, req *pb.APM_CUSTOMIZE_ALERT_SWITCH_Request) (*emptypb.Empty, error) {
	return s.client.APM_CUSTOMIZE_ALERT_SWITCH(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *mspApmAlertWrapper) APM_CUSTOMIZE_ALERT_UPDATE(ctx context.Context, req *pb.APM_CUSTOMIZE_ALERT_UPDATE_Request) (*emptypb.Empty, error) {
	return s.client.APM_CUSTOMIZE_ALERT_UPDATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
