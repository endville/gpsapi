package controllers

import (
	"encoding/json"
	"gpsapi/models"
	"gpsapi/tools"

	"github.com/astaxie/beego/orm"
)

// 车队相关
type GroupController struct {
	BaseController
}

// @Title Post
// @Description 添加一个新的车队
// @Param	body		body 	models.Group	true		"车队信息"
// @Success 200 {int} models.Group.Id
// @Failure 403 body is empty
// @router / [post]
func (this *GroupController) Post() {
	var group models.Group
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &group)
	if err != nil {
		this.ResponseErrorJSON(400, errorFormat(ErrorBadJson_400, err.Error()))
	} else {
		group.Password = tools.MD5(group.Password)
		if id, err := models.AddGroup(group); err != nil {
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

// @Title 查询接口
// @Description 根据关键字查询车队
// @Param	key	query	string	false	"关键字"
// @Success 200 {object} models.GroupSearchModel
// @Failure 400 请求的参数不正确
// @router /search [get]
func (this *GroupController) Search() {
	key := this.GetString("key", "")

	if len(key) < 3 {
		this.Data["json"] = map[string]interface{}{
			"code": 1,
			"msg":  "关键字过短",
		}
	} else {
		groups := models.SearchGroup("%"+key+"%", 25)
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"data": groups,
		}
	}

	this.ServeJson()
}

// @Title 获取车队列表
// @Description 获取车队列表
// @Param	pid	query	int	false	"父车队ID"
// @Param	pageIndex	query	int	false	"页码, 默认1"
// @Param	pageSize	query	int	false	"每页显示条数, 默认30"
// @Success 200 {object} models.Group
// @Failure 400 请求的参数不正确
// @router / [get]
func (this *GroupController) GetAll() {
	cond := orm.NewCondition()
	if pid, _ := this.GetInt64("pid", -1); pid != -1 {
		cond = cond.And("ParentId", pid)
	}
	pageIndex, _ := this.GetInt("pageIndex", 1)
	pageSize, _ := this.GetInt("pageSize", 30)
	groups, total, err := models.GetAllGroups(cond, pageIndex, pageSize)
	if err != nil {
		this.ResponseErrorJSON(400, errorFormat(ErrorBadRequest_400, err.Error()))
	}
	this.Data["json"] = map[string]interface{}{
		"code":  0,
		"data":  groups,
		"total": total,
	}
	this.ServeJson()
}

// @Title 根据车队唯一ID号获取车队信息
// @Description 根据车队唯一ID号获取车队信息
// @Param	id		path 	int	true		"车队唯一ID号"
// @Success 200 {object} models.Group
// @Failure 404 未找到对应的记录
// @router /:id [get]
func (this *GroupController) Get() {
	id, _ := this.GetInt64(":id")
	group, err := models.GetGroup(id)
	if err != nil {
		this.ResponseErrorJSON(404, errorFormat(ErrorDataNotFound_404, err.Error()))
	} else {
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"data": group,
		}
	}
	this.ServeJson()
}

// @Title Get Profile
// @Description 根据车队唯一ID号获取车队Profile信息
// @Param	id		path 	int	true		"车队唯一ID号(注意不是Profile的ID)"
// @Success 200 {object} models.GroupProfile
// @Failure 404 未找到对应的记录
// @router /:id/profile [get]
func (this *GroupController) GetProfile() {
	id, _ := this.GetInt64(":id")
	profile, err := models.GetGroupProfile(id)
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

// @Title 获取车队的终端设备列表
// @Description 获取车队的终端设备列表
// @Param	id		path 	int	true		车队唯一ID号
// @Success 200 {object} models.Terminal
// @Failure 400 请求的参数不正确
// @router /:id/terminal [get]
func (this *GroupController) GetTerminals() {
	id, _ := this.GetInt64(":id")
	cond := orm.NewCondition().And("group__id", id)
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

// @Title 获取车队的角色列表
// @Description 获取车队的角色列表
// @Param	id		path 	int	true		车队唯一ID号
// @Success 200 {object} models.Role
// @Failure 400 请求的参数不正确
// @router /:id/roles [get]
func (this *GroupController) GetRoles() {
	id, _ := this.GetInt64(":id")
	list, err := models.GetGroupRoles(id)
	if err != nil {
		this.ResponseErrorJSON(400, errorFormat(ErrorBadRequest_400, err.Error()))
	}
	this.Data["json"] = map[string]interface{}{
		"code": 0,
		"data": list,
	}
	this.ServeJson()
}

// @Title update
// @Description 更新车队信息
// @Param	id		path 	int	true		"车队唯一ID"
// @Param	body		body 	models.Group	true		"车队信息"
// @Success 200 {object} models.Group
// @Failure 400 请求的参数不正确
// @router /:id [put]
func (this *GroupController) Put() {
	id, _ := this.GetInt64(":id")
	var group models.Group
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &group); err != nil {
		this.ResponseErrorJSON(400, errorFormat(ErrorBadJson_400, err.Error()))
	} else {
		if group.Password != "" {
			group.Password = tools.MD5(group.Password)
		}
		update, err := models.UpdateGroup(id, &group)
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

// @Title 更新车队Profile信息
// @Description 更新车队Profile信息
// @Param	id		path 	int	true		"车队唯一ID(注意不是Profile的ID)"
// @Param	body		body 	models.GroupProfile	true		"车队信息"
// @Success 200 {object} models.GroupProfile
// @Failure 400 请求的参数不正确
// @router /:id/profile [put]
func (this *GroupController) PutProfile() {
	id, _ := this.GetInt64(":id")
	var profile models.GroupProfile
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &profile); err != nil {
		this.ResponseErrorJSON(400, errorFormat(ErrorBadJson_400, err.Error()))
	} else {
		update, err := models.UpdateGroupProfile(id, &profile)
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

// @Title 更新车队Roles信息
// @Description 更新车队Roles信息
// @Param	id		path 	int	true		"车队唯一ID"
// @Param	roles	body 	models.Role	true		"角色列表，如[{'id':1},{'id':2}]"
// @Success 200 {object} models.GroupProfile
// @Failure 400 请求的参数不正确
// @router /:id/role [put]
func (this *GroupController) PutRole() {
	id, _ := this.GetInt64(":id")
	var roles []*models.Role
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &roles); err != nil {
		this.ResponseErrorJSON(400, errorFormat(ErrorBadJson_400, err.Error()))
	} else {
		update, err := models.UpdateGroupRoles(id, roles)
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

// @Title delete
// @Description 删除车队
// @Param	id		path 	int	true		"车队唯一ID"
// @Success 200 {int} 返回记录ID
// @Failure 404 未找到对应的记录
// @router /:id [delete]
func (this *GroupController) Delete() {
	id, _ := this.GetInt64(":id")
	if err := models.DeleteGroup(id); err != nil {
		this.ResponseErrorJSON(404, errorFormat(ErrorDataNotFound_404, err.Error()))
	}
	this.Data["json"] = map[string]interface{}{
		"code": 0,
		"data": id,
	}
	this.ServeJson()
}
