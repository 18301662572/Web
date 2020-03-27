package main

import (
	logic "code.oldbody.com/studygolang/web/06layout/logic"
	order "code.oldbody.com/studygolang/web/06layout/model"
	pb "code.oldbody.com/studygolang/web/06layout/proto"
	"context"
)

//order微服务

type OrderServer struct {
}

func (o *OrderServer) PostOrder(c context.Context, req *pb.OrderReq) (respresult *pb.OrderResp, err error) {
	//把参数解析到自己的结构体中
	serverreq := new(order.OrderServerReq)
	serverreq.UserId = req.UserId
	serverreq.OrderId = req.OrderId
	serverreq.ProductId = req.ProductId
	serverreq.Addr = req.Addr
	result := logic.PostOrder(serverreq)
	respresult = new(pb.OrderResp)
	respresult.Msg = result
	return
}
