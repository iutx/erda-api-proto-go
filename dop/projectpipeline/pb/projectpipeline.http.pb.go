// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: projectpipeline.proto

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

// ProjectPipelineServiceHandler is the server API for ProjectPipelineService service.
type ProjectPipelineServiceHandler interface {
	// POST /api/project-pipeline
	Create(context.Context, *CreateProjectPipelineRequest) (*CreateProjectPipelineResponse, error)
	// GET /api/project-pipeline/actions/get-my-apps
	ListApp(context.Context, *ListAppRequest) (*ListAppResponse, error)
	// GET /api/project-pipeline/actions/get-pipeline-yml-list
	ListPipelineYml(context.Context, *ListAppPipelineYmlRequest) (*ListAppPipelineYmlResponse, error)
	// GET /api/project-pipeline/actions/name-pre-check
	CreateNamePreCheck(context.Context, *CreateProjectPipelineNamePreCheckRequest) (*CreateProjectPipelineNamePreCheckResponse, error)
	// GET /api/project-pipeline/actions/source-pre-check
	CreateSourcePreCheck(context.Context, *CreateProjectPipelineSourcePreCheckRequest) (*CreateProjectPipelineSourcePreCheckResponse, error)
	// GET /api/project-pipeline/actions/list-category
	ListPipelineCategory(context.Context, *ListPipelineCategoryRequest) (*ListPipelineCategoryResponse, error)
	// PUT /api/project-pipeline/definitions/{pipelineDefinitionID}
	Update(context.Context, *UpdateProjectPipelineRequest) (*UpdateProjectPipelineResponse, error)
	// POST /api/project-pipeline/definitions/{pipelineDefinitionID}/actions/run
	Run(context.Context, *RunProjectPipelineRequest) (*RunProjectPipelineResponse, error)
	// POST /api/project-pipeline/definitions/{pipelineDefinitionID}/actions/rerun
	Rerun(context.Context, *RerunProjectPipelineRequest) (*RerunProjectPipelineResponse, error)
	// POST /api/project-pipeline/definitions/{pipelineDefinitionID}/actions/rerun-failed
	RerunFailed(context.Context, *RerunFailedProjectPipelineRequest) (*RerunFailedProjectPipelineResponse, error)
	// POST /api/project-pipeline/definitions/{pipelineDefinitionID}/actions/cancel
	Cancel(context.Context, *CancelProjectPipelineRequest) (*CancelProjectPipelineResponse, error)
	// POST /api/project-pipeline/actions/one-click-create
	OneClickCreate(context.Context, *OneClickCreateProjectPipelineRequest) (*OneClickCreateProjectPipelineResponse, error)
	// POST /api/project-pipeline/actions/create-by-gittar-push-hook
	BatchCreateByGittarPushHook(context.Context, *GittarPushPayloadEvent) (*BatchCreateProjectPipelineResponse, error)
	// GET /api/project-pipeline/actions/exec-history
	ListExecHistory(context.Context, *ListPipelineExecHistoryRequest) (*ListPipelineExecHistoryResponse, error)
	// DELETE /api/project-pipeline/actions/delete-by-app
	DeleteByApp(context.Context, *DeleteByAppRequest) (*DeleteByAppResponse, error)
}

