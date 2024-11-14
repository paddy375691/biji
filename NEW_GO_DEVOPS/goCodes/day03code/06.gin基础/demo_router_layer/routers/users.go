package routers

import "github.com/gin-gonic/gin"

// 分层注册路由
func LoadUsers(r *gin.Engine) {
	r.GET("/user", UserHandler)
	//e.GET("/user", UserHandler)
	//e.GET("/user", UserHandler)
	//e.GET("/user", UserHandler)
	//e.GET("/user", UserHandler)
	//e.GET("/user", UserHandler)
	//e.GET("/user", UserHandler)
	//e.GET("/user", UserHandler)
}

func UserHandler(c *gin.Context) {
	c.String(200, "用户模块")
}
