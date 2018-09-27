package controller

import (
	"fmt"
	"log"
	"thinkgo/common"
	"time"

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
	//通过goroutine独立运行
	newCtx := ctx.Copy()
	go func() {
		// time.Sleep(12 * time.Second)
		// log.Println(newCtx)
		log.Println(newCtx.ClientIP())
	}()

	go func() {
		client := common.GetRedisClient("default")
		defer client.Close()

		log.Println(client.Do("hsetnx", "myhash", "daheige", time.Now().Format("2016-01-02 15:04:05")))
		//设置过期时间

		timeStr := time.Now().Format("2006-01-02")
		fmt.Println("timeStr:", timeStr)
		t, _ := time.Parse("2006-01-02", timeStr)
		timeNumber := t.Unix()
		fmt.Println("timeNumber:", timeNumber)

		//获取本地location
		toBeCharge := timeStr                                           //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
		timeLayout := "2006-01-02"                                      //转化所需模板
		loc, _ := time.LoadLocation("PRC")                              //重要：获取时区
		theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
		sr := theTime.Unix()                                            //转化为时间戳 类型是int64
		fmt.Println(theTime)                                            //打印输出theTime 2015-01-01 15:15:00 +0800 CST
		fmt.Println(sr)

		client.Do("expireAt", "myhash", sr+86400*3) //设置过期时间
		log.Println("do once")
		log.Println("xxx")
	}()

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
