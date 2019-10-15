package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/nilorg/pkg/logger"

	pb "github.com/DeyiXu/istio-example/proto/helloworld"
	"github.com/nilorg/istio"
)

func init() {
	logger.Init()
}
func main() {
	client := istio.NewGrpcClient("127.0.0.1", 9000)
	greeterClient := pb.NewGreeterClient(client.GetConn())

	go func() {
		for {
			time.Sleep(time.Second)
			r, err := greeterClient.SayHello(context.Background(), &pb.HelloRequest{Name: "xudeyi"})
			if err != nil {
				logger.Infof("could not greet: %v", err)
				continue
			}
			logger.Infof("Greeting: %s", r.Message)
		}
	}()

	// 等待中断信号优雅地关闭服务器
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	defer client.Close()
}
