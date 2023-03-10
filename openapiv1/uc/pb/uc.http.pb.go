// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: uc.proto

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

// UcHandler is the server API for Uc service.
type UcHandler interface {
	// GET /api/users/actions/get-pwd-security-config
	UC_PWD_SECURITY_CONFIG_GET(context.Context, *UC_PWD_SECURITY_CONFIG_GET_Request) (*PwdSecurityConfigGetResponse, error)
	// POST /api/users/actions/update-pwd-security-config
	UC_PWD_SECURITY_CONFIG_UPDATE(context.Context, *PwdSecurityConfigUpdateRequest) (*PwdSecurityConfigUpdateResponse, error)
	// PUT /api/users/actions/batch-freeze
	UC_USER_BATCH_FREEZE(context.Context, *UserBatchFreezeRequest) (*UserBatchFreezeResponse, error)
	// PUT /api/users/actions/batch-unfreeze
	UC_USER_BATCH_UNFREEZE(context.Context, *UserBatchUnFreezeRequest) (*UserBatchUnFreezeResponse, error)
	// POST /api/users/actions/batch-update-login-method
	UC_USER_BATCH_UPDATE_LOGIN_METHOD(context.Context, *UserBatchUpdateLoginMethodRequest) (*UserBatchUpdateLoginMethodResponse, error)
	// POST /api/users
	UC_USER_CREATE(context.Context, *UserCreateRequest) (*UserCreateResponse, error)
	// GET /api/users/actions/export
	UC_USER_EXPORT(context.Context, *UserPagingRequest) (*emptypb.Empty, error)
	// PUT /api/users/{userID}/actions/freeze
	UC_USER_FREEZE(context.Context, *UserFreezeRequest) (*UserFreezeResponse, error)
	// GET /api/users/actions/list-login-method
	UC_USER_LIST_LOGIN_METHOD(context.Context, *UC_USER_LIST_LOGIN_METHOD_Request) (*UserListLoginMethodResponse, error)
	// GET /api/users/actions/paging
	UC_USER_PAGING(context.Context, *UserPagingRequest) (*UserPagingResponse, error)
	// PUT /api/users/{userID}/actions/unfreeze
	UC_USER_UNFREEZE(context.Context, *UserUnfreezeRequest) (*UserUnfreezeResponse, error)
	// POST /api/users/{userID}/actions/update-login-method
	UC_USER_UPDATE_LOGIN_METHOD(context.Context, *UserUpdateLoginMethodRequest) (*UserUpdateLoginMethodResponse, error)
	// PUT /api/user/admin/update-userinfo
	UC_USER_UPDATE_USERINFO(context.Context, *UserUpdateInfoRequset) (*UserUpdateInfoResponse, error)
}

