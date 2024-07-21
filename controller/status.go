package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetStatus 获取状态
// @GET /status
func GetStatus(c *gin.Context) {
	// 当访问 localhost:8081/status 会执行这个函数
	c.JSON(http.StatusOK, map[string]string{
		"name": "张三",
	})
	// http:example.com:8081/status?name=xxx 怎么能拿到xxx
	// http:example.com:8081/status/1231	怎么拿到1231

}

func PostStatus(c *gin.Context) {
	// http:example.com:8081/post/status/ -b {name:'xxxx'} 如何获取到 name:xxxx 这个请求体
	// 怎么给接口响应一个json、txt、或者一段mp3
}
