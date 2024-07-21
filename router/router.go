package router

import (
	"gin_demo/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// 注册一个get请求
	r.GET("/status", controller.GetStatus)
	// 怎么用
	r.POST("")
	// 怎么用
	r.Group("")
}
