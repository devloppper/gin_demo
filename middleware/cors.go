package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Cors处理")
		c.Next()
	}
}
