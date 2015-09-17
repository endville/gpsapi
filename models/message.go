package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"regexp"
	"strings"
	"time"
)

var (
	messageRegex = regexp.MustCompile(`(\[.*?\])`) // 报文规则
)

type Message struct {
	Id          int64     `json:"id"`
	TerminalSn  string    `json:"sn"`
	RemoteAddr  string    `json:"addr"`
	MessageType string    `json:"type"`
	MessageBody string    `json:"body"`
	SendBy      string    `json:"sendBy"`
	SendOn      time.Time `json:"sendOn"`
	FeedBack    string    `json:"reply"`
	FeedBackOn  time.Time `json:"backOn"`
}

// 消息发送请求数据结构
type MessageRequest struct {
	SendBy       string
	TerminalSn   string
	MessageType  string
	NeedFeedback bool
	Params       []string
}

type ServerMessage struct {
	DateTime    string
	MessageType string
	Params      []string
}

func (this ServerMessage) String() string {
	if len(this.Params) == 0 {
		return fmt.Sprintf("[%s,%s]", this.DateTime, this.MessageType)
	}
	return fmt.Sprintf("[%s,%s,%s]", this.DateTime, this.MessageType, strings.Join(this.Params, ","))
}

func (this *ServerMessage) Parse(message string) error {
	msgs := SplitMessage(message)
	if msgs != nil && len(msgs) > 0 {
		message = msgs[0]
	}

	message = message[1 : len(message)-1] //去掉前后中括号
	params := strings.Split(message, ",")
	if len(params) < 2 {
		return errors.New("解析服务器消息时发现参数不足")
	}

	this.DateTime = params[0]
	this.MessageType = params[1]

	if len(params) > 2 {
		this.Params = params[2:]
	} else {
		this.Params = make([]string, 0)
	}

	return nil
}

type TerminalMessage struct {
	DateTime     string
	TerminalType string
	Version      string
	TerminalSn   string
	MessageType  string
	Params       []string
}

func (this TerminalMessage) String() string {
	if len(this.Params) == 0 {
		return fmt.Sprintf("[%s,%s,%s,%s,%s]", this.DateTime, this.MessageType, this.Version, this.TerminalSn, this.MessageType)
	}
	return fmt.Sprintf("[%s,%s,%s,%s,%s,%s]", this.DateTime, this.TerminalType, this.Version, this.TerminalSn, this.MessageType, strings.Join(this.Params, ","))
}

func (this *TerminalMessage) Parse(message string) error {
	msgs := SplitMessage(message)
	if msgs != nil && len(msgs) > 0 {
		message = msgs[0]
	}

	message = message[1 : len(message)-1] //去掉前后中括号
	params := strings.Split(message, ",")
	if len(params) < 2 {
		return errors.New("参数不足")
	}

	this.DateTime = params[0]
	this.TerminalType = params[1]
	this.Version = params[2]
	this.TerminalSn = params[3]
	this.MessageType = params[4]

	if len(params) > 5 {
		this.Params = params[5:]
	} else {
		this.Params = make([]string, 0)
	}
	return nil
}

func AddMessage(obj Message) (int64, error) {
	orm := GetOrm()
	id, err := orm.Insert(&obj)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func GetMessage(id int64) (*Message, error) {
	orm := GetOrm()
	message := new(Message)
	if err := orm.QueryTable("message").Filter("Id", id).One(message); err != nil {
		return nil, err
	}

	return message, nil
}

func GetAllMessages(cond *orm.Condition, pageIndex, pageSize int, order ...string) (*[]Message, int64, error) {
	orm := GetOrm()
	var messages *[]Message = new([]Message)
	total, err := orm.QueryTable("message").SetCond(cond).Count()
	if err != nil {
		return nil, 0, err
	}
	_, err = orm.QueryTable("message").SetCond(cond).Limit(pageSize, (pageIndex-1)*pageSize).OrderBy(order...).All(messages)
	if err != nil {
		return nil, 0, err
	}
	return messages, total, nil
}

func UpdateMessage(id int64, update *Message) (*Message, error) {
	if obj, err := GetMessage(id); err == nil {
		if update.FeedBack != "" {
			obj.FeedBack = update.FeedBack
			obj.FeedBackOn = time.Now()
		}
		orm := GetOrm()
		_, err := orm.Update(obj)
		if err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, err
	}
}

func DeleteMessage(id int64) error {
	orm := GetOrm()
	rows, err := orm.QueryTable("message").Filter("Id", id).Delete()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ERROR_NOT_FOUND
	}
	return nil
}

/*
将一段字符串分解成多个报文
*/
func SplitMessage(message string) []string {
	list := messageRegex.FindAllString(message, -1)
	if list != nil {
		return list
	}
	return nil
}
