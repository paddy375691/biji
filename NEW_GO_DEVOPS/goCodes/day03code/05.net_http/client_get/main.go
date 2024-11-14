package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

/*
使用net/http模块，作为web的客户端发送get请求
应用场景:
	1) 爬虫，获取页面数据
	2) 调用其他服务中的接口
*/
func main() {
	//// 1、直接通过url拼接处url字符串
	//apiUrl := "http://127.0.0.1:8005/req/get?name=zhangsan"
	apiUrl := "http://127.0.0.1:8005/req/get"

	// 2、通过url进行解析
	// http://127.0.0.1:8005/req/get?name=zhangsan
	// http://v5blog.cn/pages/d19d5a/#_01-%E6%89%93%E5%BC%80%E5%92%8C%E5%85%B3%E9%97%AD%E6%96%87%E4%BB%B6
	data := url.Values{}
	data.Set("name", "zhangsan")
	u, _ := url.ParseRequestURI(apiUrl)
	u.RawQuery = data.Encode()
	fmt.Println(u.String())

	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
