package controller

import (
	"github.com/gin-gonic/gin"
)

//初始化router类型的对象，首字母大写，用于跨包调用
var Router router
//声明一个router的结构体
type router struct{}

//初始化路由规则
func(r *router) InitApiRouter(router *gin.Engine) {
	router.
		//
		GET("/api/k8s/workflows", Workflow.GetList).
		GET("/api/k8s/workflow/detail", Workflow.GetById).
		POST("/api/k8s/workflow/create", Workflow.Create).
		DELETE("/api/k8s/workflow/del", Workflow.DelById).
		//
		GET("/api/k8s/pods", Pod.GetPods).
		GET("/api/k8s/pod/detail", Pod.GetPodDetail).
		DELETE("/api/k8s/pod/del", Pod.DeletePod).
		PUT("/api/k8s/pod/update", Pod.UpdatePod).
		GET("/api/k8s/pod/container", Pod.GetPodContainer).
		GET("/api/k8s/pod/log", Pod.GetPodLog).
		GET("/api/k8s/pod/numnp", Pod.GetPodNumPerNp).
		//namespace操作
		GET("/api/k8s/namespaces", Namespace.GetNamespaces).
		GET("/api/k8s/namespace/detail", Namespace.GetNamespaceDetail).
		DELETE("/api/k8s/namespace/del", Namespace.DeleteNamespace)
}