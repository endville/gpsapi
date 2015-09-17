package controllers

import (
	"encoding/json"
	"gpsapi/models"
	"gpsapi/tools"
	"time"
)

// 地理位置相关
type GeoController struct {
	BaseController
}

// @Title 获取该位置详细信息
// @Description 获取该位置详细信息
// @Param	geoId	path 	string	true	"地理位置点ID"
// @Success 200 {object} models.Geo
// @Failure 404 :not found
// @router /:geoId [get]
func (this *GeoController) Get() {
	geoId := this.GetString(":geoId")
	geo, err := models.GetGeo(geoId)
	if err != nil {
		this.ResponseErrorJSON(404, err.Error())
	} else {
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"data": geo,
		}
	}
	this.ServeJson()
}

// @Title 根据终端ID获取地理位置信息
// @Description 根据终端ID获取地理位置信息
// @Param	terminalId	path 	int	true	"终端ID"
// @Param	timeBegin	query 	string	false	"起始时间（默认昨天的现在时间点）"
// @Param	timeEnd	query 	string	false	"终止时间（默认现在）"
// @Success 200 {object} models.GeoSp
// @Failure 403 :uid is empty
// @router /terminal/:terminalId [get]
func (this *GeoController) TerminalGeo() {
	terminalId, _ := this.GetInt64(":terminalId")
	if terminalId != 0 {
		var timeBegin, timeEnd int64
		var err error
		if timeBegin, err = tools.ParseDatetime(this.GetString("timeBegin")); err != nil {
			timeBegin = time.Now().AddDate(0, 0, -1).Unix()
		}
		if timeEnd, err = tools.ParseDatetime(this.GetString("timeEnd")); err != nil {
			timeEnd = time.Now().Unix()
		}
		geos, err := models.GetTerminalGeos(timeBegin, timeEnd, terminalId)
		if err != nil {
			this.ResponseErrorJSON(403, err.Error())
		} else {
			this.Data["json"] = map[string]interface{}{
				"code":  0,
				"data":  geos,
				"total": len(*geos),
			}
		}
	} else {
		this.ResponseErrorJSON(403, "terminal id is wrong")
	}
	this.ServeJson()
}

// @Title 根据用户ID获取地理位置信息
// @Description 根据用户ID获取地理位置信息
// @Param	userId	path 	int	true	"用户ID"
// @Param	timeBegin	query 	string	false	"起始时间（默认昨天的现在时间点）"
// @Param	timeEnd	query 	string	false	"终止时间（默认现在）"
// @Success 200 {object} models.GeoSp
// @Failure 403 :uid is empty
// @router /user/:userId [get]
func (this *GeoController) UserGeo() {
	userId, _ := this.GetInt64(":userId")
	if userId != 0 {
		var timeBegin, timeEnd int64
		var err error
		if timeBegin, err = tools.ParseDatetime(this.GetString("timeBegin")); err != nil {
			timeBegin = time.Now().AddDate(0, 0, -1).Unix()
		}
		if timeEnd, err = tools.ParseDatetime(this.GetString("timeEnd")); err != nil {
			timeEnd = time.Now().Unix()
		}
		geos, err := models.GetUserGeos(timeBegin, timeEnd, userId)
		if err != nil {
			this.ResponseErrorJSON(403, err.Error())
		} else {
			this.Data["json"] = map[string]interface{}{
				"code":  0,
				"data":  geos,
				"total": len(*geos),
			}
		}
	} else {
		this.ResponseErrorJSON(403, "userId id is wrong")
	}
	this.ServeJson()
}

// @Title 根据车队ID获取地理位置信息
// @Description 根据车队ID获取地理位置信息
// @Param	groupId	path 	int	true	"车队ID"
// @Param	timeBegin	query 	string	false	"起始时间（默认昨天的现在时间点）"
// @Param	timeEnd	query 	string	false	"终止时间（默认现在）"
// @Success 200 {object} models.GeoSp
// @Failure 403 :uid is empty
// @router /group/:groupId [get]
func (this *GeoController) GroupGeo() {
	groupId, _ := this.GetInt64(":groupId")
	if groupId != 0 {
		var timeBegin, timeEnd int64
		var err error
		if timeBegin, err = tools.ParseDatetime(this.GetString("timeBegin")); err != nil {
			timeBegin = time.Now().AddDate(0, 0, -1).Unix()
		}
		if timeEnd, err = tools.ParseDatetime(this.GetString("timeEnd")); err != nil {
			timeEnd = time.Now().Unix()
		}
		geos, err := models.GetGroupGeos(timeBegin, timeEnd, groupId)
		if err != nil {
			this.ResponseErrorJSON(403, err.Error())
		} else {
			this.Data["json"] = map[string]interface{}{
				"code":  0,
				"data":  geos,
				"total": len(*geos),
			}
		}
	} else {
		this.ResponseErrorJSON(403, "groupId id is wrong")
	}
	this.ServeJson()
}

