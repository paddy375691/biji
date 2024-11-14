package main

import "fmt"

func main() {
	// 1、定义一个单号字符串
	s1 := "hello"
	fmt.Println(s1)
	// 2、定义多行字符串
	s2 := `
	第一行
	第二行
	第三行
	`
	fmt.Println(s2)
	// 3、go中str类型其实是一个只读的数组，可以使用切片
	s := "hello world"
	fmt.Printf("%#v \n", s[:2]) // "he"
	//s[0] = "abc"  // 不可以直接切片赋值

	// 4、如何修改字符串
	/*
		注：byte类型还是rune类型都是我们go中的字符串
		byte: ASCII码（ASCII不可以表示中文）
		rune: UTF8码（每个中文3个字节，每个英文是一个字节）
	*/
	ss := "美国第一"
	/*
		第一步：需要将字符串转换成 rune的 byte数组： []rune(ss)
		第二步：切片获取原有数据 ss_rune[2:]
		第三步：string方法将 byte数组转换成字符串： string(ss_rune[2:])
	*/
	// ss_rune 不是一个字符串，已经变成了一个 rune 切片
	ss_rune := []rune(ss)
	fmt.Println(ss_rune)
	// string：将其他数据类型转换成 string类型
	// :2 切片（表示数组中第二个元素往后的所有数据）
	fmt.Println(string(ss_rune[:2]))
	s_zh := "中国" + string(ss_rune[2:])
	fmt.Println(s_zh)
}