// RegisterUcHandler register UcHandler to http.Router.
func RegisterUcHandler(r http.Router, srv UcHandler, opts ...http.HandleOption) {
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

	add_UC_PWD_SECURITY_CONFIG_GET := func(method, path string, fn func(context.Context, *UC_PWD_SECURITY_CONFIG_GET_Request) (*PwdSecurityConfigGetResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*UC_PWD_SECURITY_CONFIG_GET_Request))
		}
		var UC_PWD_SECURITY_CONFIG_GET_info transport.ServiceInfo
		if h.Interceptor != nil {
			UC_PWD_SECURITY_CONFIG_GET_info = transport.NewServiceInfo("erda.openapiv1.uc.uc", "UC_PWD_SECURITY_CONFIG_GET", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UC_PWD_SECURITY_CONFIG_GET_info)
				}
				r = r.WithContext(ctx)
				var in UC_PWD_SECURITY_CONFIG_GET_Request
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

	add_UC_PWD_SECURITY_CONFIG_UPDATE := func(method, path string, fn func(context.Context, *PwdSecurityConfigUpdateRequest) (*PwdSecurityConfigUpdateResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PwdSecurityConfigUpdateRequest))
		}
		var UC_PWD_SECURITY_CONFIG_UPDATE_info transport.ServiceInfo
		if h.Interceptor != nil {
			UC_PWD_SECURITY_CONFIG_UPDATE_info = transport.NewServiceInfo("erda.openapiv1.uc.uc", "UC_PWD_SECURITY_CONFIG_UPDATE", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UC_PWD_SECURITY_CONFIG_UPDATE_info)
				}
				r = r.WithContext(ctx)
				var in PwdSecurityConfigUpdateRequest
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

	add_UC_USER_BATCH_FREEZE := func(method, path string, fn func(context.Context, *UserBatchFreezeRequest) (*UserBatchFreezeResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*UserBatchFreezeRequest))
		}
		var UC_USER_BATCH_FREEZE_info transport.ServiceInfo
		if h.Interceptor != nil {
			UC_USER_BATCH_FREEZE_info = transport.NewServiceInfo("erda.openapiv1.uc.uc", "UC_USER_BATCH_FREEZE", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UC_USER_BATCH_FREEZE_info)
				}
				r = r.WithContext(ctx)
				var in UserBatchFreezeRequest
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

	add_UC_USER_BATCH_UNFREEZE := func(method, path string, fn func(context.Context, *UserBatchUnFreezeRequest) (*UserBatchUnFreezeResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*UserBatchUnFreezeRequest))
		}
		var UC_USER_BATCH_UNFREEZE_info transport.ServiceInfo
		if h.Interceptor != nil {
			UC_USER_BATCH_UNFREEZE_info = transport.NewServiceInfo("erda.openapiv1.uc.uc", "UC_USER_BATCH_UNFREEZE", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UC_USER_BATCH_UNFREEZE_info)
				}
				r = r.WithContext(ctx)
				var in UserBatchUnFreezeRequest
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

	add_UC_USER_BATCH_UPDATE_LOGIN_METHOD := func(method, path string, fn func(context.Context, *UserBatchUpdateLoginMethodRequest) (*UserBatchUpdateLoginMethodResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*UserBatchUpdateLoginMethodRequest))
		}
		var UC_USER_BATCH_UPDATE_LOGIN_METHOD_info transport.ServiceInfo
		if h.Interceptor != nil {
			UC_USER_BATCH_UPDATE_LOGIN_METHOD_info = transport.NewServiceInfo("erda.openapiv1.uc.uc", "UC_USER_BATCH_UPDATE_LOGIN_METHOD", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UC_USER_BATCH_UPDATE_LOGIN_METHOD_info)
				}
				r = r.WithContext(ctx)
				var in UserBatchUpdateLoginMethodRequest
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

	add_UC_USER_CREATE := func(method, path string, fn func(context.Context, *UserCreateRequest) (*UserCreateResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*UserCreateRequest))
		}
		var UC_USER_CREATE_info transport.ServiceInfo
		if h.Interceptor != nil {
			UC_USER_CREATE_info = transport.NewServiceInfo("erda.openapiv1.uc.uc", "UC_USER_CREATE", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UC_USER_CREATE_info)
				}
				r = r.WithContext(ctx)
				var in UserCreateRequest
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

	add_UC_USER_EXPORT := func(method, path string, fn func(context.Context, *UserPagingRequest) (*emptypb.Empty, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*UserPagingRequest))
		}
		var UC_USER_EXPORT_info transport.ServiceInfo
		if h.Interceptor != nil {
			UC_USER_EXPORT_info = transport.NewServiceInfo("erda.openapiv1.uc.uc", "UC_USER_EXPORT", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UC_USER_EXPORT_info)
				}
				r = r.WithContext(ctx)
				var in UserPagingRequest
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

	add_UC_USER_FREEZE := func(method, path string, fn func(context.Context, *UserFreezeRequest) (*UserFreezeResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*UserFreezeRequest))
		}
		var UC_USER_FREEZE_info transport.ServiceInfo
		if h.Interceptor != nil {
			UC_USER_FREEZE_info = transport.NewServiceInfo("erda.openapiv1.uc.uc", "UC_USER_FREEZE", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UC_USER_FREEZE_info)
				}
				r = r.WithContext(ctx)
				var in UserFreezeRequest
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
						case "userID":
							in.UserID = val
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

	add_UC_USER_LIST_LOGIN_METHOD := func(method, path string, fn func(context.Context, *UC_USER_LIST_LOGIN_METHOD_Request) (*UserListLoginMethodResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*UC_USER_LIST_LOGIN_METHOD_Request))
		}
		var UC_USER_LIST_LOGIN_METHOD_info transport.ServiceInfo
		if h.Interceptor != nil {
			UC_USER_LIST_LOGIN_METHOD_info = transport.NewServiceInfo("erda.openapiv1.uc.uc", "UC_USER_LIST_LOGIN_METHOD", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UC_USER_LIST_LOGIN_METHOD_info)
				}
				r = r.WithContext(ctx)
				var in UC_USER_LIST_LOGIN_METHOD_Request
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

	add_UC_USER_PAGING := func(method, path string, fn func(context.Context, *UserPagingRequest) (*UserPagingResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*UserPagingRequest))
		}
		var UC_USER_PAGING_info transport.ServiceInfo
		if h.Interceptor != nil {
			UC_USER_PAGING_info = transport.NewServiceInfo("erda.openapiv1.uc.uc", "UC_USER_PAGING", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UC_USER_PAGING_info)
				}
				r = r.WithContext(ctx)
				var in UserPagingRequest
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

	add_UC_USER_UNFREEZE := func(method, path string, fn func(context.Context, *UserUnfreezeRequest) (*UserUnfreezeResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*UserUnfreezeRequest))
		}
		var UC_USER_UNFREEZE_info transport.ServiceInfo
		if h.Interceptor != nil {
			UC_USER_UNFREEZE_info = transport.NewServiceInfo("erda.openapiv1.uc.uc", "UC_USER_UNFREEZE", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UC_USER_UNFREEZE_info)
				}
				r = r.WithContext(ctx)
				var in UserUnfreezeRequest
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
						case "userID":
							in.UserID = val
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

	add_UC_USER_UPDATE_LOGIN_METHOD := func(method, path string, fn func(context.Context, *UserUpdateLoginMethodRequest) (*UserUpdateLoginMethodResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*UserUpdateLoginMethodRequest))
		}
		var UC_USER_UPDATE_LOGIN_METHOD_info transport.ServiceInfo
		if h.Interceptor != nil {
			UC_USER_UPDATE_LOGIN_METHOD_info = transport.NewServiceInfo("erda.openapiv1.uc.uc", "UC_USER_UPDATE_LOGIN_METHOD", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UC_USER_UPDATE_LOGIN_METHOD_info)
				}
				r = r.WithContext(ctx)
				var in UserUpdateLoginMethodRequest
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
						case "userID":
							in.UserID = val
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

	add_UC_USER_UPDATE_USERINFO := func(method, path string, fn func(context.Context, *UserUpdateInfoRequset) (*UserUpdateInfoResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*UserUpdateInfoRequset))
		}
		var UC_USER_UPDATE_USERINFO_info transport.ServiceInfo
		if h.Interceptor != nil {
			UC_USER_UPDATE_USERINFO_info = transport.NewServiceInfo("erda.openapiv1.uc.uc", "UC_USER_UPDATE_USERINFO", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UC_USER_UPDATE_USERINFO_info)
				}
				r = r.WithContext(ctx)
				var in UserUpdateInfoRequset
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

	add_UC_PWD_SECURITY_CONFIG_GET("GET", "/api/users/actions/get-pwd-security-config", srv.UC_PWD_SECURITY_CONFIG_GET)
	add_UC_PWD_SECURITY_CONFIG_UPDATE("POST", "/api/users/actions/update-pwd-security-config", srv.UC_PWD_SECURITY_CONFIG_UPDATE)
	add_UC_USER_BATCH_FREEZE("PUT", "/api/users/actions/batch-freeze", srv.UC_USER_BATCH_FREEZE)
	add_UC_USER_BATCH_UNFREEZE("PUT", "/api/users/actions/batch-unfreeze", srv.UC_USER_BATCH_UNFREEZE)
	add_UC_USER_BATCH_UPDATE_LOGIN_METHOD("POST", "/api/users/actions/batch-update-login-method", srv.UC_USER_BATCH_UPDATE_LOGIN_METHOD)
	add_UC_USER_CREATE("POST", "/api/users", srv.UC_USER_CREATE)
	add_UC_USER_EXPORT("GET", "/api/users/actions/export", srv.UC_USER_EXPORT)
	add_UC_USER_FREEZE("PUT", "/api/users/{userID}/actions/freeze", srv.UC_USER_FREEZE)
	add_UC_USER_LIST_LOGIN_METHOD("GET", "/api/users/actions/list-login-method", srv.UC_USER_LIST_LOGIN_METHOD)
	add_UC_USER_PAGING("GET", "/api/users/actions/paging", srv.UC_USER_PAGING)
	add_UC_USER_UNFREEZE("PUT", "/api/users/{userID}/actions/unfreeze", srv.UC_USER_UNFREEZE)
	add_UC_USER_UPDATE_LOGIN_METHOD("POST", "/api/users/{userID}/actions/update-login-method", srv.UC_USER_UPDATE_LOGIN_METHOD)
	add_UC_USER_UPDATE_USERINFO("PUT", "/api/user/admin/update-userinfo", srv.UC_USER_UPDATE_USERINFO)
}
