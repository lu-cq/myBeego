package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/myBeego/routers"
	"github.com/astaxie/beego"
	_ "github.com/Go-SQL-Driver/MySQL"
)

func main() {
	orm.RegisterDataBase("default", "mysql","root@tcp(127.0.0.1:3306)/mybeego?charset=utf8")
	beego.Run()
}

