package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
	"time"
)

var (
	dbUser     = beego.AppConfig.String("dbuser")
	dbPassword = beego.AppConfig.String("dbpass")
	dbHost     = beego.AppConfig.String("dbhost")
	dbPort     = beego.AppConfig.String("dbport")
	dbName     = beego.AppConfig.String("dbname")
)

var (
	ERROR_NOT_FOUND = errors.New("未找到匹配记录")
)

func GetOrm() orm.Ormer {
	o := orm.NewOrm()
	o.Using("default")
	return o
}

func init() {
	dbUser = beego.AppConfig.String("dbuser")
	dbPassword = beego.AppConfig.String("dbpass")
	dbHost = beego.AppConfig.String("dbhost")
	dbPort = beego.AppConfig.String("dbport")
	dbName = beego.AppConfig.String("dbname")
	orm.DefaultTimeLoc = time.Local
	// 需要在init中注册定义的model
	orm.RegisterModel(
		new(User),
		new(UserProfile),
		new(Group),
		new(GroupProfile),
		new(Log),
		new(Message),
		new(Terminal),
		new(TerminalProfile),
		new(TerminalCarrier),
		new(Warning),
		new(Role),
		new(Right),
	)
	maxIdle := 30 //(可选)  设置最大空闲连接
	maxConn := 50 //(可选)  设置最大数据库连接 (go >= 1.2)

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	) + "&loc=" + url.QueryEscape("Local")

	orm.RegisterDataBase("default", "mysql",
		connStr,
		maxIdle,
		maxConn,
	)

	if beego.RunMode == "dev" {
		orm.Debug = true
	}
}
