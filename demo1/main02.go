package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)


type UserInfo struct {
	ID uint
	Name string
	Gender string
	Hobby string
}
type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name string
	Age int64
}
//唯一指定表名
func (Animal) TableName() string {
	return "qimi"
}
type User1 struct {
	ID uint
	Name string
	Age sql.NullInt64 `gorm:"column:user_age"`
	Birthday time.Time
	CreateAt time.Time
}
func main() {
	//修改默认的表名
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "SMS_"+defaultTableName
	}
	//链接MySQL数据库
	db,err :=gorm.Open("mysql","root:root@(127.0.0.1:3306)/db1?charset=utf8mb4")
	if err!=nil {
		fmt.Println(err)
		panic(err)
	}
	defer db.Close()
	db.SingularTable(true)//禁用复数
	//创建表 自动迁移
	//db.AutoMigrate(&UserInfo{})

	//使用User结构体创建名叫xiaowangzi的表
	db.Table("xiaowangzi").CreateTable(&User{})

	//创建数据行
	//u1:=UserInfo{1,"qimi","男","蛙泳"}
	//db.Create(&u1)
/*
	//查询
	var u UserInfo
	db.First(&u)
	fmt.Printf("u:%#v",u)
	//更新
	db.Model(&u).Update("hobby","双色球")
	//删除
	db.Delete(&u)
*/
}
