// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: sync.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterIssueSyncServiceImp sync.proto
func RegisterIssueSyncServiceImp(regester transport.Register, srv IssueSyncServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterIssueSyncServiceHandler(regester, IssueSyncServiceHandler(srv), _ops.HTTP...)
	RegisterIssueSyncServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.dop.issue.sync.IssueSyncService",
	)
}

var (
	issueSyncServiceClientType  = reflect.TypeOf((*IssueSyncServiceClient)(nil)).Elem()
	issueSyncServiceServerType  = reflect.TypeOf((*IssueSyncServiceServer)(nil)).Elem()
	issueSyncServiceHandlerType = reflect.TypeOf((*IssueSyncServiceHandler)(nil)).Elem()
)

// IssueSyncServiceClientType .
func IssueSyncServiceClientType() reflect.Type { return issueSyncServiceClientType }

// IssueSyncServiceServerType .
func IssueSyncServiceServerType() reflect.Type { return issueSyncServiceServerType }

// IssueSyncServiceHandlerType .
func IssueSyncServiceHandlerType() reflect.Type { return issueSyncServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		issueSyncServiceClientType,
		// server types
		issueSyncServiceServerType,
		// handler types
		issueSyncServiceHandlerType,
	}
}
