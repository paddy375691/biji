package main

import "fmt"

func main() {
	// 定义map
	userInfo := map[string]int{
		"zhangsan": 24,
		"lisi":     28,
		"wangwu":   32,
	}

	//// 1、使用for range只遍历key
	//for k := range userInfo {
	//	if val, ok := userInfo[k]; ok {
	//		fmt.Println(k, val)
	//	}
	//}
	//
	// 2、遍历key和value
	for key1, value1 := range userInfo {
		fmt.Println(key1, value1)
	}
}
