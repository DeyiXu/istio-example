package server

import (
	"context"

	pb "github.com/DeyiXu/istio-example/proto/helloworld"
)

// Server is used to implement helloworld.GreeterServer.
type Server struct{}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