// @Title 根据车队ID获取地理位置信息
// @Description 根据车队ID获取地理位置信息
// @Param	groupId	path 	int	true	"车队ID"
// @Param	timeBegin	query 	string	false	"起始时间（默认昨天的现在时间点）"
// @Param	timeEnd	query 	string	false	"终止时间（默认现在）"
// @Param	onlyTerminal	query 	bool	false	"只获取终端号（默认true）"
// @Param	longitude	query 	float	true	"中心经度"
// @Param	latitude	query 	float	true	"中心纬度"
// @Param	maxDistance	query 	int	false	"最大半径(单位:米, 默认200)"
// @Success 200 {object} models.GeoSp2
// @Failure 403 :params error
// @router /group/round/:groupId [get]
func (this *GeoController) GetGroupGeosRound() {
	groupId, _ := this.GetInt64(":groupId")
	if groupId != 0 {
		onlyTerminal, _ := this.GetBool("onlyTerminal", true)
		longitude, errLng := this.GetFloat("longitude")
		if errLng != nil {
			this.ResponseErrorJSON(403, "经度数据不合法")
		}
		latitude, errLat := this.GetFloat("latitude")
		if errLat != nil {
			this.ResponseErrorJSON(403, "纬度数据不合法")
		}
		maxDistance, _ := this.GetInt("maxDistance", 200)

		var timeBegin, timeEnd int64
		var err error
		if timeBegin, err = tools.ParseDatetime(this.GetString("timeBegin")); err != nil {
			timeBegin = time.Now().AddDate(0, 0, -1).Unix()
		}
		if timeEnd, err = tools.ParseDatetime(this.GetString("timeEnd")); err != nil {
			timeEnd = time.Now().Unix()
		}

		if onlyTerminal {
			ids, err := models.GetGroupGeosRoundOnlyTerminalId(timeBegin, timeEnd, groupId, models.Location{float32(longitude), float32(latitude)}, maxDistance)
			if err != nil {
				this.ResponseErrorJSON(403, err.Error())
			} else {
				this.Data["json"] = map[string]interface{}{
					"code":  0,
					"data":  ids,
					"total": len(*ids),
				}
			}
		} else {
			geos, err := models.GetGroupGeosRound(timeBegin, timeEnd, groupId, models.Location{float32(longitude), float32(latitude)}, maxDistance)
			if err != nil {
				this.ResponseErrorJSON(403, err.Error())
			} else {
				this.Data["json"] = map[string]interface{}{
					"code":  0,
					"data":  geos,
					"total": len(*geos),
				}
			}
		}

	} else {
		this.ResponseErrorJSON(403, "terminal id is wrong")
	}
	this.ServeJson()
}

