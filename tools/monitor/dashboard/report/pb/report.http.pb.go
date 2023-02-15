// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: report.proto

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
	pb "github.com/erda-project/erda-proto-go/common/pb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// ReportServiceHandler is the server API for ReportService service.
type ReportServiceHandler interface {
	// GET /apis/report/v1/tasks
	ListTasks(context.Context, *ListTasksRequest) (*ListTasksResponse, error)
	// POST /apis/report/v1/tasks
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	// PUT /apis/report/v1/tasks/{id}
	UpdateTask(context.Context, *UpdateTaskRequest) (*pb.VoidResponse, error)
	// PUT /apis/report/v1/tasks/{id}/switch
	SwitchTask(context.Context, *SwitchTaskRequest) (*pb.VoidResponse, error)
	// GET /apis/report/v1/tasks/{id}
	GetTask(context.Context, *GetTaskRequest) (*ReportTaskDTO, error)
	// DELETE /apis/report/v1/tasks/{id}
	DeleteTask(context.Context, *DeleteTaskRequest) (*pb.VoidResponse, error)
	// GET /apis/report/v1/types
	ListTypes(context.Context, *ListTypesRequest) (*ListTypesResponse, error)
	// GET /apis/report/v1/histories
	ListHistories(context.Context, *ListHistoriesRequest) (*ListHistoriesResponse, error)
	// POST /apis/report/v1/histories
	CreateHistory(context.Context, *CreateHistoryRequest) (*CreateHistoryResponse, error)
	// GET /apis/report/v1/histories/{id}
	GetHistory(context.Context, *GetHistoryRequest) (*GetHistoryResponse, error)
	// DELETE /apis/report/v1/histories/{id}
	DeleteHistory(context.Context, *DeleteHistoryRequest) (*pb.VoidResponse, error)
}

