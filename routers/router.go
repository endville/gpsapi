// @APIVersion 1.0.0
// @Title 云电池项目 API
// @Description 云电池项目API
// @Contact chenjq@endville.com
// @TermsOfServiceUrl http://www.endville.com/
package routers

import (
	"gpsapi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/resource",
			beego.NSInclude(
				&controllers.ResourceController{},
			),
		),
		beego.NSNamespace("/group",
			beego.NSInclude(
				&controllers.GroupController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/terminal",
			beego.NSInclude(
				&controllers.TerminalController{},
			),
		),
		beego.NSNamespace("/log",
			beego.NSInclude(
				&controllers.LogController{},
			),
		),
		beego.NSNamespace("/geo",
			beego.NSInclude(
				&controllers.GeoController{},
			),
		),
		beego.NSNamespace("/message",
			beego.NSInclude(
				&controllers.MessageController{},
			),
		),
		beego.NSNamespace("/gate",
			beego.NSInclude(
				&controllers.GateController{},
			),
		),
		beego.NSNamespace("/session",
			beego.NSInclude(
				&controllers.SessionController{},
			),
		),
		beego.NSNamespace("/auth",
			beego.NSInclude(
				&controllers.AuthController{},
			),
		),
		beego.NSNamespace("/role",
			beego.NSInclude(
				&controllers.RoleController{},
			),
		),
		beego.NSNamespace("/right",
			beego.NSInclude(
				&controllers.RightController{},
			),
		),
		beego.NSNamespace("/warning",
			beego.NSInclude(
				&controllers.WarningController{},
			),
		),
		beego.NSNamespace("/statistic",
			beego.NSInclude(
				&controllers.StatisticController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
