package main

import (
	pb "code.oldbody.com/studygolang/web/06layout/proto"
	conf "code.oldbody.com/studygolang/web/06layout/system/conf"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

//运行main
//go run main.go order.go product.go

//注册rpc服务
func main() {
	//获取addr信息
	addr := conf.ConfigObj.AddrConf.OrderserverAddr
	fmt.Println("addr:", addr)

	//1.监听
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("监听异常：%s \n", err)
		return
	}
	fmt.Printf("开始监听：%s \n", addr)
	//2、实例化gRPC
	s := grpc.NewServer()
	//3.在gRPC上注册微服务

	//注册Order服务
	var orderServer = new(OrderServer)
	pb.RegisterOrderServer(s, orderServer)
	//注册Product服务
	var productServer = new(ProductServer)
	pb.RegisterProductServer(s, productServer)

	//4.启动gRPC服务端
	s.Serve(lis)
}
