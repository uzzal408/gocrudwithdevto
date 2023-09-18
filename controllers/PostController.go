package controllers

import (
	"cruddevto/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func CreatePost(c *gin.Context) {
	var input CreatePostInput

	if err := c.ShouldBind(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
	}

	models.DB.Create(&post)
	c.JSON(http.StatusOK, gin.H{
		"data":    post,
		"message": "Post Successfully created",
	})

}

func FindPosts(c *gin.Context) {
	var posts []models.Post
	models.DB.Find(&posts)
	c.JSON(http.StatusOK, gin.H{
		"data":    posts,
		"message": "Success",
	})
}

func FindPostByID(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id=?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    post,
		"message": "Success",
	})
}

type UpdatePostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func UpdatePost(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id=?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}
	var input UpdatePostInput
	if err := c.ShouldBind(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	updatedPost := models.Post{
		Title:   input.Title,
		Content: input.Content,
	}

	models.DB.Model(&post).Updates(&updatedPost)
	c.JSON(http.StatusOK, gin.H{
		"data":    post,
		"message": "Success",
	})

}

func DeletePost(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id=?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}
	models.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}
