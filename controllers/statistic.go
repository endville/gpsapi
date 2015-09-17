package controllers

import (
	// "encoding/json"
	// "github.com/astaxie/beego/orm"
	"fmt"
	"gpsapi/models"
	"time"
)

// 统计相关
type StatisticController struct {
	BaseController
}

// @Title 设备服务到期情况及状态
// @Description 设备服务到期情况及状态
// @Param 	type 		query 	int true	"1:全部 2:行驶设备 3:静止设备 4:在线设备 5:离线设备 6:离线60天 7:7天过期设备 8:60天过期 9:已过期设备"
// @Param	pageIndex	query	int	false	"页码, 默认1"
// @Param	pageSize	query	int	false	"每页显示条数, 默认30"
// @Success 200 {object} models.StatisticTerminalServiceExpire
// @Failure 400 请求的参数不正确
// @router /serviceExpire [get]
func (this *StatisticController) Terminal() {
	o := models.GetOrm()

	const (
		STATISTIC_DEVICE_TYPE_ALL = iota + 1
		STATISTIC_DEVICE_TYPE_MOVING
		STATISTIC_DEVICE_TYPE_STAY
		STATISTIC_DEVICE_TYPE_ONLINE
		STATISTIC_DEVICE_TYPE_OFFLINE
		STATISTIC_DEVICE_TYPE_OFFLINE_60
		STATISTIC_DEVICE_TYPE_EXPIRE_7
		STATISTIC_DEVICE_TYPE_EXPIRE_60
		STATISTIC_DEVICE_TYPE_EXPIRE_ALREADY
	)

	var data []models.StatisticTerminalServiceExpire

	pageIndex, _ := this.GetInt("pageIndex", 1)
	pageSize, _ := this.GetInt("pageSize", 30)
	statisticType, _ := this.GetInt("type", 1)
	where := ""

	switch statisticType {
	case STATISTIC_DEVICE_TYPE_ALL:
		where = ""
	case STATISTIC_DEVICE_TYPE_MOVING:
		where = ""
	case STATISTIC_DEVICE_TYPE_STAY:
		where = ""
	case STATISTIC_DEVICE_TYPE_ONLINE:
		where = "where t.online_on >= t.offline_on"
	case STATISTIC_DEVICE_TYPE_OFFLINE:
		where = "where t.online_on < t.offline_on"
	case STATISTIC_DEVICE_TYPE_OFFLINE_60:
		where = "where t.online_on < t.offline_on and t.offline_on < ADDDate(now(),-60)"
	case STATISTIC_DEVICE_TYPE_EXPIRE_7:
		where = "where p.expire_on > now() and p.expire_on < ADDDATE(now(),-7)"
	case STATISTIC_DEVICE_TYPE_EXPIRE_60:
		where = "where p.expire_on > now() and p.expire_on < ADDDATE(NOW(),-60)"
	case STATISTIC_DEVICE_TYPE_EXPIRE_ALREADY:
		where = "where p.expire_on < now()"
	default:
		break
	}

	sql := fmt.Sprintf("select t.id,p.imei as name,t.terminal_sn,p.activate_on,p.expire_on,t.offline_on,t.online_on from terminal t left join terminal_profile p on t.terminal_profile_id = p.id %s limit %d,%d", where, (pageIndex-1)*pageSize, pageSize)
	sqlTotal := fmt.Sprintf("select count(*) from terminal t left join terminal_profile p on t.terminal_profile_id = p.id %s", where)

	_, err := o.Raw(sql).QueryRows(&data)
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"code": 1,
			"msg":  err.Error(),
		}
	} else {
		var total int64
		o.Raw(sqlTotal).QueryRow(&total)
		for i, v := range data {
			data[i].ActivateOnTs = v.ActivateOn.Unix()
			data[i].ExpireOnTs = v.ExpireOn.Unix()
			data[i].OfflineOnTs = v.OfflineOn.Unix()
			data[i].OnlineOnTs = v.OnlineOn.Unix()
			data[i].State = data[i].OfflineOnTs > data[i].OnlineOnTs
		}
		this.Data["json"] = map[string]interface{}{
			"code":  0,
			"total": total,
			"data":  data,
		}
	}
	this.ServeJson()
}

