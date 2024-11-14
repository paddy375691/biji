package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 第四步：定义一个全局中间件（所有路由都会使用）
/*
MiddleWare ： 只是一个函数名字，可以随意定义
gin.HandlerFunc ：中间必须要返回的方法（如果不返回这个方法不是一个正确的中间件）
*/
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我是一个全局中间件")
	}
}

func main() {
	// 1、实例化引擎
	r := gin.Default()

	//// 中间件应用在 请求到达路由层之前
	//// 第五步：全局使用中间件
	//r.Use(MiddleWare())

	// 2、路由
	r.GET("/home", func(c *gin.Context) {
		fmt.Println("执行home")
		c.JSON(200, gin.H{"msg": "success"})
	})
	// 第五步：局部中间件（只对指定的路由进行限制）
	r.GET("/hello", MiddleWare(), func(c *gin.Context) {
		fmt.Println("执行hello")
		c.JSON(200, gin.H{"msg": "success"})
	})

	fmt.Println("http://127.0.0.1:8000/hello")
	// 3、启动服务
	r.Run(":8000")
}
