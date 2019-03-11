package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	name := c.GetString("name")
	c.Data["Website"] = "beego.me" + name
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
