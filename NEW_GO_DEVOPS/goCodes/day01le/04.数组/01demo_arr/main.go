package main

import "fmt"

func main() {
	// 1、定义一个数组，工作中基本都是切片
	var a [5]int
	fmt.Println(a) // [0 0 0 0 0]不是 []

	// 2、向数组中指定索引添加一个元素
	a[0] = 100
	a[1] = 200
	// 数组是有长度限制的，不会自增
	//a[10] = 200  // Invalid array index '10' (out of bounds for the 5-element array)
	fmt.Println(a) // [100 0 0 0 0]

	// 3、...让编译器自动识别我们数组长度
	var aa = [...]string{"bj", "sh", "gz"}
	fmt.Println(aa)
}
