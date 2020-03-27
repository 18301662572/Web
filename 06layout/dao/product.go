package dao

import "code.oldbody.com/studygolang/web/06layout/system/conf"

//数据访问层

type ProductModel struct {
	ProductId int32  `gorm:"primary_key;auto;column:ProductId"`
	Name      string `gorm:"column:Name"`
}

func CreateProduct(name string) *ProductModel {
	product := new(ProductModel)
	product.Name = name
	//连接mysql
	db := conf.InitDB("product")
	db.Table("product").Create(&product)
	//关闭连接
	db.Close()
	return product
}
