// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: officer.proto

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

// OfficerHandler is the server API for Officer service.
type OfficerHandler interface {
	// GET /api/script/{name}
	OFFICER_FETCH_SCRIPT(context.Context, *OFFICER_FETCH_SCRIPT_Request) (*emptypb.Empty, error)
	// POST /api/resource/aliyun/price
	OFFICER_GET_PRICE(context.Context, *OFFICER_GET_PRICE_Request) (*emptypb.Empty, error)
	// GET /api/script/info
	OFFICER_GET_SCRIPT_INFO(context.Context, *OFFICER_GET_SCRIPT_INFO_Request) (*ScriptInfo, error)
}

// RegisterOfficerHandler register OfficerHandler to http.Router.
func RegisterOfficerHandler(r http.Router, srv OfficerHandler, opts ...http.HandleOption) {
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

	add_OFFICER_FETCH_SCRIPT := func(method, path string, fn func(context.Context, *OFFICER_FETCH_SCRIPT_Request) (*emptypb.Empty, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*OFFICER_FETCH_SCRIPT_Request))
		}
		var OFFICER_FETCH_SCRIPT_info transport.ServiceInfo
		if h.Interceptor != nil {
			OFFICER_FETCH_SCRIPT_info = transport.NewServiceInfo("erda.openapiv1.officer.officer", "OFFICER_FETCH_SCRIPT", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, OFFICER_FETCH_SCRIPT_info)
				}
				r = r.WithContext(ctx)
				var in OFFICER_FETCH_SCRIPT_Request
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
						case "name":
							in.Name = val
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

	add_OFFICER_GET_PRICE := func(method, path string, fn func(context.Context, *OFFICER_GET_PRICE_Request) (*emptypb.Empty, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*OFFICER_GET_PRICE_Request))
		}
		var OFFICER_GET_PRICE_info transport.ServiceInfo
		if h.Interceptor != nil {
			OFFICER_GET_PRICE_info = transport.NewServiceInfo("erda.openapiv1.officer.officer", "OFFICER_GET_PRICE", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, OFFICER_GET_PRICE_info)
				}
				r = r.WithContext(ctx)
				var in OFFICER_GET_PRICE_Request
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

	add_OFFICER_GET_SCRIPT_INFO := func(method, path string, fn func(context.Context, *OFFICER_GET_SCRIPT_INFO_Request) (*ScriptInfo, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*OFFICER_GET_SCRIPT_INFO_Request))
		}
		var OFFICER_GET_SCRIPT_INFO_info transport.ServiceInfo
		if h.Interceptor != nil {
			OFFICER_GET_SCRIPT_INFO_info = transport.NewServiceInfo("erda.openapiv1.officer.officer", "OFFICER_GET_SCRIPT_INFO", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, OFFICER_GET_SCRIPT_INFO_info)
				}
				r = r.WithContext(ctx)
				var in OFFICER_GET_SCRIPT_INFO_Request
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

	add_OFFICER_FETCH_SCRIPT("GET", "/api/script/{name}", srv.OFFICER_FETCH_SCRIPT)
	add_OFFICER_GET_PRICE("POST", "/api/resource/aliyun/price", srv.OFFICER_GET_PRICE)
	add_OFFICER_GET_SCRIPT_INFO("GET", "/api/script/info", srv.OFFICER_GET_SCRIPT_INFO)
}
