package controllers

import (
	"crm-api/initializers"
	"crm-api/models"
	"github.com/gin-gonic/gin"
)

// CreatePost		godoc
// @Summary			Create a new post
// @Description		this method can create a new post
// @Param			post body models.PostRequest true "Create post"
// @Produce			application/json
// @Tags			Post
// @Router			/post [post]
func CreatePost(c *gin.Context) {

	var body models.PostRequest

	err := c.Bind(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error,
		})
	}

	if body.Title == "" || body.Title == "null" || body.Body == "" || body.Body == "null" {
		c.JSON(400, gin.H{
			"error": "Title and Body is required",
		})
		return
	}

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": post,
	})
}

// GetPosts		godoc
// @Summary			Get all posts
// @Description		this method can get all post
// @Produce			application/json
// @Tags			Get all post
// @Router			/post [get]
func GetPosts(c *gin.Context) {

	var posts []models.Post

	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

// GetPostById		godoc
// @Summary			Get post by id
// @Description		this method can get post by id
// @Produce			application/json
// @Tags			Get all post
// @Router			/post/id [get]
func GetPostById(c *gin.Context) {

	var post models.Post

	initializers.DB.First(&post, c.Param("id"))

	if post == (models.Post{}) {
		c.JSON(400, gin.H{
			"error": "Post not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

// UpdatePost		godoc
// @Summary			Update a post
// @Description		this method can update a post
// @Param			post body models.PostRequest true "Update post"
// @Produce			application/json
// @Tags			Post
// @Router			/post/:id [patch]
func UpdatePost(c *gin.Context) {

	id := c.Param("id")

	var post models.Post

	initializers.DB.First(&post, id)

	if post == (models.Post{}) {
		c.JSON(400, gin.H{
			"error": "Post not found",
		})
		return
	}

	var body models.PostRequest

	err := c.Bind(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error,
		})
		return
	}

	if body.Title == "" || body.Title == "null" || body.Body == "" || body.Body == "null" {
		c.JSON(400, gin.H{
			"error": "Title and Body is required",
		})
		return
	}

	post.Title = body.Title
	post.Body = body.Body

	initializers.DB.Model(&post).Updates(post)

	c.JSON(200, gin.H{
		"post": post,
	})
}

// DeletePostById		godoc
// @Summary			Delete a post
// @Description		this method can delete a post
// @Produce			application/json
// @Tags			Post
// @Router			/post/:id [delete]
func DeletePostById(c *gin.Context) {

	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)
}
