package app

import (
	"context"
	"fmt"
	pb "github.com/rapita/demo-example-foo-svc/pkg/api/example/v1/foo"
)

type FooServer struct {
}

func NewFooServer() *FooServer {
	return &FooServer{}
}

func (s *FooServer) Say(ctx context.Context, in *pb.SayRequest) (*pb.SayResponse, error) {
	return &pb.SayResponse{
		Text: fmt.Sprintf("Foo saying: %s", in.GetText()),
	}, nil
}
