// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: onedata.proto

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

// OnedataClient is the client API for Onedata service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OnedataClient interface {
	ONEDATA_ANALYSIS(ctx context.Context, in *OneDataAnalysisRequest, opts ...grpc.CallOption) (*OneDataAnalysisResponse, error)
	ONEDATA_ANALYSIS_BUSSPROC(ctx context.Context, in *OneDataAnalysisBussProcRequest, opts ...grpc.CallOption) (*OneDataAnalysisBussProcResponse, error)
	ONEDATA_ANALYSIS_BUSSPROCS(ctx context.Context, in *OneDataAnalysisBussProcsRequest, opts ...grpc.CallOption) (*OneDataAnalysisBussProcsResponse, error)
	ONEDATA_ANALYSIS_DIM(ctx context.Context, in *OneDataAnalysisDimRequest, opts ...grpc.CallOption) (*OneDataAnalysisDimRequest, error)
	ONEDATA_ANALYSIS_FUZZYATTRS(ctx context.Context, in *OneDataAnalysisFuzzyAttrsRequest, opts ...grpc.CallOption) (*OneDataAnalysisFuzzyAttrsResponse, error)
	ONEDATA_ANALYSIS_OUTPUTTABLES(ctx context.Context, in *OneDataAnalysisOutputTablesRequest, opts ...grpc.CallOption) (*OneDataAnalysisOutputTablesResponse, error)
	ONEDATA_ANALYSIS_STAR(ctx context.Context, in *OneDataAnalysisStarRequest, opts ...grpc.CallOption) (*OneDataAnalysisStarResponse, error)
}

type onedataClient struct {
	cc grpc1.ClientConnInterface
}

func NewOnedataClient(cc grpc1.ClientConnInterface) OnedataClient {
	return &onedataClient{cc}
}

