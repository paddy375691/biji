package service

import (
	"github.com/wonderivan/logger"
	"k8s-demo/config"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var K8s k8s

type k8s struct {
	ClientSet *kubernetes.Clientset
}

func(k *k8s) Init() {
	conf, err := clientcmd.BuildConfigFromFlags("", config.Kubeconfig)
	if err != nil {
		logger.Error("创建k8s配置失败, " + err.Error())
	}

	clientSet, err := kubernetes.NewForConfig(conf)
	if err != nil {
		logger.Error("创建k8s clientSet失败, " + err.Error())
	} else {
		logger.Info("创建k8s clientSet成功")
	}

	k.ClientSet = clientSet
}
