package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type LogWare struct{}

//uri access log
func (this *LogWare) AccessUri() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uri := ctx.Request.RequestURI
		fmt.Println("当前访问uri: ", uri)
		ctx.Next()
		return
	}
}