// RegisterProjectPipelineServiceHandler register ProjectPipelineServiceHandler to http.Router.
func RegisterProjectPipelineServiceHandler(r http.Router, srv ProjectPipelineServiceHandler, opts ...http.HandleOption) {
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

	add_Create := func(method, path string, fn func(context.Context, *CreateProjectPipelineRequest) (*CreateProjectPipelineResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CreateProjectPipelineRequest))
		}
		var Create_info transport.ServiceInfo
		if h.Interceptor != nil {
			Create_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "Create", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, Create_info)
				}
				r = r.WithContext(ctx)
				var in CreateProjectPipelineRequest
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

	add_ListApp := func(method, path string, fn func(context.Context, *ListAppRequest) (*ListAppResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ListAppRequest))
		}
		var ListApp_info transport.ServiceInfo
		if h.Interceptor != nil {
			ListApp_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "ListApp", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ListApp_info)
				}
				r = r.WithContext(ctx)
				var in ListAppRequest
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

	add_ListPipelineYml := func(method, path string, fn func(context.Context, *ListAppPipelineYmlRequest) (*ListAppPipelineYmlResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ListAppPipelineYmlRequest))
		}
		var ListPipelineYml_info transport.ServiceInfo
		if h.Interceptor != nil {
			ListPipelineYml_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "ListPipelineYml", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ListPipelineYml_info)
				}
				r = r.WithContext(ctx)
				var in ListAppPipelineYmlRequest
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

	add_CreateNamePreCheck := func(method, path string, fn func(context.Context, *CreateProjectPipelineNamePreCheckRequest) (*CreateProjectPipelineNamePreCheckResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CreateProjectPipelineNamePreCheckRequest))
		}
		var CreateNamePreCheck_info transport.ServiceInfo
		if h.Interceptor != nil {
			CreateNamePreCheck_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "CreateNamePreCheck", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CreateNamePreCheck_info)
				}
				r = r.WithContext(ctx)
				var in CreateProjectPipelineNamePreCheckRequest
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

	add_CreateSourcePreCheck := func(method, path string, fn func(context.Context, *CreateProjectPipelineSourcePreCheckRequest) (*CreateProjectPipelineSourcePreCheckResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CreateProjectPipelineSourcePreCheckRequest))
		}
		var CreateSourcePreCheck_info transport.ServiceInfo
		if h.Interceptor != nil {
			CreateSourcePreCheck_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "CreateSourcePreCheck", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CreateSourcePreCheck_info)
				}
				r = r.WithContext(ctx)
				var in CreateProjectPipelineSourcePreCheckRequest
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

	add_ListPipelineCategory := func(method, path string, fn func(context.Context, *ListPipelineCategoryRequest) (*ListPipelineCategoryResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ListPipelineCategoryRequest))
		}
		var ListPipelineCategory_info transport.ServiceInfo
		if h.Interceptor != nil {
			ListPipelineCategory_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "ListPipelineCategory", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ListPipelineCategory_info)
				}
				r = r.WithContext(ctx)
				var in ListPipelineCategoryRequest
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

	add_Update := func(method, path string, fn func(context.Context, *UpdateProjectPipelineRequest) (*UpdateProjectPipelineResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*UpdateProjectPipelineRequest))
		}
		var Update_info transport.ServiceInfo
		if h.Interceptor != nil {
			Update_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "Update", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, Update_info)
				}
				r = r.WithContext(ctx)
				var in UpdateProjectPipelineRequest
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
						case "pipelineDefinitionID":
							in.PipelineDefinitionID = val
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

	add_Run := func(method, path string, fn func(context.Context, *RunProjectPipelineRequest) (*RunProjectPipelineResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*RunProjectPipelineRequest))
		}
		var Run_info transport.ServiceInfo
		if h.Interceptor != nil {
			Run_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "Run", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, Run_info)
				}
				r = r.WithContext(ctx)
				var in RunProjectPipelineRequest
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
						case "pipelineDefinitionID":
							in.PipelineDefinitionID = val
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

	add_Rerun := func(method, path string, fn func(context.Context, *RerunProjectPipelineRequest) (*RerunProjectPipelineResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*RerunProjectPipelineRequest))
		}
		var Rerun_info transport.ServiceInfo
		if h.Interceptor != nil {
			Rerun_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "Rerun", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, Rerun_info)
				}
				r = r.WithContext(ctx)
				var in RerunProjectPipelineRequest
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
						case "pipelineDefinitionID":
							in.PipelineDefinitionID = val
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

	add_RerunFailed := func(method, path string, fn func(context.Context, *RerunFailedProjectPipelineRequest) (*RerunFailedProjectPipelineResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*RerunFailedProjectPipelineRequest))
		}
		var RerunFailed_info transport.ServiceInfo
		if h.Interceptor != nil {
			RerunFailed_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "RerunFailed", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, RerunFailed_info)
				}
				r = r.WithContext(ctx)
				var in RerunFailedProjectPipelineRequest
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
						case "pipelineDefinitionID":
							in.PipelineDefinitionID = val
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

	add_Cancel := func(method, path string, fn func(context.Context, *CancelProjectPipelineRequest) (*CancelProjectPipelineResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CancelProjectPipelineRequest))
		}
		var Cancel_info transport.ServiceInfo
		if h.Interceptor != nil {
			Cancel_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "Cancel", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, Cancel_info)
				}
				r = r.WithContext(ctx)
				var in CancelProjectPipelineRequest
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
						case "pipelineDefinitionID":
							in.PipelineDefinitionID = val
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

	add_OneClickCreate := func(method, path string, fn func(context.Context, *OneClickCreateProjectPipelineRequest) (*OneClickCreateProjectPipelineResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*OneClickCreateProjectPipelineRequest))
		}
		var OneClickCreate_info transport.ServiceInfo
		if h.Interceptor != nil {
			OneClickCreate_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "OneClickCreate", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, OneClickCreate_info)
				}
				r = r.WithContext(ctx)
				var in OneClickCreateProjectPipelineRequest
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

	add_BatchCreateByGittarPushHook := func(method, path string, fn func(context.Context, *GittarPushPayloadEvent) (*BatchCreateProjectPipelineResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GittarPushPayloadEvent))
		}
		var BatchCreateByGittarPushHook_info transport.ServiceInfo
		if h.Interceptor != nil {
			BatchCreateByGittarPushHook_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "BatchCreateByGittarPushHook", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, BatchCreateByGittarPushHook_info)
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

	add_ListExecHistory := func(method, path string, fn func(context.Context, *ListPipelineExecHistoryRequest) (*ListPipelineExecHistoryResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ListPipelineExecHistoryRequest))
		}
		var ListExecHistory_info transport.ServiceInfo
		if h.Interceptor != nil {
			ListExecHistory_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "ListExecHistory", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ListExecHistory_info)
				}
				r = r.WithContext(ctx)
				var in ListPipelineExecHistoryRequest
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

	add_DeleteByApp := func(method, path string, fn func(context.Context, *DeleteByAppRequest) (*DeleteByAppResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*DeleteByAppRequest))
		}
		var DeleteByApp_info transport.ServiceInfo
		if h.Interceptor != nil {
			DeleteByApp_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "DeleteByApp", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DeleteByApp_info)
				}
				r = r.WithContext(ctx)
				var in DeleteByAppRequest
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

	add_Create("POST", "/api/project-pipeline", srv.Create)
	add_ListApp("GET", "/api/project-pipeline/actions/get-my-apps", srv.ListApp)
	add_ListPipelineYml("GET", "/api/project-pipeline/actions/get-pipeline-yml-list", srv.ListPipelineYml)
	add_CreateNamePreCheck("GET", "/api/project-pipeline/actions/name-pre-check", srv.CreateNamePreCheck)
	add_CreateSourcePreCheck("GET", "/api/project-pipeline/actions/source-pre-check", srv.CreateSourcePreCheck)
	add_ListPipelineCategory("GET", "/api/project-pipeline/actions/list-category", srv.ListPipelineCategory)
	add_Update("PUT", "/api/project-pipeline/definitions/{pipelineDefinitionID}", srv.Update)
	add_Run("POST", "/api/project-pipeline/definitions/{pipelineDefinitionID}/actions/run", srv.Run)
	add_Rerun("POST", "/api/project-pipeline/definitions/{pipelineDefinitionID}/actions/rerun", srv.Rerun)
	add_RerunFailed("POST", "/api/project-pipeline/definitions/{pipelineDefinitionID}/actions/rerun-failed", srv.RerunFailed)
	add_Cancel("POST", "/api/project-pipeline/definitions/{pipelineDefinitionID}/actions/cancel", srv.Cancel)
	add_OneClickCreate("POST", "/api/project-pipeline/actions/one-click-create", srv.OneClickCreate)
	add_BatchCreateByGittarPushHook("POST", "/api/project-pipeline/actions/create-by-gittar-push-hook", srv.BatchCreateByGittarPushHook)
	add_ListExecHistory("GET", "/api/project-pipeline/actions/exec-history", srv.ListExecHistory)
	add_DeleteByApp("DELETE", "/api/project-pipeline/actions/delete-by-app", srv.DeleteByApp)
}
