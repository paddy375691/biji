package main

import "fmt"

func main() {
	// 1、声明引用类型，必须要使用make或者new方法申请内存空间
	// 我们只是声明了一个变量，但是没有到内存中申请空间
	/*
		var userInfo map[string]int
		// panic: assignment to entry in nil map

		userInfo["zhangsan"] = 24
		fmt.Println(userInfo)
	*/

	// 2、使用make方法申请空间(返回值类型)
	/*
		第一个 []int : 指定数据类型是一个 int切片
		第二 3 10 ： 切片默认长度是3， 默认容量是 10
	*/
	a := make([]int, 3, 10)
	a = append(a, 1)
	fmt.Println(a)
	fmt.Println(len(a), cap(a)) // 3 10

	// 3、通过new方法申请空间（指针类型）
	b := new([]int)
	*b = append(*b, 2)
	fmt.Printf("make:%T  ==> new: %T \n", a, b)
	fmt.Println(b)
}
