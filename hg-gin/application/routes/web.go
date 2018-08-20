package routes

import (
	"fmt"
	"hg-gin/hg-gin/application/controller"
	"hg-gin/hg-gin/application/middleware"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type res struct {
	Id    int
	Books []string
	Desc  string
}

func WebRoute(router *gin.Engine) {
	homeController := &controller.HomeController{}
	router.GET("/json", homeController.Index)

	//设置分组路由
	v1 := router.Group("v1").Use(homeController.Before()) //采用中间件
	v1.GET("/test", homeController.Test)

	// http://mygo.com/v1/hg
	//{"data":[{"Id":1,"Books":["bo","php"],"Desc":"program"},{"Id":2,"Books":["golang","php"],"Desc":"study notes"}]}
	v1.GET("/hg", func(ctx *gin.Context) {
		data := []res{
			{
				Id:    1,
				Books: []string{"bo", "php"},
				Desc:  "program",
			},
			{
				Id:    2,
				Books: []string{"golang", "php"},
				Desc:  "study notes",
			},
		}

		ctx.JSON(controller.HTTP_SUCCESS_CODE, gin.H{
			"data": data,
		})

	})

	v1.GET("/user-info", func(ctx *gin.Context) {
		homeController.Success(ctx, "", []string{"golang", "javascript"})
	})

	v1.GET("/get-user", homeController.GetUserInfo)
	v1.GET("/set-user", homeController.SetUserInfo)

	router.GET("/error", func(ctx *gin.Context) {
		homeController.Error(ctx, controller.HTTP_ERROR_CODE, "参数错误,错误的请求")
	})

	//log ware
	logware := middleware.LogWare{}
	router.GET("/", logware.AccessUri(), func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "welcome hg-gin page",
			"data": []string{
				"php",
				"go",
			},
		})
	})

	//http://mygo.com/form?username=daheige&sex=1
	router.GET("/form", func(ctx *gin.Context) {
		username := ctx.DefaultQuery("username", "")
		age := ctx.DefaultQuery("age", "0")
		sex, err := strconv.Atoi(ctx.Query("sex")) //convert to int
		if err != nil {
			ctx.String(500, "age err: %s", err)
			return
		}

		fmt.Printf("age type is %T", age)
		fmt.Printf("sex type is %T", sex)
		ctx.String(http.StatusOK, "username is %s ,age is %s", username, age)
	})

}
