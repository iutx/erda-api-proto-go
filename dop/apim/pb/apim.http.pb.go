// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: apim.proto

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

// ExportRecordsHandler is the server API for ExportRecords service.
type ExportRecordsHandler interface {
	// POST /api/apim/export
	CreateExportRecord(context.Context, *CreateExportRecordsReq) (*CreateExportRecordsResp, error)
	// GET /api/apim/export
	ListExportRecords(context.Context, *ListExportRecordsReq) (*ListExportRecordsResp, error)
	// DELETE /api/apim/export/{id}
	DeleteExportRecord(context.Context, *DeleteExportRecordReq) (*Empty, error)
}

// RegisterExportRecordsHandler register ExportRecordsHandler to http.Router.
func RegisterExportRecordsHandler(r http.Router, srv ExportRecordsHandler, opts ...http.HandleOption) {
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

	add_CreateExportRecord := func(method, path string, fn func(context.Context, *CreateExportRecordsReq) (*CreateExportRecordsResp, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CreateExportRecordsReq))
		}
		var CreateExportRecord_info transport.ServiceInfo
		if h.Interceptor != nil {
			CreateExportRecord_info = transport.NewServiceInfo("erda.dop.apim.ExportRecords", "CreateExportRecord", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CreateExportRecord_info)
				}
				r = r.WithContext(ctx)
				var in CreateExportRecordsReq
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

	add_ListExportRecords := func(method, path string, fn func(context.Context, *ListExportRecordsReq) (*ListExportRecordsResp, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ListExportRecordsReq))
		}
		var ListExportRecords_info transport.ServiceInfo
		if h.Interceptor != nil {
			ListExportRecords_info = transport.NewServiceInfo("erda.dop.apim.ExportRecords", "ListExportRecords", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ListExportRecords_info)
				}
				r = r.WithContext(ctx)
				var in ListExportRecordsReq
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

	add_DeleteExportRecord := func(method, path string, fn func(context.Context, *DeleteExportRecordReq) (*Empty, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*DeleteExportRecordReq))
		}
		var DeleteExportRecord_info transport.ServiceInfo
		if h.Interceptor != nil {
			DeleteExportRecord_info = transport.NewServiceInfo("erda.dop.apim.ExportRecords", "DeleteExportRecord", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DeleteExportRecord_info)
				}
				r = r.WithContext(ctx)
				var in DeleteExportRecordReq
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

	add_CreateExportRecord("POST", "/api/apim/export", srv.CreateExportRecord)
	add_ListExportRecords("GET", "/api/apim/export", srv.ListExportRecords)
	add_DeleteExportRecord("DELETE", "/api/apim/export/{id}", srv.DeleteExportRecord)
}
