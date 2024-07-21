package main

import (
	"fmt"
	"gin_demo/config"
	"gin_demo/middleware"
	"gin_demo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// 获取一个ENGINE
	r := gin.Default()
	// 注册中间件
	middleware.InitMD(r)
	// 注册一个路由
	router.InitRouter(r)
	// 启动
	_ = r.Run(fmt.Sprintf(":%d", config.Default.Server.Port))
}
