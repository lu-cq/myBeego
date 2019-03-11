package routers

import (
	"github.com/astaxie/beego"
	"github.com/myBeego/controllers"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    beego.Include(&controllers.UserController{})
}
