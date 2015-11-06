package controllers

import (
	"encoding/json"
	"gpsapi/models"
	"gpsapi/tools"

	"github.com/astaxie/beego/orm"
)

// 用户相关
type UserController struct {
	BaseController
}

// @Title 查询接口
// @Description 根据关键字查询用户
// @Param	key	query	string	false	"关键字"
// @Success 200 {object} models.UserSearchModel
// @Failure 400 请求的参数不正确
// @router /search [get]
func (this *UserController) Search() {
	key := this.GetString("key", "")

	if len(key) < 3 {
		this.Data["json"] = map[string]interface{}{
			"code": 1,
			"msg":  "关键字过短",
		}
	} else {
		users := models.SearchUser("%"+key+"%", 25)

		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"data": users,
		}
	}

	this.ServeJson()
}

// @Title 添加一个新的用户
// @Description 添加一个新的用户
// @Param	body		body 	models.User	true		"用户信息"
// @Success 200 {int} models.User.Id
// @Failure 400	请求的参数不正确
// @router / [post]
func (this *UserController) Post() {
	user := models.User{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	if err != nil {
		this.ResponseErrorJSON(400, errorFormat(ErrorBadJson_400, err.Error()))
	} else {
		user.Password = tools.MD5(user.Password)
		if id, err := models.AddUser(user); err != nil {
			this.ResponseErrorJSON(400, errorFormat(ErrorBadParam_400, err.Error()))
		} else {
			this.Data["json"] = map[string]interface{}{
				"code": 0,
				"data": id,
			}
		}
	}
	this.ServeJson()
}

// @Title 获取用户列表
// @Description 获取用户列表
// @Param	pageIndex	query	int	false	"页码, 默认1"
// @Param	pageSize	query	int	false	"每页显示条数, 默认30"
// @Success 200 {object} models.User
// @Failure 400 请求的参数不正确
// @router / [get]
func (this *UserController) GetAll() {
	cond := orm.NewCondition()

	pageIndex, _ := this.GetInt("pageIndex", 1)
	pageSize, _ := this.GetInt("pageSize", 30)
	users, total, err := models.GetAllUsers(cond, pageIndex, pageSize)
	if err != nil {
		this.ResponseErrorJSON(400, errorFormat(ErrorBadRequest_400, err.Error()))
	}
	this.Data["json"] = map[string]interface{}{
		"code":  0,
		"data":  users,
		"total": total,
	}
	this.ServeJson()
}

// @Title 根据用户唯一ID号获取用户信息
// @Description 根据用户唯一ID号获取用户信息
// @Param	id		path 	int	true		用户唯一ID号
// @Success 200 {object} models.User
// @Failure 404 未找到对应的记录
// @router /:id [get]
func (this *UserController) Get() {
	id, _ := this.GetInt64(":id")
	user, err := models.GetUser(id)
	if err != nil {
		this.ResponseErrorJSON(404, errorFormat(ErrorDataNotFound_404, err.Error()))
	} else {
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"data": user,
		}
	}
	this.ServeJson()
}

// @Title 根据用户唯一ID号获取用户Profile信息
// @Description 根据用户唯一ID号获取用户Profile信息
// @Param	id		path 	int	true		用户唯一ID号(注意不是Profile的ID)
// @Success 200 {object} models.User
// @Failure 404 未找到对应的记录
// @router /:id/profile [get]
func (this *UserController) GetProfile() {
	id, _ := this.GetInt64(":id")
	profile, err := models.GetUserProfile(id)
	if err != nil {
		this.ResponseErrorJSON(404, errorFormat(ErrorDataNotFound_404, err.Error()))
	} else {
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"data": profile,
		}
	}
	this.ServeJson()
}

// @Title 获取用户的终端设备列表
// @Description 获取用户的终端设备列表
// @Param	id		path 	int	true		用户唯一ID号
// @Success 200 {object} models.Terminal
// @Failure 400 请求的参数不正确
// @router /:id/terminal [get]
func (this *UserController) GetTerminals() {
	id, _ := this.GetInt64(":id")

	cond := orm.NewCondition().And("user__id", id)
	list, total, err := models.GetAllTerminals(cond, 1, 1000)
	if err != nil {
		this.ResponseErrorJSON(400, errorFormat(ErrorBadRequest_400, err.Error()))
	}
	this.Data["json"] = map[string]interface{}{
		"code":  0,
		"data":  list,
		"total": total,
	}
	this.ServeJson()
}

// @Title 更新用户信息
// @Description 更新用户信息
// @Param	id		path 	int	true		用户唯一ID
// @Param	body		body 	models.User	true		用户信息
// @Success 200 {object} models.User
// @Failure 400	请求的参数不正确
// @router /:id [put]
func (this *UserController) Put() {
	id, _ := this.GetInt64(":id")
	var user models.User
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &user); err != nil {
		this.ResponseErrorJSON(400, errorFormat(ErrorBadJson_400, err.Error()))
	} else {
		if user.Password != "" {
			user.Password = tools.MD5(user.Password)
		}
		update, err := models.UpdateUser(id, &user)
		if err != nil {
			this.ResponseErrorJSON(400, errorFormat(ErrorBadParam_400, err.Error()))
		} else {
			this.Data["json"] = map[string]interface{}{
				"code": 0,
				"data": update,
			}
		}
	}
	this.ServeJson()
}

// @Title 更新用户Profile信息
// @Description 更新用户Profile信息
// @Param	id		path 	int	true		用户唯一ID(注意不是Profile的ID)
// @Param	body	body 	models.UserProfile	true		用户Profile信息
// @Success 200 {object} models.UserProfile
// @Failure 400	请求的参数不正确
// @router /:id/profile [put]
func (this *UserController) PutProfile() {
	id, _ := this.GetInt64(":id")
	var profile models.UserProfile
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &profile); err != nil {
		this.ResponseErrorJSON(400, errorFormat(ErrorBadJson_400, err.Error()))
	} else {
		update, err := models.UpdateUserProfile(id, &profile)
		if err != nil {
			this.ResponseErrorJSON(400, errorFormat(ErrorBadParam_400, err.Error()))
		} else {
			this.Data["json"] = map[string]interface{}{
				"code": 0,
				"data": update,
			}
		}
	}
	this.ServeJson()
}

// @Title 删除用户
// @Description 删除用户
// @Param	id		path 	int	true		"用户唯一ID"
// @Success 200 {int} 返回记录ID
// @Failure 404 未找到对应的记录
// @router /:id [delete]
func (this *UserController) Delete() {
	id, _ := this.GetInt64(":id")
	if err := models.DeleteUser(id); err != nil {
		this.ResponseErrorJSON(404, errorFormat(ErrorDataNotFound_404, err.Error()))
	}
	this.Data["json"] = map[string]interface{}{
		"code": 0,
		"data": id,
	}
	this.ServeJson()
}
