package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/nilorg/pkg/logger"

	pb "github.com/DeyiXu/istio-example/proto/helloworld"
	"github.com/nilorg/istio"
)

func init() {
	logger.Init()
}
func main() {
	client := istio.NewGrpcClient("server", 9000)
	greeterClient := pb.NewGreeterClient(client.GetConn())

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		r, err := greeterClient.SayHello(context.Background(), &pb.HelloRequest{Name: "xudeyi"})
		if err != nil {
			logger.Infof("could not greet: %v", err)
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		logger.Infof("Greeting: %s", r.Message)
		c.JSON(200, gin.H{
			"message": r.Message,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
	defer client.Close()
}
