package routers

import (
	"github.com/astaxie/beego"

	"routing-sample/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