// @Title 根据车队ID获取地理位置信息
// @Description 根据车队ID获取地理位置信息
// @Param	groupId	path 	int	true	"车队ID"
// @Param	timeBegin	query 	string	false	"起始时间（默认昨天的现在时间点）"
// @Param	timeEnd	query 	string	false	"终止时间（默认现在）"
// @Param	onlyTerminal	query 	bool	false	"只获取终端号（默认true）"
// @Param	bottomLeftLongitude	query 	float	true	"左下角经度"
// @Param	bottomLeftLatitude	query 	float	true	"左下角纬度"
// @Param	upperRightLongitude	query 	float	true	"右上角经度"
// @Param	upperRightLatitude	query 	float	true	"右上角纬度"
// @Success 200 {object} models.GeoSp2
// @Failure 403 :params error
// @router /group/box/:groupId [get]
func (this *GeoController) GetGroupGeosBox() {
	groupId, _ := this.GetInt64(":groupId")
	if groupId != 0 {
		onlyTerminal, _ := this.GetBool("onlyTerminal", true)
		bottomLeftLongitude, errLng1 := this.GetFloat("bottomLeftLongitude")
		if errLng1 != nil {
			this.ResponseErrorJSON(403, "经度数据不合法")
		}
		bottomLeftLatitude, errLat1 := this.GetFloat("bottomLeftLatitude")
		if errLat1 != nil {
			this.ResponseErrorJSON(403, "纬度数据不合法")
		}
		upperRightLongitude, errLng2 := this.GetFloat("upperRightLongitude")
		if errLng2 != nil {
			this.ResponseErrorJSON(403, "经度数据不合法")
		}
		upperRightLatitude, errLat2 := this.GetFloat("upperRightLatitude")
		if errLat2 != nil {
			this.ResponseErrorJSON(403, "纬度数据不合法")
		}

		var timeBegin, timeEnd int64
		var err error
		if timeBegin, err = tools.ParseDatetime(this.GetString("timeBegin")); err != nil {
			timeBegin = time.Now().AddDate(0, 0, -1).Unix()
		}
		if timeEnd, err = tools.ParseDatetime(this.GetString("timeEnd")); err != nil {
			timeEnd = time.Now().Unix()
		}

		if onlyTerminal {
			ids, err := models.GetGroupGeosBoxOnlyTerminalId(timeBegin, timeEnd, groupId, models.Location{float32(bottomLeftLongitude), float32(bottomLeftLatitude)}, models.Location{float32(upperRightLongitude), float32(upperRightLatitude)})
			if err != nil {
				this.ResponseErrorJSON(403, err.Error())
			} else {
				this.Data["json"] = map[string]interface{}{
					"code":  0,
					"data":  ids,
					"total": len(*ids),
				}
			}
		} else {
			geos, err := models.GetGroupGeosBox(timeBegin, timeEnd, groupId, models.Location{float32(bottomLeftLongitude), float32(bottomLeftLatitude)}, models.Location{float32(upperRightLongitude), float32(upperRightLatitude)})
			if err != nil {
				this.ResponseErrorJSON(403, err.Error())
			} else {
				this.Data["json"] = map[string]interface{}{
					"code":  0,
					"data":  geos,
					"total": len(*geos),
				}
			}
		}
	} else {
		this.ResponseErrorJSON(403, "terminal id is wrong")
	}
	this.ServeJson()
}

// @Title 根据车队ID获取地理位置信息
// @Description 根据车队ID获取地理位置信息
// @Param	groupId	path 	int	true	"车队ID"
// @Param	timeBegin	query 	string	false	"起始时间（默认昨天的现在时间点）"
// @Param	timeEnd	query 	string	false	"终止时间（默认现在）"
// @Param	onlyTerminal	query 	bool	false	"只获取终端号（默认true）"
// @Param	locations	body 	[]models.Location	true	"点集合"
// @Success 200 {object} models.GeoSp2
// @Failure 403 :params error
// @router /group/polygon/:groupId [post]
func (this *GeoController) GetGroupGeosPolygon() {
	groupId, _ := this.GetInt64(":groupId")
	if groupId != 0 {
		onlyTerminal, _ := this.GetBool("onlyTerminal", true)
		var timeBegin, timeEnd int64
		var err error
		if timeBegin, err = tools.ParseDatetime(this.GetString("timeBegin")); err != nil {
			timeBegin = time.Now().AddDate(0, 0, -1).Unix()
		}
		if timeEnd, err = tools.ParseDatetime(this.GetString("timeEnd")); err != nil {
			timeEnd = time.Now().Unix()
		}

		locations := new([]models.Location)
		if err := json.Unmarshal(this.Ctx.Input.RequestBody, locations); err != nil {
			this.ResponseErrorJSON(400, errorFormat(ErrorBadJson_400, err.Error()))
		}

		if onlyTerminal {
			ids, err := models.GetGroupGeosPolygonOnlyTerminalId(timeBegin, timeEnd, groupId, (*locations)...)
			if err != nil {
				this.ResponseErrorJSON(403, err.Error())
			} else {
				this.Data["json"] = map[string]interface{}{
					"code":  0,
					"data":  ids,
					"total": len(*ids),
				}
			}
		} else {
			geos, err := models.GetGroupGeosPolygon(timeBegin, timeEnd, groupId, (*locations)...)
			if err != nil {
				this.ResponseErrorJSON(403, err.Error())
			} else {
				this.Data["json"] = map[string]interface{}{
					"code":  0,
					"data":  geos,
					"total": len(*geos),
				}
			}
		}
	} else {
		this.ResponseErrorJSON(403, "terminal id is wrong")
	}
	this.ServeJson()
}
