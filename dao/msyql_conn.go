package dao

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 参数1        数据库的别名，用来在ORM中切换数据库使用
	// 参数2        driverName
	// 参数3        对应的链接字符串
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(localhost:3306)/iax?charset=utf8&loc=Local")
	orm.SetMaxIdleConns("default", 10)
	orm.SetMaxOpenConns("default", 50)
	orm.Debug = true
}
