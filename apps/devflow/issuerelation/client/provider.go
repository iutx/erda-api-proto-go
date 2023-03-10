// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: issuerelation.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/apps/devflow/issuerelation/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.apps.devflow.issuerelation",
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
	issueRelationServiceClientType = reflect.TypeOf((*pb.IssueRelationServiceClient)(nil)).Elem()
	issueRelationServiceServerType = reflect.TypeOf((*pb.IssueRelationServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.apps.devflow.issuerelation-client":
		return p.client
	case "erda.apps.devflow.issuerelation.IssueRelationService":
		return &issueRelationServiceWrapper{client: p.client.IssueRelationService(), opts: opts}
	case "erda.apps.devflow.issuerelation.IssueRelationService.client":
		return p.client.IssueRelationService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case issueRelationServiceClientType:
		return p.client.IssueRelationService()
	case issueRelationServiceServerType:
		return &issueRelationServiceWrapper{client: p.client.IssueRelationService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.apps.devflow.issuerelation-client", &servicehub.Spec{
		Services: []string{
			"erda.apps.devflow.issuerelation.IssueRelationService",
			"erda.apps.devflow.issuerelation.IssueRelationService.client",
			"erda.apps.devflow.issuerelation-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			issueRelationServiceClientType,
			// server types
			issueRelationServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
