package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

/*
模拟美团外卖地址
一个用户可以填写多个地址
一个地址可能被多个用户填写
*/

// 第二步：自定义第三张表
type Person struct {
	ID   int
	Name string `gorm:"many2many:person_addresses"`
}

type Address struct {
	ID   uint
	Name string
}

// 自定义第三张表
type PersonAddress struct {
	PersonID  int `gorm:"primaryKey"` // 对应表的主键
	AddressId int `gorm:"primaryKey"`
	CreateAt  time.Time
}

func main() {
	// 第一步：创建grom的连接
	dsn := "root:chnsys@2016@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db)

	// 第三步：创建第三张表
	db.AutoMigrate(Person{}, Address{}, PersonAddress{})
}
