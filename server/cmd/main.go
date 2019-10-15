package main

import (
	"os"
	"os/signal"

	"github.com/nilorg/pkg/logger"

	pb "github.com/DeyiXu/istio-example/proto/helloworld"
	"github.com/DeyiXu/istio-example/server"
	"github.com/nilorg/istio"
)

func init() {
	logger.Init()
}

func main() {
	grpcServer := istio.NewGrpcServer("helloword", ":9000")

	grpcServer.Start()
	pb.RegisterGreeterServer(grpcServer.GetSrv(), &server.Server{})
	logger.Infof("go grpc server ...")
	// 等待中断信号优雅地关闭服务器
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	defer grpcServer.Stop()
}
