// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: msp_apm_trace.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/openapiv1/msp/apm/trace/pb"
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
	clientsType           = reflect.TypeOf((*Client)(nil)).Elem()
	mspApmTraceClientType = reflect.TypeOf((*pb.MspApmTraceClient)(nil)).Elem()
	mspApmTraceServerType = reflect.TypeOf((*pb.MspApmTraceServer)(nil)).Elem()
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
	case "erda.openapiv1.msp.MspApmTrace":
		return &mspApmTraceWrapper{client: p.client.MspApmTrace(), opts: opts}
	case "erda.openapiv1.msp.MspApmTrace.client":
		return p.client.MspApmTrace()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case mspApmTraceClientType:
		return p.client.MspApmTrace()
	case mspApmTraceServerType:
		return &mspApmTraceWrapper{client: p.client.MspApmTrace(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.openapiv1.msp-client", &servicehub.Spec{
		Services: []string{
			"erda.openapiv1.msp.MspApmTrace",
			"erda.openapiv1.msp.MspApmTrace.client",
			"erda.openapiv1.msp-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			mspApmTraceClientType,
			// server types
			mspApmTraceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
