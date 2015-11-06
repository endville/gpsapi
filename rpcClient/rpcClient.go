package rpcClient

import (
	"github.com/astaxie/beego"
	"gpsapi/models"
	"net/rpc"
	"time"
)

const (
	RPC_SIGNAL_TEST = iota + 1
	RPC_SIGNAL_TEST_REPLY
)

var gpsGateServerAddress = beego.AppConfig.String("gpsGateServer")

//用于RPC调用时使用的model
type RPCSendMessageModel struct {
	TerminalSn   string
	Message      models.ServerMessage
	NeedFeedback bool
	Timeout      time.Duration
}

type RPCSessionModel struct {
	TerminalSn string    // 终端Sn号
	TerminalId int64     // 终端在数据库中的ID,方便查询数据库
	UserId     int64     // 终端所属用户的ID,方便查询数据库
	GroupId    int64     // 终端所属车队的ID,方便查询数据库
	ConnectOn  time.Time // 终端开始连接时间
	RemoteAddr string    // Session ID
}

func SendMessage(model RPCSendMessageModel, result *models.TerminalMessage) error {
	return call("TerminalDelegate.SendMessage", model, result)
}

func GetOnlineTerminalCount(param int, count *int) error {
	return call("TerminalDelegate.GetOnlineTerminalCount", param, count)
}

func GetOnlineTerminalSession(terminalSn string, session *RPCSessionModel) error {
	return call("TerminalDelegate.GetOnlineTerminalSession", terminalSn, session)
}

func GetOnlineTerminalSessions(param int, sessions *[]RPCSessionModel) error {
	return call("TerminalDelegate.GetOnlineTerminalSessions", param, sessions)
}

func IsAlive(sessionID string, result *int) error {
	return call("TerminalDelegate.IsAlive", sessionID, result)
}

func Kick(sessionID string, result *int) error {
	return call("TerminalDelegate.Kick", sessionID, result)
}

func Test(signal int, reply *int) error {
	return call("TerminalDelegate.Test", signal, reply)
}

func call(method string, args, reply interface{}) error {
	client, err := rpc.DialHTTP("tcp", gpsGateServerAddress)
	if err != nil {
		return err
	}
	defer client.Close()

	err = client.Call(method, args, reply)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	gpsGateServerAddress = beego.AppConfig.String("gpsGateServer")
}
