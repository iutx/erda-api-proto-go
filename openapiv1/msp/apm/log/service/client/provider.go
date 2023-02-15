// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: msp_apm_log_service.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/openapiv1/msp/apm/log/service/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.openapiv1.msp",
	"grpc-client",
}

// +provider
type provider struct {
	client Client
}

func (p *provider) Init(ctx servicehub.Context) error {
	var conn grpc.ClientConnInterface
	for _, dep := range dependencies {
		c, ok := ctx.Service(dep).(grpc.ClientConnInterface)
		if ok {
			conn = c
			break
		}
	}
	if conn == nil {
		return fmt.Errorf("not found connector in (%s)", strings.Join(dependencies, ", "))
	}
	p.client = New(conn)
	return nil
}

var (
	clientsType                = reflect.TypeOf((*Client)(nil)).Elem()
	mspApmLogServiceClientType = reflect.TypeOf((*pb.MspApmLogServiceClient)(nil)).Elem()
	mspApmLogServiceServerType = reflect.TypeOf((*pb.MspApmLogServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.openapiv1.msp-client":
		return p.client
	case "erda.openapiv1.msp.MspApmLogService":
		return &mspApmLogServiceWrapper{client: p.client.MspApmLogService(), opts: opts}
	case "erda.openapiv1.msp.MspApmLogService.client":
		return p.client.MspApmLogService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case mspApmLogServiceClientType:
		return p.client.MspApmLogService()
	case mspApmLogServiceServerType:
		return &mspApmLogServiceWrapper{client: p.client.MspApmLogService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.openapiv1.msp-client", &servicehub.Spec{
		Services: []string{
			"erda.openapiv1.msp.MspApmLogService",
			"erda.openapiv1.msp.MspApmLogService.client",
			"erda.openapiv1.msp-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			mspApmLogServiceClientType,
			// server types
			mspApmLogServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}