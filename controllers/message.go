package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"gpsapi/models"
	rpc "gpsapi/rpcClient"
	"strings"

	"time"
)

// 命令请求相关
type MessageController struct {
	BaseController
}

// @Title Post
// @Description 发送命令请求
// @Param	body	body	models.MessageRequest	true	命令请求信息
// @Success 200 {object} models.MessageRequest
// @Failure 400	请求的参数不正确
// @Failure 502	错误网关
// @Failure 504	网关超时
// @router / [post]
func (this *MessageController) Post() {
	var r models.MessageRequest
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &r); err != nil {
		this.ResponseErrorJSON(400, errorFormat(ErrorBadJson_400, err.Error()))
	} else {
		terminals := strings.Split(r.TerminalSn, ",")
		if len(terminals) > 1 {
			r.NeedFeedback = false
		}
		for _, terminalSn := range terminals {
			//构建网关待发送的消息
			serverMessage := models.ServerMessage{
				DateTime:    time.Now().Format("2006-01-02 15:04:05"),
				MessageType: r.MessageType,
				Params:      r.Params,
			}

			// 插入数据库做记录
			message := models.Message{
				TerminalSn:  terminalSn,
				MessageType: serverMessage.MessageType,
				MessageBody: serverMessage.String(),
				SendBy:      r.SendBy,
				SendOn:      time.Now(),
			}
			if messageId, err := models.AddMessage(message); err != nil {
				this.ResponseErrorJSON(400, errorFormat(ErrorBadParam_400, err.Error()))
			} else {
				//构建RPC服务所需的消息
				request := rpc.RPCSendMessageModel{
					TerminalSn:   terminalSn,
					Message:      serverMessage,
					NeedFeedback: r.NeedFeedback,
					Timeout:      time.Second * 10,
				}
				reply := new(models.TerminalMessage)
				if err := rpc.SendMessage(request, reply); err != nil {
					this.ResponseErrorJSON(502, errorFormat(ErrorBadGateway_502, err.Error()))
				} else {
					if update, err := models.UpdateMessage(messageId, &models.Message{FeedBack: reply.String()}); err != nil {
						this.ResponseErrorJSON(400, errorFormat(ErrorBadParam_400, err.Error()))
					} else {
						this.Data["json"] = map[string]interface{}{
							"code": 0,
							"data": update,
						}
					}
				}
			}
		}
	}
	this.ServeJson()
}

// @Title Get
// @Description 获取命令请求列表
// @Param	sn	query	int	false	"终端编号"
// @Param	time_begin	query	string	false	"消息发送起始时间"
// @Param	time_end	query	string	false	"消息发送终止时间"
// @Param	pageIndex	query	int	false	"页码, 默认1"
// @Param	pageSize	query	int	false	"每页显示条数, 默认30"
// @Success 200 {object} models.Message
// @router / [get]
func (this *MessageController) GetAll() {
	cond := orm.NewCondition()
	if sn := this.GetString("sn"); sn != "" {
		cond = cond.And("TerminalSn", sn)
	}
	if on_begin := this.GetString("time_begin"); on_begin != "" {
		cond = cond.And("SendOn__gt", on_begin)
	}
	if on_end := this.GetString("time_end"); on_end != "" {
		cond = cond.And("SendOn__lt", on_end)
	}
	pageIndex, _ := this.GetInt("pageIndex", 1)
	pageSize, _ := this.GetInt("pageSize", 30)
	messages, total, err := models.GetAllMessages(cond, pageIndex, pageSize)
	if err != nil {
		this.ResponseErrorJSON(400, errorFormat(ErrorBadRequest_400, err.Error()))
	}
	this.Data["json"] = map[string]interface{}{
		"code":  0,
		"data":  messages,
		"total": total,
	}
	this.ServeJson()
}

// @Title Get
// @Description 根据命令请求唯一ID号获取命令请求信息
// @Param	id		path 	int	true		"命令请求唯一ID号"
// @Success 200 {object} models.Message
// @Failure 400	请求的参数不正确
// @router /:id [get]
func (this *MessageController) Get() {
	id, _ := this.GetInt64(":id")
	message, err := models.GetMessage(id)
	if err != nil {
		this.ResponseErrorJSON(400, errorFormat(ErrorBadRequest_400, err.Error()))
	} else {
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"data": message,
		}
	}
	this.ServeJson()
}

// @Title delete
// @Description 删除命令请求
// @Param	id		path 	int	true		"命令请求唯一ID"
// @Success 200 {int} 返回记录ID
// @Failure 403 id is empty
// @router /:id [delete]
func (this *MessageController) Delete() {
	id, _ := this.GetInt64(":id")
	if err := models.DeleteMessage(id); err != nil {
		this.ResponseErrorJSON(403, err.Error())
	}
	this.Data["json"] = map[string]interface{}{
		"code": 0,
		"data": id,
	}
	this.ServeJson()
}
