package routers

import "github.com/gin-gonic/gin"

func LoadBooks(r *gin.Engine) {
	// 这个r就是main函数中 gin.Default() 返回的
	r.GET("/book", BookHandler)
}

func BookHandler(c *gin.Context) {
	c.String(200, "书籍模块")
}
