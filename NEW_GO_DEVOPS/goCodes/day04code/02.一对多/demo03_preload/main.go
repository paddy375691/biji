package main

import (
	"encoding/json"
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
	CreditCards []CreditCard // CreditCards只是一个名字(随便写)
	// []CreditCard 类型（sting）
}

// 2.2 定义一个Card表
type CreditCard struct {
	gorm.Model
	Number string
	UserID uint // 这个就是与User表关联的外键 名字是  结构体+主键(gorm的规定)
}

/*
使用 `Preload` 方法, 在查询 `User` 时先去获取 `CreditCard` 的记录
*/

func main() {
	// 第一步：创建grom的连接
	dsn := "root:chnsys@2016@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db)

	// 第三步：通过preload方法进行一对多的查询
	users := []User{}
	db.Preload("CreditCards").Find(&users)
	strUser, _ := json.Marshal(&users)
	fmt.Println(string(strUser))

}
