// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: stream.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/dop/issue/stream/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.dop.issue.stream",
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
	clientsType                         = reflect.TypeOf((*Client)(nil)).Elem()
	commentIssueStreamServiceClientType = reflect.TypeOf((*pb.CommentIssueStreamServiceClient)(nil)).Elem()
	commentIssueStreamServiceServerType = reflect.TypeOf((*pb.CommentIssueStreamServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.dop.issue.stream-client":
		return p.client
	case "erda.dop.issue.stream.CommentIssueStreamService":
		return &commentIssueStreamServiceWrapper{client: p.client.CommentIssueStreamService(), opts: opts}
	case "erda.dop.issue.stream.CommentIssueStreamService.client":
		return p.client.CommentIssueStreamService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case commentIssueStreamServiceClientType:
		return p.client.CommentIssueStreamService()
	case commentIssueStreamServiceServerType:
		return &commentIssueStreamServiceWrapper{client: p.client.CommentIssueStreamService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.dop.issue.stream-client", &servicehub.Spec{
		Services: []string{
			"erda.dop.issue.stream.CommentIssueStreamService",
			"erda.dop.issue.stream.CommentIssueStreamService.client",
			"erda.dop.issue.stream-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			commentIssueStreamServiceClientType,
			// server types
			commentIssueStreamServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
