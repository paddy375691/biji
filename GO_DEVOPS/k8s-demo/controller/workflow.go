package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"k8s-demo/service"
	"net/http"
)

var Workflow workflow

type workflow struct {}

//创建工作流
func(w *workflow) Create(ctx *gin.Context) {
	var (
		wc = new(service.WorkflowCreate)
		err error
	)

	if err = ctx.ShouldBindJSON(wc); err != nil {
		logger.Error("Bind请求参数dc失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}

	if err = service.Workflow.CreateWorkflow(wc); err != nil {
		logger.Error("创建Deployment失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "创建工作流成功",
		"data": nil,
	})

}