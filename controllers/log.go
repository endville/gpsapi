package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"gpsapi/models"
)

// 日志相关
type LogController struct {
	BaseController
}

// @Title 添加一个新的日志
// @Description 添加一个新的日志
// @Param	body	body	models.Log	true		"日志信息"
// @Success 200 {int} models.Log.ID
// @Failure 400 请求的参数不正确
// @router / [post]
func (this *LogController) Post() {
	var log models.Log
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &log); err != nil {
		this.ResponseErrorJSON(400, errorFormat(ErrorBadJson_400, err.Error()))
	} else {
		if id, err := models.AddLog(log); err != nil {
			this.ResponseErrorJSON(400, errorFormat(ErrorBadRequest_400, err.Error()))
		} else {
			this.Data["json"] = map[string]interface{}{
				"code": 0,
				"data": id,
			}
		}
	}
	this.ServeJson()
}

// @Title 获取日志列表
// @Description 获取日志列表
// @Param	level	query	int	false	"日志级别"
// @Param	type	query	int	false	"日志类型"
// @Param	by	query	string	false	"记录产生者"
// @Param	on_begin	query	string	false	"日志记录起始时间"
// @Param	on_end	query	string	false	"日志记录终止时间"
// @Param	pageIndex	query	int	false	"页码, 默认1"
// @Param	pageSize	query	int	false	"每页显示条数, 默认30"
// @Success 200 {object} models.Log
// @Failure 400 请求的参数不正确
// @router / [get]
func (this *LogController) GetAll() {
	cond := orm.NewCondition()
	if level, _ := this.GetInt64("level", -1); level != -1 {
		cond = cond.And("Level", level)
	}
	if typi, _ := this.GetInt64("type", -1); typi != -1 {
		cond = cond.And("Type", typi)
	}
	if by := this.GetString("by"); by != "" {
		cond = cond.And("LogBy", by)
	}
	if on_begin := this.GetString("on_begin"); on_begin != "" {
		cond = cond.And("LogOn__gt", on_begin)
	}
	if on_end := this.GetString("on_end"); on_end != "" {
		cond = cond.And("LogOn__lt", on_end)
	}
	pageIndex, _ := this.GetInt("pageIndex", 1)
	pageSize, _ := this.GetInt("pageSize", 30)
	logs, total, err := models.GetAllLogs(cond, pageIndex, pageSize)
	if err != nil {
		this.ResponseErrorJSON(400, errorFormat(ErrorBadRequest_400, err.Error()))
	}
	this.Data["json"] = map[string]interface{}{
		"code":  0,
		"data":  logs,
		"total": total,
	}
	this.ServeJson()
}

// @Title 根据日志唯一ID号获取日志信息
// @Description 根据日志唯一ID号获取日志信息
// @Param	id		path 	int	true		"日志唯一ID号"
// @Success 200 {object} models.Log
// @Failure 404 未找到对应的记录
// @router /:id [get]
func (this *LogController) Get() {
	id, _ := this.GetInt64(":id")
	if id != 0 {
		log, err := models.GetLog(id)
		if err != nil {
			this.ResponseErrorJSON(404, errorFormat(ErrorDataNotFound_404, err.Error()))
		} else {
			this.Data["json"] = map[string]interface{}{
				"code": 0,
				"data": log,
			}
		}
	}
	this.ServeJson()
}

// @Title 删除日志
// @Description 删除日志
// @Param	id		path 	int	true		"日志唯一ID"
// @Success 200 {int} 返回记录ID
// @Failure 404 未找到对应的记录
// @router /:id [delete]
func (this *LogController) Delete() {
	id, _ := this.GetInt64(":id")
	if err := models.DeleteLog(id); err != nil {
		this.ResponseErrorJSON(404, errorFormat(ErrorDataNotFound_404, err.Error()))
	}
	this.Data["json"] = map[string]interface{}{
		"code": 0,
		"data": id,
	}
	this.ServeJson()
}
