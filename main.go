package main

import (
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/myBeego/controllers"
	_ "github.com/myBeego/routers"
)

func main() {
	orm.RegisterDataBase("default", "mysql","root@tcp(127.0.0.1:3306)/mybeego?charset=utf8")
	beego.Router("/item", &controllers.MainController{})
	beego.Run()
}

