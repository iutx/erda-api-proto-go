// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: report.proto

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

// ReportServiceClient is the client API for ReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportServiceClient interface {
	ListTasks(ctx context.Context, in *ListTasksRequest, opts ...grpc.CallOption) (*ListTasksResponse, error)
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*pb.VoidResponse, error)
	SwitchTask(ctx context.Context, in *SwitchTaskRequest, opts ...grpc.CallOption) (*pb.VoidResponse, error)
	GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*ReportTaskDTO, error)
	DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*pb.VoidResponse, error)
	ListTypes(ctx context.Context, in *ListTypesRequest, opts ...grpc.CallOption) (*ListTypesResponse, error)
	ListHistories(ctx context.Context, in *ListHistoriesRequest, opts ...grpc.CallOption) (*ListHistoriesResponse, error)
	CreateHistory(ctx context.Context, in *CreateHistoryRequest, opts ...grpc.CallOption) (*CreateHistoryResponse, error)
	GetHistory(ctx context.Context, in *GetHistoryRequest, opts ...grpc.CallOption) (*GetHistoryResponse, error)
	DeleteHistory(ctx context.Context, in *DeleteHistoryRequest, opts ...grpc.CallOption) (*pb.VoidResponse, error)
}

type reportServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewReportServiceClient(cc grpc1.ClientConnInterface) ReportServiceClient {
	return &reportServiceClient{cc}
}

