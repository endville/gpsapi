package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"gpsapi/models"
)

// 终端相关
type TerminalController struct {
	BaseController
}

// @Title 添加一个新的终端
// @Description 添加一个新的终端
// @Param	body		body 	models.Terminal	true		"终端信息"
// @Success 200 {int} models.Terminal.Id
// @Failure 400	请求的参数不正确
// @router / [post]
func (this *TerminalController) Post() {
	var terminal models.Terminal
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &terminal)
	if err != nil {
		this.ResponseErrorJSON(400, errorFormat(ErrorBadJson_400, err.Error()))
	} else {
		if id, err := models.AddTerminal(terminal); err != nil {
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

// @Title 获取终端列表
// @Description 获取终端列表
// @Param	uid	query	int	false	"用户ID"
// @Param	gid	query	int	false	"车队ID"
// @Param 	terminalSn	query string false "设备终端号"
// @Param	pageIndex	query	int	false	"页码, 默认1"
// @Param	pageSize	query	int	false	"每页显示条数, 默认30"
// @Success 200 {object} models.Terminal
// @Failure 400 请求的参数不正确
// @router / [get]
func (this *TerminalController) GetAll() {
	cond := orm.NewCondition()
	if uid, _ := this.GetInt64("uid", -1); uid != -1 {
		cond = cond.And("user_id", uid)
	}
	if gid, _ := this.GetInt64("gid", -1); gid != -1 {
		cond = cond.And("group_id", gid)
	}
	if terminalSn := this.GetString("terminalSn"); terminalSn != "" {
		cond = cond.And("terminal_sn", terminalSn)
	}
	pageIndex, _ := this.GetInt("pageIndex", 1)
	pageSize, _ := this.GetInt("pageSize", 30)
	terminals, total, err := models.GetAllTerminals(cond, pageIndex, pageSize)
	if err != nil {
		this.ResponseErrorJSON(400, errorFormat(ErrorBadRequest_400, err.Error()))
	}
	this.Data["json"] = map[string]interface{}{
		"code":  0,
		"data":  terminals,
		"total": total,
	}
	this.ServeJson()
}

// @Title 根据终端唯一ID号获取终端信息
// @Description 根据终端唯一ID号获取终端信息
// @Param	id		path 	int	true		"终端唯一ID号"
// @Success 200 {object} models.Terminal
// @Failure 404 未找到对应的记录
// @router /:id [get]
func (this *TerminalController) Get() {
	id, _ := this.GetInt64(":id")
	if id != 0 {
		terminal, err := models.GetTerminal(id)
		if err != nil {
			this.ResponseErrorJSON(403, err.Error())
		} else {
			this.Data["json"] = map[string]interface{}{
				"code": 0,
				"data": terminal,
			}
		}
	}
	this.ServeJson()
}

// @Title 根据终端唯一ID号获取终端Profile信息
// @Description 根据终端唯一ID号获取终端Profile信息
// @Param	id		path 	int	true		"终端唯一ID号(并非Profile的ID)"
// @Success 200 {object} models.TerminalProfile
// @Failure 404 未找到对应的记录
// @router /:id/profile [get]
func (this *TerminalController) GetProfile() {
	id, _ := this.GetInt64(":id")
	if id != 0 {
		profile, err := models.GetTerminalProfile(id)
		if err != nil {
			this.ResponseErrorJSON(403, err.Error())
		} else {
			this.Data["json"] = map[string]interface{}{
				"code": 0,
				"data": profile,
			}
		}
	}
	this.ServeJson()
}

// @Title 根据终端唯一ID号获取终端载具信息
// @Description 根据终端唯一ID号获取终端载具信息
// @Param	id		path 	int	true		"终端唯一ID号(并非载具的ID)"
// @Success 200 {object} models.TerminalCarrier
// @Failure 404 未找到对应的记录
// @router /:id/carrier [get]
func (this *TerminalController) GetCarrier() {
	id, _ := this.GetInt64(":id")
	if id != 0 {
		carrier, err := models.GetTerminalCarrier(id)
		if err != nil {
			this.ResponseErrorJSON(403, err.Error())
		} else {
			this.Data["json"] = map[string]interface{}{
				"code": 0,
				"data": carrier,
			}
		}
	}
	this.ServeJson()
}

// @Title 更新终端信息
// @Description 更新终端信息
// @Param	id		path 	int	true		"终端唯一ID"
// @Param	body		body 	models.Terminal	true		"终端信息"
// @Success 200 {object} models.Terminal
// @Failure 400 请求的参数不正确
// @router /:id [put]
func (this *TerminalController) Put() {
	id, _ := this.GetInt64(":id")
	if id != 0 {
		var terminal models.Terminal
		if err := json.Unmarshal(this.Ctx.Input.RequestBody, &terminal); err != nil {
			this.ResponseErrorJSON(400, errorFormat(ErrorBadJson_400, err.Error()))
		} else {
			update, err := models.UpdateTerminal(id, &terminal)
			if err != nil {
				this.ResponseErrorJSON(403, err.Error())
			} else {
				this.Data["json"] = map[string]interface{}{
					"code": 0,
					"data": update,
				}
			}
		}

	}
	this.ServeJson()
}

// @Title 更新终端Profile信息
// @Description 更新终端Profile信息
// @Param	id		path 	int	true		"终端唯一ID(并非Profile的ID)"
// @Param	body		body 	models.TerminalProfile	true		"终端Profile信息"
// @Success 200 {object} models.TerminalProfile
// @Failure 400 请求的参数不正确
// @router /:id/profile [put]
func (this *TerminalController) PutProfile() {
	id, _ := this.GetInt64(":id")
	if id != 0 {
		var profile models.TerminalProfile
		if err := json.Unmarshal(this.Ctx.Input.RequestBody, &profile); err != nil {
			this.ResponseErrorJSON(400, errorFormat(ErrorBadJson_400, err.Error()))
		} else {
			update, err := models.UpdateTerminalProfile(id, &profile)
			if err != nil {
				this.ResponseErrorJSON(400, errorFormat(ErrorBadParam_400, err.Error()))
			} else {
				this.Data["json"] = map[string]interface{}{
					"code": 0,
					"data": update,
				}
			}
		}
	}
	this.ServeJson()
}

// @Title 更新终端载具信息
// @Description 更新终端载具信息
// @Param	id		path 	int	true		"终端唯一ID(并非载具ID)"
// @Param	body		body 	models.TerminalCarrier	true		"终端载具信息"
// @Success 200 {object} models.TerminalCarrier
// @Failure 400 请求的参数不正确
// @router /:id/carrier [put]
func (this *TerminalController) PutCarrier() {
	id, _ := this.GetInt64(":id")
	if id != 0 {
		var carrier models.TerminalCarrier
		if err := json.Unmarshal(this.Ctx.Input.RequestBody, &carrier); err != nil {
			this.ResponseErrorJSON(400, errorFormat(ErrorBadJson_400, err.Error()))
		} else {
			update, err := models.UpdateTerminalCarrier(id, &carrier)
			if err != nil {
				this.ResponseErrorJSON(400, errorFormat(ErrorBadParam_400, err.Error()))
			} else {
				this.Data["json"] = map[string]interface{}{
					"code": 0,
					"data": update,
				}
			}
		}
	}
	this.ServeJson()
}

// @Title 删除终端
// @Description 删除终端
// @Param	id		path 	int	true		"终端唯一ID"
// @Success 200 {string} 返回记录ID
// @Failure 404 未找到对应的记录
// @router /:id [delete]
func (this *TerminalController) Delete() {
	id, _ := this.GetInt64(":id")
	if err := models.DeleteTerminal(id); err != nil {
		this.ResponseErrorJSON(403, err.Error())
	}
	this.Data["json"] = map[string]interface{}{
		"code": 0,
		"data": id,
	}
	this.ServeJson()
}
