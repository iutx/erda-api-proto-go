// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: guide.proto

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
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// GuideServiceHandler is the server API for GuideService service.
type GuideServiceHandler interface {
	// POST /api/guide/actions/create-by-gittar-push-hook
	CreateGuideByGittarPushHook(context.Context, *GittarPushPayloadEvent) (*CreateGuideResponse, error)
	// GET /api/guide
	ListGuide(context.Context, *ListGuideRequest) (*ListGuideResponse, error)
	// POST /api/guide/actions/process
	ProcessGuide(context.Context, *ProcessGuideRequest) (*ProcessGuideResponse, error)
	// POST /api/guide/actions/delete-by-gittar-push-hook
	DeleteGuideByGittarPushHook(context.Context, *GittarPushPayloadEvent) (*DeleteGuideResponse, error)
	// POST /api/guide/{ID}/actions/cancel
	CancelGuide(context.Context, *CancelGuideRequest) (*CancelGuideResponse, error)
}

// RegisterGuideServiceHandler register GuideServiceHandler to http.Router.
func RegisterGuideServiceHandler(r http.Router, srv GuideServiceHandler, opts ...http.HandleOption) {
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

	add_CreateGuideByGittarPushHook := func(method, path string, fn func(context.Context, *GittarPushPayloadEvent) (*CreateGuideResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GittarPushPayloadEvent))
		}
		var CreateGuideByGittarPushHook_info transport.ServiceInfo
		if h.Interceptor != nil {
			CreateGuideByGittarPushHook_info = transport.NewServiceInfo("erda.dop.guide.GuideService", "CreateGuideByGittarPushHook", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CreateGuideByGittarPushHook_info)
				}
				r = r.WithContext(ctx)
				var in GittarPushPayloadEvent
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

	add_ListGuide := func(method, path string, fn func(context.Context, *ListGuideRequest) (*ListGuideResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ListGuideRequest))
		}
		var ListGuide_info transport.ServiceInfo
		if h.Interceptor != nil {
			ListGuide_info = transport.NewServiceInfo("erda.dop.guide.GuideService", "ListGuide", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ListGuide_info)
				}
				r = r.WithContext(ctx)
				var in ListGuideRequest
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

	add_ProcessGuide := func(method, path string, fn func(context.Context, *ProcessGuideRequest) (*ProcessGuideResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ProcessGuideRequest))
		}
		var ProcessGuide_info transport.ServiceInfo
		if h.Interceptor != nil {
			ProcessGuide_info = transport.NewServiceInfo("erda.dop.guide.GuideService", "ProcessGuide", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ProcessGuide_info)
				}
				r = r.WithContext(ctx)
				var in ProcessGuideRequest
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

	add_DeleteGuideByGittarPushHook := func(method, path string, fn func(context.Context, *GittarPushPayloadEvent) (*DeleteGuideResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GittarPushPayloadEvent))
		}
		var DeleteGuideByGittarPushHook_info transport.ServiceInfo
		if h.Interceptor != nil {
			DeleteGuideByGittarPushHook_info = transport.NewServiceInfo("erda.dop.guide.GuideService", "DeleteGuideByGittarPushHook", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DeleteGuideByGittarPushHook_info)
				}
				r = r.WithContext(ctx)
				var in GittarPushPayloadEvent
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

	add_CancelGuide := func(method, path string, fn func(context.Context, *CancelGuideRequest) (*CancelGuideResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CancelGuideRequest))
		}
		var CancelGuide_info transport.ServiceInfo
		if h.Interceptor != nil {
			CancelGuide_info = transport.NewServiceInfo("erda.dop.guide.GuideService", "CancelGuide", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CancelGuide_info)
				}
				r = r.WithContext(ctx)
				var in CancelGuideRequest
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
						case "ID":
							in.ID = val
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

	add_CreateGuideByGittarPushHook("POST", "/api/guide/actions/create-by-gittar-push-hook", srv.CreateGuideByGittarPushHook)
	add_ListGuide("GET", "/api/guide", srv.ListGuide)
	add_ProcessGuide("POST", "/api/guide/actions/process", srv.ProcessGuide)
	add_DeleteGuideByGittarPushHook("POST", "/api/guide/actions/delete-by-gittar-push-hook", srv.DeleteGuideByGittarPushHook)
	add_CancelGuide("POST", "/api/guide/{ID}/actions/cancel", srv.CancelGuide)
}
