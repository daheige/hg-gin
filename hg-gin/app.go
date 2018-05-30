package main

import (
	"hg-gin/hg-gin/application/routes"
	"os"

	"github.com/gin-gonic/gin"
)

//初始化环境判断 release(production),debug(development),test(testing)
//touch /etc/go.env.testing #根据不同环境创建文件
func InitEnv() {
	//默认是开发环境development
	if _, err := os.Stat("/etc/go.env.production"); err == nil { //生产环境
		gin.SetMode(gin.ReleaseMode)
	} else if _, err := os.Stat("/etc/go.env.testing"); err == nil { //测试环境
		gin.SetMode(gin.DebugMode)
		// gin.SetMode(gin.TestMode)
	} else { //开发环境
		gin.SetMode(gin.DebugMode)
	}
}

func init() {
	InitEnv()
}

func main() {
	router := gin.New()

	//待完成
	routes.WebRoute(router)

	router.Run(":8080")
}
