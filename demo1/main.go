package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string
	Age int64
	Active bool
}
func main() {
	//链接MySQL数据库
	db,err :=gorm.Open("mysql","root:root@(127.0.0.1:3306)/db1?charset=utf8mb4")
	if err!=nil {
		fmt.Println(err)
		panic(err)
	}
	defer db.Close()
	//把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})
	//创建
	u1:=User{Name:"qimi",Age:18,Active:true}
	db.Create(&u1)
	u2:=User{Name:"jinzhu",Age:20,Active:false}
	db.Create(&u2)

	//查询
	var user User
	db.First(&user)

	m1:=map[string]interface{}{
		"name":"xiaoli",
		"age":29,
		"active":true,
	}
	db.Debug().Model(&user).Select("age").Update(m1)
	db.Debug().Model(&user).Omit("active").Updates(m1)
}
