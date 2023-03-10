// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: details.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/monitor/alertdetail/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.core.monitor.alertdetail",
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
	clientsType                  = reflect.TypeOf((*Client)(nil)).Elem()
	alertDetailServiceClientType = reflect.TypeOf((*pb.AlertDetailServiceClient)(nil)).Elem()
	alertDetailServiceServerType = reflect.TypeOf((*pb.AlertDetailServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.core.monitor.alertdetail-client":
		return p.client
	case "erda.core.monitor.alertdetail.AlertDetailService":
		return &alertDetailServiceWrapper{client: p.client.AlertDetailService(), opts: opts}
	case "erda.core.monitor.alertdetail.AlertDetailService.client":
		return p.client.AlertDetailService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case alertDetailServiceClientType:
		return p.client.AlertDetailService()
	case alertDetailServiceServerType:
		return &alertDetailServiceWrapper{client: p.client.AlertDetailService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.core.monitor.alertdetail-client", &servicehub.Spec{
		Services: []string{
			"erda.core.monitor.alertdetail.AlertDetailService",
			"erda.core.monitor.alertdetail.AlertDetailService.client",
			"erda.core.monitor.alertdetail-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			alertDetailServiceClientType,
			// server types
			alertDetailServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
