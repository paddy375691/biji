package service

import (
	"fmt"
	"github.com/wonderivan/logger"
	"k8s-demo/config"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"net/http"
)

var Terminal terminal

type terminal struct {}

func(t *terminal) WsHandler(w http.ResponseWriter, r *http.Request) {
	conf, err := clientcmd.BuildConfigFromFlags("", config.Kubeconfig)
	if err != nil {
		logger.Error("创建k8s配置失败, " + err.Error())
	}
	if err := r.ParseForm(); err != nil {
		return
	}
	namespace := r.Form.Get("namespace")
	podName := r.Form.Get("pod_name")
	containerName := r.Form.Get("container_name")
	logger.Info("exec pod: %s, container: %s, namespace: %s\n", podName, containerName, namespace)

	pty, err := NewTerminalSession(w, r, nil)
	if err != nil {
		logger.Error("get pty failed: %v\n", err)
		return
	}
	defer func() {
		logger.Info("close session.")
		pty.Close()
	}()

	req := K8s.ClientSet.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace(namespace).
		SubResource("exec").
		VersionedParams(&v1.PodExecOptions{
			Container: containerName,
			Command:   []string{"/bin/bash"},
			Stdin:     true,
			Stdout:    true,
			Stderr:    true,
			TTY:       true,
		}, scheme.ParameterCodec)
	fmt.Println(req.URL())
	executor, err := remotecommand.NewSPDYExecutor(conf,"POST", req.URL())

	if err != nil {
		return
	}

	err = executor.Stream(remotecommand.StreamOptions{
		Stdin:             pty.Stdin(),
		Stdout:            pty.Stdout(),
		Stderr:            pty.Stderr(),
		TerminalSizeQueue: pty,
		Tty:               pty.Tty(),
	})

	if err != nil {
		msg := fmt.Sprintf("Exec to pod error! err: %v", err)
		logger.Info(msg)
		pty.Write([]byte(msg))
		pty.Done()
	}


}