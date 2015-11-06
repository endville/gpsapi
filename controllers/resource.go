package controllers

import (
	"gpsapi/models"
)

// 资源相关相关
type ResourceController struct {
	BaseController
}

// @Title 查询接口
// @Description 根据关键字查询车队、用户、终端
// @Param	key	query	string	false	"关键字"
// @Success 200 {object} models.GroupSearchModel
// @Failure 400 请求的参数不正确
// @router /search [get]
func (this *ResourceController) Search() {
	key := this.GetString("key", "")

	if len(key) < 3 {
		this.Data["json"] = map[string]interface{}{
			"code": 1,
			"msg":  "关键字过短",
		}
	} else {
		list := make([]interface{}, 0)
		groups := models.SearchGroup("%"+key+"%", 5)
		for _, v := range *groups {
			v.SearchType = "group"
			list = append(list, v)
		}
		users := models.SearchUser("%"+key+"%", 5)
		for _, v := range *users {
			v.SearchType = "user"
			list = append(list, v)
		}
		terminals := models.SearchTerminal("%"+key+"%", 5)
		for _, v := range *terminals {
			v.SearchType = "terminal"
			list = append(list, v)
		}
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"data": list,
		}
	}

	this.ServeJson()
}
