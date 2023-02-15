// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: task.proto

package pb

import (
	context "context"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	grpc1 "github.com/erda-project/erda-infra/pkg/transport/grpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion5

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskServiceClient interface {
	// task
	PipelineTaskDetail(ctx context.Context, in *PipelineTaskDetailRequest, opts ...grpc.CallOption) (*PipelineTaskDetailResponse, error)
	PipelineTaskGetBootstrapInfo(ctx context.Context, in *PipelineTaskGetBootstrapInfoRequest, opts ...grpc.CallOption) (*PipelineTaskGetBootstrapInfoResponse, error)
}

type taskServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewTaskServiceClient(cc grpc1.ClientConnInterface) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) PipelineTaskDetail(ctx context.Context, in *PipelineTaskDetailRequest, opts ...grpc.CallOption) (*PipelineTaskDetailResponse, error) {
	out := new(PipelineTaskDetailResponse)
	err := c.cc.Invoke(ctx, "/erda.core.pipeline.task.TaskService/PipelineTaskDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) PipelineTaskGetBootstrapInfo(ctx context.Context, in *PipelineTaskGetBootstrapInfoRequest, opts ...grpc.CallOption) (*PipelineTaskGetBootstrapInfoResponse, error) {
	out := new(PipelineTaskGetBootstrapInfoResponse)
	err := c.cc.Invoke(ctx, "/erda.core.pipeline.task.TaskService/PipelineTaskGetBootstrapInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
// All implementations should embed UnimplementedTaskServiceServer
// for forward compatibility
type TaskServiceServer interface {
	// task
	PipelineTaskDetail(context.Context, *PipelineTaskDetailRequest) (*PipelineTaskDetailResponse, error)
	PipelineTaskGetBootstrapInfo(context.Context, *PipelineTaskGetBootstrapInfoRequest) (*PipelineTaskGetBootstrapInfoResponse, error)
}

// UnimplementedTaskServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (*UnimplementedTaskServiceServer) PipelineTaskDetail(context.Context, *PipelineTaskDetailRequest) (*PipelineTaskDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PipelineTaskDetail not implemented")
}
func (*UnimplementedTaskServiceServer) PipelineTaskGetBootstrapInfo(context.Context, *PipelineTaskGetBootstrapInfoRequest) (*PipelineTaskGetBootstrapInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PipelineTaskGetBootstrapInfo not implemented")
}

func RegisterTaskServiceServer(s grpc1.ServiceRegistrar, srv TaskServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_TaskService_serviceDesc(srv, opts...), srv)
}

var _TaskService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.core.pipeline.task.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "task.proto",
}

func _get_TaskService_serviceDesc(srv TaskServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_TaskService_PipelineTaskDetail_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.PipelineTaskDetail(ctx, req.(*PipelineTaskDetailRequest))
	}
	var _TaskService_PipelineTaskDetail_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TaskService_PipelineTaskDetail_info = transport.NewServiceInfo("erda.core.pipeline.task.TaskService", "PipelineTaskDetail", srv)
		_TaskService_PipelineTaskDetail_Handler = h.Interceptor(_TaskService_PipelineTaskDetail_Handler)
	}

	_TaskService_PipelineTaskGetBootstrapInfo_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.PipelineTaskGetBootstrapInfo(ctx, req.(*PipelineTaskGetBootstrapInfoRequest))
	}
	var _TaskService_PipelineTaskGetBootstrapInfo_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TaskService_PipelineTaskGetBootstrapInfo_info = transport.NewServiceInfo("erda.core.pipeline.task.TaskService", "PipelineTaskGetBootstrapInfo", srv)
		_TaskService_PipelineTaskGetBootstrapInfo_Handler = h.Interceptor(_TaskService_PipelineTaskGetBootstrapInfo_Handler)
	}

	var serviceDesc = _TaskService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "PipelineTaskDetail",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(PipelineTaskDetailRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TaskServiceServer).PipelineTaskDetail(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TaskService_PipelineTaskDetail_info)
				}
				if interceptor == nil {
					return _TaskService_PipelineTaskDetail_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.pipeline.task.TaskService/PipelineTaskDetail",
				}
				return interceptor(ctx, in, info, _TaskService_PipelineTaskDetail_Handler)
			},
		},
		{
			MethodName: "PipelineTaskGetBootstrapInfo",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(PipelineTaskGetBootstrapInfoRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TaskServiceServer).PipelineTaskGetBootstrapInfo(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TaskService_PipelineTaskGetBootstrapInfo_info)
				}
				if interceptor == nil {
					return _TaskService_PipelineTaskGetBootstrapInfo_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.pipeline.task.TaskService/PipelineTaskGetBootstrapInfo",
				}
				return interceptor(ctx, in, info, _TaskService_PipelineTaskGetBootstrapInfo_Handler)
			},
		},
	}
	return &serviceDesc
}