package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Role struct {
	Id          int64     `json:"id"`
	ParentId    int64     `json:"pid"`
	RoleName    string    `json:"name"`
	Description string    `json:"desc"`
	CreateOn    time.Time `json:"create_on"`
	ModifyOn    time.Time `json:"modify_on"`

	Rights []*Right `json:"rights" orm:"rel(m2m)"`
}

func AddRole(obj Role) (int64, error) {
	obj.CreateOn = time.Now()
	obj.ModifyOn = time.Now()
	o := GetOrm()

	if err := o.Begin(); err == nil {
		id, err := o.Insert(&obj)
		if err != nil {
			o.Rollback()
			return 0, err
		} else {
			obj.Id = id
			if obj.Rights != nil && len(obj.Rights) > 0 {
				m2m := o.QueryM2M(&obj, "Rights")
				if _, err := m2m.Add(obj.Rights); err != nil {
					o.Rollback()
					return 0, err
				}
			}
		}
		if err := o.Commit(); err != nil {
			o.Rollback()
			return 0, err
		}
		return id, nil
	} else {
		return 0, err
	}
}

func GetRole(id int64) (*Role, error) {
	o := GetOrm()
	role := &Role{Id: id}
	if err := o.Read(role); err != nil {
		return nil, err
	}
	if _, err := o.LoadRelated(role, "Rights"); err != nil {
		return nil, err
	}
	return role, nil
}

func GetAllRoles(cond *orm.Condition, pageIndex, pageSize int, order ...string) (*[]Role, int64, error) {
	o := GetOrm()
	var roles *[]Role = new([]Role)
	total, err := o.QueryTable("role").SetCond(cond).Count()
	if err != nil {
		return nil, 0, err
	}
	_, err = o.QueryTable("role").SetCond(cond).Limit(pageSize, (pageIndex-1)*pageSize).OrderBy(order...).All(roles)
	if err != nil {
		return nil, 0, err
	}
	return roles, total, nil
}

func UpdateRole(id int64, update *Role) (*Role, error) {
	if obj, err := GetRole(id); err == nil {
		if update.ParentId != 0 {
			obj.ParentId = update.ParentId
		}
		if update.RoleName != "" {
			obj.RoleName = update.RoleName
		}
		if update.Description != "" {
			obj.Description = update.Description
		}
		if update.Rights != nil {
			obj.Rights = update.Rights
		}
		obj.ModifyOn = time.Now()
		o := GetOrm()
		if err := o.Begin(); err == nil {
			_, err := o.Update(obj)
			if err != nil {
				o.Rollback()
				return nil, err
			}
			if obj.Rights != nil && len(obj.Rights) > 0 {
				m2m := o.QueryM2M(obj, "Rights")
				if _, err := m2m.Clear(); err != nil {
					o.Rollback()
					return nil, err
				}
				if _, err := m2m.Add(obj.Rights); err != nil {
					o.Rollback()
					return nil, err
				}
			}
			if err := o.Commit(); err != nil {
				o.Rollback()
				return nil, err
			}
			return obj, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func DeleteRole(id int64) error {
	o := GetOrm()
	o.Begin()
	role := Role{Id: id}
	m2m := o.QueryM2M(&role, "Rights")
	if _, err := m2m.Clear(); err != nil {
		o.Rollback()
		return err
	}
	_, err := o.QueryTable("role").Filter("Id", id).Delete()
	if err != nil {
		o.Rollback()
		return err
	}
	if err := o.Commit(); err != nil {
		o.Rollback()
		return err
	}
	return nil
}
