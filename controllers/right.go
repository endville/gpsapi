package controllers

import (
	"gpsapi/models"
)

// 权限相关
type RightController struct {
	BaseController
}

// @Title 获取权限列表
// @Description 获取权限列表
// @Success 200 {object} models.Right
// @Failure 400 请求的参数不正确
// @router / [get]
func (this *RightController) GetAll() {
	roles, total, err := models.GetAllRights(nil, 1, 30)
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
