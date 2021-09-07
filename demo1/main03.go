package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
//1.定义模型
type User03 struct {
	ID int64
	Name *string `gorm:"default:'小王子'"`
	Age int64
}
func main() {
	//链接MySQL数据库
	db,err :=gorm.Open("mysql","root:root@(127.0.0.1:3306)/db1?charset=utf8mb4")
	if err!=nil {
		fmt.Println(err)
		panic(err)
	}
	defer db.Close()
	//2.把模型与数据库中的表对应起来jiao'yan
	db.AutoMigrate(&User{})
	//3.创建
	//u:=User{Name:new(string),Age:28}
	//fmt.Println(db.NewRecord(&u))
	//db.Debug().Create(&u)
	//fmt.Println(db.NewRecord(&u))

	//4.查询
	var user User//user:=new(User)
	db.First(&user)
	fmt.Printf("user:%#v\n",user)

}
