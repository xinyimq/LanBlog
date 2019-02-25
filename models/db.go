package models

import "github.com/astaxie/beego/orm"
import _"github.com/go-sql-driver/mysql"



const (
	DB_DRIVER="mysql"
	DB_URL="root:cxx15755599659@tcp(127.0.0.1:3306)/myblog"
)

func RegisterDB() {

	orm.RegisterDriver(DB_DRIVER, orm.DRMySQL)
	orm.RegisterModel(new(Topic),new(Label),new(Category))
	orm.RegisterDataBase("default", "mysql",DB_URL )
}