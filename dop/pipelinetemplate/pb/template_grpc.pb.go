// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: template.proto

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

// TemplateServiceClient is the client API for TemplateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TemplateServiceClient interface {
	ApplyPipelineTemplate(ctx context.Context, in *PipelineTemplateApplyRequest, opts ...grpc.CallOption) (*PipelineTemplateCreateResponse, error)
	QueryPipelineTemplates(ctx context.Context, in *PipelineTemplateQueryRequest, opts ...grpc.CallOption) (*PipelineTemplateQueryResponse, error)
	RenderPipelineTemplate(ctx context.Context, in *PipelineTemplateRenderRequest, opts ...grpc.CallOption) (*PipelineTemplateRenderResponse, error)
	RenderPipelineTemplateBySpec(ctx context.Context, in *PipelineTemplateRenderSpecRequest, opts ...grpc.CallOption) (*PipelineTemplateRenderResponse, error)
	GetPipelineTemplateVersion(ctx context.Context, in *PipelineTemplateVersionGetRequest, opts ...grpc.CallOption) (*PipelineTemplateVersionGetResponse, error)
	QueryPipelineTemplateVersions(ctx context.Context, in *PipelineTemplateVersionQueryRequest, opts ...grpc.CallOption) (*PipelineTemplateVersionQueryResponse, error)
	QuerySnippetYml(ctx context.Context, in *QuerySnippetYmlRequest, opts ...grpc.CallOption) (*QuerySnippetYmlResponse, error)
}

type templateServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewTemplateServiceClient(cc grpc1.ClientConnInterface) TemplateServiceClient {
	return &templateServiceClient{cc}
}

