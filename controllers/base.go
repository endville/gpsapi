package controllers

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
)

// 400   Bad Request（错误请求） 服务器不理解请求的语法。
// 401   Unauthorized（未授权） 请求要求身份验证。 对于需要登录的网页，服务器可能返回此响应。
// 403   Forbidden（禁止） 服务器拒绝请求。
// 404   Not Found（未找到） 服务器找不到请求的网页。
//
// 500   Internal Server Error（服务器内部错误）  服务器遇到错误，无法完成请求。
// 501   Not Implemented（尚未实施） 服务器不具备完成请求的功能。
// 502   Bad Gateway（错误网关） 服务器作为网关或代理，从上游服务器收到无效响应。
// 504   Gateway Timeout（网关超时）  服务器作为网关或代理，但是没有及时从上游服务器收到请求。

var (
	ErrorBadRequest_400   = errors.New("请求不正确")
	ErrorBadJson_400      = errors.New("请求的Json格式不正确")
	ErrorBadParam_400     = errors.New("请求的参数未通过验证")
	ErrorUnauthorized_401 = errors.New("该请求未授权")
	ErrorForbidden_403    = errors.New("该请求被禁止")
	ErrorDataNotFound_404 = errors.New("服务器找不到对应的数据")
	ErrorPageNotFound_404 = errors.New("服务器找不到该请求")

	ErrorInternalServerError_500 = errors.New("服务器内部发生错误，请联系管理员")
	ErrorNotImplemented_501      = errors.New("该功能未实现")
	ErrorBadGateway_502          = errors.New("终端网关未响应")
	ErrorGatewayTimeout_504      = errors.New("终端网关超时")
)

func errorFormat(err error, details string) string {
	return fmt.Sprintf("%s. 详情:%s.\n", err.Error(), details)
}

// 网关数据监控相关
type BaseController struct {
	beego.Controller
}

func (this *BaseController) ResponseErrorJSON(code int, msg string, other ...map[string]interface{}) {
	response := map[string]interface{}{
		"code": code,
		"msg":  msg,
	}

	if other != nil && len(other) > 0 {
		for k, v := range other[0] {
			if k == "code" || k == "msg" {
				continue
			}
			response[k] = v
		}
	}
	this.Data["json"] = response
	this.ServeJson()
	this.StopRun()
}
