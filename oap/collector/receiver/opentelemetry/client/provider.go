// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: opentelemetry.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/oap/collector/receiver/opentelemetry/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.oap.collector.receiver.opentelemetry",
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
	clientsType                    = reflect.TypeOf((*Client)(nil)).Elem()
	openTelemetryServiceClientType = reflect.TypeOf((*pb.OpenTelemetryServiceClient)(nil)).Elem()
	openTelemetryServiceServerType = reflect.TypeOf((*pb.OpenTelemetryServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.oap.collector.receiver.opentelemetry-client":
		return p.client
	case "erda.oap.collector.receiver.opentelemetry.OpenTelemetryService":
		return &openTelemetryServiceWrapper{client: p.client.OpenTelemetryService(), opts: opts}
	case "erda.oap.collector.receiver.opentelemetry.OpenTelemetryService.client":
		return p.client.OpenTelemetryService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case openTelemetryServiceClientType:
		return p.client.OpenTelemetryService()
	case openTelemetryServiceServerType:
		return &openTelemetryServiceWrapper{client: p.client.OpenTelemetryService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.oap.collector.receiver.opentelemetry-client", &servicehub.Spec{
		Services: []string{
			"erda.oap.collector.receiver.opentelemetry.OpenTelemetryService",
			"erda.oap.collector.receiver.opentelemetry.OpenTelemetryService.client",
			"erda.oap.collector.receiver.opentelemetry-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			openTelemetryServiceClientType,
			// server types
			openTelemetryServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