// @Title 设备运行总览
// @Description 设备运行总览
// @Param	groupId		query 	int	true	"车队ID"
// @Param	terminalId	query 	int	true	"终端ID"
// @Param	timeBegin	query 	string	false	"起始时间（默认昨天的现在时间点）"
// @Param	timeEnd		query 	string	false	"终止时间（默认现在）"
// @Success 200 {object} models.StatisticTerminalRunInfo
// @Failure 400 请求的参数不正确
// @router /terminalRunInfo [get]
func (this *StatisticController) TerminalRunInfo() {
	o := models.GetOrm()
	gid, _ := this.GetInt("groupId")
	tid, _ := this.GetInt("terminalId")
	timeBegin := this.GetString("timeBegin", time.Now().AddDate(0, 0, -1).Format("2006-01-02 15:04:05"))
	timeEnd := this.GetString("timeEnd", time.Now().Format("2006-01-02 15:04:05"))

	sql := ""

	var data []models.StatisticTerminalRunInfo
	_, err := o.Raw(sql).QueryRows(&data)
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"code": 1,
			"msg":  err.Error(),
		}
	} else {
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"data": data,
		}
	}
	this.ServeJson()
}

// @Title 设备里程统计
// @Description 设备里程统计
// @Param	groupId		query 	int	true	"车队ID"
// @Param	terminalId	query 	int	true	"终端ID"
// @Param	timeBegin	query 	string	false	"起始时间（默认昨天的现在时间点）"
// @Param	timeEnd		query 	string	false	"终止时间（默认现在）"
// @Success 200 {object} models.StatisticTerminalMileage
// @Failure 400 请求的参数不正确
// @router /terminalMileage [get]
func (this *StatisticController) TerminalMileage() {
	o := models.GetOrm()
	gid, _ := this.GetInt("groupId")
	tid, _ := this.GetInt("terminalId")
	timeBegin := this.GetString("timeBegin", time.Now().AddDate(0, 0, -1).Format("2006-01-02 15:04:05"))
	timeEnd := this.GetString("timeEnd", time.Now().Format("2006-01-02 15:04:05"))

	sql := ""

	var data []models.StatisticTerminalMileage
	_, err := o.Raw(sql).QueryRows(&data)
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"code": 1,
			"msg":  err.Error(),
		}
	} else {
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"data": data,
		}
	}
	this.ServeJson()
}

// @Title 设备停留详情
// @Description 设备停留详情
// @Param	groupId		query 	int	true	"车队ID"
// @Param	terminalId	query 	int	true	"终端ID"
// @Param	timeBegin	query 	string	false	"起始时间（默认昨天的现在时间点）"
// @Param	timeEnd		query 	string	false	"终止时间（默认现在）"
// @Param	minTime		query 	int	false	"最短停留时间 (单位分钟，默认10分钟)"
// @Success 200 {object} models.StatisticTerminalStay
// @Failure 400 请求的参数不正确
// @router /terminalStay [get]
func (this *StatisticController) TerminalStayInfo() {
	o := models.GetOrm()
	gid, _ := this.GetInt("groupId")
	tid, _ := this.GetInt("terminalId")
	timeBegin := this.GetString("timeBegin", time.Now().AddDate(0, 0, -1).Format("2006-01-02 15:04:05"))
	timeEnd := this.GetString("timeEnd", time.Now().Format("2006-01-02 15:04:05"))
	minTime, _ := this.GetInt("minTime", 10)
	sql := ""

	var data []models.StatisticTerminalStay
	_, err := o.Raw(sql).QueryRows(&data)
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"code": 1,
			"msg":  err.Error(),
		}
	} else {
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"data": data,
		}
	}
	this.ServeJson()
}

// @Title 设备警告详情
// @Description 设备警告详情
// @Param	groupId		query 	int	true	"车队ID"
// @Param	terminalId	query 	int	true	"终端ID"
// @Param	timeBegin	query 	string	false	"起始时间（默认昨天的现在时间点）"
// @Param	timeEnd		query 	string	false	"终止时间（默认现在）"
// @Param	type		query 	int	true	"报警类型"
// @Success 200 {object} models.StatisticTerminalWarning
// @Failure 400 请求的参数不正确
// @router /terminalWarning [get]
func (this *StatisticController) TerminalWarning() {
	o := models.GetOrm()
	gid, _ := this.GetInt("groupId")
	tid, _ := this.GetInt("terminalId")
	timeBegin := this.GetString("timeBegin", time.Now().AddDate(0, 0, -1).Format("2006-01-02 15:04:05"))
	timeEnd := this.GetString("timeEnd", time.Now().Format("2006-01-02 15:04:05"))
	warningType, _ := this.GetInt("type", 0)

	switch warningType {
	default:
		break
	}

	sql := ""

	var data []models.StatisticTerminalWarning
	_, err := o.Raw(sql).QueryRows(&data)
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"code": 1,
			"msg":  err.Error(),
		}
	} else {
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"data": data,
		}
	}
	this.ServeJson()
}

