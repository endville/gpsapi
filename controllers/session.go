package controllers

import (
	rpc "gpsapi/rpcClient"
)

// 网关Session监控相关
type SessionController struct {
	BaseController
}

// @Title Get
// @Description 获取某个终端设备的session
// @Param	terminalSN		path 	string	true		"终端sn号"
// @Success 200 {object} rpcClient.RPCSessionModel
// @Failure 403 :gate is out of server
// @router /:terminalSN [get]
func (this *SessionController) Get() {
	sn := this.GetString(":terminalSN")
	if sn != "" {
		session := new(rpc.RPCSessionModel)
		if err := rpc.GetOnlineTerminalSession(sn, session); err == nil {
			this.Data["json"] = map[string]interface{}{
				"code": 0,
				"data": session,
			}
		} else {
			this.ResponseErrorJSON(500, err.Error())
		}
	}
	this.ServeJson()
}

// @Title GetAll
// @Description 获取终端设备session列表
// @Success 200 {object} rpcClient.RPCSessionModel
// @Failure 403 :gate is out of server
// @router / [get]
func (this *SessionController) GetAll() {
	sessionList := new([]rpc.RPCSessionModel)
	if err := rpc.GetOnlineTerminalSessions(0, sessionList); err == nil {
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"data": sessionList,
		}
	} else {
		this.ResponseErrorJSON(500, err.Error())
	}
	this.ServeJson()
}

// @Title delete
// @Description 删除(踢出)Session
// @Param	uid		path 	string	true		"session id.(terminalSN)"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (this *SessionController) Delete() {
	uid := this.GetString(":uid")
	result := new(int)
	if err := rpc.Kick(uid, result); err == nil {
		if (*result) == 1 {
			this.Data["json"] = map[string]interface{}{
				"code": 0,
				"msg":  "Kick success!",
			}
		} else {
			this.ResponseErrorJSON(403, "Kick failed.")
		}
	} else {
		this.ResponseErrorJSON(500, err.Error())
	}
	this.ServeJson()
}
