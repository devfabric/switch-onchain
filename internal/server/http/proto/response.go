package proto

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	FabUserOrgKey = "FabUserOrgKey"
)

const (
	INVALID_RES_AUTH = -3 //资源授权类型错误
	INVALID_PARAMS   = -2 //请求参数错误
	ERROR            = -1 //失败
	SUCCESS          = 0  //成功
)

var RespCodeFlags = map[int]string{
	SUCCESS:          "OK",
	INVALID_PARAMS:   "请求参数错误",
	ERROR:            "ERROR",
	INVALID_RES_AUTH: "资源授权类型错误",
}

type Response struct {
	Code    int         `json:"code" example:"0"`               //返回0成功 其中Data为其返回数据； 其他失败
	Message string      `json:"message,omitempty" example:"OK"` //成功"OK"； 失败为错误信息
	Data    interface{} `json:"data,omitempty"`
	Warn    string      `json:"warn,omitempty"`
}

func GetErrMsg(code int) string {
	msg, ok := RespCodeFlags[code]
	if ok {
		return msg
	}

	return RespCodeFlags[ERROR]
}

func NewResponse(code int, data interface{}) Response {
	return Response{
		Code:    code,
		Message: GetErrMsg(code),
		Data:    data,
	}
}

func NewErrorResponse(code int, err error) Response {
	return Response{
		Code:    code,
		Message: fmt.Sprintf("%s:%s", GetErrMsg(code), err.Error()),
		Data:    nil,
	}
}

func NewWarnResponse(code int, data interface{}, WarnMsg string) Response {
	return Response{
		Code:    code,
		Message: GetErrMsg(code),
		Data:    data,
		Warn:    WarnMsg,
	}
}

func ResponseAbortLogic(c *gin.Context, code int, detail error) {
	// log.Logger.Infof("uri:%s, err:%s", c.Request.RequestURI, detail.Error())
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, NewErrorResponse(code, detail))
}

func ResponseSuccessLogic(c *gin.Context, code int, data interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, NewResponse(code, data))
}

func ResponseSuccessAndWarnLogic(c *gin.Context, code int, data interface{}, warn string) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, NewWarnResponse(code, data, warn))
}
