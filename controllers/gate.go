package controllers

import (
	rpc "gpsapi/rpcClient"
)

// 网关数据监控相关
type GateController struct {
	BaseController
}

// @Title SessionCount
// @Description 获取当前网关连接设备数
// @Success 200 {int} count
// @Failure 403 :gate is out of server
// @router /sessionCount [get]
func (this *GateController) SessionCount() {
	count := new(int)
	if err := rpc.GetOnlineTerminalCount(0, count); err == nil {
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"data": count,
		}
	} else {
		this.ResponseErrorJSON(500, err.Error())
	}
	this.ServeJson()
}

// @Title State
// @Description 获取当前网关状态
// @Success 200 {string} state
// @Failure 401 :gate is out of server
// @router /state [get]
func (this *GateController) State() {
	signal := rpc.RPC_SIGNAL_TEST
	reply := new(int)
	if err := rpc.Test(signal, reply); err == nil {
		if (*reply) == rpc.RPC_SIGNAL_TEST_REPLY {
			this.Data["json"] = map[string]interface{}{
				"code": 0,
				"data": "ok",
			}
		} else {
			this.ResponseErrorJSON(401, "not ok")
		}
	} else {
		this.ResponseErrorJSON(500, err.Error())
	}
	this.ServeJson()
}
