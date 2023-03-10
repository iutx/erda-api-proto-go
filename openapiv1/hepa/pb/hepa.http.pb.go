// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: hepa.proto

package pb

import (
	context "context"
	http1 "net/http"
	strconv "strconv"
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

// HepaHandler is the server API for Hepa service.
type HepaHandler interface {
	// GET /api/domains
	HEPA_DOMAINS_GET(context.Context, *HEPA_DOMAINS_GET_Request) (*emptypb.Empty, error)
	// GET /api/runtimes/{runtimeID}/k8s-domains
	HEPA_RUNTIME_DOMAIN_GET(context.Context, *DomainListRequest) (*DomainListResponse, error)
	// PUT /api/runtimes/{runtimeId}/services/{serviceName}/domains
	HEPA_RUNTIME_SERVICE_DOMAIN_UPDATE(context.Context, *HEPA_RUNTIME_SERVICE_DOMAIN_UPDATE_Request) (*emptypb.Empty, error)
}

// RegisterHepaHandler register HepaHandler to http.Router.
func RegisterHepaHandler(r http.Router, srv HepaHandler, opts ...http.HandleOption) {
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

	add_HEPA_DOMAINS_GET := func(method, path string, fn func(context.Context, *HEPA_DOMAINS_GET_Request) (*emptypb.Empty, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*HEPA_DOMAINS_GET_Request))
		}
		var HEPA_DOMAINS_GET_info transport.ServiceInfo
		if h.Interceptor != nil {
			HEPA_DOMAINS_GET_info = transport.NewServiceInfo("erda.openapiv1.hepa.hepa", "HEPA_DOMAINS_GET", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, HEPA_DOMAINS_GET_info)
				}
				r = r.WithContext(ctx)
				var in HEPA_DOMAINS_GET_Request
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

	add_HEPA_RUNTIME_DOMAIN_GET := func(method, path string, fn func(context.Context, *DomainListRequest) (*DomainListResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*DomainListRequest))
		}
		var HEPA_RUNTIME_DOMAIN_GET_info transport.ServiceInfo
		if h.Interceptor != nil {
			HEPA_RUNTIME_DOMAIN_GET_info = transport.NewServiceInfo("erda.openapiv1.hepa.hepa", "HEPA_RUNTIME_DOMAIN_GET", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, HEPA_RUNTIME_DOMAIN_GET_info)
				}
				r = r.WithContext(ctx)
				var in DomainListRequest
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
						case "runtimeID":
							val, err := strconv.ParseUint(val, 10, 64)
							if err != nil {
								return nil, err
							}
							in.RuntimeID = val
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

	add_HEPA_RUNTIME_SERVICE_DOMAIN_UPDATE := func(method, path string, fn func(context.Context, *HEPA_RUNTIME_SERVICE_DOMAIN_UPDATE_Request) (*emptypb.Empty, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*HEPA_RUNTIME_SERVICE_DOMAIN_UPDATE_Request))
		}
		var HEPA_RUNTIME_SERVICE_DOMAIN_UPDATE_info transport.ServiceInfo
		if h.Interceptor != nil {
			HEPA_RUNTIME_SERVICE_DOMAIN_UPDATE_info = transport.NewServiceInfo("erda.openapiv1.hepa.hepa", "HEPA_RUNTIME_SERVICE_DOMAIN_UPDATE", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, HEPA_RUNTIME_SERVICE_DOMAIN_UPDATE_info)
				}
				r = r.WithContext(ctx)
				var in HEPA_RUNTIME_SERVICE_DOMAIN_UPDATE_Request
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
						case "runtimeId":
							in.RuntimeId = val
						case "serviceName":
							in.ServiceName = val
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

	add_HEPA_DOMAINS_GET("GET", "/api/domains", srv.HEPA_DOMAINS_GET)
	add_HEPA_RUNTIME_DOMAIN_GET("GET", "/api/runtimes/{runtimeID}/k8s-domains", srv.HEPA_RUNTIME_DOMAIN_GET)
	add_HEPA_RUNTIME_SERVICE_DOMAIN_UPDATE("PUT", "/api/runtimes/{runtimeId}/services/{serviceName}/domains", srv.HEPA_RUNTIME_SERVICE_DOMAIN_UPDATE)
}
