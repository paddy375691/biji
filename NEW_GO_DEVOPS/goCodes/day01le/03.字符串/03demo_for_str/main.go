package main

import "fmt"

func main() {
	s := "hello 张三"

	// 1、for遍历字符串常见问题
	// %d取数字 %c取字符串
	for i := 0; i < len(s); i++ {
		fmt.Printf("%T: %v: %c \n", s[i], s[i], s[i])
	}

	//// 2、使用range遍历字符串
	// _下划线的意思是，占位符，下面不引用不会报错
	//for _, val := range s {
	//	//fmt.Println(index, val)
	//	fmt.Printf("%T: %v: %c \n", val, val, val)
	//}
}