func (c *templateServiceClient) ApplyPipelineTemplate(ctx context.Context, in *PipelineTemplateApplyRequest, opts ...grpc.CallOption) (*PipelineTemplateCreateResponse, error) {
	out := new(PipelineTemplateCreateResponse)
	err := c.cc.Invoke(ctx, "/erda.dop.pipelinetemplate.TemplateService/ApplyPipelineTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateServiceClient) QueryPipelineTemplates(ctx context.Context, in *PipelineTemplateQueryRequest, opts ...grpc.CallOption) (*PipelineTemplateQueryResponse, error) {
	out := new(PipelineTemplateQueryResponse)
	err := c.cc.Invoke(ctx, "/erda.dop.pipelinetemplate.TemplateService/QueryPipelineTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateServiceClient) RenderPipelineTemplate(ctx context.Context, in *PipelineTemplateRenderRequest, opts ...grpc.CallOption) (*PipelineTemplateRenderResponse, error) {
	out := new(PipelineTemplateRenderResponse)
	err := c.cc.Invoke(ctx, "/erda.dop.pipelinetemplate.TemplateService/RenderPipelineTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateServiceClient) RenderPipelineTemplateBySpec(ctx context.Context, in *PipelineTemplateRenderSpecRequest, opts ...grpc.CallOption) (*PipelineTemplateRenderResponse, error) {
	out := new(PipelineTemplateRenderResponse)
	err := c.cc.Invoke(ctx, "/erda.dop.pipelinetemplate.TemplateService/RenderPipelineTemplateBySpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateServiceClient) GetPipelineTemplateVersion(ctx context.Context, in *PipelineTemplateVersionGetRequest, opts ...grpc.CallOption) (*PipelineTemplateVersionGetResponse, error) {
	out := new(PipelineTemplateVersionGetResponse)
	err := c.cc.Invoke(ctx, "/erda.dop.pipelinetemplate.TemplateService/GetPipelineTemplateVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateServiceClient) QueryPipelineTemplateVersions(ctx context.Context, in *PipelineTemplateVersionQueryRequest, opts ...grpc.CallOption) (*PipelineTemplateVersionQueryResponse, error) {
	out := new(PipelineTemplateVersionQueryResponse)
	err := c.cc.Invoke(ctx, "/erda.dop.pipelinetemplate.TemplateService/QueryPipelineTemplateVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateServiceClient) QuerySnippetYml(ctx context.Context, in *QuerySnippetYmlRequest, opts ...grpc.CallOption) (*QuerySnippetYmlResponse, error) {
	out := new(QuerySnippetYmlResponse)
	err := c.cc.Invoke(ctx, "/erda.dop.pipelinetemplate.TemplateService/QuerySnippetYml", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemplateServiceServer is the server API for TemplateService service.
// All implementations should embed UnimplementedTemplateServiceServer
// for forward compatibility
type TemplateServiceServer interface {
	ApplyPipelineTemplate(context.Context, *PipelineTemplateApplyRequest) (*PipelineTemplateCreateResponse, error)
	QueryPipelineTemplates(context.Context, *PipelineTemplateQueryRequest) (*PipelineTemplateQueryResponse, error)
	RenderPipelineTemplate(context.Context, *PipelineTemplateRenderRequest) (*PipelineTemplateRenderResponse, error)
	RenderPipelineTemplateBySpec(context.Context, *PipelineTemplateRenderSpecRequest) (*PipelineTemplateRenderResponse, error)
	GetPipelineTemplateVersion(context.Context, *PipelineTemplateVersionGetRequest) (*PipelineTemplateVersionGetResponse, error)
	QueryPipelineTemplateVersions(context.Context, *PipelineTemplateVersionQueryRequest) (*PipelineTemplateVersionQueryResponse, error)
	QuerySnippetYml(context.Context, *QuerySnippetYmlRequest) (*QuerySnippetYmlResponse, error)
}

// UnimplementedTemplateServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTemplateServiceServer struct {
}

func (*UnimplementedTemplateServiceServer) ApplyPipelineTemplate(context.Context, *PipelineTemplateApplyRequest) (*PipelineTemplateCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyPipelineTemplate not implemented")
}
func (*UnimplementedTemplateServiceServer) QueryPipelineTemplates(context.Context, *PipelineTemplateQueryRequest) (*PipelineTemplateQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPipelineTemplates not implemented")
}
func (*UnimplementedTemplateServiceServer) RenderPipelineTemplate(context.Context, *PipelineTemplateRenderRequest) (*PipelineTemplateRenderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenderPipelineTemplate not implemented")
}
func (*UnimplementedTemplateServiceServer) RenderPipelineTemplateBySpec(context.Context, *PipelineTemplateRenderSpecRequest) (*PipelineTemplateRenderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenderPipelineTemplateBySpec not implemented")
}
func (*UnimplementedTemplateServiceServer) GetPipelineTemplateVersion(context.Context, *PipelineTemplateVersionGetRequest) (*PipelineTemplateVersionGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPipelineTemplateVersion not implemented")
}
func (*UnimplementedTemplateServiceServer) QueryPipelineTemplateVersions(context.Context, *PipelineTemplateVersionQueryRequest) (*PipelineTemplateVersionQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPipelineTemplateVersions not implemented")
}
func (*UnimplementedTemplateServiceServer) QuerySnippetYml(context.Context, *QuerySnippetYmlRequest) (*QuerySnippetYmlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySnippetYml not implemented")
}

func RegisterTemplateServiceServer(s grpc1.ServiceRegistrar, srv TemplateServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_TemplateService_serviceDesc(srv, opts...), srv)
}

var _TemplateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.dop.pipelinetemplate.TemplateService",
	HandlerType: (*TemplateServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "template.proto",
}

func _get_TemplateService_serviceDesc(srv TemplateServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_TemplateService_ApplyPipelineTemplate_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ApplyPipelineTemplate(ctx, req.(*PipelineTemplateApplyRequest))
	}
	var _TemplateService_ApplyPipelineTemplate_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TemplateService_ApplyPipelineTemplate_info = transport.NewServiceInfo("erda.dop.pipelinetemplate.TemplateService", "ApplyPipelineTemplate", srv)
		_TemplateService_ApplyPipelineTemplate_Handler = h.Interceptor(_TemplateService_ApplyPipelineTemplate_Handler)
	}

	_TemplateService_QueryPipelineTemplates_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.QueryPipelineTemplates(ctx, req.(*PipelineTemplateQueryRequest))
	}
	var _TemplateService_QueryPipelineTemplates_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TemplateService_QueryPipelineTemplates_info = transport.NewServiceInfo("erda.dop.pipelinetemplate.TemplateService", "QueryPipelineTemplates", srv)
		_TemplateService_QueryPipelineTemplates_Handler = h.Interceptor(_TemplateService_QueryPipelineTemplates_Handler)
	}

	_TemplateService_RenderPipelineTemplate_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.RenderPipelineTemplate(ctx, req.(*PipelineTemplateRenderRequest))
	}
	var _TemplateService_RenderPipelineTemplate_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TemplateService_RenderPipelineTemplate_info = transport.NewServiceInfo("erda.dop.pipelinetemplate.TemplateService", "RenderPipelineTemplate", srv)
		_TemplateService_RenderPipelineTemplate_Handler = h.Interceptor(_TemplateService_RenderPipelineTemplate_Handler)
	}

	_TemplateService_RenderPipelineTemplateBySpec_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.RenderPipelineTemplateBySpec(ctx, req.(*PipelineTemplateRenderSpecRequest))
	}
	var _TemplateService_RenderPipelineTemplateBySpec_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TemplateService_RenderPipelineTemplateBySpec_info = transport.NewServiceInfo("erda.dop.pipelinetemplate.TemplateService", "RenderPipelineTemplateBySpec", srv)
		_TemplateService_RenderPipelineTemplateBySpec_Handler = h.Interceptor(_TemplateService_RenderPipelineTemplateBySpec_Handler)
	}

	_TemplateService_GetPipelineTemplateVersion_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetPipelineTemplateVersion(ctx, req.(*PipelineTemplateVersionGetRequest))
	}
	var _TemplateService_GetPipelineTemplateVersion_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TemplateService_GetPipelineTemplateVersion_info = transport.NewServiceInfo("erda.dop.pipelinetemplate.TemplateService", "GetPipelineTemplateVersion", srv)
		_TemplateService_GetPipelineTemplateVersion_Handler = h.Interceptor(_TemplateService_GetPipelineTemplateVersion_Handler)
	}

	_TemplateService_QueryPipelineTemplateVersions_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.QueryPipelineTemplateVersions(ctx, req.(*PipelineTemplateVersionQueryRequest))
	}
	var _TemplateService_QueryPipelineTemplateVersions_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TemplateService_QueryPipelineTemplateVersions_info = transport.NewServiceInfo("erda.dop.pipelinetemplate.TemplateService", "QueryPipelineTemplateVersions", srv)
		_TemplateService_QueryPipelineTemplateVersions_Handler = h.Interceptor(_TemplateService_QueryPipelineTemplateVersions_Handler)
	}

	_TemplateService_QuerySnippetYml_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.QuerySnippetYml(ctx, req.(*QuerySnippetYmlRequest))
	}
	var _TemplateService_QuerySnippetYml_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TemplateService_QuerySnippetYml_info = transport.NewServiceInfo("erda.dop.pipelinetemplate.TemplateService", "QuerySnippetYml", srv)
		_TemplateService_QuerySnippetYml_Handler = h.Interceptor(_TemplateService_QuerySnippetYml_Handler)
	}

	var serviceDesc = _TemplateService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "ApplyPipelineTemplate",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(PipelineTemplateApplyRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TemplateServiceServer).ApplyPipelineTemplate(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TemplateService_ApplyPipelineTemplate_info)
				}
				if interceptor == nil {
					return _TemplateService_ApplyPipelineTemplate_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.dop.pipelinetemplate.TemplateService/ApplyPipelineTemplate",
				}
				return interceptor(ctx, in, info, _TemplateService_ApplyPipelineTemplate_Handler)
			},
		},
		{
			MethodName: "QueryPipelineTemplates",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(PipelineTemplateQueryRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TemplateServiceServer).QueryPipelineTemplates(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TemplateService_QueryPipelineTemplates_info)
				}
				if interceptor == nil {
					return _TemplateService_QueryPipelineTemplates_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.dop.pipelinetemplate.TemplateService/QueryPipelineTemplates",
				}
				return interceptor(ctx, in, info, _TemplateService_QueryPipelineTemplates_Handler)
			},
		},
		{
			MethodName: "RenderPipelineTemplate",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(PipelineTemplateRenderRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TemplateServiceServer).RenderPipelineTemplate(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TemplateService_RenderPipelineTemplate_info)
				}
				if interceptor == nil {
					return _TemplateService_RenderPipelineTemplate_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.dop.pipelinetemplate.TemplateService/RenderPipelineTemplate",
				}
				return interceptor(ctx, in, info, _TemplateService_RenderPipelineTemplate_Handler)
			},
		},
		{
			MethodName: "RenderPipelineTemplateBySpec",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(PipelineTemplateRenderSpecRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TemplateServiceServer).RenderPipelineTemplateBySpec(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TemplateService_RenderPipelineTemplateBySpec_info)
				}
				if interceptor == nil {
					return _TemplateService_RenderPipelineTemplateBySpec_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.dop.pipelinetemplate.TemplateService/RenderPipelineTemplateBySpec",
				}
				return interceptor(ctx, in, info, _TemplateService_RenderPipelineTemplateBySpec_Handler)
			},
		},
		{
			MethodName: "GetPipelineTemplateVersion",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(PipelineTemplateVersionGetRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TemplateServiceServer).GetPipelineTemplateVersion(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TemplateService_GetPipelineTemplateVersion_info)
				}
				if interceptor == nil {
					return _TemplateService_GetPipelineTemplateVersion_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.dop.pipelinetemplate.TemplateService/GetPipelineTemplateVersion",
				}
				return interceptor(ctx, in, info, _TemplateService_GetPipelineTemplateVersion_Handler)
			},
		},
		{
			MethodName: "QueryPipelineTemplateVersions",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(PipelineTemplateVersionQueryRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TemplateServiceServer).QueryPipelineTemplateVersions(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TemplateService_QueryPipelineTemplateVersions_info)
				}
				if interceptor == nil {
					return _TemplateService_QueryPipelineTemplateVersions_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.dop.pipelinetemplate.TemplateService/QueryPipelineTemplateVersions",
				}
				return interceptor(ctx, in, info, _TemplateService_QueryPipelineTemplateVersions_Handler)
			},
		},
		{
			MethodName: "QuerySnippetYml",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(QuerySnippetYmlRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TemplateServiceServer).QuerySnippetYml(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TemplateService_QuerySnippetYml_info)
				}
				if interceptor == nil {
					return _TemplateService_QuerySnippetYml_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.dop.pipelinetemplate.TemplateService/QuerySnippetYml",
				}
				return interceptor(ctx, in, info, _TemplateService_QuerySnippetYml_Handler)
			},
		},
	}
	return &serviceDesc
}
