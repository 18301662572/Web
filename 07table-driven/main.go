package _7table_driven

//table-driven 开发

//实例：
//func entry() {
//	var bi BusinessInstance
//	switch businessType {
//	case TravelBusiness:
//		bi = travelorder.New()
//	case MarketBusiness:
//		bi = marketorder.New()
//	default:
//		return errors.New("not supported business")
//	}
//}

//修改为：
//var businessInstanceMap = map[int]BusinessInstance {
//	TravelBusiness : travelorder.New(),
//	MarketBusiness : marketorder.New(),
//}
//func entry() {
//	bi := businessInstanceMap[businessType]
//}
