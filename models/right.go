package models

import (
	"github.com/astaxie/beego/orm"
)

type Right struct {
	Id          int64  `json:"id"`
	RightName   string `json:"name"`
	Description string `json:"desc"`
}

func GetAllRights(cond *orm.Condition, pageIndex, pageSize int, order ...string) (*[]Right, int64, error) {
	orm := GetOrm()
	var roles *[]Right = new([]Right)
	total, err := orm.QueryTable("right").SetCond(cond).Count()
	if err != nil {
		return nil, 0, err
	}
	_, err = orm.QueryTable("right").SetCond(cond).Limit(pageSize, (pageIndex-1)*pageSize).OrderBy(order...).All(roles)
	if err != nil {
		return nil, 0, err
	}
	return roles, total, nil
}
