package main

import "fmt"

func main() {
	// map 是一种无序的基于 key-value 的数据结构
	//1、定义一个map字典
	/*
		[string] : 表示 key的数据类型
		第二个string：表示value值的数据类型
	*/
	userInfo := map[string]int{ // key: str   value: int
		"zhangsan": 24,
		"lisi":     28,
	}
	fmt.Println(userInfo)

	//2、判断一个key是否在map中(唯一依据是以 ok1第二个返回值true，false来判断)
	/*
		会返回两个数据
		val: 返回时查询可以的值
		ok1: 返回是一个bool值，true表示存在这个key， false不存在这个key
	*/
	val1, ok1 := userInfo["zhangsan"]
	// val=24 取出 key="zhangsan"的值， ok=true true表示有这个key，false表示没有key
	fmt.Println(val1, ok1) // 24 true
	val2, ok2 := userInfo["wangwu"]
	fmt.Println(val2, ok2) // 0 false

	// 3、和if一起常用写法
	if _, ok := userInfo["zhangsan"]; ok {
		fmt.Println(userInfo["zhangsan"])
	} else {
		fmt.Println("map中不存在这个key")
	}

	// 4、删除map中的一个数据
	fmt.Println(userInfo)
	delete(userInfo, "lisi")
	fmt.Println(userInfo)
}
