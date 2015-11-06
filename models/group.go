package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Group struct {
	Id        int64     `json:"id"`
	ParentId  int64     `json:"pid"`
	GroupName string    `json:"name"`
	Password  string    `json:"pwd"`
	CreateOn  time.Time `json:"createOn"`
	ModifyOn  time.Time `json:"modifyOn"`

	GroupProfile *GroupProfile `json:"profile" orm:"rel(one)"`
	Roles        []*Role       `json:"roles" orm:"rel(m2m)"`
	// Terminals []*Terminal `json:"terminals" orm:"reverse(many)"`
}

type GroupProfile struct {
	Id            int64     `json:"id"`
	GroupRealName string    `json:"realname"`
	ContactName   string    `json:"contact"`
	ContactPhone  string    `json:"phone"`
	CreateOn      time.Time `json:"createOn"`
	ModifyOn      time.Time `json:"modifyOn"`
}

type GroupSearchModel struct {
	GroupId       int64  `json:"id"`
	GroupName     string `json:"name"`
	GroupRealName string `json:"realname"`
	ContactName   string `json:"contact"`
	ContactPhone  string `json:"phone"`

	SearchType string `json:"search_type"`
}

func SearchGroup(key string, limit int) *[]GroupSearchModel {
	var groups []GroupSearchModel
	// 获取 QueryBuilder 对象. 需要指定数据库驱动参数。
	// 第二个返回值是错误对象，在这里略过
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select(
		"`group`.id group_id",
		"`group`.group_name",
		"group_profile.*",
	).
		From(
		"`group`",
	).
		InnerJoin(
		"group_profile",
	).
		On(
		"`group`.group_profile_id = group_profile.id",
	).
		Where("`group`.group_name like ? or group_profile.group_real_name like ? or group_profile.contact_phone like ? or group_profile.contact_name like ?").
		Limit(limit).Offset(0)

	// 导出SQL语句
	sql := qb.String()

	// 执行SQL语句
	o := orm.NewOrm()
	o.Raw(sql, key, key, key, key).QueryRows(&groups)

	return &groups
}

func AddGroup(obj Group) (int64, error) {
	var id int64
	var profileId int64
	var err error
	o := GetOrm()
	if err = o.Begin(); err == nil {
		obj.CreateOn = time.Now()
		obj.ModifyOn = time.Now()
		if obj.GroupProfile == nil {
			obj.GroupProfile = &GroupProfile{
				CreateOn: time.Now(),
				ModifyOn: time.Now(),
			}
		} else {
			obj.GroupProfile.CreateOn = time.Now()
			obj.GroupProfile.ModifyOn = time.Now()
		}

		// 插入profile
		if profileId, err = o.Insert(obj.GroupProfile); err != nil {
			o.Rollback()
			return 0, err
		} else {
			obj.GroupProfile.Id = profileId
		}

		//插入group
		if id, err = o.Insert(&obj); err != nil {
			o.Rollback()
			return 0, err
		} else {
			obj.Id = id
			if obj.Roles != nil && len(obj.Roles) > 0 {
				m2m := o.QueryM2M(&obj, "Roles")
				if _, err := m2m.Add(obj.Roles); err != nil {
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

func GetGroup(id int64) (*Group, error) {
	o := GetOrm()
	group := new(Group)
	if err := o.QueryTable("group").Filter("Id", id).RelatedSel().One(group); err != nil {
		return nil, err
	} else {
		if _, err := o.LoadRelated(group, "Roles"); err != nil {
			return nil, err
		}
	}

	return group, nil
}

func GetGroupProfile(id int64) (*GroupProfile, error) {
	o := GetOrm()
	group := new(Group)
	if err := o.QueryTable("group").Filter("Id", id).One(group, "GroupProfile"); err != nil {
		return nil, err
	}
	profile := new(GroupProfile)
	if err := o.QueryTable("group_profile").Filter("Id", group.GroupProfile.Id).One(profile); err != nil {
		return nil, err
	}
	return profile, nil
}

func GetGroupRoles(id int64) ([]*Role, error) {
	o := GetOrm()
	group := Group{Id: id}
	if _, err := o.LoadRelated(&group, "Roles"); err != nil {
		return nil, err
	}
	return group.Roles, nil
}

func GetAllGroups(cond *orm.Condition, pageIndex, pageSize int, order ...string) (*[]Group, int64, error) {
	o := GetOrm()
	var groups *[]Group = new([]Group)
	total, err := o.QueryTable("group").SetCond(cond).Count()
	if err != nil {
		return nil, 0, err
	}
	_, err = o.QueryTable("group").SetCond(cond).Limit(pageSize, (pageIndex-1)*pageSize).OrderBy(order...).All(groups)
	if err != nil {
		return nil, 0, err
	}
	return groups, total, nil
}

func UpdateGroup(id int64, update *Group) (*Group, error) {
	if obj, err := GetGroup(id); err == nil {
		if update.ParentId != 0 {
			obj.ParentId = update.ParentId
		}
		if update.GroupName != "" {
			obj.GroupName = update.GroupName
		}
		if update.Password != "" {
			obj.Password = update.Password
		}

		if update.GroupProfile != nil && obj.GroupProfile != nil {
			if profile, err := UpdateGroupProfile(id, update.GroupProfile); err == nil {
				obj.GroupProfile = profile
			}
		}

		if update.Roles != nil {
			obj.Roles = update.Roles
		}

		obj.ModifyOn = time.Now()

		o := GetOrm()
		if err := o.Begin(); err == nil {
			_, err := o.Update(obj)
			if err != nil {
				o.Rollback()
				return nil, err
			}
			if obj.Roles != nil && len(obj.Roles) > 0 {
				m2m := o.QueryM2M(obj, "Roles")
				if _, err := m2m.Clear(); err != nil {
					o.Rollback()
					return nil, err
				}
				if _, err := m2m.Add(obj.Roles); err != nil {
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

func UpdateGroupProfile(id int64, update *GroupProfile) (*GroupProfile, error) {
	if obj, err := GetGroupProfile(id); err == nil {
		if update.GroupRealName != "" {
			obj.GroupRealName = update.GroupRealName
		}
		if update.ContactName != "" {
			obj.ContactName = update.ContactName
		}
		if update.ContactPhone != "" {
			obj.ContactPhone = update.ContactPhone
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

func UpdateGroupRoles(id int64, roles []*Role) ([]*Role, error) {
	o := GetOrm()
	if err := o.Begin(); err == nil {
		group := Group{Id: id}
		m2m := o.QueryM2M(&group, "Roles")
		if _, err := m2m.Clear(); err != nil {
			o.Rollback()
			return nil, err
		}
		if len(roles) > 0 {
			if _, err := m2m.Add(roles); err != nil {
				o.Rollback()
				return nil, err
			}
		}
		o.Commit()
		return roles, nil
	} else {
		return nil, err
	}
}

func DeleteGroup(id int64) error {
	o := GetOrm()
	if err := o.Begin(); err == nil {
		group := Group{Id: id}
		m2m := o.QueryM2M(&group, "Roles")
		if _, err := m2m.Clear(); err != nil {
			o.Rollback()
			return err
		}
		_, err := o.QueryTable("group").Filter("Id", id).Delete()
		if err != nil {
			o.Rollback()
			return err
		}
		if err := o.Commit(); err != nil {
			o.Rollback()
			return err
		}
		return nil
	} else {
		return err
	}
}
