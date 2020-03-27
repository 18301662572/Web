package main

import (
	proto "code.oldbody.com/studygolang/web/06layout/proto"
	conf "code.oldbody.com/studygolang/web/06layout/system/conf"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"net/http"
)

//服务网关

//注： 使用postman验证，必须使用json格式
//curl -X POST http://127.0.0.1:8080/postorder -d "{\"user_id\":3,\"order_id\":33,\"product_id\":444,\"addr\":\"\"}"
//curl -X POST http://127.0.0.1:8080/postproduct -d "{\"name\":\"dada\"}"

func main() {

	//获取addr信息
	gateway := conf.ConfigObj.AddrConf.GateAddr
	server := conf.ConfigObj.AddrConf.OrderserverAddr

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	//创建路由处理器
	mux := runtime.NewServeMux()

	//创建一个[]grpc.DialOption
	opts := []grpc.DialOption{grpc.WithInsecure()}

	//RestService服务相关的REST接口中转到后面的GRPC服务
	//OrderGRPC
	proto.RegisterOrderHandlerFromEndpoint(ctx, mux, server, opts)
	//ProductGRPC
	proto.RegisterProductHandlerFromEndpoint(ctx, mux, server, opts)

	fmt.Println("addr:", gateway)
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	http.ListenAndServe(gateway, mux)
}
