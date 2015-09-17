package tools

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

func MD5(in string) string {
	hash := md5.New()
	hash.Write([]byte(in))
	return hex.EncodeToString(hash.Sum(nil))
}

func ParseDatetime(datetimeStr string) (int64, error) {
	if datetime, err := strconv.ParseInt(datetimeStr, 10, 64); err == nil {
		return datetime, nil
	}
	var timeLayout string
	if len(datetimeStr) == 10 {
		timeLayout = "2006-01-02" //转化所需模板
	} else {
		timeLayout = "2006-01-02 15:04:05" //转化所需模板
	}
	//获取本地location
	loc, err := time.LoadLocation("Local") //重要：获取时区
	if err != nil {
		return time.Now().Unix(), err
	}
	theTime, parseErr := time.ParseInLocation(timeLayout, datetimeStr, loc) //使用模板在对应时区转化为time.time类型
	if parseErr != nil {
		return time.Now().Unix(), parseErr
	}
	return theTime.Unix(), nil
}
