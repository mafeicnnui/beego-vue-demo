package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"student/controllers"
)

func init() {
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/:id", &controllers.UserControllerPar{})
	beego.Router("/user/add", &controllers.UserControllerAdd{})
	beego.Router("/user/search", &controllers.UserControllerSearch{})
	beego.Router("/user/search_", &controllers.UserControllerSearch_{})
}
