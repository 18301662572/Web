package main

const (
	getAllByProductIDAndCustomerID = `select * from p_orders where product_id in (:product_id) and customer_id=:customer_id`
)
// GetAllByProductIDAndCustomerID
// @param driver_id
// @param rate_date
// @return []Order, error
func GetAllByProductIDAndCustomerID(ctx context.Context, productIDs []uint64, customerID uint64) ([]Order, error) {
	var orderList []Order
	params := map[string]interface{}{
		"product_id" : productIDs,
		"customer_id": customerID,
	}
	// getAllByProductIDAndCustomerID 是 const 类型的 sql 字符串
	sql, args, err := sqlutil.Named(getAllByProductIDAndCustomerID, params)
	if err != nil {
		return nil, err
	}
	err = dao.QueryList(ctx, sqldbInstance, sql, args, &orderList)
	if err != nil {
		return nil, err
	}
	return orderList, err
}

//像这样的代码，在上线之前把DAO层的变更集的 const 部分直接拿给 dba 来进行审核，就比较方便了。
// 代码中的 sqlutil.Named 是类似于 sqlx 中的 Named 函数，同时支持 where 表达式中的比较操作符和 in。
