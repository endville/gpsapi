package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Terminal struct {
	Id              int64            `json:"id"`
	TerminalSn      string           `json:"sn"`
	Password        string           `json:"pwd"`
	User            *User            `json:"user" orm:"null;rel(fk)"`
	Group           *Group           `json:"group" orm:"null;rel(fk)"`
	TerminalProfile *TerminalProfile `json:"profile" orm:"null;rel(one);on_delete(set_null)"`
	TerminalCarrier *TerminalCarrier `json:"carrier" orm:"null;rel(one);on_delete(set_null)"`
	CreateOn        time.Time        `json:"createOn"`
	ModifyOn        time.Time        `json:"modifyOn"`
	OfflineOn       time.Time        `json:"offlineOn"`
	OnlineOn        time.Time        `json:"onlineOn"`
}

type TerminalProfile struct {
	Id          int64     `json:"id"`
	TerminalSn  string    `json:"sn"`
	Tmsisdn     string    `json:"tmsisdn"`
	Pmsisdn     string    `json:"pmsisdn"`
	Imsi        string    `json:"imsi"`
	Imei        string    `json:"imei"`
	ProductCode string    `json:"pcode"`
	IsActivated int       `json:"activated"`
	Mileage     int64     `json:"mileage"` // 记录里程
	ActivateOn  time.Time `json:"activateOn`
	ExpireOn    time.Time `json:"expireOn"`
	CreateOn    time.Time `json:"createOn"`
	ModifyOn    time.Time `json:"modifyOn"`
}

type TerminalCarrier struct {
	Id                          int64     `json:"id"`
	LicensePlateNumber          string    `json:"lpn"`   // 车牌号
	VehicleIdentificationNumber string    `json:"vin"`   // 车架号
	CarrierType                 string    `json:"type"`  // 载具类型
	Brand                       string    `json:"brand"` // 品牌
	Color                       string    `json:"color"`
	Picture                     string    `json:"picture"`
	CreateOn                    time.Time `json:"createOn"`
	ModifyOn                    time.Time `json:"modifyOn"`
}

type TerminalSearchModel struct {
	TerminalId                  int64     `json:"id"`
	TerminalSn                  string    `json:"sn"`
	Tmsisdn                     string    `json:"-"`
	Pmsisdn                     string    `json:"-"`
	Imsi                        string    `json:"-"`
	Imei                        string    `json:"-"`
	ProductCode                 string    `json:"-"`
	IsActivated                 int       `json:"-"`
	Mileage                     int64     `json:"-"` // 记录里程
	LicensePlateNumber          string    `json:"-"` // 车牌号
	VehicleIdentificationNumber string    `json:"-"` // 车架号
	CarrierType                 string    `json:"-"` // 载具类型
	Brand                       string    `json:"-"` // 品牌
	Color                       string    `json:"-"`
	Picture                     string    `json:"-"`
	OfflineOn                   time.Time `json:"offlineOn"`
	OnlineOn                    time.Time `json:"onlineOn"`

	SearchType string `json:"search_type"`
}

func SearchTerminal(key string, limit int) *[]TerminalSearchModel {
	var terminals []TerminalSearchModel
	// 获取 QueryBuilder 对象. 需要指定数据库驱动参数。
	// 第二个返回值是错误对象，在这里略过
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select(
		"terminal.id terminal_id",
		"terminal.offline_on",
		"terminal.online_on",
		"terminal_profile.*",
		"terminal_carrier.*",
		"terminal.terminal_sn",
	).
		From(
		"terminal",
	).
		InnerJoin(
		"terminal_profile",
	).
		On(
		"terminal.terminal_profile_id = terminal_profile.id",
	).
		InnerJoin(
		"terminal_carrier",
	).
		On(
		"terminal.terminal_carrier_id = terminal_carrier.id",
	).
		Where("terminal.terminal_sn like ?").
		Limit(limit).Offset(0)

	// 导出SQL语句
	sql := qb.String()

	// 执行SQL语句
	o := orm.NewOrm()
	o.Raw(sql, key).QueryRows(&terminals)

	return &terminals
}

func AddTerminal(obj Terminal) (int64, error) {
	var id int64
	var profileId int64
	var carrierId int64
	var err error
	orm := GetOrm()
	if err = orm.Begin(); err == nil {
		obj.CreateOn = time.Now()
		obj.ModifyOn = time.Now()
		obj.OfflineOn = time.Now()
		//profile
		if obj.TerminalProfile == nil {
			obj.TerminalProfile = &TerminalProfile{
				CreateOn: time.Now(),
				ModifyOn: time.Now(),
			}
		} else {
			obj.TerminalProfile.CreateOn = time.Now()
			obj.TerminalProfile.ModifyOn = time.Now()
		}

		if profileId, err = orm.Insert(obj.TerminalProfile); err == nil {
			obj.TerminalProfile.Id = profileId
		} else {
			orm.Rollback()
			return 0, err
		}
		//profile end

		//carrier
		if obj.TerminalCarrier == nil {
			obj.TerminalCarrier = &TerminalCarrier{
				CreateOn: time.Now(),
				ModifyOn: time.Now(),
			}
		} else {
			obj.TerminalCarrier.CreateOn = time.Now()
			obj.TerminalCarrier.ModifyOn = time.Now()
		}

		if carrierId, err = orm.Insert(obj.TerminalCarrier); err == nil {
			obj.TerminalCarrier.Id = carrierId
		} else {
			orm.Rollback()
			return 0, err
		}
		//carrier end

		if id, err = orm.Insert(&obj); err == nil {
			orm.Commit()
		} else {
			orm.Rollback()
			return 0, err
		}
	} else {
		return 0, err
	}

	return id, nil
}

