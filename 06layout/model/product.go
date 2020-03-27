package model

//接收结构体
type ProductServerReq struct {
	Name string
}

//返回结构体
type ProductResp struct {
	ProductId int32
	Name      string
}
