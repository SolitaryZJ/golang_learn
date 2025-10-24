package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/task/web/docs"
	"github.com/task/web/facade"
	"github.com/task/web/middleware"
)

func main() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(middleware.RequestLogger())
	user := router.Group("/user")
	{
		user.POST("/login", facade.Login)
		user.POST("/register", facade.Register)
	}
	post := router.Group("/post", middleware.JWTAuth())
	{
		post.POST("/create", facade.CreatePost)
		post.GET("/read/:id", facade.GetPost)
		post.PUT("/update", facade.UpdatePost)
		post.DELETE("/delete", facade.DeletePost)
		post.GET("/list", facade.ListPost)
	}
	comment := router.Group("/comment", middleware.JWTAuth())
	{
		comment.POST("/create", facade.CreateComment)
		comment.GET("/list", facade.ListComments)
	}
	router.Run()
}
