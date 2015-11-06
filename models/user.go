package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id       int64     `json:"id"`
	UserName string    `json:"username"`
	Password string    `json:"password"`
	CreateOn time.Time `json:"createOn"`
	ModifyOn time.Time `json:"modifyOn"`

	UserProfile *UserProfile `json:"profile" orm:"rel(one)"`
}

type UserProfile struct {
	Id        int64     `json:"id"`
	Identity  string    `json:"identity"` //身份证号
	RealName  string    `json:"realname"` //真实姓名
	Gender    int       `json:"gender"`   //性别
	Address   string    `json:"address"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	SimNumber string    `json:"sim"`
	CreateOn  time.Time `json:"createOn"`
	ModifyOn  time.Time `json:"modifyOn"`
}

type UserSearchModel struct {
	UserId    int64  `json:"id"`
	UserName  string `json:"username"`
	Identity  string `json:"identity"` //身份证号
	RealName  string `json:"realname"` //真实姓名
	Gender    int    `json:"-"`        //性别
	Address   string `json:"-"`
	Email     string `json:"-"`
	Phone     string `json:"phone"`
	SimNumber string `json:"-"`

	SearchType string `json:"search_type"`
}

func SearchUser(key string, limit int) *[]UserSearchModel {
	var users []UserSearchModel
	// 获取 QueryBuilder 对象. 需要指定数据库驱动参数。
	// 第二个返回值是错误对象，在这里略过
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select(
		"user.id user_id",
		"user.user_name",
		"user_profile.*",
	).
		From(
		"user",
	).
		InnerJoin(
		"user_profile",
	).
		On(
		"user.user_profile_id = user_profile.id",
	).
		Where("user_profile.identity like ? or user_profile.real_name like ? or user.user_name like ? or user_profile.phone like ?").
		Limit(limit).Offset(0)

	// 导出SQL语句
	sql := qb.String()

	// 执行SQL语句
	o := orm.NewOrm()
	o.Raw(sql, key, key, key, key).QueryRows(&users)

	return &users
}

func AddUser(obj User) (int64, error) {
	var id int64
	var profileId int64
	var err error
	orm := GetOrm()
	if err = orm.Begin(); err == nil {
		obj.CreateOn = time.Now()
		obj.ModifyOn = time.Now()
		if obj.UserProfile == nil {
			obj.UserProfile = &UserProfile{
				CreateOn: time.Now(),
				ModifyOn: time.Now(),
			}
		} else {
			obj.UserProfile.CreateOn = time.Now()
			obj.UserProfile.ModifyOn = time.Now()
		}
		if profileId, err = orm.Insert(obj.UserProfile); err == nil {
			obj.UserProfile.Id = profileId
			if id, err = orm.Insert(&obj); err == nil {
				orm.Commit()
			} else {
				orm.Rollback()
				return 0, err
			}
		} else {
			orm.Rollback()
			return 0, err
		}
	} else {
		return 0, err
	}

	return id, nil
}

func GetUser(id int64) (*User, error) {
	orm := GetOrm()
	user := new(User)
	if err := orm.QueryTable("user").Filter("id", id).RelatedSel().One(user); err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserProfile(id int64) (*UserProfile, error) {
	orm := GetOrm()
	user := new(User)
	if err := orm.QueryTable("user").Filter("id", id).One(user, "UserProfile"); err != nil {
		return nil, err
	}
	profile := new(UserProfile)
	if err := orm.QueryTable("user_profile").Filter("id", user.UserProfile.Id).One(profile); err != nil {
		return nil, err
	}
	return profile, nil
}

func GetAllUsers(cond *orm.Condition, pageIndex, pageSize int, order ...string) (*[]User, int64, error) {
	orm := GetOrm()
	var users *[]User = new([]User)
	total, err := orm.QueryTable("user").SetCond(cond).Count()
	if err != nil {
		return nil, 0, err
	}
	_, err = orm.QueryTable("user").SetCond(cond).Limit(pageSize, (pageIndex-1)*pageSize).OrderBy(order...).All(users)
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func UpdateUser(id int64, update *User) (*User, error) {
	if obj, err := GetUser(id); err == nil {
		if update.UserName != "" {
			obj.UserName = update.UserName
		}
		if update.Password != "" {
			obj.Password = update.Password
		}
		if update.UserProfile != nil && obj.UserProfile != nil {
			if profile, err := UpdateUserProfile(id, update.UserProfile); err == nil {
				obj.UserProfile = profile
			}
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

func UpdateUserProfile(id int64, update *UserProfile) (*UserProfile, error) {
	if obj, err := GetUserProfile(id); err == nil {
		if update.Identity != "" {
			obj.Identity = update.Identity
		}
		if update.RealName != "" {
			obj.RealName = update.RealName
		}
		if update.Gender != 0 {
			obj.Gender = update.Gender
		}
		if update.Email != "" {
			obj.Email = update.Email
		}
		if update.Address != "" {
			obj.Address = update.Address
		}
		if update.Phone != "" {
			obj.Phone = update.Phone
		}
		if update.SimNumber != "" {
			obj.SimNumber = update.SimNumber
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

func DeleteUser(id int64) error {
	orm := GetOrm()
	rows, err := orm.QueryTable("user").Filter("Id", id).Delete()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ERROR_NOT_FOUND
	}
	return nil
}
