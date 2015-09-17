package models

import (
	"errors"
	"fmt"
	"labix.org/v2/mgo/bson"
	"time"
)

type Geo struct {
	Id          string   `json:"id" bson:"_id"`
	GroupId     int64    `json:"gid" bson:"gid"`
	UserId      int64    `json:"uid" bson:"uid"`
	TerminalId  int64    `json:"tid" bson:"tid"`
	Location    Location `json:"loc" bson:"loc"`
	TimeStamp   int64    `json:"timestamp" bson:"ts"`
	Direction   float32  `json:"direct" bson:"direct"`
	Status      int      `json:"status" bson:"status"`
	CellId      string   `json:"cell" bson:"cell"`
	Speed       float32  `json:"s" bson:"s"` // 速度
	Info        string   `json:"i" bson:"i"` // 三个数字：卫星数量、信号强度、电量
	Voltage     float32  `json:"v" bson:"v"` // 电压
	Temperature int      `json:"t" bson:"t"` // 温度
}

type GeoSp struct {
	Id        string   `json:"id" bson:"_id"`
	Location  Location `json:"loc" bson:"loc"`
	TimeStamp int64    `json:"timestamp" bson:"ts"`
}

type GeoSp2 struct {
	Id         string   `json:"id" bson:"_id"`
	Location   Location `json:"loc" bson:"loc"`
	TimeStamp  int64    `json:"timestamp" bson:"ts"`
	TerminalId int64    `json:"tid" bson:"tid"`
}

type Location struct {
	Longitude float32 `json:"lng" bson:"lng"`
	Latitude  float32 `json:"lat" bson:"lat"`
}

func GetGeo(id string) (*Geo, error) {
	var geo *Geo = new(Geo)

	sess, err := getMongoDbSession()
	if err != nil {
		return nil, err
	}
	defer sess.Close()

	if bson.IsObjectIdHex(id) {
		objectId := bson.ObjectIdHex(id)
		c := sess.DB(fmt.Sprintf("endville-gps-%s", objectId.Time().Format("200601"))).C("geo")
		if err := c.FindId(id).One(geo); err == nil {
			return geo, nil
		} else {
			return nil, err
		}
	} else {
		return nil, errors.New("id不合法")
	}
}

func GetTerminalGeos(timestampBegin, timestampEnd, terminalId int64) (*[]GeoSp, error) {
	endFlag := time.Unix(timestampEnd, 0).Format("200601")
	beginFlag := time.Unix(timestampBegin, 0).Format("200601")
	var geos *[]GeoSp = new([]GeoSp)

	sess, err := getMongoDbSession()
	if err != nil {
		return nil, err
	}
	defer sess.Close()

	c := sess.DB(fmt.Sprintf("endville-gps-%s", endFlag)).C("geo")

	query := bson.M{
		"ts": bson.M{
			"$gte": timestampBegin,
			"$lt":  timestampEnd,
		},
		"tid": terminalId,
	}

	if err := c.Find(query).Select(bson.M{"_id": 1, "loc": 1, "ts": 1}).Sort("ts").Limit(250).All(geos); err == nil {
		if beginFlag != endFlag {
			var geos2 *[]GeoSp = new([]GeoSp)
			c = sess.DB(fmt.Sprintf("endville-gps-%s", beginFlag)).C("geo")

			if err := c.Find(query).Select(bson.M{"_id": 1, "loc": 1, "ts": 1}).Sort("ts").Limit(250).All(geos2); err == nil {
				for _, v := range *geos2 {
					*geos = append(*geos, v)
				}
			} else {
				return nil, err
			}
		}
	} else {
		return nil, err
	}

	return geos, nil
}