// RegisterReportServiceHandler register ReportServiceHandler to http.Router.
func RegisterReportServiceHandler(r http.Router, srv ReportServiceHandler, opts ...http.HandleOption) {
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

	add_ListTasks := func(method, path string, fn func(context.Context, *ListTasksRequest) (*ListTasksResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ListTasksRequest))
		}
		var ListTasks_info transport.ServiceInfo
		if h.Interceptor != nil {
			ListTasks_info = transport.NewServiceInfo("erda.tools.monitor.dashboard.report.ReportService", "ListTasks", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ListTasks_info)
				}
				r = r.WithContext(ctx)
				var in ListTasksRequest
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

	add_CreateTask := func(method, path string, fn func(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CreateTaskRequest))
		}
		var CreateTask_info transport.ServiceInfo
		if h.Interceptor != nil {
			CreateTask_info = transport.NewServiceInfo("erda.tools.monitor.dashboard.report.ReportService", "CreateTask", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CreateTask_info)
				}
				r = r.WithContext(ctx)
				var in CreateTaskRequest
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

	add_UpdateTask := func(method, path string, fn func(context.Context, *UpdateTaskRequest) (*pb.VoidResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*UpdateTaskRequest))
		}
		var UpdateTask_info transport.ServiceInfo
		if h.Interceptor != nil {
			UpdateTask_info = transport.NewServiceInfo("erda.tools.monitor.dashboard.report.ReportService", "UpdateTask", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UpdateTask_info)
				}
				r = r.WithContext(ctx)
				var in UpdateTaskRequest
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
							val, err := strconv.ParseUint(val, 10, 64)
							if err != nil {
								return nil, err
							}
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

	add_SwitchTask := func(method, path string, fn func(context.Context, *SwitchTaskRequest) (*pb.VoidResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*SwitchTaskRequest))
		}
		var SwitchTask_info transport.ServiceInfo
		if h.Interceptor != nil {
			SwitchTask_info = transport.NewServiceInfo("erda.tools.monitor.dashboard.report.ReportService", "SwitchTask", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, SwitchTask_info)
				}
				r = r.WithContext(ctx)
				var in SwitchTaskRequest
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
							val, err := strconv.ParseUint(val, 10, 64)
							if err != nil {
								return nil, err
							}
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

	add_GetTask := func(method, path string, fn func(context.Context, *GetTaskRequest) (*ReportTaskDTO, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetTaskRequest))
		}
		var GetTask_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetTask_info = transport.NewServiceInfo("erda.tools.monitor.dashboard.report.ReportService", "GetTask", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetTask_info)
				}
				r = r.WithContext(ctx)
				var in GetTaskRequest
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
							val, err := strconv.ParseUint(val, 10, 64)
							if err != nil {
								return nil, err
							}
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

	add_DeleteTask := func(method, path string, fn func(context.Context, *DeleteTaskRequest) (*pb.VoidResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*DeleteTaskRequest))
		}
		var DeleteTask_info transport.ServiceInfo
		if h.Interceptor != nil {
			DeleteTask_info = transport.NewServiceInfo("erda.tools.monitor.dashboard.report.ReportService", "DeleteTask", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DeleteTask_info)
				}
				r = r.WithContext(ctx)
				var in DeleteTaskRequest
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
							val, err := strconv.ParseUint(val, 10, 64)
							if err != nil {
								return nil, err
							}
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

	add_ListTypes := func(method, path string, fn func(context.Context, *ListTypesRequest) (*ListTypesResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ListTypesRequest))
		}
		var ListTypes_info transport.ServiceInfo
		if h.Interceptor != nil {
			ListTypes_info = transport.NewServiceInfo("erda.tools.monitor.dashboard.report.ReportService", "ListTypes", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ListTypes_info)
				}
				r = r.WithContext(ctx)
				var in ListTypesRequest
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

	add_ListHistories := func(method, path string, fn func(context.Context, *ListHistoriesRequest) (*ListHistoriesResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ListHistoriesRequest))
		}
		var ListHistories_info transport.ServiceInfo
		if h.Interceptor != nil {
			ListHistories_info = transport.NewServiceInfo("erda.tools.monitor.dashboard.report.ReportService", "ListHistories", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ListHistories_info)
				}
				r = r.WithContext(ctx)
				var in ListHistoriesRequest
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

	add_CreateHistory := func(method, path string, fn func(context.Context, *CreateHistoryRequest) (*CreateHistoryResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CreateHistoryRequest))
		}
		var CreateHistory_info transport.ServiceInfo
		if h.Interceptor != nil {
			CreateHistory_info = transport.NewServiceInfo("erda.tools.monitor.dashboard.report.ReportService", "CreateHistory", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CreateHistory_info)
				}
				r = r.WithContext(ctx)
				var in CreateHistoryRequest
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

	add_GetHistory := func(method, path string, fn func(context.Context, *GetHistoryRequest) (*GetHistoryResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetHistoryRequest))
		}
		var GetHistory_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetHistory_info = transport.NewServiceInfo("erda.tools.monitor.dashboard.report.ReportService", "GetHistory", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetHistory_info)
				}
				r = r.WithContext(ctx)
				var in GetHistoryRequest
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
							val, err := strconv.ParseUint(val, 10, 64)
							if err != nil {
								return nil, err
							}
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

	add_DeleteHistory := func(method, path string, fn func(context.Context, *DeleteHistoryRequest) (*pb.VoidResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*DeleteHistoryRequest))
		}
		var DeleteHistory_info transport.ServiceInfo
		if h.Interceptor != nil {
			DeleteHistory_info = transport.NewServiceInfo("erda.tools.monitor.dashboard.report.ReportService", "DeleteHistory", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DeleteHistory_info)
				}
				r = r.WithContext(ctx)
				var in DeleteHistoryRequest
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
							val, err := strconv.ParseUint(val, 10, 64)
							if err != nil {
								return nil, err
							}
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

	add_ListTasks("GET", "/apis/report/v1/tasks", srv.ListTasks)
	add_CreateTask("POST", "/apis/report/v1/tasks", srv.CreateTask)
	add_UpdateTask("PUT", "/apis/report/v1/tasks/{id}", srv.UpdateTask)
	add_SwitchTask("PUT", "/apis/report/v1/tasks/{id}/switch", srv.SwitchTask)
	add_GetTask("GET", "/apis/report/v1/tasks/{id}", srv.GetTask)
	add_DeleteTask("DELETE", "/apis/report/v1/tasks/{id}", srv.DeleteTask)
	add_ListTypes("GET", "/apis/report/v1/types", srv.ListTypes)
	add_ListHistories("GET", "/apis/report/v1/histories", srv.ListHistories)
	add_CreateHistory("POST", "/apis/report/v1/histories", srv.CreateHistory)
	add_GetHistory("GET", "/apis/report/v1/histories/{id}", srv.GetHistory)
	add_DeleteHistory("DELETE", "/apis/report/v1/histories/{id}", srv.DeleteHistory)
}