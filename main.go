package main

import (
	"fmt"
	"gohub/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	//初始化 Gin 实例
	router := gin.New()

	//初始化路由绑定
	bootstrap.SetupRoute(router)

	//运行服务
	err := router.Run(":3000")
	if err != nil {
		//错误处理， 端口被占用了或者其他错误
		fmt.Println(err.Error())
	}

	// //注册中间件
	// r.Use(gin.Logger(), gin.Recovery())

	// //注册一个路由
	// r.GET("/", func(c *gin.Context) {

	// 	//以json格式响应
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"Hello": "World",
	// 	})
	// })

	// //处理404请求
	// r.NoRoute(func(c *gin.Context) {
	// 	//获取标头信息的 Accept 信息
	// 	acceptString := c.Request.Header.Get("Accept")
	// 	if strings.Contains(acceptString, "text/html") {
	// 		//如果是HTML的话
	// 		c.String(http.StatusNotFound, "页面返回404")
	// 	} else {
	// 		//默认返回json
	// 		c.JSON(http.StatusNotFound, gin.H{
	// 			"error_code":    404,
	// 			"error_message": "路由未定义， 请确认 url 和请求方法是否正确",
	// 		})
	// 	}

	// })

	// //运行服务
	// r.Run(":8009")
}
