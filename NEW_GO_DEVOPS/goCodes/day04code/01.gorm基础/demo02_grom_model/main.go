package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id           int64      `gorm:"primary_key"  json:"id"`
	Name         string     // 默认字符串对应的是数据库的text文本类型
	CreatedAt    *time.Time `json:"createdAt" gorm:"column:create_at"`
	Email        string     `gorm:"type:varchar(100);unique_index"` // 唯一索引
	Role         string     `gorm:"size:255"`                       //设置字段的大小为255个字节
	MemberNumber *string    `gorm:"unique;not null"`                // 设置 memberNumber 字段唯一且不为空
	Num          int        `gorm:"AUTO_INCREMENT"`                 // 设置 Num字段自增
	Address      string     `gorm:"index:addr"`                     // 给Address 创建一个名字是  `addr`的索引
	IgnoreMe     int        `gorm:"-"`                              //忽略这个字段
}

func main() {
	// 第一步：创建grom的连接
	dsn := "root:chnsys@2016@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(User{})

}
