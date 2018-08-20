package controller

import (
	"fmt"
	"thinkgo/common"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

type HomeController struct {
	BaseController
}

//前置操作,返回处理器函数(中间件)
func (ctrl *HomeController) Before() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uri := ctx.Request.RequestURI
		fmt.Println(uri)
		ctx.Next()
		return
	}
}

// action
func (ctrl *HomeController) Index(ctx *gin.Context) {
	ctx.JSON(HTTP_SUCCESS_CODE, gin.H{
		"code":    200,
		"message": "ok",
		"data":    "this is test",
	})
}

func (ctrl *HomeController) Test(ctx *gin.Context) {
	ctx.String(200, "this is test page %s", "daheige")
}

//set redis
func (ctrl *HomeController) GetUserInfo(ctx *gin.Context) {
	client := common.GetRedisClient("default")
	defer client.Close()

	value, err := redis.String(client.Do("get", "myname"))
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "get redis data error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "get redis ok",
		"data":    value,
	})

}

//get redis data
func (ctrl *HomeController) SetUserInfo(ctx *gin.Context) {
	client := common.GetRedisClient("default")
	defer client.Close()

	_, err := client.Do("set", "myname", "daheige")
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "set redis data error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "set redis ok",
	})
}
