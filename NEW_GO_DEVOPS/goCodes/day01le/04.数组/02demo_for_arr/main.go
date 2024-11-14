package main

import "fmt"

func main() {
	// 1、定义一个数组
	var a = [...]string{"bj", "sh", "gz"}
	fmt.Println(a)
	// 2、普通方法通过索引遍历数组
	/*
		i := 0 =》 定义一个变量i，初始值为 0
		i < len(a)  当i的值小于数组长度，就继续循环
		i++   等价于 i = i + 1
	*/
	//for i := 0; i < len(a); i++ {
	//	fmt.Println(i, a[i])
	//}

	// 3、for range方法循环数组
	/*
		i: 是索引，只是一个变量（数组的索引）
		value： 获取的值
		:= range 是go中自己实现的方法，会返回两个值
		   - i
		   - value
	*/
	//for i, value := range a {
	for a1, b1 := range a {
		fmt.Println(a1, b1)
	}
}
