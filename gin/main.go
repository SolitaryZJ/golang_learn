package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `json:"u" uri:"u" form:"u" binding:"required"`
	Password int    `json:"p" uri:"p" form:"p"`
}

func main() {
	router := gin.Default()
	// 调用方式
	//router.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//
	//router.Any("/login", func(c *gin.Context) {
	//	c.String(200, "login")
	//})
	//
	//grp1 := router.Group("v1")
	//{
	//	grp1.GET("/v1test", func(c *gin.Context) {
	//		c.String(200, "v1test")
	//	})
	//}

	// 加载静态文件
	//router.Static("/static", "./static")
	//router.Static("/static", http.Dir("static"))
	//router.StaticFile("favicon.ico", "./favicon.ico")

	// 模板
	//router.LoadHTMLGlob("templates/**/*")
	//router.GET("/index", func(c *gin.Context) {
	//	c.HTML(200, "sub1/user.impl", gin.H{
	//		"title": "index test",
	//	})
	//})

	// 参数传递
	// 1. 路径参数 /user/111
	//router.GET("/user/:name", func(c *gin.Context) {
	//	name := c.Param("name")
	//	c.String(200, "Hello %s", name)
	//})
	// 2. query参数 /user?name=111
	router.GET("/user", func(c *gin.Context) {
		name := c.Query("name")
		//user := c.ShouldBindQuery(&User{})
		c.String(200, "Hello %s", name)
	})

	// 3. post参数
	router.POST("/user", func(c *gin.Context) {
		name := c.PostForm("name")
		c.String(200, "Hello %s", name)
	})

	err := router.Run() // 监听并在 0.0.0.0:8080 上启动服务
	if err != nil {
		panic(err)
	}
}
