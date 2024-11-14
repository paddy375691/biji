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

func main() {
	// 第一步：创建grom的连接
	dsn := "root:chnsys@2016@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db)

	// 第三步：查询关表数据
	/*
		1)使用 `Association` 方法, 需要把把 `User` 查询好,
		2)然后根据 `User` 定义中指定的 `AssociationForeignKey` 去查找 `CreditCard`
	*/
	// 第一步：查user表数据
	u := User{Username: "zhangsan"}
	db.First(&u) // 查询单条数据
	//fmt.Printf("%v", u.Username)
	/*
		Association("CreditCards")   管理外键的字段名
		Find(&u.CreditCards)        []CreditCard  AssociationForeignKey
	*/
	// 第二步：根据 AssociationForeignKey 外键字段进行查找
	err := db.Model(&u).Association("CreditCards").Find(&u.CreditCards)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(u)
	strUser, _ := json.Marshal(&u)
	fmt.Println(string(strUser))

}
