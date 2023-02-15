// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: msp_apm_block.proto

package pb

import (
	context "context"
	http1 "net/http"
	strings "strings"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	httprule "github.com/erda-project/erda-infra/pkg/transport/http/httprule"
	runtime "github.com/erda-project/erda-infra/pkg/transport/http/runtime"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// MspApmBlockHandler is the server API for MspApmBlock service.
type MspApmBlockHandler interface {
	// POST /api/tmc/dashboard/blocks
	CREATE_BLOCK(context.Context, *CREATE_BLOCK_Request) (*emptypb.Empty, error)
	// DELETE /api/tmc/dashboard/blocks/{id}
	DELETE_BLOCK(context.Context, *DELETE_BLOCK_Request) (*emptypb.Empty, error)
	// GET /api/tmc/dashboard/blocks/{id}
	GET_BLOCK(context.Context, *GET_BLOCK_Request) (*emptypb.Empty, error)
	// GET /api/tmc/dashboard/blocks
	LIST_BLOCKS(context.Context, *LIST_BLOCKS_Request) (*emptypb.Empty, error)
	// PUT /api/tmc/dashboard/blocks/{id}
	TMC_METRIC_DASHBOARD_UPDATE(context.Context, *TMC_METRIC_DASHBOARD_UPDATE_Request) (*emptypb.Empty, error)
}

// RegisterMspApmBlockHandler register MspApmBlockHandler to http.Router.
func RegisterMspApmBlockHandler(r http.Router, srv MspApmBlockHandler, opts ...http.HandleOption) {
	h := http.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}
	encodeFunc := func(fn func(http1.ResponseWriter, *http1.Request) (interface{}, error)) http.HandlerFunc {
		handler := func(w http1.ResponseWriter, r *http1.Request) {
			out, err := fn(w, r)
			if err != nil {
				h.Error(w, r, err)
				return
			}
			if err := h.Encode(w, r, out); err != nil {
				h.Error(w, r, err)
			}
		}
		if h.HTTPInterceptor != nil {
			handler = h.HTTPInterceptor(handler)
		}
		return handler
	}

	add_CREATE_BLOCK := func(method, path string, fn func(context.Context, *CREATE_BLOCK_Request) (*emptypb.Empty, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CREATE_BLOCK_Request))
		}
		var CREATE_BLOCK_info transport.ServiceInfo
		if h.Interceptor != nil {
			CREATE_BLOCK_info = transport.NewServiceInfo("erda.openapiv1.msp.msp_apm_block", "CREATE_BLOCK", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CREATE_BLOCK_info)
				}
				r = r.WithContext(ctx)
				var in CREATE_BLOCK_Request
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_DELETE_BLOCK := func(method, path string, fn func(context.Context, *DELETE_BLOCK_Request) (*emptypb.Empty, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*DELETE_BLOCK_Request))
		}
		var DELETE_BLOCK_info transport.ServiceInfo
		if h.Interceptor != nil {
			DELETE_BLOCK_info = transport.NewServiceInfo("erda.openapiv1.msp.msp_apm_block", "DELETE_BLOCK", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DELETE_BLOCK_info)
				}
				r = r.WithContext(ctx)
				var in DELETE_BLOCK_Request
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "id":
							in.Id = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GET_BLOCK := func(method, path string, fn func(context.Context, *GET_BLOCK_Request) (*emptypb.Empty, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GET_BLOCK_Request))
		}
		var GET_BLOCK_info transport.ServiceInfo
		if h.Interceptor != nil {
			GET_BLOCK_info = transport.NewServiceInfo("erda.openapiv1.msp.msp_apm_block", "GET_BLOCK", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GET_BLOCK_info)
				}
				r = r.WithContext(ctx)
				var in GET_BLOCK_Request
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "id":
							in.Id = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_LIST_BLOCKS := func(method, path string, fn func(context.Context, *LIST_BLOCKS_Request) (*emptypb.Empty, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*LIST_BLOCKS_Request))
		}
		var LIST_BLOCKS_info transport.ServiceInfo
		if h.Interceptor != nil {
			LIST_BLOCKS_info = transport.NewServiceInfo("erda.openapiv1.msp.msp_apm_block", "LIST_BLOCKS", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, LIST_BLOCKS_info)
				}
				r = r.WithContext(ctx)
				var in LIST_BLOCKS_Request
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_TMC_METRIC_DASHBOARD_UPDATE := func(method, path string, fn func(context.Context, *TMC_METRIC_DASHBOARD_UPDATE_Request) (*emptypb.Empty, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*TMC_METRIC_DASHBOARD_UPDATE_Request))
		}
		var TMC_METRIC_DASHBOARD_UPDATE_info transport.ServiceInfo
		if h.Interceptor != nil {
			TMC_METRIC_DASHBOARD_UPDATE_info = transport.NewServiceInfo("erda.openapiv1.msp.msp_apm_block", "TMC_METRIC_DASHBOARD_UPDATE", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, TMC_METRIC_DASHBOARD_UPDATE_info)
				}
				r = r.WithContext(ctx)
				var in TMC_METRIC_DASHBOARD_UPDATE_Request
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "id":
							in.Id = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_CREATE_BLOCK("POST", "/api/tmc/dashboard/blocks", srv.CREATE_BLOCK)
	add_DELETE_BLOCK("DELETE", "/api/tmc/dashboard/blocks/{id}", srv.DELETE_BLOCK)
	add_GET_BLOCK("GET", "/api/tmc/dashboard/blocks/{id}", srv.GET_BLOCK)
	add_LIST_BLOCKS("GET", "/api/tmc/dashboard/blocks", srv.LIST_BLOCKS)
	add_TMC_METRIC_DASHBOARD_UPDATE("PUT", "/api/tmc/dashboard/blocks/{id}", srv.TMC_METRIC_DASHBOARD_UPDATE)
}
