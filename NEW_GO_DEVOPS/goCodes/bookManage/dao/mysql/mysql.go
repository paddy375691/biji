package mysql

import (
	"bookManage/model"
	"fmt"
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 为什么要将变量定义在全局，是应为在其他包里需要使用
var DB *gorm.DB

func InitMysql() {
	// 1、连接数据库
	dsn := "root:chnsys@2016@tcp(127.0.0.1:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(gmysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("初始化mysql连接失败", err)
	}
	DB = db

	if err := DB.AutoMigrate(model.User{}, model.Book{}); err != nil {
		fmt.Println("创建表结构失败")
	}
}