func (c *onedataClient) ONEDATA_ANALYSIS(ctx context.Context, in *OneDataAnalysisRequest, opts ...grpc.CallOption) (*OneDataAnalysisResponse, error) {
	out := new(OneDataAnalysisResponse)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.onedata.onedata/ONEDATA_ANALYSIS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onedataClient) ONEDATA_ANALYSIS_BUSSPROC(ctx context.Context, in *OneDataAnalysisBussProcRequest, opts ...grpc.CallOption) (*OneDataAnalysisBussProcResponse, error) {
	out := new(OneDataAnalysisBussProcResponse)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.onedata.onedata/ONEDATA_ANALYSIS_BUSSPROC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onedataClient) ONEDATA_ANALYSIS_BUSSPROCS(ctx context.Context, in *OneDataAnalysisBussProcsRequest, opts ...grpc.CallOption) (*OneDataAnalysisBussProcsResponse, error) {
	out := new(OneDataAnalysisBussProcsResponse)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.onedata.onedata/ONEDATA_ANALYSIS_BUSSPROCS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onedataClient) ONEDATA_ANALYSIS_DIM(ctx context.Context, in *OneDataAnalysisDimRequest, opts ...grpc.CallOption) (*OneDataAnalysisDimRequest, error) {
	out := new(OneDataAnalysisDimRequest)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.onedata.onedata/ONEDATA_ANALYSIS_DIM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onedataClient) ONEDATA_ANALYSIS_FUZZYATTRS(ctx context.Context, in *OneDataAnalysisFuzzyAttrsRequest, opts ...grpc.CallOption) (*OneDataAnalysisFuzzyAttrsResponse, error) {
	out := new(OneDataAnalysisFuzzyAttrsResponse)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.onedata.onedata/ONEDATA_ANALYSIS_FUZZYATTRS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onedataClient) ONEDATA_ANALYSIS_OUTPUTTABLES(ctx context.Context, in *OneDataAnalysisOutputTablesRequest, opts ...grpc.CallOption) (*OneDataAnalysisOutputTablesResponse, error) {
	out := new(OneDataAnalysisOutputTablesResponse)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.onedata.onedata/ONEDATA_ANALYSIS_OUTPUTTABLES", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onedataClient) ONEDATA_ANALYSIS_STAR(ctx context.Context, in *OneDataAnalysisStarRequest, opts ...grpc.CallOption) (*OneDataAnalysisStarResponse, error) {
	out := new(OneDataAnalysisStarResponse)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.onedata.onedata/ONEDATA_ANALYSIS_STAR", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OnedataServer is the server API for Onedata service.
// All implementations should embed UnimplementedOnedataServer
// for forward compatibility
type OnedataServer interface {
	ONEDATA_ANALYSIS(context.Context, *OneDataAnalysisRequest) (*OneDataAnalysisResponse, error)
	ONEDATA_ANALYSIS_BUSSPROC(context.Context, *OneDataAnalysisBussProcRequest) (*OneDataAnalysisBussProcResponse, error)
	ONEDATA_ANALYSIS_BUSSPROCS(context.Context, *OneDataAnalysisBussProcsRequest) (*OneDataAnalysisBussProcsResponse, error)
	ONEDATA_ANALYSIS_DIM(context.Context, *OneDataAnalysisDimRequest) (*OneDataAnalysisDimRequest, error)
	ONEDATA_ANALYSIS_FUZZYATTRS(context.Context, *OneDataAnalysisFuzzyAttrsRequest) (*OneDataAnalysisFuzzyAttrsResponse, error)
	ONEDATA_ANALYSIS_OUTPUTTABLES(context.Context, *OneDataAnalysisOutputTablesRequest) (*OneDataAnalysisOutputTablesResponse, error)
	ONEDATA_ANALYSIS_STAR(context.Context, *OneDataAnalysisStarRequest) (*OneDataAnalysisStarResponse, error)
}

// UnimplementedOnedataServer should be embedded to have forward compatible implementations.
type UnimplementedOnedataServer struct {
}

func (*UnimplementedOnedataServer) ONEDATA_ANALYSIS(context.Context, *OneDataAnalysisRequest) (*OneDataAnalysisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ONEDATA_ANALYSIS not implemented")
}
func (*UnimplementedOnedataServer) ONEDATA_ANALYSIS_BUSSPROC(context.Context, *OneDataAnalysisBussProcRequest) (*OneDataAnalysisBussProcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ONEDATA_ANALYSIS_BUSSPROC not implemented")
}
func (*UnimplementedOnedataServer) ONEDATA_ANALYSIS_BUSSPROCS(context.Context, *OneDataAnalysisBussProcsRequest) (*OneDataAnalysisBussProcsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ONEDATA_ANALYSIS_BUSSPROCS not implemented")
}
func (*UnimplementedOnedataServer) ONEDATA_ANALYSIS_DIM(context.Context, *OneDataAnalysisDimRequest) (*OneDataAnalysisDimRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ONEDATA_ANALYSIS_DIM not implemented")
}
func (*UnimplementedOnedataServer) ONEDATA_ANALYSIS_FUZZYATTRS(context.Context, *OneDataAnalysisFuzzyAttrsRequest) (*OneDataAnalysisFuzzyAttrsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ONEDATA_ANALYSIS_FUZZYATTRS not implemented")
}
func (*UnimplementedOnedataServer) ONEDATA_ANALYSIS_OUTPUTTABLES(context.Context, *OneDataAnalysisOutputTablesRequest) (*OneDataAnalysisOutputTablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ONEDATA_ANALYSIS_OUTPUTTABLES not implemented")
}
func (*UnimplementedOnedataServer) ONEDATA_ANALYSIS_STAR(context.Context, *OneDataAnalysisStarRequest) (*OneDataAnalysisStarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ONEDATA_ANALYSIS_STAR not implemented")
}

func RegisterOnedataServer(s grpc1.ServiceRegistrar, srv OnedataServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_Onedata_serviceDesc(srv, opts...), srv)
}

var _Onedata_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.openapiv1.onedata.onedata",
	HandlerType: (*OnedataServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "onedata.proto",
}

