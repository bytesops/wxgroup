package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"wxgroup/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/group", &controllers.GroupController{})
	beego.Router("/qrcode", &controllers.QrcodeController{})
}
