package main

import (
	logic "code.oldbody.com/studygolang/web/06layout/logic"
	product "code.oldbody.com/studygolang/web/06layout/model"
	pb "code.oldbody.com/studygolang/web/06layout/proto"
	"context"
	"errors"
	"fmt"
	"time"
)

//微服务

type ProductServer struct {
}

func (p *ProductServer) PostProduct(c context.Context, req *pb.ProductReq) (resp *pb.ProductResp, err error) {
	//把参数解析到自己的结构体中
	serverreq := new(product.ProductServerReq)
	serverreq.Name = req.Name
	result := logic.PostProduct(serverreq)
	if result != nil {
		resp = new(pb.ProductResp)
		resp.ProductId = result.ProductId
		resp.Name = result.Name
		fmt.Printf("[%s %s]\n", time.Now().Format("2006-01-02 15:04:05"), "CreateProduct success!")
	} else {
		err = errors.New("create product failed!")
	}
	return
}
