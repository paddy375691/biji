package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
第一：路由的四种写法
第二：获取路由传参的方式以及获取post请求的数据
？连接数据库查询
第三：如何返回数据 返回json、返回string数据
*/

func main() {
	r := gin.Default()
	r.GET("/response", ResponseHandler)
	r.GET("/response/json", ResponseJsonHandler)
	r.GET("/response/redirect", ResponseRedirectHandler)
	r.Run(":8000")
}

// 1、响应一个普通的String字符串
func ResponseHandler(c *gin.Context) {
	c.String(200, "响应一个string字符串")
}

// 2、返回一个json数据（最常用）
func ResponseJsonHandler(c *gin.Context) {
	//type Data struct {
	//	Msg  string `json:"msg"`
	//	Code int    `json:"code"`
	//}
	//// 从数据库查到的数据
	//d := Data{
	//	Msg:  "Success",
	//	Code: 1001,
	//}
	//c.JSON(200, d)
	c.JSON(200, gin.H{
		"msg":  "success",
		"code": 1001,
	})
}

// 3、路由重定向(基本不怎么用)
func ResponseRedirectHandler(c *gin.Context) {
	time.Sleep(time.Second * 3)
	c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
}
