package main

import (
	"flag"
	"fmt"
)

/*
// 1、查看帮助
go run main.go --help

// 2、非flag 命令行传参
>go run main.go zs
[zs]

// 3、flag命令行传参
>go run main.go -name "zs" 1 2
zs
[1 2]
*/
func main() {
	var name string
	/* go run main.go --help
	-name string
	      姓名 (default "zhangsan")
	*/
	flag.StringVar(&name, "name", "zhangsan", "姓名")
	/*
		&name : 传递是变量的指针类型，传入的数据赋值给他了
		"name"：run main.go -name "zs" 命令行里的名字
		"zhangsan" ： 默认值，不传递就是他
		"姓名"   --help中的提示信息
	*/
	flag.Parse()
	// 1、获取指定参数
	fmt.Println(name)

	// 2、flag.Args()：获取其他参数
	fmt.Println(flag.Args())
}
