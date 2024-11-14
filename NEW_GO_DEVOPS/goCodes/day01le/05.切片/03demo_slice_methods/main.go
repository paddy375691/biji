package main

import "fmt"

func main() {
	// 定义切片
	nums := []int{1, 2, 3, 4}
	fmt.Println(nums)
	fmt.Printf("长度：%v, 容量：%v \n", len(nums), cap(nums))
	// 1、apend添加一个元素
	nums = append(nums, 5)
	fmt.Println(nums)
	// 切片扩容机制：小于 1024成倍扩容，大于1024 ...
	// 长度：5, 容量：8
	fmt.Printf("长度：%v, 容量：%v \n", len(nums), cap(nums))

	// 2、apend添加多个元素
	nums = append(nums, 11, 12, 13)
	fmt.Println(nums)

	// 3、切片合并
	num1 := []int{1, 2, 3}
	num2 := []int{3, 4, 5}
	// num1.apend(num2)
	num1 = append(num1, num2...) // [1 2 3 3 4 5]
	fmt.Println(num1)

	// 4、删除切片中的第二个元素
	num3 := []int{1, 2, 3, 4, 5}
	//num3 = append(num3[:2], num3[3:]...)
	//fmt.Println(num3)
	num3 = append(num3[:1], num3[1+3:]...)
	fmt.Println(num3) // [1 5]
}
