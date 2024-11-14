package main

import "fmt"

func main() {
	// 变量和常量的区别
	// 多行注释 /* */
	// 单号注释  //
	// 快捷键注释：control + /
	/*
		变量： 值可以改变的 var
		常量： 不可以改变  const
	*/

	// 1、定义变量
	// num 变量名字； int 变量类型；100 变量值
	var num int = 100 // 方法1
	num2 := 100       // 方法2（最常用的）
	var num3 int      // 方法3（先定义类型然后赋值，全局变量初始化）
	num3 = 100
	fmt.Println(num, num2, num3)

	// 2、一次定义多个变量
	var num4, num5 int
	num4 = 100
	num5 = 100
	fmt.Println(num4, num5)

	// 3、批量声明
	var (
		name string
		age  int
	)
	name = "zhangsan"
	age = 24
	fmt.Println(name, age)

	// 4、定义常量
	const num10 = 100
	//num10 = 200  // 常量一旦定义，不可修改
	num11 := 100 // 变量可以修改同一种类型数据
	num11 = 200
	//num11 = "zhangsan"  // 不能将一个int类型的数据，修改成其他类型
	fmt.Println(num10, num11)

	// 批量定义常量
	const (
		// 在同一个代码块里，相同的变量名字不能同时指定为变量和常量
		//name string = "zhangsan"
		age1 = 26
	)
	fmt.Println(name, age)
}
