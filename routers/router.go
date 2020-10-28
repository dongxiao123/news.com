package routers

import (
	"github.com/astaxie/beego"
	"news.com/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
