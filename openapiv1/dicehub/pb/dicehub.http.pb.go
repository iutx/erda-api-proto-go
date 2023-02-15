// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: dicehub.proto

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

// DicehubHandler is the server API for Dicehub service.
type DicehubHandler interface {
	// GET /api/releases/{releaseId}/actions/download
	DICEHUB_RELEASES_DOWNLOAD(context.Context, *DICEHUB_RELEASES_DOWNLOAD_Request) (*emptypb.Empty, error)
	// GET /api/releases/{releaseID}/actions/get-dice
	DICEHUB_RELEASES_YAML_GET(context.Context, *ReleaseGetDiceYmlRequest) (*emptypb.Empty, error)
	// GET /api/images
	IMAGEHUB_IMAGE_LIST(context.Context, *ReleaseListRequest) (*ReleaseListResponse, error)
	// POST /api/release-rules
	RELEASE_RULE_CREATE(context.Context, *CreateUpdateDeleteReleaseRuleRequest) (*CreateBranchRuleResponse, error)
	// GET /api/release-rules
	RELEASE_RULE_LIST(context.Context, *RELEASE_RULE_LIST_Request) (*emptypb.Empty, error)
	// PUT /api/release-rules/{id}
	RELEASE_RULE_UPDATE(context.Context, *RELEASE_RULE_UPDATE_Request) (*emptypb.Empty, error)
	// DELETE /api/release-rules/{id}
	RELEASE_RULEDelete(context.Context, *RELEASE_RULEDelete_Request) (*emptypb.Empty, error)
}

// RegisterDicehubHandler register DicehubHandler to http.Router.
func RegisterDicehubHandler(r http.Router, srv DicehubHandler, opts ...http.HandleOption) {
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

	add_DICEHUB_RELEASES_DOWNLOAD := func(method, path string, fn func(context.Context, *DICEHUB_RELEASES_DOWNLOAD_Request) (*emptypb.Empty, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*DICEHUB_RELEASES_DOWNLOAD_Request))
		}
		var DICEHUB_RELEASES_DOWNLOAD_info transport.ServiceInfo
		if h.Interceptor != nil {
			DICEHUB_RELEASES_DOWNLOAD_info = transport.NewServiceInfo("erda.openapiv1.dicehub.dicehub", "DICEHUB_RELEASES_DOWNLOAD", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DICEHUB_RELEASES_DOWNLOAD_info)
				}
				r = r.WithContext(ctx)
				var in DICEHUB_RELEASES_DOWNLOAD_Request
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
						case "releaseId":
							in.ReleaseId = val
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

	add_DICEHUB_RELEASES_YAML_GET := func(method, path string, fn func(context.Context, *ReleaseGetDiceYmlRequest) (*emptypb.Empty, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ReleaseGetDiceYmlRequest))
		}
		var DICEHUB_RELEASES_YAML_GET_info transport.ServiceInfo
		if h.Interceptor != nil {
			DICEHUB_RELEASES_YAML_GET_info = transport.NewServiceInfo("erda.openapiv1.dicehub.dicehub", "DICEHUB_RELEASES_YAML_GET", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DICEHUB_RELEASES_YAML_GET_info)
				}
				r = r.WithContext(ctx)
				var in ReleaseGetDiceYmlRequest
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
						case "releaseID":
							in.ReleaseID = val
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

	add_IMAGEHUB_IMAGE_LIST := func(method, path string, fn func(context.Context, *ReleaseListRequest) (*ReleaseListResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ReleaseListRequest))
		}
		var IMAGEHUB_IMAGE_LIST_info transport.ServiceInfo
		if h.Interceptor != nil {
			IMAGEHUB_IMAGE_LIST_info = transport.NewServiceInfo("erda.openapiv1.dicehub.dicehub", "IMAGEHUB_IMAGE_LIST", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, IMAGEHUB_IMAGE_LIST_info)
				}
				r = r.WithContext(ctx)
				var in ReleaseListRequest
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

	add_RELEASE_RULE_CREATE := func(method, path string, fn func(context.Context, *CreateUpdateDeleteReleaseRuleRequest) (*CreateBranchRuleResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CreateUpdateDeleteReleaseRuleRequest))
		}
		var RELEASE_RULE_CREATE_info transport.ServiceInfo
		if h.Interceptor != nil {
			RELEASE_RULE_CREATE_info = transport.NewServiceInfo("erda.openapiv1.dicehub.dicehub", "RELEASE_RULE_CREATE", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, RELEASE_RULE_CREATE_info)
				}
				r = r.WithContext(ctx)
				var in CreateUpdateDeleteReleaseRuleRequest
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

	add_RELEASE_RULE_LIST := func(method, path string, fn func(context.Context, *RELEASE_RULE_LIST_Request) (*emptypb.Empty, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*RELEASE_RULE_LIST_Request))
		}
		var RELEASE_RULE_LIST_info transport.ServiceInfo
		if h.Interceptor != nil {
			RELEASE_RULE_LIST_info = transport.NewServiceInfo("erda.openapiv1.dicehub.dicehub", "RELEASE_RULE_LIST", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, RELEASE_RULE_LIST_info)
				}
				r = r.WithContext(ctx)
				var in RELEASE_RULE_LIST_Request
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

	add_RELEASE_RULE_UPDATE := func(method, path string, fn func(context.Context, *RELEASE_RULE_UPDATE_Request) (*emptypb.Empty, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*RELEASE_RULE_UPDATE_Request))
		}
		var RELEASE_RULE_UPDATE_info transport.ServiceInfo
		if h.Interceptor != nil {
			RELEASE_RULE_UPDATE_info = transport.NewServiceInfo("erda.openapiv1.dicehub.dicehub", "RELEASE_RULE_UPDATE", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, RELEASE_RULE_UPDATE_info)
				}
				r = r.WithContext(ctx)
				var in RELEASE_RULE_UPDATE_Request
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

	add_RELEASE_RULEDelete := func(method, path string, fn func(context.Context, *RELEASE_RULEDelete_Request) (*emptypb.Empty, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*RELEASE_RULEDelete_Request))
		}
		var RELEASE_RULEDelete_info transport.ServiceInfo
		if h.Interceptor != nil {
			RELEASE_RULEDelete_info = transport.NewServiceInfo("erda.openapiv1.dicehub.dicehub", "RELEASE_RULEDelete", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, RELEASE_RULEDelete_info)
				}
				r = r.WithContext(ctx)
				var in RELEASE_RULEDelete_Request
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

	add_DICEHUB_RELEASES_DOWNLOAD("GET", "/api/releases/{releaseId}/actions/download", srv.DICEHUB_RELEASES_DOWNLOAD)
	add_DICEHUB_RELEASES_YAML_GET("GET", "/api/releases/{releaseID}/actions/get-dice", srv.DICEHUB_RELEASES_YAML_GET)
	add_IMAGEHUB_IMAGE_LIST("GET", "/api/images", srv.IMAGEHUB_IMAGE_LIST)
	add_RELEASE_RULE_CREATE("POST", "/api/release-rules", srv.RELEASE_RULE_CREATE)
	add_RELEASE_RULE_LIST("GET", "/api/release-rules", srv.RELEASE_RULE_LIST)
	add_RELEASE_RULE_UPDATE("PUT", "/api/release-rules/{id}", srv.RELEASE_RULE_UPDATE)
	add_RELEASE_RULEDelete("DELETE", "/api/release-rules/{id}", srv.RELEASE_RULEDelete)
}
