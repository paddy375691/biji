package main

import (
	// demo_router_layer 项目根路径
	// go.mod文件中指定 module demo_router_layer
	"github.com/gin-gonic/gin"
	"xxx/routers" // 导入本地的模块
)

/*
1、初始化项目
go mod init xxx
go mod tidy    // 更新项目中使用的模块
go get         // 下载包
go.mod文件中指定 module demo_router_layer （指定导入根路径）
*/

func main() {
	// r本质 实例化 *gin.Engine结构体
	r := gin.Default() // 实例化engine引擎实例
	// 注册路由
	routers.LoadUsers(r)
	routers.LoadBooks(r)

	//utils.timeConv()
	r.Run(":8000")
}
