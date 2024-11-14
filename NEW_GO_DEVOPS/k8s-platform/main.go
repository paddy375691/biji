package main

import (
	"github.com/gin-gonic/gin"
	"k8s-platform/config"
	"k8s-platform/controller"
	"k8s-platform/db"
	"k8s-platform/middle"
	"k8s-platform/service"
	"net/http"
)

func main() {
	//初始化数据库
	db.Init()
	//初始化k8s client
	service.K8s.Init()   //可以使用service.K8s.ClientSet
	//初始化gin
	r := gin.Default()
	//加载jwt中间件
	//r.Use(middle.JWTAuth())
	//加载跨域中间件
	r.Use(middle.Cors())
	//跨包调用router的初始化方法
	controller.Router.InitApiRouter(r)
	//启动websocket
	go func() {
		http.HandleFunc("/ws", service.Terminal.WsHandler)
		http.ListenAndServe(":8081", nil)
	}()
	//启动gin server
	r.Run(config.ListenAddr)

	//关闭数据库连接
	db.Close()


}