func GetUserGeos(timestampBegin, timestampEnd, userId int64) (*[]GeoSp, error) {
	endFlag := time.Unix(timestampEnd, 0).Format("200601")
	beginFlag := time.Unix(timestampBegin, 0).Format("200601")
	var geos *[]GeoSp = new([]GeoSp)

	sess, err := getMongoDbSession()
	if err != nil {
		return nil, err
	}
	defer sess.Close()

	c := sess.DB(fmt.Sprintf("endville-gps-%s", endFlag)).C("geo")

	query := bson.M{
		"ts": bson.M{
			"$gte": timestampBegin,
			"$lt":  timestampEnd,
		},
		"uid": userId,
	}

	if err := c.Find(query).Select(bson.M{"_id": 1, "loc": 1, "ts": 1}).Sort("ts").Limit(250).All(geos); err == nil {
		if beginFlag != endFlag {
			var geos2 *[]GeoSp = new([]GeoSp)
			c = sess.DB(fmt.Sprintf("endville-gps-%s", beginFlag)).C("geo")

			if err := c.Find(query).Select(bson.M{"_id": 1, "loc": 1, "ts": 1}).Sort("ts").Limit(250).All(geos2); err == nil {
				for _, v := range *geos2 {
					*geos = append(*geos, v)
				}
			} else {
				return nil, err
			}
		}
	} else {
		return nil, err
	}

	return geos, nil
}

func GetGroupGeos(timestampBegin, timestampEnd, groupId int64) (*[]GeoSp, error) {
	endFlag := time.Unix(timestampEnd, 0).Format("200601")
	beginFlag := time.Unix(timestampBegin, 0).Format("200601")

	var geos *[]GeoSp = new([]GeoSp)

	sess, err := getMongoDbSession()
	if err != nil {
		return nil, err
	}
	defer sess.Close()

	c := sess.DB(fmt.Sprintf("endville-gps-%s", endFlag)).C("geo")

	query := bson.M{
		"gid": groupId,
		"ts": bson.M{
			"$gte": timestampBegin,
			"$lt":  timestampEnd,
		},
	}

	if err := c.Find(query).Select(bson.M{"_id": 1, "loc": 1, "ts": 1}).Sort("ts").Limit(250).All(geos); err == nil {
		if beginFlag != endFlag {
			var geos2 *[]GeoSp = new([]GeoSp)
			c = sess.DB(fmt.Sprintf("endville-gps-%s", beginFlag)).C("geo")

			if err := c.Find(query).Select(bson.M{"_id": 1, "loc": 1, "ts": 1}).Sort("ts").Limit(250).All(geos2); err == nil {
				for _, v := range *geos2 {
					*geos = append(*geos, v)
				}
			} else {
				return nil, err
			}
		}
	} else {
		return nil, err
	}

	return geos, nil
}

func GetGroupGeosRound(timestampBegin, timestampEnd, groupId int64, center Location, maxDistance int) (*[]GeoSp2, error) {
	endFlag := time.Unix(timestampEnd, 0).Format("200601")
	beginFlag := time.Unix(timestampBegin, 0).Format("200601")
	var geos *[]GeoSp2 = new([]GeoSp2)

	sess, err := getMongoDbSession()
	if err != nil {
		return nil, err
	}
	defer sess.Close()

	c := sess.DB(fmt.Sprintf("endville-gps-%s", endFlag)).C("geo")

	query := bson.M{
		"ts": bson.M{
			"$gte": timestampBegin,
			"$lt":  timestampEnd,
		},
		"gid": groupId,
		"loc": bson.M{
			"$geoWithin": bson.M{
				"$centerSphere": []interface{}{
					[]float32{center.Longitude, center.Latitude},
					float32(maxDistance) / 63781000.0,
				},
			},
		},
	}

	if err := c.Find(query).Select(bson.M{"_id": 1, "tid": 1, "loc": 1, "ts": 1}).Sort("ts").Limit(250).All(geos); err == nil {
		if beginFlag != endFlag {
			var geos2 *[]GeoSp2 = new([]GeoSp2)
			c = sess.DB(fmt.Sprintf("endville-gps-%s", beginFlag)).C("geo")

			if err := c.Find(query).Select(bson.M{"_id": 1, "tid": 1, "loc": 1, "ts": 1}).Sort("ts").Limit(250).All(geos2); err == nil {
				for _, v := range *geos2 {
					*geos = append(*geos, v)
				}
			} else {
				return nil, err
			}
		}
	} else {
		return nil, err
	}

	return geos, nil
}

