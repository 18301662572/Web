package logic

import (
	dao "code.oldbody.com/studygolang/web/06layout/dao"
	model "code.oldbody.com/studygolang/web/06layout/model"
	"fmt"
	"time"
)

//业务逻辑

func PostOrder(req *model.OrderServerReq) (result string) {
	order := dao.CreateOrder(req.UserId, req.ProductId, req.Addr)
	if order != nil {
		fmt.Printf("[%s %s]\n", time.Now().Format("2006-01-02 15:04:05"), "CreateOrder success!")
		return "success"
	} else {
		return "create order failed!"
	}
}
