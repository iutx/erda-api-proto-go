// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: guide.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/dop/guide/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.dop.guide",
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
	guideServiceClientType = reflect.TypeOf((*pb.GuideServiceClient)(nil)).Elem()
	guideServiceServerType = reflect.TypeOf((*pb.GuideServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.dop.guide-client":
		return p.client
	case "erda.dop.guide.GuideService":
		return &guideServiceWrapper{client: p.client.GuideService(), opts: opts}
	case "erda.dop.guide.GuideService.client":
		return p.client.GuideService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case guideServiceClientType:
		return p.client.GuideService()
	case guideServiceServerType:
		return &guideServiceWrapper{client: p.client.GuideService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.dop.guide-client", &servicehub.Spec{
		Services: []string{
			"erda.dop.guide.GuideService",
			"erda.dop.guide.GuideService.client",
			"erda.dop.guide-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			guideServiceClientType,
			// server types
			guideServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
