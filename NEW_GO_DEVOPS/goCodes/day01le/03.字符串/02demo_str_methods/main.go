package main

import (
	"fmt"
	"strings"
)

func main() {
	// 0.0 定义一个字符串
	s := "hello world"
	s_zh := "中国"
	// 1、获取字符串的长度 len()
	fmt.Println(len(s)) // 11
	// “中国”这个汉子，字符串长度不是我们想象的2而是6
	// 一个中文占用三个字节
	fmt.Println(len(s_zh)) // 6

	// 2、字符串拼接 "+"
	s_new := s_zh + "第一"
	fmt.Println(s_new) // 中国第一

	// 3、strings.Split将字符串转换成数组
	str := "11+12+13"
	// 查看某一个方法的源码： control + 点击鼠标
	arr1 := strings.Split(str, "+")
	fmt.Printf("%T: %v \n", arr1, arr1)

	// 4、strings.Join将数组转换成字符串
	ss_new := strings.Join(arr1, "*")
	fmt.Printf("%T: %v \n", ss_new, ss_new)

	// 5、单引号智能表示一个字符
	s100 := 'a'
	//s200 := 'aa'  // 错误写法
	fmt.Println(s100)
}
