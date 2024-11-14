package service

import (
	"github.com/wonderivan/logger"
	"k8s-platform/config"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

//用于初始化k8s clientset

var K8s k8s

type k8s struct{
	ClientSet *kubernetes.Clientset
}
//初始化
func(k *k8s) Init() {
	//将kubeconfig文件转换成rest.config类型的对象
	conf, err := clientcmd.BuildConfigFromFlags("", config.Kubeconfig)
	if err != nil {
		panic("获取k8s client配置失败, " + err.Error())
	}
	//根据rest.config类型的对象，new一个clientset出来
	clientset, err := kubernetes.NewForConfig(conf)
	if err != nil {
		panic("创建k8s client失败, " + err.Error())
	} else {
		logger.Info("k8s client 初始化成功!")
	}

	k.ClientSet = clientset
}