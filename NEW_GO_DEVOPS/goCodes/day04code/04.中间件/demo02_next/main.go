package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("开始执行中间件")
		c.Next() // 执行视图函数
		fmt.Println("视图函数执行完成后再调用next()方法")
	}
}

func main() {
	r := gin.Default()
	r.Use(MiddleWare())
	r.GET("/hello", func(c *gin.Context) {
		fmt.Println("执行了Get方法")
		c.JSON(200, gin.H{"msg": "success"})
	})
	r.Run()
}

/*
开始执行中间件
执行了Get方法
视图函数执行完成后再调用next()方法
*/