func _get_Onedata_serviceDesc(srv OnedataServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_Onedata_ONEDATA_ANALYSIS_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ONEDATA_ANALYSIS(ctx, req.(*OneDataAnalysisRequest))
	}
	var _Onedata_ONEDATA_ANALYSIS_info transport.ServiceInfo
	if h.Interceptor != nil {
		_Onedata_ONEDATA_ANALYSIS_info = transport.NewServiceInfo("erda.openapiv1.onedata.onedata", "ONEDATA_ANALYSIS", srv)
		_Onedata_ONEDATA_ANALYSIS_Handler = h.Interceptor(_Onedata_ONEDATA_ANALYSIS_Handler)
	}

	_Onedata_ONEDATA_ANALYSIS_BUSSPROC_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ONEDATA_ANALYSIS_BUSSPROC(ctx, req.(*OneDataAnalysisBussProcRequest))
	}
	var _Onedata_ONEDATA_ANALYSIS_BUSSPROC_info transport.ServiceInfo
	if h.Interceptor != nil {
		_Onedata_ONEDATA_ANALYSIS_BUSSPROC_info = transport.NewServiceInfo("erda.openapiv1.onedata.onedata", "ONEDATA_ANALYSIS_BUSSPROC", srv)
		_Onedata_ONEDATA_ANALYSIS_BUSSPROC_Handler = h.Interceptor(_Onedata_ONEDATA_ANALYSIS_BUSSPROC_Handler)
	}

	_Onedata_ONEDATA_ANALYSIS_BUSSPROCS_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ONEDATA_ANALYSIS_BUSSPROCS(ctx, req.(*OneDataAnalysisBussProcsRequest))
	}
	var _Onedata_ONEDATA_ANALYSIS_BUSSPROCS_info transport.ServiceInfo
	if h.Interceptor != nil {
		_Onedata_ONEDATA_ANALYSIS_BUSSPROCS_info = transport.NewServiceInfo("erda.openapiv1.onedata.onedata", "ONEDATA_ANALYSIS_BUSSPROCS", srv)
		_Onedata_ONEDATA_ANALYSIS_BUSSPROCS_Handler = h.Interceptor(_Onedata_ONEDATA_ANALYSIS_BUSSPROCS_Handler)
	}

	_Onedata_ONEDATA_ANALYSIS_DIM_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ONEDATA_ANALYSIS_DIM(ctx, req.(*OneDataAnalysisDimRequest))
	}
	var _Onedata_ONEDATA_ANALYSIS_DIM_info transport.ServiceInfo
	if h.Interceptor != nil {
		_Onedata_ONEDATA_ANALYSIS_DIM_info = transport.NewServiceInfo("erda.openapiv1.onedata.onedata", "ONEDATA_ANALYSIS_DIM", srv)
		_Onedata_ONEDATA_ANALYSIS_DIM_Handler = h.Interceptor(_Onedata_ONEDATA_ANALYSIS_DIM_Handler)
	}

	_Onedata_ONEDATA_ANALYSIS_FUZZYATTRS_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ONEDATA_ANALYSIS_FUZZYATTRS(ctx, req.(*OneDataAnalysisFuzzyAttrsRequest))
	}
	var _Onedata_ONEDATA_ANALYSIS_FUZZYATTRS_info transport.ServiceInfo
	if h.Interceptor != nil {
		_Onedata_ONEDATA_ANALYSIS_FUZZYATTRS_info = transport.NewServiceInfo("erda.openapiv1.onedata.onedata", "ONEDATA_ANALYSIS_FUZZYATTRS", srv)
		_Onedata_ONEDATA_ANALYSIS_FUZZYATTRS_Handler = h.Interceptor(_Onedata_ONEDATA_ANALYSIS_FUZZYATTRS_Handler)
	}

	_Onedata_ONEDATA_ANALYSIS_OUTPUTTABLES_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ONEDATA_ANALYSIS_OUTPUTTABLES(ctx, req.(*OneDataAnalysisOutputTablesRequest))
	}
	var _Onedata_ONEDATA_ANALYSIS_OUTPUTTABLES_info transport.ServiceInfo
	if h.Interceptor != nil {
		_Onedata_ONEDATA_ANALYSIS_OUTPUTTABLES_info = transport.NewServiceInfo("erda.openapiv1.onedata.onedata", "ONEDATA_ANALYSIS_OUTPUTTABLES", srv)
		_Onedata_ONEDATA_ANALYSIS_OUTPUTTABLES_Handler = h.Interceptor(_Onedata_ONEDATA_ANALYSIS_OUTPUTTABLES_Handler)
	}

	_Onedata_ONEDATA_ANALYSIS_STAR_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ONEDATA_ANALYSIS_STAR(ctx, req.(*OneDataAnalysisStarRequest))
	}
	var _Onedata_ONEDATA_ANALYSIS_STAR_info transport.ServiceInfo
	if h.Interceptor != nil {
		_Onedata_ONEDATA_ANALYSIS_STAR_info = transport.NewServiceInfo("erda.openapiv1.onedata.onedata", "ONEDATA_ANALYSIS_STAR", srv)
		_Onedata_ONEDATA_ANALYSIS_STAR_Handler = h.Interceptor(_Onedata_ONEDATA_ANALYSIS_STAR_Handler)
	}

	var serviceDesc = _Onedata_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "ONEDATA_ANALYSIS",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(OneDataAnalysisRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(OnedataServer).ONEDATA_ANALYSIS(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _Onedata_ONEDATA_ANALYSIS_info)
				}
				if interceptor == nil {
					return _Onedata_ONEDATA_ANALYSIS_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.onedata.onedata/ONEDATA_ANALYSIS",
				}
				return interceptor(ctx, in, info, _Onedata_ONEDATA_ANALYSIS_Handler)
			},
		},
		{
			MethodName: "ONEDATA_ANALYSIS_BUSSPROC",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(OneDataAnalysisBussProcRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(OnedataServer).ONEDATA_ANALYSIS_BUSSPROC(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _Onedata_ONEDATA_ANALYSIS_BUSSPROC_info)
				}
				if interceptor == nil {
					return _Onedata_ONEDATA_ANALYSIS_BUSSPROC_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.onedata.onedata/ONEDATA_ANALYSIS_BUSSPROC",
				}
				return interceptor(ctx, in, info, _Onedata_ONEDATA_ANALYSIS_BUSSPROC_Handler)
			},
		},
		{
			MethodName: "ONEDATA_ANALYSIS_BUSSPROCS",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(OneDataAnalysisBussProcsRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(OnedataServer).ONEDATA_ANALYSIS_BUSSPROCS(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _Onedata_ONEDATA_ANALYSIS_BUSSPROCS_info)
				}
				if interceptor == nil {
					return _Onedata_ONEDATA_ANALYSIS_BUSSPROCS_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.onedata.onedata/ONEDATA_ANALYSIS_BUSSPROCS",
				}
				return interceptor(ctx, in, info, _Onedata_ONEDATA_ANALYSIS_BUSSPROCS_Handler)
			},
		},
		{
			MethodName: "ONEDATA_ANALYSIS_DIM",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(OneDataAnalysisDimRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(OnedataServer).ONEDATA_ANALYSIS_DIM(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _Onedata_ONEDATA_ANALYSIS_DIM_info)
				}
				if interceptor == nil {
					return _Onedata_ONEDATA_ANALYSIS_DIM_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.onedata.onedata/ONEDATA_ANALYSIS_DIM",
				}
				return interceptor(ctx, in, info, _Onedata_ONEDATA_ANALYSIS_DIM_Handler)
			},
		},
		{
			MethodName: "ONEDATA_ANALYSIS_FUZZYATTRS",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(OneDataAnalysisFuzzyAttrsRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(OnedataServer).ONEDATA_ANALYSIS_FUZZYATTRS(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _Onedata_ONEDATA_ANALYSIS_FUZZYATTRS_info)
				}
				if interceptor == nil {
					return _Onedata_ONEDATA_ANALYSIS_FUZZYATTRS_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.onedata.onedata/ONEDATA_ANALYSIS_FUZZYATTRS",
				}
				return interceptor(ctx, in, info, _Onedata_ONEDATA_ANALYSIS_FUZZYATTRS_Handler)
			},
		},
		{
			MethodName: "ONEDATA_ANALYSIS_OUTPUTTABLES",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(OneDataAnalysisOutputTablesRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(OnedataServer).ONEDATA_ANALYSIS_OUTPUTTABLES(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _Onedata_ONEDATA_ANALYSIS_OUTPUTTABLES_info)
				}
				if interceptor == nil {
					return _Onedata_ONEDATA_ANALYSIS_OUTPUTTABLES_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.onedata.onedata/ONEDATA_ANALYSIS_OUTPUTTABLES",
				}
				return interceptor(ctx, in, info, _Onedata_ONEDATA_ANALYSIS_OUTPUTTABLES_Handler)
			},
		},
		{
			MethodName: "ONEDATA_ANALYSIS_STAR",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(OneDataAnalysisStarRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(OnedataServer).ONEDATA_ANALYSIS_STAR(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _Onedata_ONEDATA_ANALYSIS_STAR_info)
				}
				if interceptor == nil {
					return _Onedata_ONEDATA_ANALYSIS_STAR_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.onedata.onedata/ONEDATA_ANALYSIS_STAR",
				}
				return interceptor(ctx, in, info, _Onedata_ONEDATA_ANALYSIS_STAR_Handler)
			},
		},
	}
	return &serviceDesc
}