func GetTerminal(id int64) (*Terminal, error) {
	orm := GetOrm()
	terminal := new(Terminal)
	if err := orm.QueryTable("terminal").Filter("Id", id).RelatedSel().One(terminal); err != nil {
		return nil, err
	}

	return terminal, nil
}

func GetTerminalProfile(id int64) (*TerminalProfile, error) {
	orm := GetOrm()
	terminal := new(Terminal)
	if err := orm.QueryTable("terminal").Filter("id", id).One(terminal, "TerminalProfile"); err != nil {
		return nil, err
	}
	profile := new(TerminalProfile)
	if err := orm.QueryTable("terminal_profile").Filter("id", terminal.TerminalProfile.Id).One(profile); err != nil {
		return nil, err
	}
	return profile, nil
}

func GetTerminalCarrier(id int64) (*TerminalCarrier, error) {
	orm := GetOrm()
	terminal := new(Terminal)
	if err := orm.QueryTable("terminal").Filter("id", id).One(terminal, "TerminalProfile"); err != nil {
		return nil, err
	}
	carrier := new(TerminalCarrier)
	if err := orm.QueryTable("terminal_carrier").Filter("id", terminal.TerminalProfile.Id).One(carrier); err != nil {
		return nil, err
	}
	return carrier, nil
}

func GetAllTerminals(cond *orm.Condition, pageIndex, pageSize int, order ...string) (*[]Terminal, int64, error) {
	orm := GetOrm()
	var terminals *[]Terminal = new([]Terminal)
	total, err := orm.QueryTable("terminal").SetCond(cond).Count()
	if err != nil {
		return nil, 0, err
	}
	_, err = orm.QueryTable("terminal").SetCond(cond).Limit(pageSize, (pageIndex-1)*pageSize).OrderBy(order...).All(terminals)
	if err != nil {
		return nil, 0, err
	}
	return terminals, total, nil
}

func UpdateTerminal(id int64, update *Terminal) (*Terminal, error) {
	if obj, err := GetTerminal(id); err == nil {
		if update.TerminalSn != "" {
			obj.TerminalSn = update.TerminalSn
		}
		if update.Password != "" {
			obj.Password = update.Password
		}
		if update.User != nil {
			obj.User.Id = update.User.Id
		}
		if update.Group != nil {
			obj.Group.Id = update.Group.Id
		}
		if update.TerminalProfile != nil && obj.TerminalProfile != nil {
			if profile, err := UpdateTerminalProfile(id, update.TerminalProfile); err == nil {
				obj.TerminalProfile = profile
			}
		}
		if update.TerminalCarrier != nil && obj.TerminalCarrier != nil {
			if carrier, err := UpdateTerminalCarrier(id, update.TerminalCarrier); err == nil {
				obj.TerminalCarrier = carrier
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

func UpdateTerminalProfile(id int64, update *TerminalProfile) (*TerminalProfile, error) {
	if obj, err := GetTerminalProfile(id); err == nil {
		if update.TerminalSn != "" {
			obj.TerminalSn = update.TerminalSn
		}
		if update.Tmsisdn != "" {
			obj.Tmsisdn = update.Tmsisdn
		}
		if update.Pmsisdn != "" {
			obj.Pmsisdn = update.Pmsisdn
		}
		if update.Imsi != "" {
			obj.Imsi = update.Imsi
		}
		if update.Imei != "" {
			obj.Imei = update.Imei
		}
		if update.ProductCode != "" {
			obj.ProductCode = update.ProductCode
		}
		if update.IsActivated != 0 {
			obj.IsActivated = update.IsActivated
		}
		if update.Mileage != 0 {
			obj.Mileage = update.Mileage
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

func UpdateTerminalCarrier(id int64, update *TerminalCarrier) (*TerminalCarrier, error) {
	if obj, err := GetTerminalCarrier(id); err == nil {
		if update.LicensePlateNumber != "" {
			obj.LicensePlateNumber = update.LicensePlateNumber
		}
		if update.VehicleIdentificationNumber != "" {
			obj.VehicleIdentificationNumber = update.VehicleIdentificationNumber
		}
		if update.CarrierType != "" {
			obj.CarrierType = update.CarrierType
		}
		if update.Brand != "" {
			obj.Brand = update.Brand
		}
		if update.Color != "" {
			obj.Color = update.Color
		}
		if update.Picture != "" {
			obj.Picture = update.Picture
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

func DeleteTerminal(id int64) error {
	orm := GetOrm()
	rows, err := orm.QueryTable("terminal").Filter("Id", id).Delete()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ERROR_NOT_FOUND
	}
	return nil
}
