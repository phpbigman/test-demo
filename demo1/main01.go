package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(c *gin.Context)  {
	c.JSON(200,gin.H{
		"message":"hello Golang",
	})
}
func main01(){
	//返回默认的路由引擎
	r:=gin.Default()
	r.GET("/hello",sayHello)
	r.GET("/web", func(c *gin.Context) {
		name :=c.Query("query")
		c.JSON(http.StatusOK,gin.H{
			"name":name,
		})
	})
	//获取URI路径参数
	r.GET("/:name/:age", func(c *gin.Context) {
		name:=c.Param("name")
		age:=c.Param("age")
		c.JSON(http.StatusOK,gin.H{
			"name":name,
			"age":age,
		})
	})
	//启动服务
	r.Run(":9090")
}
/**
func sayHello(w http.ResponseWriter, r *http.Request)  {
	_,_=fmt.Fprintln(w,"<h1>Hello Golang!</h1>");
}
func main() {
	http.HandleFunc("/hello",sayHello);
	err:=http.ListenAndServe(":9090",nil)
	if err!=nil {
		fmt.Println("htpp Server failed err")
		return
	}
}*/