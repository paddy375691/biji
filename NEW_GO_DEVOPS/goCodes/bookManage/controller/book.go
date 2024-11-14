package controller

import (
	"bookManage/dao/mysql"
	"bookManage/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 添加数据
func CreateBookHandler(c *gin.Context) {
	p := new(model.Book)
	if err := c.ShouldBind(p); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	mysql.DB.Create(p)
	c.JSON(200, gin.H{"msg": "success"})
}

type BookParams struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	UserId int64  `json:"user_id"`
}

// 添加数据
func CreateBookUserHandler(c *gin.Context) {
	p := new(BookParams)
	if err := c.ShouldBind(p); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	//Name  string `json:"name"`
	//Desc  string `json:"desc"`
	//Users []User `gorm:"many2many:book_users"`

	//mysql.DB.Create(p)
	book := model.Book{
		Name: p.Name,
		Desc: p.Desc,
		Users: []model.User{
			{Id: 4},
		},
	}
	mysql.DB.Create(&book)

	c.JSON(200, gin.H{"msg": "success"})
}

// 查看书籍列表
func GetBookListHandler(c *gin.Context) {
	books := []model.Book{}
	mysql.DB.Find(&books)
	c.JSON(200, gin.H{"books": books})
}

// 查看指定书籍 http://127.0.0.1:8000/book/4/
func GetBookDetailHandler(c *gin.Context) {
	bookIdStr := c.Param("id")
	bookId, _ := strconv.ParseInt(bookIdStr, 10, 64)
	book := model.Book{Id: bookId}
	mysql.DB.Find(&book)
	c.JSON(200, gin.H{"book": book})
}

// 修改操作
func UpdataBookHandler(c *gin.Context) {
	p := new(model.Book)
	if err := c.ShouldBind(p); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	fmt.Println("IIIIIIIIIIIIIIII", p.Id)
	oldBook := &model.Book{Id: p.Id}
	var newBook model.Book
	fmt.Println(p.Name, p.Desc)
	if p.Name != "" {
		newBook.Name = p.Name
	}
	if p.Desc != "" {
		newBook.Desc = p.Desc
	}
	fmt.Println(3333333, newBook.Name)
	mysql.DB.Model(&oldBook).Updates(newBook)
	c.JSON(200, gin.H{"book": newBook})
}