// 获取在一个圆形范围内的所有终端设备ID
func GetGroupGeosRoundOnlyTerminalId(timestampBegin, timestampEnd, groupId int64, center Location, maxDistance int) (*[]int64, error) {
	endFlag := time.Unix(timestampEnd, 0).Format("200601")
	beginFlag := time.Unix(timestampBegin, 0).Format("200601")
	var ids *[]int64 = new([]int64)

	sess, err := getMongoDbSession()
	if err != nil {
		return nil, err
	}
	defer sess.Close()

	c := sess.DB(fmt.Sprintf("endville-gps-%s", endFlag)).C("geo")

	query := bson.M{
		"ts": bson.M{
			"$gte": timestampBegin,
			"$lt":  timestampEnd,
		},
		"gid": groupId,
		"loc": bson.M{
			"$geoWithin": bson.M{
				"$centerSphere": []interface{}{
					[]float32{center.Longitude, center.Latitude},
					float32(maxDistance) / 63781000.0,
				},
			},
		},
	}

	if err := c.Find(query).Select(bson.M{"tid": 1}).Distinct("tid", ids); err == nil {
		if beginFlag != endFlag {
			var ids2 *[]int64 = new([]int64)
			c = sess.DB(fmt.Sprintf("endville-gps-%s", beginFlag)).C("geo")
			query["tid"] = bson.M{"$nin": (*ids)}
			if err := c.Find(query).Select(bson.M{"tid": 1}).Distinct("tid", ids2); err == nil {
				for _, v := range *ids2 {
					*ids = append(*ids, v)
				}
			} else {
				return nil, err
			}
		}
	} else {
		return nil, err
	}

	return ids, nil
}

func GetGroupGeosBox(timestampBegin, timestampEnd, groupId int64, bottomLeft, upperRight Location) (*[]GeoSp2, error) {
	endFlag := time.Unix(timestampEnd, 0).Format("200601")
	beginFlag := time.Unix(timestampBegin, 0).Format("200601")
	var geos *[]GeoSp2 = new([]GeoSp2)

	sess, err := getMongoDbSession()
	if err != nil {
		return nil, err
	}
	defer sess.Close()

	c := sess.DB(fmt.Sprintf("endville-gps-%s", endFlag)).C("geo")

	query := bson.M{
		"ts": bson.M{
			"$gte": timestampBegin,
			"$lt":  timestampEnd,
		},
		"gid": groupId,
		"loc": bson.M{
			"$geoWithin": bson.M{
				"$box": [][]float32{
					[]float32{bottomLeft.Longitude, bottomLeft.Latitude},
					[]float32{upperRight.Longitude, upperRight.Latitude},
				},
			},
		},
	}

	if err := c.Find(query).Select(bson.M{"_id": 1, "tid": 1, "loc": 1, "ts": 1}).Sort("ts").Limit(250).All(geos); err == nil {
		if beginFlag != endFlag {
			var geos2 *[]GeoSp2 = new([]GeoSp2)
			c = sess.DB(fmt.Sprintf("endville-gps-%s", beginFlag)).C("geo")

			if err := c.Find(query).Select(bson.M{"_id": 1, "tid": 1, "loc": 1, "ts": 1}).Sort("ts").Limit(250).All(geos2); err == nil {
				for _, v := range *geos2 {
					*geos = append(*geos, v)
				}
			} else {
				return nil, err
			}
		}
	} else {
		return nil, err
	}

	return geos, nil
}

func GetGroupGeosBoxOnlyTerminalId(timestampBegin, timestampEnd, groupId int64, bottomLeft, upperRight Location) (*[]int64, error) {
	endFlag := time.Unix(timestampEnd, 0).Format("200601")
	beginFlag := time.Unix(timestampBegin, 0).Format("200601")
	var ids *[]int64 = new([]int64)

	sess, err := getMongoDbSession()
	if err != nil {
		return nil, err
	}
	defer sess.Close()

	c := sess.DB(fmt.Sprintf("endville-gps-%s", endFlag)).C("geo")

	query := bson.M{
		"ts": bson.M{
			"$gte": timestampBegin,
			"$lt":  timestampEnd,
		},
		"gid": groupId,
		"loc": bson.M{
			"$geoWithin": bson.M{
				"$box": [][]float32{
					[]float32{bottomLeft.Longitude, bottomLeft.Latitude},
					[]float32{upperRight.Longitude, upperRight.Latitude},
				},
			},
		},
	}

	if err := c.Find(query).Select(bson.M{"tid": 1}).Distinct("tid", ids); err == nil {
		if beginFlag != endFlag {
			query["tid"] = bson.M{"$nin": (*ids)}
			var ids2 *[]int64 = new([]int64)
			c = sess.DB(fmt.Sprintf("endville-gps-%s", beginFlag)).C("geo")
			if err := c.Find(query).Select(bson.M{"tid": 1}).Distinct("tid", ids2); err == nil {
				for _, v := range *ids2 {
					*ids = append(*ids, v)
				}
			} else {
				return nil, err
			}
		}
	} else {
		return nil, err
	}

	return ids, nil
}

