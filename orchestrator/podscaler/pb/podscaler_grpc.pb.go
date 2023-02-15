// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: podscaler.proto

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

// PodScalerServiceClient is the client API for PodScalerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PodScalerServiceClient interface {
	CreateRuntimeHPARules(ctx context.Context, in *HPARuleCreateRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	ListRuntimeHPARules(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ErdaRuntimeHPARules, error)
	UpdateRuntimeHPARules(ctx context.Context, in *ErdaRuntimeHPARules, opts ...grpc.CallOption) (*CommonResponse, error)
	DeleteHPARulesByIds(ctx context.Context, in *DeleteRuntimePARulesRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	ApplyOrCancelHPARulesByIds(ctx context.Context, in *ApplyOrCancelPARulesRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	GetRuntimeBaseInfo(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*RuntimeServiceBaseInfos, error)
	ListRuntimeHPAEvents(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ErdaRuntimeHPAEvents, error)
	CreateRuntimeVPARules(ctx context.Context, in *VPARuleCreateRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	ListRuntimeVPARules(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ErdaRuntimeVPARules, error)
	UpdateRuntimeVPARules(ctx context.Context, in *ErdaRuntimeVPARules, opts ...grpc.CallOption) (*CommonResponse, error)
	DeleteVPARulesByIds(ctx context.Context, in *DeleteRuntimePARulesRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	ApplyOrCancelVPARulesByIds(ctx context.Context, in *ApplyOrCancelPARulesRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	ListRuntimeVPARecommendations(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ErdaRuntimeVPARecommendations, error)
	HPScaleManual(ctx context.Context, in *ManualHPRequest, opts ...grpc.CallOption) (*HPManualResponse, error)
}

type podScalerServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewPodScalerServiceClient(cc grpc1.ClientConnInterface) PodScalerServiceClient {
	return &podScalerServiceClient{cc}
}

func (c *podScalerServiceClient) CreateRuntimeHPARules(ctx context.Context, in *HPARuleCreateRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/erda.orchestrator.podscaler.PodScalerService/CreateRuntimeHPARules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podScalerServiceClient) ListRuntimeHPARules(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ErdaRuntimeHPARules, error) {
	out := new(ErdaRuntimeHPARules)
	err := c.cc.Invoke(ctx, "/erda.orchestrator.podscaler.PodScalerService/ListRuntimeHPARules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podScalerServiceClient) UpdateRuntimeHPARules(ctx context.Context, in *ErdaRuntimeHPARules, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/erda.orchestrator.podscaler.PodScalerService/UpdateRuntimeHPARules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podScalerServiceClient) DeleteHPARulesByIds(ctx context.Context, in *DeleteRuntimePARulesRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/erda.orchestrator.podscaler.PodScalerService/DeleteHPARulesByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podScalerServiceClient) ApplyOrCancelHPARulesByIds(ctx context.Context, in *ApplyOrCancelPARulesRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/erda.orchestrator.podscaler.PodScalerService/ApplyOrCancelHPARulesByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podScalerServiceClient) GetRuntimeBaseInfo(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*RuntimeServiceBaseInfos, error) {
	out := new(RuntimeServiceBaseInfos)
	err := c.cc.Invoke(ctx, "/erda.orchestrator.podscaler.PodScalerService/GetRuntimeBaseInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podScalerServiceClient) ListRuntimeHPAEvents(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ErdaRuntimeHPAEvents, error) {
	out := new(ErdaRuntimeHPAEvents)
	err := c.cc.Invoke(ctx, "/erda.orchestrator.podscaler.PodScalerService/ListRuntimeHPAEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podScalerServiceClient) CreateRuntimeVPARules(ctx context.Context, in *VPARuleCreateRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/erda.orchestrator.podscaler.PodScalerService/CreateRuntimeVPARules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podScalerServiceClient) ListRuntimeVPARules(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ErdaRuntimeVPARules, error) {
	out := new(ErdaRuntimeVPARules)
	err := c.cc.Invoke(ctx, "/erda.orchestrator.podscaler.PodScalerService/ListRuntimeVPARules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podScalerServiceClient) UpdateRuntimeVPARules(ctx context.Context, in *ErdaRuntimeVPARules, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/erda.orchestrator.podscaler.PodScalerService/UpdateRuntimeVPARules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podScalerServiceClient) DeleteVPARulesByIds(ctx context.Context, in *DeleteRuntimePARulesRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/erda.orchestrator.podscaler.PodScalerService/DeleteVPARulesByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podScalerServiceClient) ApplyOrCancelVPARulesByIds(ctx context.Context, in *ApplyOrCancelPARulesRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/erda.orchestrator.podscaler.PodScalerService/ApplyOrCancelVPARulesByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podScalerServiceClient) ListRuntimeVPARecommendations(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ErdaRuntimeVPARecommendations, error) {
	out := new(ErdaRuntimeVPARecommendations)
	err := c.cc.Invoke(ctx, "/erda.orchestrator.podscaler.PodScalerService/ListRuntimeVPARecommendations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podScalerServiceClient) HPScaleManual(ctx context.Context, in *ManualHPRequest, opts ...grpc.CallOption) (*HPManualResponse, error) {
	out := new(HPManualResponse)
	err := c.cc.Invoke(ctx, "/erda.orchestrator.podscaler.PodScalerService/HPScaleManual", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PodScalerServiceServer is the server API for PodScalerService service.
// All implementations should embed UnimplementedPodScalerServiceServer
// for forward compatibility
type PodScalerServiceServer interface {
	CreateRuntimeHPARules(context.Context, *HPARuleCreateRequest) (*CommonResponse, error)
	ListRuntimeHPARules(context.Context, *ListRequest) (*ErdaRuntimeHPARules, error)
	UpdateRuntimeHPARules(context.Context, *ErdaRuntimeHPARules) (*CommonResponse, error)
	DeleteHPARulesByIds(context.Context, *DeleteRuntimePARulesRequest) (*CommonResponse, error)
	ApplyOrCancelHPARulesByIds(context.Context, *ApplyOrCancelPARulesRequest) (*CommonResponse, error)
	GetRuntimeBaseInfo(context.Context, *ListRequest) (*RuntimeServiceBaseInfos, error)
	ListRuntimeHPAEvents(context.Context, *ListRequest) (*ErdaRuntimeHPAEvents, error)
	CreateRuntimeVPARules(context.Context, *VPARuleCreateRequest) (*CommonResponse, error)
	ListRuntimeVPARules(context.Context, *ListRequest) (*ErdaRuntimeVPARules, error)
	UpdateRuntimeVPARules(context.Context, *ErdaRuntimeVPARules) (*CommonResponse, error)
	DeleteVPARulesByIds(context.Context, *DeleteRuntimePARulesRequest) (*CommonResponse, error)
	ApplyOrCancelVPARulesByIds(context.Context, *ApplyOrCancelPARulesRequest) (*CommonResponse, error)
	ListRuntimeVPARecommendations(context.Context, *ListRequest) (*ErdaRuntimeVPARecommendations, error)
	HPScaleManual(context.Context, *ManualHPRequest) (*HPManualResponse, error)
}

// UnimplementedPodScalerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPodScalerServiceServer struct {
}

func (*UnimplementedPodScalerServiceServer) CreateRuntimeHPARules(context.Context, *HPARuleCreateRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRuntimeHPARules not implemented")
}
func (*UnimplementedPodScalerServiceServer) ListRuntimeHPARules(context.Context, *ListRequest) (*ErdaRuntimeHPARules, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRuntimeHPARules not implemented")
}
func (*UnimplementedPodScalerServiceServer) UpdateRuntimeHPARules(context.Context, *ErdaRuntimeHPARules) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRuntimeHPARules not implemented")
}
func (*UnimplementedPodScalerServiceServer) DeleteHPARulesByIds(context.Context, *DeleteRuntimePARulesRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHPARulesByIds not implemented")
}
func (*UnimplementedPodScalerServiceServer) ApplyOrCancelHPARulesByIds(context.Context, *ApplyOrCancelPARulesRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyOrCancelHPARulesByIds not implemented")
}
func (*UnimplementedPodScalerServiceServer) GetRuntimeBaseInfo(context.Context, *ListRequest) (*RuntimeServiceBaseInfos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRuntimeBaseInfo not implemented")
}
func (*UnimplementedPodScalerServiceServer) ListRuntimeHPAEvents(context.Context, *ListRequest) (*ErdaRuntimeHPAEvents, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRuntimeHPAEvents not implemented")
}
func (*UnimplementedPodScalerServiceServer) CreateRuntimeVPARules(context.Context, *VPARuleCreateRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRuntimeVPARules not implemented")
}
func (*UnimplementedPodScalerServiceServer) ListRuntimeVPARules(context.Context, *ListRequest) (*ErdaRuntimeVPARules, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRuntimeVPARules not implemented")
}
func (*UnimplementedPodScalerServiceServer) UpdateRuntimeVPARules(context.Context, *ErdaRuntimeVPARules) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRuntimeVPARules not implemented")
}
func (*UnimplementedPodScalerServiceServer) DeleteVPARulesByIds(context.Context, *DeleteRuntimePARulesRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVPARulesByIds not implemented")
}
func (*UnimplementedPodScalerServiceServer) ApplyOrCancelVPARulesByIds(context.Context, *ApplyOrCancelPARulesRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyOrCancelVPARulesByIds not implemented")
}
func (*UnimplementedPodScalerServiceServer) ListRuntimeVPARecommendations(context.Context, *ListRequest) (*ErdaRuntimeVPARecommendations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRuntimeVPARecommendations not implemented")
}
func (*UnimplementedPodScalerServiceServer) HPScaleManual(context.Context, *ManualHPRequest) (*HPManualResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HPScaleManual not implemented")
}

func RegisterPodScalerServiceServer(s grpc1.ServiceRegistrar, srv PodScalerServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_PodScalerService_serviceDesc(srv, opts...), srv)
}

var _PodScalerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.orchestrator.podscaler.PodScalerService",
	HandlerType: (*PodScalerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "podscaler.proto",
}

func _get_PodScalerService_serviceDesc(srv PodScalerServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_PodScalerService_CreateRuntimeHPARules_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CreateRuntimeHPARules(ctx, req.(*HPARuleCreateRequest))
	}
	var _PodScalerService_CreateRuntimeHPARules_info transport.ServiceInfo
	if h.Interceptor != nil {
		_PodScalerService_CreateRuntimeHPARules_info = transport.NewServiceInfo("erda.orchestrator.podscaler.PodScalerService", "CreateRuntimeHPARules", srv)
		_PodScalerService_CreateRuntimeHPARules_Handler = h.Interceptor(_PodScalerService_CreateRuntimeHPARules_Handler)
	}

	_PodScalerService_ListRuntimeHPARules_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ListRuntimeHPARules(ctx, req.(*ListRequest))
	}
	var _PodScalerService_ListRuntimeHPARules_info transport.ServiceInfo
	if h.Interceptor != nil {
		_PodScalerService_ListRuntimeHPARules_info = transport.NewServiceInfo("erda.orchestrator.podscaler.PodScalerService", "ListRuntimeHPARules", srv)
		_PodScalerService_ListRuntimeHPARules_Handler = h.Interceptor(_PodScalerService_ListRuntimeHPARules_Handler)
	}

	_PodScalerService_UpdateRuntimeHPARules_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.UpdateRuntimeHPARules(ctx, req.(*ErdaRuntimeHPARules))
	}
	var _PodScalerService_UpdateRuntimeHPARules_info transport.ServiceInfo
	if h.Interceptor != nil {
		_PodScalerService_UpdateRuntimeHPARules_info = transport.NewServiceInfo("erda.orchestrator.podscaler.PodScalerService", "UpdateRuntimeHPARules", srv)
		_PodScalerService_UpdateRuntimeHPARules_Handler = h.Interceptor(_PodScalerService_UpdateRuntimeHPARules_Handler)
	}

	_PodScalerService_DeleteHPARulesByIds_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.DeleteHPARulesByIds(ctx, req.(*DeleteRuntimePARulesRequest))
	}
	var _PodScalerService_DeleteHPARulesByIds_info transport.ServiceInfo
	if h.Interceptor != nil {
		_PodScalerService_DeleteHPARulesByIds_info = transport.NewServiceInfo("erda.orchestrator.podscaler.PodScalerService", "DeleteHPARulesByIds", srv)
		_PodScalerService_DeleteHPARulesByIds_Handler = h.Interceptor(_PodScalerService_DeleteHPARulesByIds_Handler)
	}

	_PodScalerService_ApplyOrCancelHPARulesByIds_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ApplyOrCancelHPARulesByIds(ctx, req.(*ApplyOrCancelPARulesRequest))
	}
	var _PodScalerService_ApplyOrCancelHPARulesByIds_info transport.ServiceInfo
	if h.Interceptor != nil {
		_PodScalerService_ApplyOrCancelHPARulesByIds_info = transport.NewServiceInfo("erda.orchestrator.podscaler.PodScalerService", "ApplyOrCancelHPARulesByIds", srv)
		_PodScalerService_ApplyOrCancelHPARulesByIds_Handler = h.Interceptor(_PodScalerService_ApplyOrCancelHPARulesByIds_Handler)
	}

	_PodScalerService_GetRuntimeBaseInfo_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetRuntimeBaseInfo(ctx, req.(*ListRequest))
	}
	var _PodScalerService_GetRuntimeBaseInfo_info transport.ServiceInfo
	if h.Interceptor != nil {
		_PodScalerService_GetRuntimeBaseInfo_info = transport.NewServiceInfo("erda.orchestrator.podscaler.PodScalerService", "GetRuntimeBaseInfo", srv)
		_PodScalerService_GetRuntimeBaseInfo_Handler = h.Interceptor(_PodScalerService_GetRuntimeBaseInfo_Handler)
	}

	_PodScalerService_ListRuntimeHPAEvents_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ListRuntimeHPAEvents(ctx, req.(*ListRequest))
	}
	var _PodScalerService_ListRuntimeHPAEvents_info transport.ServiceInfo
	if h.Interceptor != nil {
		_PodScalerService_ListRuntimeHPAEvents_info = transport.NewServiceInfo("erda.orchestrator.podscaler.PodScalerService", "ListRuntimeHPAEvents", srv)
		_PodScalerService_ListRuntimeHPAEvents_Handler = h.Interceptor(_PodScalerService_ListRuntimeHPAEvents_Handler)
	}

	_PodScalerService_CreateRuntimeVPARules_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CreateRuntimeVPARules(ctx, req.(*VPARuleCreateRequest))
	}
	var _PodScalerService_CreateRuntimeVPARules_info transport.ServiceInfo
	if h.Interceptor != nil {
		_PodScalerService_CreateRuntimeVPARules_info = transport.NewServiceInfo("erda.orchestrator.podscaler.PodScalerService", "CreateRuntimeVPARules", srv)
		_PodScalerService_CreateRuntimeVPARules_Handler = h.Interceptor(_PodScalerService_CreateRuntimeVPARules_Handler)
	}

	_PodScalerService_ListRuntimeVPARules_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ListRuntimeVPARules(ctx, req.(*ListRequest))
	}
	var _PodScalerService_ListRuntimeVPARules_info transport.ServiceInfo
	if h.Interceptor != nil {
		_PodScalerService_ListRuntimeVPARules_info = transport.NewServiceInfo("erda.orchestrator.podscaler.PodScalerService", "ListRuntimeVPARules", srv)
		_PodScalerService_ListRuntimeVPARules_Handler = h.Interceptor(_PodScalerService_ListRuntimeVPARules_Handler)
	}

	_PodScalerService_UpdateRuntimeVPARules_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.UpdateRuntimeVPARules(ctx, req.(*ErdaRuntimeVPARules))
	}
	var _PodScalerService_UpdateRuntimeVPARules_info transport.ServiceInfo
	if h.Interceptor != nil {
		_PodScalerService_UpdateRuntimeVPARules_info = transport.NewServiceInfo("erda.orchestrator.podscaler.PodScalerService", "UpdateRuntimeVPARules", srv)
		_PodScalerService_UpdateRuntimeVPARules_Handler = h.Interceptor(_PodScalerService_UpdateRuntimeVPARules_Handler)
	}

	_PodScalerService_DeleteVPARulesByIds_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.DeleteVPARulesByIds(ctx, req.(*DeleteRuntimePARulesRequest))
	}
	var _PodScalerService_DeleteVPARulesByIds_info transport.ServiceInfo
	if h.Interceptor != nil {
		_PodScalerService_DeleteVPARulesByIds_info = transport.NewServiceInfo("erda.orchestrator.podscaler.PodScalerService", "DeleteVPARulesByIds", srv)
		_PodScalerService_DeleteVPARulesByIds_Handler = h.Interceptor(_PodScalerService_DeleteVPARulesByIds_Handler)
	}

	_PodScalerService_ApplyOrCancelVPARulesByIds_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ApplyOrCancelVPARulesByIds(ctx, req.(*ApplyOrCancelPARulesRequest))
	}
	var _PodScalerService_ApplyOrCancelVPARulesByIds_info transport.ServiceInfo
	if h.Interceptor != nil {
		_PodScalerService_ApplyOrCancelVPARulesByIds_info = transport.NewServiceInfo("erda.orchestrator.podscaler.PodScalerService", "ApplyOrCancelVPARulesByIds", srv)
		_PodScalerService_ApplyOrCancelVPARulesByIds_Handler = h.Interceptor(_PodScalerService_ApplyOrCancelVPARulesByIds_Handler)
	}

	_PodScalerService_ListRuntimeVPARecommendations_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ListRuntimeVPARecommendations(ctx, req.(*ListRequest))
	}
	var _PodScalerService_ListRuntimeVPARecommendations_info transport.ServiceInfo
	if h.Interceptor != nil {
		_PodScalerService_ListRuntimeVPARecommendations_info = transport.NewServiceInfo("erda.orchestrator.podscaler.PodScalerService", "ListRuntimeVPARecommendations", srv)
		_PodScalerService_ListRuntimeVPARecommendations_Handler = h.Interceptor(_PodScalerService_ListRuntimeVPARecommendations_Handler)
	}

	_PodScalerService_HPScaleManual_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.HPScaleManual(ctx, req.(*ManualHPRequest))
	}
	var _PodScalerService_HPScaleManual_info transport.ServiceInfo
	if h.Interceptor != nil {
		_PodScalerService_HPScaleManual_info = transport.NewServiceInfo("erda.orchestrator.podscaler.PodScalerService", "HPScaleManual", srv)
		_PodScalerService_HPScaleManual_Handler = h.Interceptor(_PodScalerService_HPScaleManual_Handler)
	}

	var serviceDesc = _PodScalerService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "CreateRuntimeHPARules",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(HPARuleCreateRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(PodScalerServiceServer).CreateRuntimeHPARules(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _PodScalerService_CreateRuntimeHPARules_info)
				}
				if interceptor == nil {
					return _PodScalerService_CreateRuntimeHPARules_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.orchestrator.podscaler.PodScalerService/CreateRuntimeHPARules",
				}
				return interceptor(ctx, in, info, _PodScalerService_CreateRuntimeHPARules_Handler)
			},
		},
		{
			MethodName: "ListRuntimeHPARules",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ListRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(PodScalerServiceServer).ListRuntimeHPARules(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _PodScalerService_ListRuntimeHPARules_info)
				}
				if interceptor == nil {
					return _PodScalerService_ListRuntimeHPARules_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.orchestrator.podscaler.PodScalerService/ListRuntimeHPARules",
				}
				return interceptor(ctx, in, info, _PodScalerService_ListRuntimeHPARules_Handler)
			},
		},
		{
			MethodName: "UpdateRuntimeHPARules",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ErdaRuntimeHPARules)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(PodScalerServiceServer).UpdateRuntimeHPARules(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _PodScalerService_UpdateRuntimeHPARules_info)
				}
				if interceptor == nil {
					return _PodScalerService_UpdateRuntimeHPARules_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.orchestrator.podscaler.PodScalerService/UpdateRuntimeHPARules",
				}
				return interceptor(ctx, in, info, _PodScalerService_UpdateRuntimeHPARules_Handler)
			},
		},
		{
			MethodName: "DeleteHPARulesByIds",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(DeleteRuntimePARulesRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(PodScalerServiceServer).DeleteHPARulesByIds(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _PodScalerService_DeleteHPARulesByIds_info)
				}
				if interceptor == nil {
					return _PodScalerService_DeleteHPARulesByIds_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.orchestrator.podscaler.PodScalerService/DeleteHPARulesByIds",
				}
				return interceptor(ctx, in, info, _PodScalerService_DeleteHPARulesByIds_Handler)
			},
		},
		{
			MethodName: "ApplyOrCancelHPARulesByIds",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ApplyOrCancelPARulesRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(PodScalerServiceServer).ApplyOrCancelHPARulesByIds(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _PodScalerService_ApplyOrCancelHPARulesByIds_info)
				}
				if interceptor == nil {
					return _PodScalerService_ApplyOrCancelHPARulesByIds_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.orchestrator.podscaler.PodScalerService/ApplyOrCancelHPARulesByIds",
				}
				return interceptor(ctx, in, info, _PodScalerService_ApplyOrCancelHPARulesByIds_Handler)
			},
		},
		{
			MethodName: "GetRuntimeBaseInfo",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ListRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(PodScalerServiceServer).GetRuntimeBaseInfo(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _PodScalerService_GetRuntimeBaseInfo_info)
				}
				if interceptor == nil {
					return _PodScalerService_GetRuntimeBaseInfo_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.orchestrator.podscaler.PodScalerService/GetRuntimeBaseInfo",
				}
				return interceptor(ctx, in, info, _PodScalerService_GetRuntimeBaseInfo_Handler)
			},
		},
		{
			MethodName: "ListRuntimeHPAEvents",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ListRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(PodScalerServiceServer).ListRuntimeHPAEvents(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _PodScalerService_ListRuntimeHPAEvents_info)
				}
				if interceptor == nil {
					return _PodScalerService_ListRuntimeHPAEvents_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.orchestrator.podscaler.PodScalerService/ListRuntimeHPAEvents",
				}
				return interceptor(ctx, in, info, _PodScalerService_ListRuntimeHPAEvents_Handler)
			},
		},
		{
			MethodName: "CreateRuntimeVPARules",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(VPARuleCreateRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(PodScalerServiceServer).CreateRuntimeVPARules(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _PodScalerService_CreateRuntimeVPARules_info)
				}
				if interceptor == nil {
					return _PodScalerService_CreateRuntimeVPARules_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.orchestrator.podscaler.PodScalerService/CreateRuntimeVPARules",
				}
				return interceptor(ctx, in, info, _PodScalerService_CreateRuntimeVPARules_Handler)
			},
		},
		{
			MethodName: "ListRuntimeVPARules",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ListRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(PodScalerServiceServer).ListRuntimeVPARules(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _PodScalerService_ListRuntimeVPARules_info)
				}
				if interceptor == nil {
					return _PodScalerService_ListRuntimeVPARules_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.orchestrator.podscaler.PodScalerService/ListRuntimeVPARules",
				}
				return interceptor(ctx, in, info, _PodScalerService_ListRuntimeVPARules_Handler)
			},
		},
		{
			MethodName: "UpdateRuntimeVPARules",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ErdaRuntimeVPARules)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(PodScalerServiceServer).UpdateRuntimeVPARules(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _PodScalerService_UpdateRuntimeVPARules_info)
				}
				if interceptor == nil {
					return _PodScalerService_UpdateRuntimeVPARules_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.orchestrator.podscaler.PodScalerService/UpdateRuntimeVPARules",
				}
				return interceptor(ctx, in, info, _PodScalerService_UpdateRuntimeVPARules_Handler)
			},
		},
		{
			MethodName: "DeleteVPARulesByIds",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(DeleteRuntimePARulesRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(PodScalerServiceServer).DeleteVPARulesByIds(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _PodScalerService_DeleteVPARulesByIds_info)
				}
				if interceptor == nil {
					return _PodScalerService_DeleteVPARulesByIds_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.orchestrator.podscaler.PodScalerService/DeleteVPARulesByIds",
				}
				return interceptor(ctx, in, info, _PodScalerService_DeleteVPARulesByIds_Handler)
			},
		},
		{
			MethodName: "ApplyOrCancelVPARulesByIds",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ApplyOrCancelPARulesRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(PodScalerServiceServer).ApplyOrCancelVPARulesByIds(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _PodScalerService_ApplyOrCancelVPARulesByIds_info)
				}
				if interceptor == nil {
					return _PodScalerService_ApplyOrCancelVPARulesByIds_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.orchestrator.podscaler.PodScalerService/ApplyOrCancelVPARulesByIds",
				}
				return interceptor(ctx, in, info, _PodScalerService_ApplyOrCancelVPARulesByIds_Handler)
			},
		},
		{
			MethodName: "ListRuntimeVPARecommendations",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ListRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(PodScalerServiceServer).ListRuntimeVPARecommendations(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _PodScalerService_ListRuntimeVPARecommendations_info)
				}
				if interceptor == nil {
					return _PodScalerService_ListRuntimeVPARecommendations_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.orchestrator.podscaler.PodScalerService/ListRuntimeVPARecommendations",
				}
				return interceptor(ctx, in, info, _PodScalerService_ListRuntimeVPARecommendations_Handler)
			},
		},
		{
			MethodName: "HPScaleManual",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ManualHPRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(PodScalerServiceServer).HPScaleManual(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _PodScalerService_HPScaleManual_info)
				}
				if interceptor == nil {
					return _PodScalerService_HPScaleManual_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.orchestrator.podscaler.PodScalerService/HPScaleManual",
				}
				return interceptor(ctx, in, info, _PodScalerService_HPScaleManual_Handler)
			},
		},
	}
	return &serviceDesc
}
