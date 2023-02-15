// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: action.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/pipeline/action/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.core.pipeline.action",
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
	clientsType             = reflect.TypeOf((*Client)(nil)).Elem()
	actionServiceClientType = reflect.TypeOf((*pb.ActionServiceClient)(nil)).Elem()
	actionServiceServerType = reflect.TypeOf((*pb.ActionServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.core.pipeline.action-client":
		return p.client
	case "erda.core.pipeline.action.ActionService":
		return &actionServiceWrapper{client: p.client.ActionService(), opts: opts}
	case "erda.core.pipeline.action.ActionService.client":
		return p.client.ActionService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case actionServiceClientType:
		return p.client.ActionService()
	case actionServiceServerType:
		return &actionServiceWrapper{client: p.client.ActionService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.core.pipeline.action-client", &servicehub.Spec{
		Services: []string{
			"erda.core.pipeline.action.ActionService",
			"erda.core.pipeline.action.ActionService.client",
			"erda.core.pipeline.action-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			actionServiceClientType,
			// server types
			actionServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
