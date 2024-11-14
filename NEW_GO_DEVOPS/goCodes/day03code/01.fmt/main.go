package main

import "fmt"

func main() {
	m := map[string]string{
		"username": "zhangsan",
	}
	// 1、占位符
	fmt.Printf("%T %v \n", m, m)
	fmt.Printf("%T %+v \n", m, m)
	fmt.Printf("%T %#v \n", m, m)

	// Println: 自动换行
	fmt.Println("Println, 自动换行")
	fmt.Println("Println, 自动换行")

	// Print
	fmt.Print("Print, 换行")
	fmt.Print("Print, 换行")

	name := "zhangsan"
	age := 24
	// Printf : 格式化打印输出
	fmt.Printf("\n 姓名：%s age: %d \n", name, age)

	// Sprint(最常用的) ：字符串格式化
	s := fmt.Sprintf("姓名：%s age: %d", name, age)
	fmt.Println(s)
}
