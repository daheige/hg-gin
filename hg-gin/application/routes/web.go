package routes

import (
	"fmt"
	"hg-gin/hg-gin/application/controller"
	"hg-gin/hg-gin/application/middleware"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func WebRoute(router *gin.Engine) {
	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "ok",
		})
	})

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "welcome hg-gin page",
			"data": []string{
				"php",
				"go",
			},
		})
	})

	// 将相同控制器的路由放在一个Router方法中,方便管理
	homeController := &controller.HomeController{}
	homeController.Router(router)

	//log ware
	logware := middleware.LogWare{}
	// http://mygo.com/index
	router.GET("/index", logware.AccessUri(), func(ctx *gin.Context) {
		homeController.Success(ctx, "", "fefe")
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
