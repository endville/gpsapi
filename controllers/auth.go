package controllers

import (
	"github.com/astaxie/beego/orm"
	"gpsapi/models"
	"gpsapi/tools"
)

// 登录认证相关
type AuthController struct {
	BaseController
}

// @Title 登录
// @Description 用户登录，可以通过车队用户名、用户手机号、终端编号进行设备登录
// @Param	username		form 	string	true		"用户名"
// @Param	password		form 	string	true		"密码"
// @Success 200 {ret:0,type:1车队用户/2手机账户/3终端账号,msg:描述,data:对应账号ID}
// @Failure 400	请求的参数不正确
// @router / [post]
func (this *AuthController) Login() {
	userName := this.GetString("username")
	password := this.GetString("password")
	func() {
		if userName == "" {
			this.Data["json"] = map[string]interface{}{
				"code": 1,
				"msg":  "用户名不能为空",
			}
			return
		}
		if password == "" {

			this.Data["json"] = map[string]interface{}{
				"code": 1,
				"msg":  "密码不能为空",
			}
			return
		}
		md5_password := tools.MD5(password)

		var cond *orm.Condition
		// 先匹配车队
		cond = orm.NewCondition()
		cond1 := cond.And("GroupName", userName)
		if groups, count, err := models.GetAllGroups(cond1, 1, 1); err == nil {
			if count > 0 {
				if (*groups)[0].Password == md5_password {
					this.Data["json"] = map[string]interface{}{
						"code": 0,
						"msg":  "车队用户登录",
						"type": 1,
						"data": (*groups)[0].Id,
					}
					return
				}
			}
		}
		// 匹配用户
		cond = orm.NewCondition()
		cond2 := cond.And("UserName", userName)
		if users, count, err := models.GetAllUsers(cond2, 1, 1); err == nil {
			if count > 0 {
				if (*users)[0].Password == md5_password {
					this.Data["json"] = map[string]interface{}{
						"code": 0,
						"msg":  "手机号登录",
						"type": 2,
						"data": (*users)[0].Id,
					}
					return
				}
			}
		}

		// 匹配终端编号
		cond = orm.NewCondition()
		cond3 := cond.And("TerminalSn", userName)
		if terminals, count, err := models.GetAllTerminals(cond3, 1, 1); err == nil {
			if count > 0 {
				if (*terminals)[0].Password == password {
					this.Data["json"] = map[string]interface{}{
						"code": 0,
						"msg":  "终端编号登录",
						"type": 3,
						"data": (*terminals)[0].Id,
					}
					return
				}
			}
		}

		// 全部没能匹配到
		this.Data["json"] = map[string]interface{}{
			"code": 1,
			"msg":  "用户名或密码错误",
		}
	}()

	this.ServeJson()
}

// @Title 登出
// @Description 登出
// @Param	username	form 	string	true		"用户名"
// @Success 200 {int} models.User.Id
// @Failure 400	请求的参数不正确
// @router / [get]
func (this *AuthController) Logout() {
	this.Data["json"] = map[string]interface{}{
		"code": 0,
		"msg":  "退出成功",
	}
	this.ServeJson()
}
