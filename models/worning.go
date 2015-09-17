package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// 报警信息会在一定时间之后删除
type Warning struct {
	Id          int64   `json:"id"`
	TerminalSn  string  `json:"sn"`
	TerminalId  int64   `json:"tid"`
	UserId      int64   `json:"uid"`
	GroupId     int64   `json:"gid"`
	Longitude   float32 `json:"lng"`
	Latitude    float32 `json:"lat"`
	Speed       float32 `json:"speed"`
	Direction   float32 `json:"direction"`
	Status      int     `json:"status"`
	CellId      string  `json:"cellId"`
	Voltage     float32 `json:"voltage"`     // 电压
	Temperature int     `json:"temperature"` // 温度

	Type     int16     `json:"type"`  // 报警类型
	State    int16     `json:"state"` // 状态 初步设计为 未处理、已处理 2种
	Flag     int16     `json:"flag"`  // 标识
	CreateOn time.Time `json:"createOn" orm:"auto_now_add;type(datetime)"`
	ModifyOn time.Time `json:"modifyOn" orm:"auto_now;type(datetime)"`
}

func GetWarning(id int64) (*Warning, error) {
	orm := GetOrm()
	warning := new(Warning)
	if err := orm.QueryTable("warning").Filter("Id", id).RelatedSel().One(warning); err != nil {
		return nil, err
	}

	return warning, nil
}

func GetAllWarnings(cond *orm.Condition, pageIndex, pageSize int, order ...string) (*[]Warning, int64, error) {
	orm := GetOrm()
	var warnings *[]Warning = new([]Warning)
	total, err := orm.QueryTable("warning").SetCond(cond).Count()
	if err != nil {
		return nil, 0, err
	}
	_, err = orm.QueryTable("warning").SetCond(cond).Limit(pageSize, (pageIndex-1)*pageSize).OrderBy(order...).All(warnings)
	if err != nil {
		return nil, 0, err
	}
	return warnings, total, nil
}

func UpdateWarning(id int64, update *Warning) (*Warning, error) {
	if obj, err := GetWarning(id); err == nil {
		if update.State != 0 {
			obj.State = update.State
		}

		if update.Flag != 0 {
			obj.Flag = update.Flag
		}

		obj.ModifyOn = time.Now()

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

func DeleteWarning(id int64) error {
	orm := GetOrm()
	rows, err := orm.QueryTable("warning").Filter("Id", id).Delete()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ERROR_NOT_FOUND
	}
	return nil
}
