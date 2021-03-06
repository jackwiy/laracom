package main

import (
	"context"
	"github.com/micro/go-micro"
	pb "github.com/nonfu/laracom/demo-service/proto/demo"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Name("laracom.demo.cli"),
		micro.WrapClient(hystrix.NewClientWrapper()),
	)
	service.Init()

	client := pb.NewDemoServiceClient("laracom.service.demo", service.Client())
	rsp, err := client.SayHello(context.TODO(), &pb.DemoRequest{Name: "学院君"})
	if err != nil {
		log.Fatalf("服务调用失败：%v", err)
		return
	}
	log.Println(rsp.Text)
}