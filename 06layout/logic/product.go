package logic

import (
	"code.oldbody.com/studygolang/web/06layout/dao"
	"code.oldbody.com/studygolang/web/06layout/model"
)

//业务逻辑

func PostProduct(req *model.ProductServerReq) (model *dao.ProductModel) {
	return dao.CreateProduct(req.Name)
}
