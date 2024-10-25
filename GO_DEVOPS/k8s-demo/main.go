package main

import (
	"github.com/gin-gonic/gin"
	"k8s-demo/config"
	"k8s-demo/controller"
	"k8s-demo/middle"
	"k8s-demo/service"
	"net/http"
)

func main() {
	//初始化k8s clientset
	service.K8s.Init()

	//初始化路由配置
	r := gin.Default()
	//跨域配置
	r.Use(middle.Cors())
	//jwt token验证
	//r.Use(middle.JWTAuth())
	//初始化路由
	controller.Router.InitApiRouter(r)

	//终端websocket
	go func() {
		http.HandleFunc("/ws", service.Terminal.WsHandler)
		http.ListenAndServe(":8081", nil)
	}()

	//http server启动
	r.Run(config.ListenAddr)
}
