package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "http://127.0.0.1:8005/req/post"
	// 模拟form表单提交数据 contentType := "application/x-www-form-urlencoded"
	// 传json数据： json contentType := "application/json"
	contentType := "application/json"
	data := `{
		"name": "root",
		"password": "123456"
	}`
	resp, _ := http.Post(url, contentType, strings.NewReader(data))
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
}