func GetGroupGeosPolygon(timestampBegin, timestampEnd, groupId int64, points ...Location) (*[]GeoSp2, error) {
	endFlag := time.Unix(timestampEnd, 0).Format("200601")
	beginFlag := time.Unix(timestampBegin, 0).Format("200601")
	var geos *[]GeoSp2 = new([]GeoSp2)

	sess, err := getMongoDbSession()
	if err != nil {
		return nil, err
	}
	defer sess.Close()

	c := sess.DB(fmt.Sprintf("endville-gps-%s", endFlag)).C("geo")

	polygons := make([][]float32, len(points))
	for i, v := range points {
		polygons[i] = []float32{v.Longitude, v.Latitude}
	}
	query := bson.M{
		"ts": bson.M{
			"$gte": timestampBegin,
			"$lt":  timestampEnd,
		},
		"gid": groupId,
		"loc": bson.M{
			"$geoWithin": bson.M{
				"$polygon": polygons,
			},
		},
	}

	if err := c.Find(query).Select(bson.M{"_id": 1, "tid": 1, "loc": 1, "ts": 1}).Sort("ts").Limit(250).All(geos); err == nil {
		if beginFlag != endFlag {
			var geos2 *[]GeoSp2 = new([]GeoSp2)
			c = sess.DB(fmt.Sprintf("endville-gps-%s", beginFlag)).C("geo")

			if err := c.Find(query).Select(bson.M{"_id": 1, "tid": 1, "loc": 1, "ts": 1}).Sort("ts").Limit(250).All(geos2); err == nil {
				for _, v := range *geos2 {
					*geos = append(*geos, v)
				}
			} else {
				return nil, err
			}
		}
	} else {
		return nil, err
	}

	return geos, nil
}

func GetGroupGeosPolygonOnlyTerminalId(timestampBegin, timestampEnd, groupId int64, points ...Location) (*[]int64, error) {
	endFlag := time.Unix(timestampEnd, 0).Format("200601")
	beginFlag := time.Unix(timestampBegin, 0).Format("200601")
	var ids *[]int64 = new([]int64)

	sess, err := getMongoDbSession()
	if err != nil {
		return nil, err
	}
	defer sess.Close()

	c := sess.DB(fmt.Sprintf("endville-gps-%s", endFlag)).C("geo")

	polygons := make([][]float32, len(points))
	for i, v := range points {
		polygons[i] = []float32{v.Longitude, v.Latitude}
	}
	query := bson.M{
		"ts": bson.M{
			"$gte": timestampBegin,
			"$lt":  timestampEnd,
		},
		"gid": groupId,
		"loc": bson.M{
			"$geoWithin": bson.M{
				"$polygon": polygons,
			},
		},
	}

	if err := c.Find(query).Select(bson.M{"tid": 1}).Distinct("tid", ids); err == nil {
		if beginFlag != endFlag {
			query["tid"] = bson.M{"$nin": (*ids)}
			var ids2 *[]int64 = new([]int64)
			c = sess.DB(fmt.Sprintf("endville-gps-%s", beginFlag)).C("geo")
			if err := c.Find(query).Select(bson.M{"tid": 1}).Distinct("tid", ids2); err == nil {
				for _, v := range *ids2 {
					*ids = append(*ids, v)
				}
			} else {
				return nil, err
			}
		}
	} else {
		return nil, err
	}

	return ids, nil
}
