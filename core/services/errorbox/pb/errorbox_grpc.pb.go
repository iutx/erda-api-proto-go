// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: errorbox.proto

package pb

import (
	context "context"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	grpc1 "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/dop/taskerror/pb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion5

// ErrorBoxServiceClient is the client API for ErrorBoxService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ErrorBoxServiceClient interface {
	ListErrorLog(ctx context.Context, in *TaskErrorListRequest, opts ...grpc.CallOption) (*pb.ErrorLogListResponseData, error)
}

type errorBoxServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewErrorBoxServiceClient(cc grpc1.ClientConnInterface) ErrorBoxServiceClient {
	return &errorBoxServiceClient{cc}
}

func (c *errorBoxServiceClient) ListErrorLog(ctx context.Context, in *TaskErrorListRequest, opts ...grpc.CallOption) (*pb.ErrorLogListResponseData, error) {
	out := new(pb.ErrorLogListResponseData)
	err := c.cc.Invoke(ctx, "/erda.core.services.errorbox.ErrorBoxService/ListErrorLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ErrorBoxServiceServer is the server API for ErrorBoxService service.
// All implementations should embed UnimplementedErrorBoxServiceServer
// for forward compatibility
type ErrorBoxServiceServer interface {
	ListErrorLog(context.Context, *TaskErrorListRequest) (*pb.ErrorLogListResponseData, error)
}

// UnimplementedErrorBoxServiceServer should be embedded to have forward compatible implementations.
type UnimplementedErrorBoxServiceServer struct {
}

func (*UnimplementedErrorBoxServiceServer) ListErrorLog(context.Context, *TaskErrorListRequest) (*pb.ErrorLogListResponseData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListErrorLog not implemented")
}

func RegisterErrorBoxServiceServer(s grpc1.ServiceRegistrar, srv ErrorBoxServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_ErrorBoxService_serviceDesc(srv, opts...), srv)
}

var _ErrorBoxService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.core.services.errorbox.ErrorBoxService",
	HandlerType: (*ErrorBoxServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "errorbox.proto",
}

func _get_ErrorBoxService_serviceDesc(srv ErrorBoxServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_ErrorBoxService_ListErrorLog_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ListErrorLog(ctx, req.(*TaskErrorListRequest))
	}
	var _ErrorBoxService_ListErrorLog_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ErrorBoxService_ListErrorLog_info = transport.NewServiceInfo("erda.core.services.errorbox.ErrorBoxService", "ListErrorLog", srv)
		_ErrorBoxService_ListErrorLog_Handler = h.Interceptor(_ErrorBoxService_ListErrorLog_Handler)
	}

	var serviceDesc = _ErrorBoxService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "ListErrorLog",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(TaskErrorListRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ErrorBoxServiceServer).ListErrorLog(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ErrorBoxService_ListErrorLog_info)
				}
				if interceptor == nil {
					return _ErrorBoxService_ListErrorLog_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.services.errorbox.ErrorBoxService/ListErrorLog",
				}
				return interceptor(ctx, in, info, _ErrorBoxService_ListErrorLog_Handler)
			},
		},
	}
	return &serviceDesc
}
