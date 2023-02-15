// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: pipeline.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterPipelineServiceImp pipeline.proto
func RegisterPipelineServiceImp(regester transport.Register, srv PipelineServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterPipelineServiceHandler(regester, PipelineServiceHandler(srv), _ops.HTTP...)
	RegisterPipelineServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.core.pipeline.pipeline.PipelineService",
	)
}

var (
	pipelineServiceClientType  = reflect.TypeOf((*PipelineServiceClient)(nil)).Elem()
	pipelineServiceServerType  = reflect.TypeOf((*PipelineServiceServer)(nil)).Elem()
	pipelineServiceHandlerType = reflect.TypeOf((*PipelineServiceHandler)(nil)).Elem()
)

// PipelineServiceClientType .
func PipelineServiceClientType() reflect.Type { return pipelineServiceClientType }

// PipelineServiceServerType .
func PipelineServiceServerType() reflect.Type { return pipelineServiceServerType }

// PipelineServiceHandlerType .
func PipelineServiceHandlerType() reflect.Type { return pipelineServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		pipelineServiceClientType,
		// server types
		pipelineServiceServerType,
		// handler types
		pipelineServiceHandlerType,
	}
}
