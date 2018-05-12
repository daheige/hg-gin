package helper

import (
	"github.com/gin-gonic/gin"
)

func Env_str() string {
	return gin.Mode()
}
