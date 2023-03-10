// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: stream.proto

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

// CommentIssueStreamServiceHandler is the server API for CommentIssueStreamService service.
type CommentIssueStreamServiceHandler interface {
	// POST /api/issues/actions/batch-create-comment-stream
	BatchCreateIssueStream(context.Context, *CommentIssueStreamBatchCreateRequest) (*CommentIssueStreamBatchCreateResponse, error)
	// POST /api/issues/{id}/streams
	CreateIssueStream(context.Context, *IssueStreamCreateRequest) (*IssueStreamCreateResponse, error)
	// GET /api/issues/{id}/streams
	PagingIssueStreams(context.Context, *PagingIssueStreamsRequest) (*PagingIssueStreamsResponse, error)
}

// RegisterCommentIssueStreamServiceHandler register CommentIssueStreamServiceHandler to http.Router.
func RegisterCommentIssueStreamServiceHandler(r http.Router, srv CommentIssueStreamServiceHandler, opts ...http.HandleOption) {
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

	add_BatchCreateIssueStream := func(method, path string, fn func(context.Context, *CommentIssueStreamBatchCreateRequest) (*CommentIssueStreamBatchCreateResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CommentIssueStreamBatchCreateRequest))
		}
		var BatchCreateIssueStream_info transport.ServiceInfo
		if h.Interceptor != nil {
			BatchCreateIssueStream_info = transport.NewServiceInfo("erda.dop.issue.stream.CommentIssueStreamService", "BatchCreateIssueStream", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, BatchCreateIssueStream_info)
				}
				r = r.WithContext(ctx)
				var in CommentIssueStreamBatchCreateRequest
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

	add_CreateIssueStream := func(method, path string, fn func(context.Context, *IssueStreamCreateRequest) (*IssueStreamCreateResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*IssueStreamCreateRequest))
		}
		var CreateIssueStream_info transport.ServiceInfo
		if h.Interceptor != nil {
			CreateIssueStream_info = transport.NewServiceInfo("erda.dop.issue.stream.CommentIssueStreamService", "CreateIssueStream", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CreateIssueStream_info)
				}
				r = r.WithContext(ctx)
				var in IssueStreamCreateRequest
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

	add_PagingIssueStreams := func(method, path string, fn func(context.Context, *PagingIssueStreamsRequest) (*PagingIssueStreamsResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PagingIssueStreamsRequest))
		}
		var PagingIssueStreams_info transport.ServiceInfo
		if h.Interceptor != nil {
			PagingIssueStreams_info = transport.NewServiceInfo("erda.dop.issue.stream.CommentIssueStreamService", "PagingIssueStreams", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, PagingIssueStreams_info)
				}
				r = r.WithContext(ctx)
				var in PagingIssueStreamsRequest
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

	add_BatchCreateIssueStream("POST", "/api/issues/actions/batch-create-comment-stream", srv.BatchCreateIssueStream)
	add_CreateIssueStream("POST", "/api/issues/{id}/streams", srv.CreateIssueStream)
	add_PagingIssueStreams("GET", "/api/issues/{id}/streams", srv.PagingIssueStreams)
}