func (c *reportServiceClient) ListTasks(ctx context.Context, in *ListTasksRequest, opts ...grpc.CallOption) (*ListTasksResponse, error) {
	out := new(ListTasksResponse)
	err := c.cc.Invoke(ctx, "/erda.tools.monitor.dashboard.report.ReportService/ListTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportServiceClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	out := new(CreateTaskResponse)
	err := c.cc.Invoke(ctx, "/erda.tools.monitor.dashboard.report.ReportService/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportServiceClient) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*pb.VoidResponse, error) {
	out := new(pb.VoidResponse)
	err := c.cc.Invoke(ctx, "/erda.tools.monitor.dashboard.report.ReportService/UpdateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportServiceClient) SwitchTask(ctx context.Context, in *SwitchTaskRequest, opts ...grpc.CallOption) (*pb.VoidResponse, error) {
	out := new(pb.VoidResponse)
	err := c.cc.Invoke(ctx, "/erda.tools.monitor.dashboard.report.ReportService/SwitchTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportServiceClient) GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*ReportTaskDTO, error) {
	out := new(ReportTaskDTO)
	err := c.cc.Invoke(ctx, "/erda.tools.monitor.dashboard.report.ReportService/GetTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportServiceClient) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*pb.VoidResponse, error) {
	out := new(pb.VoidResponse)
	err := c.cc.Invoke(ctx, "/erda.tools.monitor.dashboard.report.ReportService/DeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportServiceClient) ListTypes(ctx context.Context, in *ListTypesRequest, opts ...grpc.CallOption) (*ListTypesResponse, error) {
	out := new(ListTypesResponse)
	err := c.cc.Invoke(ctx, "/erda.tools.monitor.dashboard.report.ReportService/ListTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportServiceClient) ListHistories(ctx context.Context, in *ListHistoriesRequest, opts ...grpc.CallOption) (*ListHistoriesResponse, error) {
	out := new(ListHistoriesResponse)
	err := c.cc.Invoke(ctx, "/erda.tools.monitor.dashboard.report.ReportService/ListHistories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportServiceClient) CreateHistory(ctx context.Context, in *CreateHistoryRequest, opts ...grpc.CallOption) (*CreateHistoryResponse, error) {
	out := new(CreateHistoryResponse)
	err := c.cc.Invoke(ctx, "/erda.tools.monitor.dashboard.report.ReportService/CreateHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportServiceClient) GetHistory(ctx context.Context, in *GetHistoryRequest, opts ...grpc.CallOption) (*GetHistoryResponse, error) {
	out := new(GetHistoryResponse)
	err := c.cc.Invoke(ctx, "/erda.tools.monitor.dashboard.report.ReportService/GetHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportServiceClient) DeleteHistory(ctx context.Context, in *DeleteHistoryRequest, opts ...grpc.CallOption) (*pb.VoidResponse, error) {
	out := new(pb.VoidResponse)
	err := c.cc.Invoke(ctx, "/erda.tools.monitor.dashboard.report.ReportService/DeleteHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportServiceServer is the server API for ReportService service.
// All implementations should embed UnimplementedReportServiceServer
// for forward compatibility
type ReportServiceServer interface {
	ListTasks(context.Context, *ListTasksRequest) (*ListTasksResponse, error)
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	UpdateTask(context.Context, *UpdateTaskRequest) (*pb.VoidResponse, error)
	SwitchTask(context.Context, *SwitchTaskRequest) (*pb.VoidResponse, error)
	GetTask(context.Context, *GetTaskRequest) (*ReportTaskDTO, error)
	DeleteTask(context.Context, *DeleteTaskRequest) (*pb.VoidResponse, error)
	ListTypes(context.Context, *ListTypesRequest) (*ListTypesResponse, error)
	ListHistories(context.Context, *ListHistoriesRequest) (*ListHistoriesResponse, error)
	CreateHistory(context.Context, *CreateHistoryRequest) (*CreateHistoryResponse, error)
	GetHistory(context.Context, *GetHistoryRequest) (*GetHistoryResponse, error)
	DeleteHistory(context.Context, *DeleteHistoryRequest) (*pb.VoidResponse, error)
}

// UnimplementedReportServiceServer should be embedded to have forward compatible implementations.
type UnimplementedReportServiceServer struct {
}

func (*UnimplementedReportServiceServer) ListTasks(context.Context, *ListTasksRequest) (*ListTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTasks not implemented")
}
func (*UnimplementedReportServiceServer) CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (*UnimplementedReportServiceServer) UpdateTask(context.Context, *UpdateTaskRequest) (*pb.VoidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (*UnimplementedReportServiceServer) SwitchTask(context.Context, *SwitchTaskRequest) (*pb.VoidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwitchTask not implemented")
}
func (*UnimplementedReportServiceServer) GetTask(context.Context, *GetTaskRequest) (*ReportTaskDTO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (*UnimplementedReportServiceServer) DeleteTask(context.Context, *DeleteTaskRequest) (*pb.VoidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (*UnimplementedReportServiceServer) ListTypes(context.Context, *ListTypesRequest) (*ListTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTypes not implemented")
}
func (*UnimplementedReportServiceServer) ListHistories(context.Context, *ListHistoriesRequest) (*ListHistoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHistories not implemented")
}
func (*UnimplementedReportServiceServer) CreateHistory(context.Context, *CreateHistoryRequest) (*CreateHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHistory not implemented")
}
func (*UnimplementedReportServiceServer) GetHistory(context.Context, *GetHistoryRequest) (*GetHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHistory not implemented")
}
func (*UnimplementedReportServiceServer) DeleteHistory(context.Context, *DeleteHistoryRequest) (*pb.VoidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHistory not implemented")
}

func RegisterReportServiceServer(s grpc1.ServiceRegistrar, srv ReportServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_ReportService_serviceDesc(srv, opts...), srv)
}

var _ReportService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.tools.monitor.dashboard.report.ReportService",
	HandlerType: (*ReportServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "report.proto",
}

func _get_ReportService_serviceDesc(srv ReportServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_ReportService_ListTasks_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ListTasks(ctx, req.(*ListTasksRequest))
	}
	var _ReportService_ListTasks_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ReportService_ListTasks_info = transport.NewServiceInfo("erda.tools.monitor.dashboard.report.ReportService", "ListTasks", srv)
		_ReportService_ListTasks_Handler = h.Interceptor(_ReportService_ListTasks_Handler)
	}

	_ReportService_CreateTask_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CreateTask(ctx, req.(*CreateTaskRequest))
	}
	var _ReportService_CreateTask_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ReportService_CreateTask_info = transport.NewServiceInfo("erda.tools.monitor.dashboard.report.ReportService", "CreateTask", srv)
		_ReportService_CreateTask_Handler = h.Interceptor(_ReportService_CreateTask_Handler)
	}

	_ReportService_UpdateTask_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.UpdateTask(ctx, req.(*UpdateTaskRequest))
	}
	var _ReportService_UpdateTask_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ReportService_UpdateTask_info = transport.NewServiceInfo("erda.tools.monitor.dashboard.report.ReportService", "UpdateTask", srv)
		_ReportService_UpdateTask_Handler = h.Interceptor(_ReportService_UpdateTask_Handler)
	}

	_ReportService_SwitchTask_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.SwitchTask(ctx, req.(*SwitchTaskRequest))
	}
	var _ReportService_SwitchTask_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ReportService_SwitchTask_info = transport.NewServiceInfo("erda.tools.monitor.dashboard.report.ReportService", "SwitchTask", srv)
		_ReportService_SwitchTask_Handler = h.Interceptor(_ReportService_SwitchTask_Handler)
	}

	_ReportService_GetTask_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetTask(ctx, req.(*GetTaskRequest))
	}
	var _ReportService_GetTask_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ReportService_GetTask_info = transport.NewServiceInfo("erda.tools.monitor.dashboard.report.ReportService", "GetTask", srv)
		_ReportService_GetTask_Handler = h.Interceptor(_ReportService_GetTask_Handler)
	}

	_ReportService_DeleteTask_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.DeleteTask(ctx, req.(*DeleteTaskRequest))
	}
	var _ReportService_DeleteTask_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ReportService_DeleteTask_info = transport.NewServiceInfo("erda.tools.monitor.dashboard.report.ReportService", "DeleteTask", srv)
		_ReportService_DeleteTask_Handler = h.Interceptor(_ReportService_DeleteTask_Handler)
	}

	_ReportService_ListTypes_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ListTypes(ctx, req.(*ListTypesRequest))
	}
	var _ReportService_ListTypes_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ReportService_ListTypes_info = transport.NewServiceInfo("erda.tools.monitor.dashboard.report.ReportService", "ListTypes", srv)
		_ReportService_ListTypes_Handler = h.Interceptor(_ReportService_ListTypes_Handler)
	}

	_ReportService_ListHistories_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ListHistories(ctx, req.(*ListHistoriesRequest))
	}
	var _ReportService_ListHistories_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ReportService_ListHistories_info = transport.NewServiceInfo("erda.tools.monitor.dashboard.report.ReportService", "ListHistories", srv)
		_ReportService_ListHistories_Handler = h.Interceptor(_ReportService_ListHistories_Handler)
	}

	_ReportService_CreateHistory_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CreateHistory(ctx, req.(*CreateHistoryRequest))
	}
	var _ReportService_CreateHistory_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ReportService_CreateHistory_info = transport.NewServiceInfo("erda.tools.monitor.dashboard.report.ReportService", "CreateHistory", srv)
		_ReportService_CreateHistory_Handler = h.Interceptor(_ReportService_CreateHistory_Handler)
	}

	_ReportService_GetHistory_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetHistory(ctx, req.(*GetHistoryRequest))
	}
	var _ReportService_GetHistory_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ReportService_GetHistory_info = transport.NewServiceInfo("erda.tools.monitor.dashboard.report.ReportService", "GetHistory", srv)
		_ReportService_GetHistory_Handler = h.Interceptor(_ReportService_GetHistory_Handler)
	}

	_ReportService_DeleteHistory_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.DeleteHistory(ctx, req.(*DeleteHistoryRequest))
	}
	var _ReportService_DeleteHistory_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ReportService_DeleteHistory_info = transport.NewServiceInfo("erda.tools.monitor.dashboard.report.ReportService", "DeleteHistory", srv)
		_ReportService_DeleteHistory_Handler = h.Interceptor(_ReportService_DeleteHistory_Handler)
	}

	var serviceDesc = _ReportService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "ListTasks",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ListTasksRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ReportServiceServer).ListTasks(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ReportService_ListTasks_info)
				}
				if interceptor == nil {
					return _ReportService_ListTasks_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.tools.monitor.dashboard.report.ReportService/ListTasks",
				}
				return interceptor(ctx, in, info, _ReportService_ListTasks_Handler)
			},
		},
		{
			MethodName: "CreateTask",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CreateTaskRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ReportServiceServer).CreateTask(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ReportService_CreateTask_info)
				}
				if interceptor == nil {
					return _ReportService_CreateTask_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.tools.monitor.dashboard.report.ReportService/CreateTask",
				}
				return interceptor(ctx, in, info, _ReportService_CreateTask_Handler)
			},
		},
		{
			MethodName: "UpdateTask",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(UpdateTaskRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ReportServiceServer).UpdateTask(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ReportService_UpdateTask_info)
				}
				if interceptor == nil {
					return _ReportService_UpdateTask_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.tools.monitor.dashboard.report.ReportService/UpdateTask",
				}
				return interceptor(ctx, in, info, _ReportService_UpdateTask_Handler)
			},
		},
		{
			MethodName: "SwitchTask",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(SwitchTaskRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ReportServiceServer).SwitchTask(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ReportService_SwitchTask_info)
				}
				if interceptor == nil {
					return _ReportService_SwitchTask_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.tools.monitor.dashboard.report.ReportService/SwitchTask",
				}
				return interceptor(ctx, in, info, _ReportService_SwitchTask_Handler)
			},
		},
		{
			MethodName: "GetTask",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetTaskRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ReportServiceServer).GetTask(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ReportService_GetTask_info)
				}
				if interceptor == nil {
					return _ReportService_GetTask_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.tools.monitor.dashboard.report.ReportService/GetTask",
				}
				return interceptor(ctx, in, info, _ReportService_GetTask_Handler)
			},
		},
		{
			MethodName: "DeleteTask",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(DeleteTaskRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ReportServiceServer).DeleteTask(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ReportService_DeleteTask_info)
				}
				if interceptor == nil {
					return _ReportService_DeleteTask_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.tools.monitor.dashboard.report.ReportService/DeleteTask",
				}
				return interceptor(ctx, in, info, _ReportService_DeleteTask_Handler)
			},
		},
		{
			MethodName: "ListTypes",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ListTypesRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ReportServiceServer).ListTypes(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ReportService_ListTypes_info)
				}
				if interceptor == nil {
					return _ReportService_ListTypes_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.tools.monitor.dashboard.report.ReportService/ListTypes",
				}
				return interceptor(ctx, in, info, _ReportService_ListTypes_Handler)
			},
		},
		{
			MethodName: "ListHistories",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ListHistoriesRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ReportServiceServer).ListHistories(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ReportService_ListHistories_info)
				}
				if interceptor == nil {
					return _ReportService_ListHistories_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.tools.monitor.dashboard.report.ReportService/ListHistories",
				}
				return interceptor(ctx, in, info, _ReportService_ListHistories_Handler)
			},
		},
		{
			MethodName: "CreateHistory",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CreateHistoryRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ReportServiceServer).CreateHistory(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ReportService_CreateHistory_info)
				}
				if interceptor == nil {
					return _ReportService_CreateHistory_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.tools.monitor.dashboard.report.ReportService/CreateHistory",
				}
				return interceptor(ctx, in, info, _ReportService_CreateHistory_Handler)
			},
		},
		{
			MethodName: "GetHistory",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetHistoryRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ReportServiceServer).GetHistory(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ReportService_GetHistory_info)
				}
				if interceptor == nil {
					return _ReportService_GetHistory_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.tools.monitor.dashboard.report.ReportService/GetHistory",
				}
				return interceptor(ctx, in, info, _ReportService_GetHistory_Handler)
			},
		},
		{
			MethodName: "DeleteHistory",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(DeleteHistoryRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ReportServiceServer).DeleteHistory(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ReportService_DeleteHistory_info)
				}
				if interceptor == nil {
					return _ReportService_DeleteHistory_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.tools.monitor.dashboard.report.ReportService/DeleteHistory",
				}
				return interceptor(ctx, in, info, _ReportService_DeleteHistory_Handler)
			},
		},
	}
	return &serviceDesc
}