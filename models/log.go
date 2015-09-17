package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

const (
	LOG_LEVEL_DEBUG = 1 + iota
	LOG_LEVEL_INFO
	LOG_LEVEL_WORNING
	LOG_LEVEL_ERROR

	LOG_TYPE_TERMINAL = 1 + iota
	LOG_TYPE_USER
	LOG_TYPE_MANAGER
)

type Log struct {
	Id      int64     `json:"id"`
	Level   int       `json:"level"`
	Type    int       `json:"type"`
	Content string    `json:"content"`
	LogBy   string    `json:"by"`
	LogOn   time.Time `json:"on"`
}

func AddLog(obj Log) (int64, error) {
	obj.LogOn = time.Now()
	orm := GetOrm()
	id, err := orm.Insert(&obj)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func GetLog(id int64) (*Log, error) {
	orm := GetOrm()
	log := new(Log)
	if err := orm.QueryTable("log").Filter("Id", id).RelatedSel().One(log); err != nil {
		return nil, err
	}

	return log, nil
}

func GetAllLogs(cond *orm.Condition, pageIndex, pageSize int, order ...string) (*[]Log, int64, error) {
	orm := GetOrm()
	var logs *[]Log = new([]Log)
	total, err := orm.QueryTable("log").SetCond(cond).Count()
	if err != nil {
		return nil, 0, err
	}
	_, err = orm.QueryTable("log").SetCond(cond).Limit(pageSize, (pageIndex-1)*pageSize).OrderBy(order...).All(logs)
	if err != nil {
		return nil, 0, err
	}
	return logs, total, nil
}

func DeleteLog(id int64) error {
	orm := GetOrm()
	rows, err := orm.QueryTable("log").Filter("Id", id).Delete()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ERROR_NOT_FOUND
	}
	return nil
}
