package controller

import (
	"github.com/gin-gonic/gin"
	"time"
)

const (
	HTTP_SUCCESS_CODE = 200
	HTTP_ERROR_CODE   = 500
)

//api 响应的结果
type ApiResponseRes struct {
	Code         int         `json:"code"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
	ResponseTime int64       `json:response_time` //响应时间戳
}

type ApiData struct {
	Message string
	Data    interface{}
}

type BaseController struct{}

func (ctrl *BaseController) ajaxReturn(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(HTTP_SUCCESS_CODE, gin.H{
		"code":     code,
		"message":  message,
		"data":     data,
		"req_time": time.Now().Unix(),
	})
}

//请求成功返回结果
//message,data
func (ctrl *BaseController) Success(ctx *gin.Context, message string, data interface{}) {
	if len([]rune(message)) == 0 {
		message = "ok"
	}

	ctrl.ajaxReturn(ctx, HTTP_SUCCESS_CODE, message, data)
}

//错误处理code,message
func (ctrl *BaseController) Error(ctx *gin.Context, code int, message string) {
	if code <= 0 {
		code = HTTP_ERROR_CODE
	}

	ctrl.ajaxReturn(ctx, HTTP_SUCCESS_CODE, message, nil)
}
