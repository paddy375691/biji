package main

import (
	"bookManage/dao/mysql"
	"bookManage/router"
)

func main() {
	// 初始化mysql
	mysql.InitMysql()
	// 将实例化router服务的方法拆分到router文件件
	r := router.InitRouter()
	// 3、启动服务
	r.Run(":8000")
}
