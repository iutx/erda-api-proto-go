// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: opentelemetry.proto

package pb

import (
	context "context"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	grpc1 "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/common/pb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion5

// OpenTelemetryServiceClient is the client API for OpenTelemetryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpenTelemetryServiceClient interface {
	Export(ctx context.Context, in *PostSpansRequest, opts ...grpc.CallOption) (*pb.VoidResponse, error)
}

type openTelemetryServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewOpenTelemetryServiceClient(cc grpc1.ClientConnInterface) OpenTelemetryServiceClient {
	return &openTelemetryServiceClient{cc}
}

func (c *openTelemetryServiceClient) Export(ctx context.Context, in *PostSpansRequest, opts ...grpc.CallOption) (*pb.VoidResponse, error) {
	out := new(pb.VoidResponse)
	err := c.cc.Invoke(ctx, "/erda.oap.collector.receiver.opentelemetry.OpenTelemetryService/Export", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpenTelemetryServiceServer is the server API for OpenTelemetryService service.
// All implementations should embed UnimplementedOpenTelemetryServiceServer
// for forward compatibility
type OpenTelemetryServiceServer interface {
	Export(context.Context, *PostSpansRequest) (*pb.VoidResponse, error)
}

// UnimplementedOpenTelemetryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOpenTelemetryServiceServer struct {
}

func (*UnimplementedOpenTelemetryServiceServer) Export(context.Context, *PostSpansRequest) (*pb.VoidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Export not implemented")
}

func RegisterOpenTelemetryServiceServer(s grpc1.ServiceRegistrar, srv OpenTelemetryServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_OpenTelemetryService_serviceDesc(srv, opts...), srv)
}

var _OpenTelemetryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.oap.collector.receiver.opentelemetry.OpenTelemetryService",
	HandlerType: (*OpenTelemetryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "opentelemetry.proto",
}

func _get_OpenTelemetryService_serviceDesc(srv OpenTelemetryServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_OpenTelemetryService_Export_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.Export(ctx, req.(*PostSpansRequest))
	}
	var _OpenTelemetryService_Export_info transport.ServiceInfo
	if h.Interceptor != nil {
		_OpenTelemetryService_Export_info = transport.NewServiceInfo("erda.oap.collector.receiver.opentelemetry.OpenTelemetryService", "Export", srv)
		_OpenTelemetryService_Export_Handler = h.Interceptor(_OpenTelemetryService_Export_Handler)
	}

	var serviceDesc = _OpenTelemetryService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "Export",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(PostSpansRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(OpenTelemetryServiceServer).Export(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _OpenTelemetryService_Export_info)
				}
				if interceptor == nil {
					return _OpenTelemetryService_Export_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.oap.collector.receiver.opentelemetry.OpenTelemetryService/Export",
				}
				return interceptor(ctx, in, info, _OpenTelemetryService_Export_Handler)
			},
		},
	}
	return &serviceDesc
}
