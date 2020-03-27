package conf

import (
	model "code.oldbody.com/studygolang/web/06layout/model"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/ini.v1"
)

//数据库初始化
var db *gorm.DB
var ConfigObj *model.Collect
var err error

func init() {
	//0.读配置文件 `go-ini`包
	ConfigObj = new(model.Collect)
	err := ini.MapTo(ConfigObj, "g:/gowork/src/code.oldbody.com/studygolang/web/06layout/conf/conf.ini")

	if err != nil {
		fmt.Printf("ini load config failed,err:%s\n", err)
		errors.New("ini load config failed")
		return
	}
}

func InitDB(dbname string) *gorm.DB {
	//连接名称
	dbAlias := ConfigObj.MysqlConf.Alias
	//数据库名称
	dbName := dbname
	//数据库连接用户
	dbUser := ConfigObj.MysqlConf.User
	//数据库连接密码
	dbPwd := ConfigObj.MysqlConf.Pwd
	//数据库IP(域名)
	dbHost := ConfigObj.MysqlConf.Host
	//数据库端口 (127.0.0.1:通过网卡，有防火墙限制； localhost:不经过网卡，没有防火墙限制)
	dbPort := ConfigObj.MysqlConf.Port
	//charset
	dbCharset := ConfigObj.MysqlConf.Charset

	//1.获取数据库连接信息，连接数据库
	db, err = gorm.Open(dbAlias, dbUser+":"+dbPwd+"@tcp"+"("+dbHost+":"+dbPort+")/"+dbName+"?charset="+dbCharset+"&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("connect to mysql error, err:%s\n", err)
		db.Close()
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	fmt.Println(dbUser + ":" + dbPwd + "@tcp" + "(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=" + dbCharset + "&parseTime=True&loc=Local")
	return db
}
