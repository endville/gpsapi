package models

import (
	"time"
)

type StatisticTerminalServiceExpire struct {
	Id           int64     `json:"id"`
	Name         string    `json:"name"`
	TerminalSn   string    `json:"sn"`
	ActivateOn   time.Time `json:"-"`
	ExpireOn     time.Time `json:"-"`
	OfflineOn    time.Time `json:"-"`
	OnlineOn     time.Time `json:"-"`
	ActivateOnTs int64     `json:"activateOn"`
	ExpireOnTs   int64     `json:"expireOn"`
	OfflineOnTs  int64     `json:"offlineOn"`
	OnlineOnTs   int64     `json:"onlineOn"`
	State        bool      `json:"state"`
}

type StatisticTerminalRunInfo struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	TerminalSn    string `json:"sn"`
	Mileage       int64  `json:"mileage"`
	SpeedingTimes int    `json:"speedingTimes"`
	StayTimes     int    `json:"stayTimes"`
}

type StatisticTerminalMileage struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	TerminalSn    string `json:"sn"`
	Mileage       int64  `json:"mileage"`
	SpeedingTimes int    `json:"speedingTimes"`
	StayTimes     int    `json:"stayTimes"`
}

type StatisticTerminalStay struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	TerminalSn  string    `json:"sn"`
	TimeBegin   time.Time `json:"-"`
	TimeEnd     time.Time `json:"-"`
	TimeBeginTs int64     `json:"timeBegin"`
	TimeEndTs   int64     `json:"timeEnd"`
}

type StatisticTerminalWarning struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	TerminalSn string    `json:"sn"`
	Time       time.Time `json:"-"`
	TimeTs     int64     `json:"time"`
	Longitude  float64   `json:"lng"`
	Latitude   float64   `json:"lat"`
	Type       int       `json:"type"`
}

type StatisticTerminalOnline struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type StatisticTerminalOnlineDetails struct {
	Id           int64     `json:"id"`
	Name         string    `json:"name"`
	TerminalSn   string    `json:"sn"`
	IMEI         string    `json:"imei"`
	PhoneNumber  string    `json:"phone"`
	OnlineTime   time.Time `json:"-"`
	OnlineTimeTs int64     `json:"time"`
}

type StatisticTerminalOffline struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	TerminalSn  string    `json:"sn"`
	TimeBegin   time.Time `json:"-"`
	TimeEnd     time.Time `json:"-"`
	TimeBeginTs int64     `json:"timeBegin"`
	TimeEndTs   int64     `json:"timeEnd"`
}

type StatisticTerminalQuery struct {
	Id               int64  `json:"id"`
	Name             string `json:"name"`
	TerminalSn       string `json:"sn"`
	PhoneNumber      string `json:"phone"`
	OwnerPhoneNumber string `json:"ownerPhone"`
}
