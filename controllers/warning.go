package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"gpsapi/models"
)

// 警告相关
type WarningController struct {
	BaseController
}

// @Title 获取警告列表
// @Description 获取警告列表
// @Param	tid	query	int	false	"终端ID"
// @Param	uid	query	int	false	"用户ID"
// @Param	gid	query	int	false	"车队ID"
// @Param	time_begin	query	string	false	"警告记录起始时间"
// @Param	time_end	query	string	false	"警告记录终止时间"
// @Param	pageIndex	query	int	false	"页码, 默认1"
// @Param	pageSize	query	int	false	"每页显示条数, 默认30"
// @Success 200 {object} models.Warning
// @Failure 400 请求的参数不正确
// @router / [get]
func (this *WarningController) GetAll() {
	cond := orm.NewCondition()
	if tid, _ := this.GetInt64("tid", -1); tid != -1 {
		cond = cond.And("TerminalId", tid)
	}
	if uid, _ := this.GetInt64("uid", -1); uid != -1 {
		cond = cond.And("UserId", uid)
	}
	if gid, _ := this.GetInt64("gid", -1); gid != -1 {
		cond = cond.And("GroupId", gid)
	}
	if on_begin := this.GetString("time_begin"); on_begin != "" {
		cond = cond.And("CreateOn__gt", on_begin)
	}
	if on_end := this.GetString("time_end"); on_end != "" {
		cond = cond.And("CreateOn__lt", on_end)
	}
	pageIndex, _ := this.GetInt("pageIndex", 1)
	pageSize, _ := this.GetInt("pageSize", 30)
	warnings, total, err := models.GetAllWarnings(cond, pageIndex, pageSize)
	if err != nil {
		this.ResponseErrorJSON(400, errorFormat(ErrorBadRequest_400, err.Error()))
	}
	this.Data["json"] = map[string]interface{}{
		"code":  0,
		"data":  warnings,
		"total": total,
	}
	this.ServeJson()
}

// @Title 根据警告唯一ID号获取警告信息
// @Description 根据警告唯一ID号获取警告信息
// @Param	id		path 	int	true		"警告唯一ID号"
// @Success 200 {object} models.Warning
// @Failure 404 未找到对应的记录
// @router /:id [get]
func (this *WarningController) Get() {
	id, _ := this.GetInt64(":id")
	warning, err := models.GetWarning(id)
	if err != nil {
		this.ResponseErrorJSON(404, errorFormat(ErrorDataNotFound_404, err.Error()))
	} else {
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"data": warning,
		}
	}
	this.ServeJson()
}

// @Title 更新警告信息
// @Description 更新警告信息
// @Param	id		path 	int	true		"警告唯一ID"
// @Param	body		body 	models.Warning	true		"警告信息"
// @Success 200 {object} models.Warning
// @Failure 400 请求的参数不正确
// @router /:id [put]
func (this *WarningController) Put() {
	id, _ := this.GetInt64(":id")
	var warning models.Warning
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &warning); err != nil {
		this.ResponseErrorJSON(400, errorFormat(ErrorBadJson_400, err.Error()))
	} else {
		update, err := models.UpdateWarning(id, &warning)
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

// @Title 删除警告
// @Description 删除警告
// @Param	id		path 	int	true		"警告唯一ID"
// @Success 200 {int} 返回记录ID
// @Failure 404 未找到对应的记录
// @router /:id [delete]
func (this *WarningController) Delete() {
	id, _ := this.GetInt64(":id")
	if err := models.DeleteWarning(id); err != nil {
		this.ResponseErrorJSON(404, errorFormat(ErrorDataNotFound_404, err.Error()))
	}
	this.Data["json"] = map[string]interface{}{
		"code": 0,
		"data": id,
	}
	this.ServeJson()
}
