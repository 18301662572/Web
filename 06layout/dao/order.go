package dao

import (
	conf "code.oldbody.com/studygolang/web/06layout/system/conf"
)

//数据访问层

type OrderModel struct {
	OrderId   int32  `gorm:"primary_key;auto;column:OrderId"`
	UserId    int32  `gorm:"column:UserId"`
	ProductId int32  `gorm:"column:ProductId"`
	Addr      string `gorm:"default:"";column:Addr"`
}

func CreateOrder(userid int32, productid int32, addr string) *OrderModel {
	order := new(OrderModel)
	order.UserId = userid
	order.ProductId = productid
	order.Addr = addr
	//连接mysql
	db := conf.InitDB("order")
	db.Table("order").Create(&order)
	//关闭连接
	db.Close()
	return order
}
