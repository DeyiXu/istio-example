package server

import (
	"context"
	"fmt"
	"os"

	pb "github.com/DeyiXu/istio-example/proto/helloworld"
	"github.com/nilorg/pkg/logger"
)

// Server is used to implement helloworld.GreeterServer.
type Server struct{}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	hostname, _ := os.Hostname()
	msg := fmt.Sprintf("%s: Hello %s", hostname, in.Name)
	logger.Debugln(msg)
	return &pb.HelloReply{Message: msg}, nil
}
