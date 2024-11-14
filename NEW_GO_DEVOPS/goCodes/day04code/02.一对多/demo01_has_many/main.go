package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 第二步定义一对多表结构
// 2.1 定义一个User表
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"column:username"`
	// 添加外键关联
	CreditCards []CreditCard
}

// 2.2 定义一个Card表
type CreditCard struct {
	gorm.Model
	Number string
	UserID uint // 这个就是与User表关联的外键 名字是  结构体+主键(gorm的规定)
}

func main() {
	// 第一步：创建grom的连接
	dsn := "root:chnsys@2016@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db)

	//// 第三步：创建表结构
	//db.AutoMigrate(User{}, CreditCard{})

	// 第四步：创建一对多
	//// 创建一对多
	//user := User{
	//	Username: "zhangsan",
	//	CreditCards: []CreditCard{
	//		{Number: "0001"},
	//		{Number: "0002"},
	//	},
	//}
	//db.Create(&user)

	// 给zhangsan添加信用卡
	u := User{Username: "zhangsan"}
	db.First(&u)
	db.Model(&u).Association("CreditCards").Append([]CreditCard{
		{Number: "0003"},
	})
}
