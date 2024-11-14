package main

import "fmt"

func main() {
	/*
		第一步：声明了一个变量a
		第二步：在内存中申请了一块内存空间，用来存放 100这个值
		第三步：要想用a变量取出100这个值，实际先获取到的是100的内存地址
	*/
	a := 100

	// 1、指针地址概念
	fmt.Println(&a) // &a打印是内存地址 0xc00000a0b8
	fmt.Printf("%d \n", &a)
	b := &a

	// 2、*&a 取出值
	fmt.Println(*b) // 100

	// 3、获取指针类型
	// *int（指针类型）  0xc00000a0b8 内存地址位置
	fmt.Printf("%T: %v", &a, &a)
}
