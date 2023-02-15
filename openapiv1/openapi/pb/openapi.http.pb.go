// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: openapi.proto

package pb

import (
	context "context"
	http1 "net/http"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// OpenapiHandler is the server API for Openapi service.
type OpenapiHandler interface {
	// GET /metadata.json
	DICE_METADATA(context.Context, *DICE_METADATA_Request) (*emptypb.Empty, error)
	// GET /api/openapi/swagger.json
	DOC_JSON(context.Context, *DOC_JSON_Request) (*emptypb.Empty, error)
	// GET /api/openapi-doc
	OPENAPI_DOC(context.Context, *OPENAPI_DOC_Request) (*emptypb.Empty, error)
	// GET /api/openapi-event-doc
	OPENAPI_EVENT_DOC(context.Context, *OPENAPI_EVENT_DOC_Request) (*emptypb.Empty, error)
	// POST /api/openapi/client-token
	OPENAPI_GEN_CLIENT_TOKEN(context.Context, *OPENAPI_GEN_CLIENT_TOKEN_Request) (*emptypb.Empty, error)
	// GET /api/openapi/manager/clients
	OPENAPI_LIST_CLIENT(context.Context, *OPENAPI_LIST_CLIENT_Request) (*emptypb.Empty, error)
	// GET /api/openapi/metrics
	OPENAPI_METRICS(context.Context, *OPENAPI_METRICS_Request) (*emptypb.Empty, error)
	// POST /api/openapi/manager/clients
	OPENAPI_NEW_CLIENT(context.Context, *OPENAPI_NEW_CLIENT_Request) (*emptypb.Empty, error)
	// GET /api/openapi/stat
	OPENAPI_STAT(context.Context, *OPENAPI_STAT_Request) (*emptypb.Empty, error)
	// GET /api/openapi/version
	OPENAPI_VERSION(context.Context, *OPENAPI_VERSION_Request) (*emptypb.Empty, error)
}

// RegisterOpenapiHandler register OpenapiHandler to http.Router.
func RegisterOpenapiHandler(r http.Router, srv OpenapiHandler, opts ...http.HandleOption) {
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

	add_DICE_METADATA := func(method, path string, fn func(context.Context, *DICE_METADATA_Request) (*emptypb.Empty, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*DICE_METADATA_Request))
		}
		var DICE_METADATA_info transport.ServiceInfo
		if h.Interceptor != nil {
			DICE_METADATA_info = transport.NewServiceInfo("erda.openapiv1.openapi.openapi", "DICE_METADATA", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DICE_METADATA_info)
				}
				r = r.WithContext(ctx)
				var in DICE_METADATA_Request
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

	add_DOC_JSON := func(method, path string, fn func(context.Context, *DOC_JSON_Request) (*emptypb.Empty, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*DOC_JSON_Request))
		}
		var DOC_JSON_info transport.ServiceInfo
		if h.Interceptor != nil {
			DOC_JSON_info = transport.NewServiceInfo("erda.openapiv1.openapi.openapi", "DOC_JSON", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DOC_JSON_info)
				}
				r = r.WithContext(ctx)
				var in DOC_JSON_Request
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

	add_OPENAPI_DOC := func(method, path string, fn func(context.Context, *OPENAPI_DOC_Request) (*emptypb.Empty, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*OPENAPI_DOC_Request))
		}
		var OPENAPI_DOC_info transport.ServiceInfo
		if h.Interceptor != nil {
			OPENAPI_DOC_info = transport.NewServiceInfo("erda.openapiv1.openapi.openapi", "OPENAPI_DOC", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, OPENAPI_DOC_info)
				}
				r = r.WithContext(ctx)
				var in OPENAPI_DOC_Request
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

	add_OPENAPI_EVENT_DOC := func(method, path string, fn func(context.Context, *OPENAPI_EVENT_DOC_Request) (*emptypb.Empty, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*OPENAPI_EVENT_DOC_Request))
		}
		var OPENAPI_EVENT_DOC_info transport.ServiceInfo
		if h.Interceptor != nil {
			OPENAPI_EVENT_DOC_info = transport.NewServiceInfo("erda.openapiv1.openapi.openapi", "OPENAPI_EVENT_DOC", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, OPENAPI_EVENT_DOC_info)
				}
				r = r.WithContext(ctx)
				var in OPENAPI_EVENT_DOC_Request
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

	add_OPENAPI_GEN_CLIENT_TOKEN := func(method, path string, fn func(context.Context, *OPENAPI_GEN_CLIENT_TOKEN_Request) (*emptypb.Empty, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*OPENAPI_GEN_CLIENT_TOKEN_Request))
		}
		var OPENAPI_GEN_CLIENT_TOKEN_info transport.ServiceInfo
		if h.Interceptor != nil {
			OPENAPI_GEN_CLIENT_TOKEN_info = transport.NewServiceInfo("erda.openapiv1.openapi.openapi", "OPENAPI_GEN_CLIENT_TOKEN", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, OPENAPI_GEN_CLIENT_TOKEN_info)
				}
				r = r.WithContext(ctx)
				var in OPENAPI_GEN_CLIENT_TOKEN_Request
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

	add_OPENAPI_LIST_CLIENT := func(method, path string, fn func(context.Context, *OPENAPI_LIST_CLIENT_Request) (*emptypb.Empty, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*OPENAPI_LIST_CLIENT_Request))
		}
		var OPENAPI_LIST_CLIENT_info transport.ServiceInfo
		if h.Interceptor != nil {
			OPENAPI_LIST_CLIENT_info = transport.NewServiceInfo("erda.openapiv1.openapi.openapi", "OPENAPI_LIST_CLIENT", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, OPENAPI_LIST_CLIENT_info)
				}
				r = r.WithContext(ctx)
				var in OPENAPI_LIST_CLIENT_Request
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

	add_OPENAPI_METRICS := func(method, path string, fn func(context.Context, *OPENAPI_METRICS_Request) (*emptypb.Empty, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*OPENAPI_METRICS_Request))
		}
		var OPENAPI_METRICS_info transport.ServiceInfo
		if h.Interceptor != nil {
			OPENAPI_METRICS_info = transport.NewServiceInfo("erda.openapiv1.openapi.openapi", "OPENAPI_METRICS", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, OPENAPI_METRICS_info)
				}
				r = r.WithContext(ctx)
				var in OPENAPI_METRICS_Request
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

	add_OPENAPI_NEW_CLIENT := func(method, path string, fn func(context.Context, *OPENAPI_NEW_CLIENT_Request) (*emptypb.Empty, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*OPENAPI_NEW_CLIENT_Request))
		}
		var OPENAPI_NEW_CLIENT_info transport.ServiceInfo
		if h.Interceptor != nil {
			OPENAPI_NEW_CLIENT_info = transport.NewServiceInfo("erda.openapiv1.openapi.openapi", "OPENAPI_NEW_CLIENT", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, OPENAPI_NEW_CLIENT_info)
				}
				r = r.WithContext(ctx)
				var in OPENAPI_NEW_CLIENT_Request
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

	add_OPENAPI_STAT := func(method, path string, fn func(context.Context, *OPENAPI_STAT_Request) (*emptypb.Empty, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*OPENAPI_STAT_Request))
		}
		var OPENAPI_STAT_info transport.ServiceInfo
		if h.Interceptor != nil {
			OPENAPI_STAT_info = transport.NewServiceInfo("erda.openapiv1.openapi.openapi", "OPENAPI_STAT", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, OPENAPI_STAT_info)
				}
				r = r.WithContext(ctx)
				var in OPENAPI_STAT_Request
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

	add_OPENAPI_VERSION := func(method, path string, fn func(context.Context, *OPENAPI_VERSION_Request) (*emptypb.Empty, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*OPENAPI_VERSION_Request))
		}
		var OPENAPI_VERSION_info transport.ServiceInfo
		if h.Interceptor != nil {
			OPENAPI_VERSION_info = transport.NewServiceInfo("erda.openapiv1.openapi.openapi", "OPENAPI_VERSION", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, OPENAPI_VERSION_info)
				}
				r = r.WithContext(ctx)
				var in OPENAPI_VERSION_Request
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

	add_DICE_METADATA("GET", "/metadata.json", srv.DICE_METADATA)
	add_DOC_JSON("GET", "/api/openapi/swagger.json", srv.DOC_JSON)
	add_OPENAPI_DOC("GET", "/api/openapi-doc", srv.OPENAPI_DOC)
	add_OPENAPI_EVENT_DOC("GET", "/api/openapi-event-doc", srv.OPENAPI_EVENT_DOC)
	add_OPENAPI_GEN_CLIENT_TOKEN("POST", "/api/openapi/client-token", srv.OPENAPI_GEN_CLIENT_TOKEN)
	add_OPENAPI_LIST_CLIENT("GET", "/api/openapi/manager/clients", srv.OPENAPI_LIST_CLIENT)
	add_OPENAPI_METRICS("GET", "/api/openapi/metrics", srv.OPENAPI_METRICS)
	add_OPENAPI_NEW_CLIENT("POST", "/api/openapi/manager/clients", srv.OPENAPI_NEW_CLIENT)
	add_OPENAPI_STAT("GET", "/api/openapi/stat", srv.OPENAPI_STAT)
	add_OPENAPI_VERSION("GET", "/api/openapi/version", srv.OPENAPI_VERSION)
}
