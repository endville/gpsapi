package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["gpsapi/controllers:TerminalController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:TerminalController"],
		beego.ControllerComments{
			"Search",
			`/search`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:TerminalController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:TerminalController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:TerminalController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:TerminalController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:TerminalController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:TerminalController"],
		beego.ControllerComments{
			"Get",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:TerminalController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:TerminalController"],
		beego.ControllerComments{
			"GetProfile",
			`/:id/profile`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:TerminalController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:TerminalController"],
		beego.ControllerComments{
			"GetCarrier",
			`/:id/carrier`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:TerminalController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:TerminalController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:TerminalController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:TerminalController"],
		beego.ControllerComments{
			"PutProfile",
			`/:id/profile`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:TerminalController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:TerminalController"],
		beego.ControllerComments{
			"PutCarrier",
			`/:id/carrier`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:TerminalController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:TerminalController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:RoleController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:RoleController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:RoleController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:RoleController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:RoleController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:RoleController"],
		beego.ControllerComments{
			"Get",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:RoleController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:RoleController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:RoleController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:RoleController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:MessageController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:MessageController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:MessageController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:MessageController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:MessageController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:MessageController"],
		beego.ControllerComments{
			"Get",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:MessageController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:MessageController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:GateController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:GateController"],
		beego.ControllerComments{
			"SessionCount",
			`/sessionCount`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:GateController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:GateController"],
		beego.ControllerComments{
			"State",
			`/state`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:UserController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:UserController"],
		beego.ControllerComments{
			"Search",
			`/search`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:UserController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:UserController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:UserController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:UserController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:UserController"],
		beego.ControllerComments{
			"GetProfile",
			`/:id/profile`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:UserController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:UserController"],
		beego.ControllerComments{
			"GetTerminals",
			`/:id/terminal`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:UserController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:UserController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:UserController"],
		beego.ControllerComments{
			"PutProfile",
			`/:id/profile`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:UserController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:WarningController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:WarningController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:WarningController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:WarningController"],
		beego.ControllerComments{
			"Get",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:WarningController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:WarningController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:WarningController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:WarningController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:ResourceController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:ResourceController"],
		beego.ControllerComments{
			"Search",
			`/search`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:StatisticController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:StatisticController"],
		beego.ControllerComments{
			"Terminal",
			`/serviceExpire`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:StatisticController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:StatisticController"],
		beego.ControllerComments{
			"TerminalRunInfo",
			`/terminalRunInfo`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:StatisticController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:StatisticController"],
		beego.ControllerComments{
			"TerminalMileage",
			`/terminalMileage`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:StatisticController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:StatisticController"],
		beego.ControllerComments{
			"TerminalStayInfo",
			`/terminalStay`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:StatisticController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:StatisticController"],
		beego.ControllerComments{
			"TerminalWarning",
			`/terminalWarning`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:StatisticController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:StatisticController"],
		beego.ControllerComments{
			"TerminalOnline",
			`/terminalOnline`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:StatisticController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:StatisticController"],
		beego.ControllerComments{
			"TerminalOnlineDetails",
			`/terminalOnlineDetails`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:StatisticController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:StatisticController"],
		beego.ControllerComments{
			"TerminalOffline",
			`/terminalOffline`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:StatisticController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:StatisticController"],
		beego.ControllerComments{
			"TerminalQuery",
			`/terminalInfo`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:GeoController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:GeoController"],
		beego.ControllerComments{
			"Get",
			`/:geoId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:GeoController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:GeoController"],
		beego.ControllerComments{
			"TerminalGeo",
			`/terminal/:terminalId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:GeoController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:GeoController"],
		beego.ControllerComments{
			"UserGeo",
			`/user/:userId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:GeoController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:GeoController"],
		beego.ControllerComments{
			"GroupGeo",
			`/group/:groupId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:GeoController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:GeoController"],
		beego.ControllerComments{
			"GetGroupGeosRound",
			`/group/round/:groupId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:GeoController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:GeoController"],
		beego.ControllerComments{
			"GetGroupGeosBox",
			`/group/box/:groupId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:GeoController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:GeoController"],
		beego.ControllerComments{
			"GetGroupGeosPolygon",
			`/group/polygon/:groupId`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:RightController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:RightController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:SessionController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:SessionController"],
		beego.ControllerComments{
			"Get",
			`/:terminalSN`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:SessionController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:SessionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:SessionController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:SessionController"],
		beego.ControllerComments{
			"Delete",
			`/:uid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:AuthController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:AuthController"],
		beego.ControllerComments{
			"Login",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:AuthController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:AuthController"],
		beego.ControllerComments{
			"Logout",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:LogController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:LogController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:LogController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:LogController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:LogController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:LogController"],
		beego.ControllerComments{
			"Get",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:LogController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:LogController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:GroupController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:GroupController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:GroupController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:GroupController"],
		beego.ControllerComments{
			"Search",
			`/search`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:GroupController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:GroupController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:GroupController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:GroupController"],
		beego.ControllerComments{
			"Get",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:GroupController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:GroupController"],
		beego.ControllerComments{
			"GetProfile",
			`/:id/profile`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:GroupController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:GroupController"],
		beego.ControllerComments{
			"GetTerminals",
			`/:id/terminal`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:GroupController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:GroupController"],
		beego.ControllerComments{
			"GetRoles",
			`/:id/roles`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:GroupController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:GroupController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:GroupController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:GroupController"],
		beego.ControllerComments{
			"PutProfile",
			`/:id/profile`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:GroupController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:GroupController"],
		beego.ControllerComments{
			"PutRole",
			`/:id/role`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["gpsapi/controllers:GroupController"] = append(beego.GlobalControllerRouter["gpsapi/controllers:GroupController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
