package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin"
)

func main() {
	// 1、创建路由engine
	// r就是 *Engine 结构体
	r := gin.Default()
	// 2、路由绑定
	/*
		"/" 路由
		func(c *gin.Context)  处理函数
	*/
	r.GET("/", func(c *gin.Context) {
		// 第一：解析get/post请求的参数
		// 第二：根据参数去查询数据库（操作数据库：添加数据、删除数据）
		// 第三：返回从数据库查询的数据
		c.String(200, "hello world")
	})
	fmt.Println("http://127.0.0.1:8000")
	// 3、启动监听端口
	// 对 net/http服务的封装，替换了 http.ListenAndServe(address, engine)
	r.Run(":8000")
}
