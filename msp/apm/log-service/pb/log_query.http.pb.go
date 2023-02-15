// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: log_query.proto

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

// LogServiceHandler is the server API for LogService service.
type LogServiceHandler interface {
	// GET /api/log-service/{addon}/statistics/histogram
	HistogramAggregation(context.Context, *HistogramAggregationRequest) (*HistogramAggregationResponse, error)
	// GET /api/log-service/{addon}/statistics/bucket
	BucketAggregation(context.Context, *BucketAggregationRequest) (*BucketAggregationResponse, error)
	// GET /api/log-service/{addon}/search/paged
	PagedSearch(context.Context, *PagedSearchRequest) (*PagedSearchResponse, error)
	// GET /api/log-service/{addon}/search/sequential
	SequentialSearch(context.Context, *SequentialSearchRequest) (*SequentialSearchResponse, error)
	// GET /api/log-service/{addon}/settings/fields
	GetFieldSettings(context.Context, *GetFieldSettingsRequest) (*GetFieldSettingsResponse, error)
	// GET /api/log-service/action/parse-regexp
	ParseRegexp(context.Context, *ParseRegexpRequest) (*ParseRegexpResponse, error)
}

// RegisterLogServiceHandler register LogServiceHandler to http.Router.
func RegisterLogServiceHandler(r http.Router, srv LogServiceHandler, opts ...http.HandleOption) {
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

	add_HistogramAggregation := func(method, path string, fn func(context.Context, *HistogramAggregationRequest) (*HistogramAggregationResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*HistogramAggregationRequest))
		}
		var HistogramAggregation_info transport.ServiceInfo
		if h.Interceptor != nil {
			HistogramAggregation_info = transport.NewServiceInfo("erda.msp.apm.log_service.LogService", "HistogramAggregation", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, HistogramAggregation_info)
				}
				r = r.WithContext(ctx)
				var in HistogramAggregationRequest
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
						case "addon":
							in.Addon = val
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

	add_BucketAggregation := func(method, path string, fn func(context.Context, *BucketAggregationRequest) (*BucketAggregationResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*BucketAggregationRequest))
		}
		var BucketAggregation_info transport.ServiceInfo
		if h.Interceptor != nil {
			BucketAggregation_info = transport.NewServiceInfo("erda.msp.apm.log_service.LogService", "BucketAggregation", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, BucketAggregation_info)
				}
				r = r.WithContext(ctx)
				var in BucketAggregationRequest
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
						case "addon":
							in.Addon = val
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

	add_PagedSearch := func(method, path string, fn func(context.Context, *PagedSearchRequest) (*PagedSearchResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PagedSearchRequest))
		}
		var PagedSearch_info transport.ServiceInfo
		if h.Interceptor != nil {
			PagedSearch_info = transport.NewServiceInfo("erda.msp.apm.log_service.LogService", "PagedSearch", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, PagedSearch_info)
				}
				r = r.WithContext(ctx)
				var in PagedSearchRequest
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
						case "addon":
							in.Addon = val
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

	add_SequentialSearch := func(method, path string, fn func(context.Context, *SequentialSearchRequest) (*SequentialSearchResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*SequentialSearchRequest))
		}
		var SequentialSearch_info transport.ServiceInfo
		if h.Interceptor != nil {
			SequentialSearch_info = transport.NewServiceInfo("erda.msp.apm.log_service.LogService", "SequentialSearch", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, SequentialSearch_info)
				}
				r = r.WithContext(ctx)
				var in SequentialSearchRequest
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
						case "addon":
							in.Addon = val
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

	add_GetFieldSettings := func(method, path string, fn func(context.Context, *GetFieldSettingsRequest) (*GetFieldSettingsResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetFieldSettingsRequest))
		}
		var GetFieldSettings_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetFieldSettings_info = transport.NewServiceInfo("erda.msp.apm.log_service.LogService", "GetFieldSettings", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetFieldSettings_info)
				}
				r = r.WithContext(ctx)
				var in GetFieldSettingsRequest
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
						case "addon":
							in.Addon = val
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

	add_ParseRegexp := func(method, path string, fn func(context.Context, *ParseRegexpRequest) (*ParseRegexpResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ParseRegexpRequest))
		}
		var ParseRegexp_info transport.ServiceInfo
		if h.Interceptor != nil {
			ParseRegexp_info = transport.NewServiceInfo("erda.msp.apm.log_service.LogService", "ParseRegexp", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ParseRegexp_info)
				}
				r = r.WithContext(ctx)
				var in ParseRegexpRequest
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

	add_HistogramAggregation("GET", "/api/log-service/{addon}/statistics/histogram", srv.HistogramAggregation)
	add_BucketAggregation("GET", "/api/log-service/{addon}/statistics/bucket", srv.BucketAggregation)
	add_PagedSearch("GET", "/api/log-service/{addon}/search/paged", srv.PagedSearch)
	add_SequentialSearch("GET", "/api/log-service/{addon}/search/sequential", srv.SequentialSearch)
	add_GetFieldSettings("GET", "/api/log-service/{addon}/settings/fields", srv.GetFieldSettings)
	add_ParseRegexp("GET", "/api/log-service/action/parse-regexp", srv.ParseRegexp)
}
