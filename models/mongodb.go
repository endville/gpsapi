package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

var (
	DB_NAME         = beego.AppConfig.String("mongoname")
	DB_HOST         = beego.AppConfig.String("mongohost")
	DB_PORT         = beego.AppConfig.String("mongoport")
	MAX_SESSION_NUM = 128
	TIME_OUT        = 15 * time.Second
)

var mongoDbSess *mgo.Session = nil
var sessPool chan bool

type mongoDbSession struct {
	*mgo.Session
}

func getMongoDbSession() (*mongoDbSession, error) {
	if mongoDbSess == nil {
		if err := dialMongoDb(); err != nil {
			return nil, err
		}
	}
	select {
	case sessPool <- true:
		session := &mongoDbSession{
			mongoDbSess.New(),
		}
		return session, nil
	case <-time.After(TIME_OUT):
		errmsg := fmt.Sprintf("数据库连接数超过最大值(%s)", MAX_SESSION_NUM)
		return nil, errors.New(errmsg)
	}
}

func (this *mongoDbSession) Close() {
	if this.Session != nil {
		this.Session.Close()
		<-sessPool
	}
}

func dialMongoDb() (err error) {
	err = nil
	mongoDbSess, err = mgo.Dial(fmt.Sprintf("%s:%s", DB_HOST, DB_PORT))
	if err != nil {
		log.Println(err.Error())
	}
	return err
}

func init() {
	sessPool = make(chan bool, MAX_SESSION_NUM)
	if err := dialMongoDb(); err != nil {
		log.Println(err.Error())
	}
}
