package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//初始化 Gin 实例
	r := gin.New()

	//注册中间件
	r.Use(gin.Logger(), gin.Recovery())

	//注册一个路由
	r.GET("/", func(c *gin.Context) {

		//以json格式响应
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World",
		})
	})

	//运行服务
	r.Run()
}
