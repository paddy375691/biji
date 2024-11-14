package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 第二步：定义表结构（模型定义）
type User struct {
	// gorm:"primary_key" 主键索引，标记当前这个Id是自增的（1.2.3...）
	Id       int64 `json:"id" gorm:"primary_key"`
	Username string
	Password string
}

func main() {
	// 第一步：创建grom的连接
	// 配置信息要求：知道参数的意义
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	/*
		root:chnsys@2016   mysql用户名：密码
		127.0.0.1:3306     mysql 地址：端口号
		test_db            mysql中创建数据库的名字
	*/
	dsn := "root:chnsys@2016@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// err = nil 证明连接正常，否则连接失败
	if err != nil {
		panic(err) // 异常处理，有问题抛出异常
	}

	// 第三步：自动创建表结构
	// 迁移 https://gorm.io/zh_CN/docs/migration.html
	db.AutoMigrate(User{})

	//// 第四步：增删改成
	//// 4.1 增：向User表添加一条数据
	////https://gorm.io/zh_CN/docs/index.html
	//db.Create(&User{
	//	//Id: 1,  // 因为Id设置了自增，所以无需传递
	//	Username: "lisi",
	//	Password: "123456",
	//})

	//// 4.2 改：修改表的某一个字段
	//db.Model(User{
	//	Id: 1,
	//}).Update("password", "123456")

	//// 4.3 查询
	//// 过滤单体数据：First
	//u := User{Id: 1}
	//db.First(&u)
	//fmt.Printf("%#v \n", u) // main.User{Id:1, Username:"zhangsan", Password:"123456"}
	//
	//// 查询所有数据
	//users := []User{} // 定义一个User结构体的切片来接收
	//db.Find(&users)
	//fmt.Printf("%#v \n", users)
	//// []main.User{main.User{Id:1, Username:"zhangsan", Password:"123456"}, main.User{Id:2, Username:"l
	//// isi", Password:"123456"}}

	//4.4 删除
	//// 根据主键删除
	//db.Delete(&User{Id: 1})
	// 条件删除
	//db.Where("username = ?", "lisi").Delete(&User{})

}
