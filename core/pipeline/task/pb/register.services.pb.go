// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: task.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterTaskServiceImp task.proto
func RegisterTaskServiceImp(regester transport.Register, srv TaskServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterTaskServiceHandler(regester, TaskServiceHandler(srv), _ops.HTTP...)
	RegisterTaskServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.core.pipeline.task.TaskService",
	)
}

var (
	taskServiceClientType  = reflect.TypeOf((*TaskServiceClient)(nil)).Elem()
	taskServiceServerType  = reflect.TypeOf((*TaskServiceServer)(nil)).Elem()
	taskServiceHandlerType = reflect.TypeOf((*TaskServiceHandler)(nil)).Elem()
)

// TaskServiceClientType .
func TaskServiceClientType() reflect.Type { return taskServiceClientType }

// TaskServiceServerType .
func TaskServiceServerType() reflect.Type { return taskServiceServerType }

// TaskServiceHandlerType .
func TaskServiceHandlerType() reflect.Type { return taskServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		taskServiceClientType,
		// server types
		taskServiceServerType,
		// handler types
		taskServiceHandlerType,
	}
}