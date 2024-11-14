package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 中间件进行身份验证
func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {
		fmt.Println("执行Auth-------------------》")
		// 通过 c.Request的结构体我们可以取出http header中的token信息
		token := c.Request.Header.Get("token")
		fmt.Println("获取token为：", token)
		if token != "twgdh" {
			c.String(403, "身份验证不通过")
			c.Abort() // 终止当前请求，不会讲请求转发给路由，所以 处理函数不会执行
			return
		}
		c.Next() // 可以向下执行
		fmt.Println("当处理函数执行完成，然后的一些钩子函数")
	}
}

// 中间件进行身份验证
func AuthNew() gin.HandlerFunc {

	return func(c *gin.Context) {
		fmt.Println("执行AuthNew-------------------》")
		// 通过 c.Request的结构体我们可以取出http header中的token信息
		token := c.Request.Header.Get("token")
		fmt.Println("获取token为：", token)
		if token != "twgdh" {
			c.String(403, "身份验证不通过")
			c.Abort() // 终止当前请求，不会讲请求转发给路由，所以 处理函数不会执行
			return
		}
		c.Next() // 可以向下执行
		fmt.Println("当处理函数执行完成，然后的一些钩子函数")
	}
}

func main() {
	r := gin.Default()

	// 首页，无需登录直接访问
	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "index无需登录直接访问"})
	})
	// home页面，需要登录才能访问
	r.GET("/home", Auth(), AuthNew(), func(c *gin.Context) {
		fmt.Println("--------------》执行了home函数")
		c.JSON(200, gin.H{"msg": "home需要登录验证token"})
	})
	r.Run(":8000")
}
