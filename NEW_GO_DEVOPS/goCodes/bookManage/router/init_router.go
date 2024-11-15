package router

import "github.com/gin-gonic/gin"

/*
加载其他路由文件中的路由
*/

// 这个方法作用：初始化其他文件中的路由
func InitRouter() *gin.Engine {
	// 1、初始化gin服务
	r := gin.Default()
	// LoadTestRouter 和当前的方法在同一个包里，可以不用导入
	LoadTestRouter(r)
	LoadAPIRouter(r)
	return r
}
