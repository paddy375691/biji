package main

import (
	"fmt"
	"os"
)

func main() {
	//// 1、获取目录
	//// C:\_online\goCodes\day03code
	//fmt.Println(os.Getwd())
	//// 2、cd 切换路径
	//os.Chdir("/Users")
	//// C:\Users
	//fmt.Println(os.Getwd())
	//// 3、创建文件夹
	//// test_dir: 文件夹名字  0777：可读可写可执行
	//os.Mkdir("test_dir", 0777)
	//// 4、删除文件夹
	//os.Remove("test_dir")
	//// 5、重命名文件
	//os.Rename("test_dir", "xxx")
	//// 6、新建文件夹
	//os.Create("./file.txt")

	// 7、打开并写入文件
	/*
		O_RDONLY	打开只读文件
		O_WRONLY	打开只写文件
		O_RDWR	打开既可以读取又可以写入文件
		O_APPEND	写入文件时将数据追加到文件尾部
		O_CREATE	如果文件不存在，则创建一个新的文件
	*/
	file, err := os.OpenFile("file.txt", os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.WriteString("hello world")
	if err != nil {
		fmt.Println(err)
	}
}
