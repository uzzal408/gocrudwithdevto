package main

import (
	"cruddevto/controllers"
	"cruddevto/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.FindPosts)
	r.GET("/posts/:id", controllers.FindPostByID)
	r.PATCH("/posts/:id", controllers.UpdatePost)
	r.DELETE("posts/:id", controllers.DeletePost)
	r.Run(":8080")
}
