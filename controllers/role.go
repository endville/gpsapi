package controllers

import (
	"encoding/json"
	"gpsapi/models"
)

// 角色相关
type RoleController struct {
	BaseController
}

// @Title 添加一个新的角色
// @Description 添加一个新的角色
// @Param	body		body 	models.Role	true		"角色信息"
// @Success 200 {int} models.Role.Id
// @Failure 403 body is empty
// @router / [post]
func (this *RoleController) Post() {
	var role models.Role
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &role)
	if err != nil {
		this.ResponseErrorJSON(400, errorFormat(ErrorBadJson_400, err.Error()))
	} else {
		if id, err := models.AddRole(role); err != nil {
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

// @Title 获取角色列表
// @Description 获取角色列表
// @Success 200 {object} models.Role
// @Failure 400 请求的参数不正确
// @router / [get]
func (this *RoleController) GetAll() {
	roles, total, err := models.GetAllRoles(nil, 1, 30)
	if err != nil {
		this.ResponseErrorJSON(400, errorFormat(ErrorBadRequest_400, err.Error()))
	}
	this.Data["json"] = map[string]interface{}{
		"code":  0,
		"data":  roles,
		"total": total,
	}
	this.ServeJson()
}

// @Title 根据角色唯一ID号获取角色信息
// @Description 根据角色唯一ID号获取角色信息
// @Param	id		path 	int	true		"角色唯一ID号"
// @Success 200 {object} models.Role
// @Failure 404 未找到对应的记录
// @router /:id [get]
func (this *RoleController) Get() {
	id, _ := this.GetInt64(":id")
	role, err := models.GetRole(id)
	if err != nil {
		this.ResponseErrorJSON(404, errorFormat(ErrorDataNotFound_404, err.Error()))
	} else {
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"data": role,
		}
	}
	this.ServeJson()
}

// @Title 更新角色信息
// @Description 更新角色信息
// @Param	id		path 	int	true		"角色唯一ID"
// @Param	body		body 	models.Role	true		"角色信息"
// @Success 200 {object} models.Role
// @Failure 400 请求的参数不正确
// @router /:id [put]
func (this *RoleController) Put() {
	id, _ := this.GetInt64(":id")
	var role models.Role
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &role); err != nil {
		this.ResponseErrorJSON(400, errorFormat(ErrorBadJson_400, err.Error()))
	} else {
		update, err := models.UpdateRole(id, &role)
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

// @Title 删除角色
// @Description 删除角色
// @Param	id		path 	int	true		"角色唯一ID"
// @Success 200 {int} 返回记录ID
// @Failure 404 未找到对应的记录
// @router /:id [delete]
func (this *RoleController) Delete() {
	id, _ := this.GetInt64(":id")
	if err := models.DeleteRole(id); err != nil {
		this.ResponseErrorJSON(404, errorFormat(ErrorDataNotFound_404, err.Error()))
	}
	this.Data["json"] = map[string]interface{}{
		"code": 0,
		"data": id,
	}
	this.ServeJson()
}
