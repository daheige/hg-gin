package routes

import (
	"hg-gin/hg-gin/application/controller"
	"hg-gin/hg-gin/application/middleware"

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
	router.GET("/index", logware.AccessUri(), func(ctx *gin.Context) {
		homeController.Success(ctx, "", "fefe")
	})

}
