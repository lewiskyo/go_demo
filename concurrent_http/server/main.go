package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func queryPhoneLocation(c *gin.Context) {
	numStr := c.Query("number") // 是 c.Request.URL.Query().Get("lastname") 的简写
	number, _ := strconv.Atoi(numStr)
	res := number % 10
	location := "unknown"
	switch res {
	case 0:
		location = "guangdong"
	case 1:
		location = "fujian"
	case 2:
		location = "henan"
	case 3:
		location = "guangxi"
	case 4:
		location = "tianjin"
	case 5:
		location = "hubei"
	case 6:
		location = "hunan"
	case 7:
		location = "beijing"
	case 8:
		location = "shanghai"
	case 9:
		location = "shanxi"
	}

	// 通过请求上下文对象Context, 直接往客户端返回一个json
	c.JSON(200, gin.H{
		"data": location,
	})
}

// 入口函数
func main() {
	// 初始化一个http服务对象
	r := gin.Default()

	// 设置一个get请求的路由，url为/hello, 处理函数（或者叫控制器函数）是一个闭包函数。
	r.GET("/query", queryPhoneLocation)

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
