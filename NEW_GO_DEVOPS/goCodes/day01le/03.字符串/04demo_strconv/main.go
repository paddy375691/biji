package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1、将 int 转换成字符串
	num := 100
	fmt.Printf("%T %d \n", num, num) // int 100
	strNum := strconv.Itoa(num)
	fmt.Printf("%T %v \n", strNum, strNum) // string 100

	// 2、将字符串转int
	// intNum接收转换后的数据，err接收转换如果报错结果
	intNum, err := strconv.Atoi(strNum)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T: %v \n", intNum, intNum)
}
