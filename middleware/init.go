package middleware

import "github.com/gin-gonic/gin"

func InitMD(r *gin.Engine) {
	r.Use(Cors())
}
