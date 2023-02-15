// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: build.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/pipeline/build/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.core.pipeline.build",
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
	clientsType            = reflect.TypeOf((*Client)(nil)).Elem()
	buildServiceClientType = reflect.TypeOf((*pb.BuildServiceClient)(nil)).Elem()
	buildServiceServerType = reflect.TypeOf((*pb.BuildServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.core.pipeline.build-client":
		return p.client
	case "erda.core.pipeline.build.BuildService":
		return &buildServiceWrapper{client: p.client.BuildService(), opts: opts}
	case "erda.core.pipeline.build.BuildService.client":
		return p.client.BuildService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case buildServiceClientType:
		return p.client.BuildService()
	case buildServiceServerType:
		return &buildServiceWrapper{client: p.client.BuildService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.core.pipeline.build-client", &servicehub.Spec{
		Services: []string{
			"erda.core.pipeline.build.BuildService",
			"erda.core.pipeline.build.BuildService.client",
			"erda.core.pipeline.build-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			buildServiceClientType,
			// server types
			buildServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