// @Title 设备上线统计
// @Description 设备上线统计
// @Param	groupId		query 	int	true	"车队ID"
// @Param	timeBegin	query 	string	false	"起始时间（默认昨天的现在时间点）"
// @Param	timeEnd		query 	string	false	"终止时间（默认现在）"
// @Success 200 {object} models.StatisticTerminalOnline
// @Failure 400 请求的参数不正确
// @router /terminalOnline [get]
func (this *StatisticController) TerminalOnline() {
	o := models.GetOrm()
	gid, _ := this.GetInt("groupId")
	timeBegin := this.GetString("timeBegin", time.Now().AddDate(0, 0, -1).Format("2006-01-02 15:04:05"))
	timeEnd := this.GetString("timeEnd", time.Now().Format("2006-01-02 15:04:05"))

	sql := ""

	var data []models.StatisticTerminalOnline
	_, err := o.Raw(sql).QueryRows(&data)
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"code": 1,
			"msg":  err.Error(),
		}
	} else {
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"data": data,
		}
	}
	this.ServeJson()
}

// @Title 设备上线统计详情
// @Description 设备上线统计详情
// @Param	groupId		query 	int	true	"车队ID"
// @Param	timeBegin	query 	string	false	"起始时间（默认昨天的现在时间点）"
// @Param	timeEnd		query 	string	false	"终止时间（默认现在）"
// @Success 200 {object} models.StatisticTerminalOnlineDetails
// @Failure 400 请求的参数不正确
// @router /terminalOnlineDetails [get]
func (this *StatisticController) TerminalOnlineDetails() {
	o := models.GetOrm()
	gid, _ := this.GetInt("groupId")
	timeBegin := this.GetString("timeBegin", time.Now().AddDate(0, 0, -1).Format("2006-01-02 15:04:05"))
	timeEnd := this.GetString("timeEnd", time.Now().Format("2006-01-02 15:04:05"))

	sql := ""

	var data []models.StatisticTerminalOnlineDetails
	_, err := o.Raw(sql).QueryRows(&data)
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"code": 1,
			"msg":  err.Error(),
		}
	} else {
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"data": data,
		}
	}
	this.ServeJson()
}

// @Title 设备离线统计
// @Description 设备运行总览
// @Param	groupId		query 	int	true	"车队ID"
// @Param	terminalId	query 	int	true	"终端ID"
// @Param	timeBegin	query 	string	false	"起始时间（默认昨天的现在时间点）"
// @Param	timeEnd		query 	string	false	"终止时间（默认现在）"
// @Success 200 {object} models.StatisticTerminalOffline
// @Failure 400 请求的参数不正确
// @router /terminalOffline [get]
func (this *StatisticController) TerminalOffline() {
	o := models.GetOrm()
	gid, _ := this.GetInt("groupId")
	tid, _ := this.GetInt("terminalId")
	timeBegin := this.GetString("timeBegin", time.Now().AddDate(0, 0, -1).Format("2006-01-02 15:04:05"))
	timeEnd := this.GetString("timeEnd", time.Now().Format("2006-01-02 15:04:05"))

	sql := ""

	var data []models.StatisticTerminalOffline
	_, err := o.Raw(sql).QueryRows(&data)
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"code": 1,
			"msg":  err.Error(),
		}
	} else {
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"data": data,
		}
	}
	this.ServeJson()
}

// @Title 设备查询
// @Description 设备运行总览
// @Param	groupId		query 	int	true	"车队ID"
// @Param	terminalId	query 	int	true	"终端ID"
// @Success 200 {object} models.StatisticTerminalQuery
// @Failure 400 请求的参数不正确
// @router /terminalInfo [get]
func (this *StatisticController) TerminalQuery() {
	o := models.GetOrm()
	gid, _ := this.GetInt("groupId")
	tid, _ := this.GetInt("terminalId")

	sql := ""

	var data []models.StatisticTerminalQuery
	_, err := o.Raw(sql).QueryRows(&data)
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"code": 1,
			"msg":  err.Error(),
		}
	} else {
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"data": data,
		}
	}
	this.ServeJson()
}
