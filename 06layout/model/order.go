package model

//接收结构体
type OrderServerReq struct {
	OrderId   int32
	UserId    int32
	ProductId int32
	Addr      string
}

//返回结构体
type OrderServerResp struct {
	Msg string
}
