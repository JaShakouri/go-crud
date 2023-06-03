package main

import (
	"crm-api/controllers"
	"crm-api/docs"
	"crm-api/initializers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

// @title 	Crm crud project
// @version	1.0.0
// @description A test sample for tutorial golang

// @host 	localhost:3000
// @BasePath /api/v1

func main() {
	r := gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	docs.SwaggerInfo.Title = "Crm app"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Description = "A test sample for tutorial golang"

	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.BasePath = "/api/v1/"

	v1 := r.Group("/api")
	{
		eg := v1.Group("/v1")
		{
			// Create a new post
			eg.POST("/post", controllers.CreatePost)

			// Get all post list
			eg.GET("/post", controllers.GetPosts)

			// Get one post by id
			eg.GET("/post/:id", controllers.GetPostById)

			// Delete post by id
			eg.DELETE("/post/:id", controllers.DeletePostById)

			// Update post by id
			eg.PATCH("/post/:id", controllers.UpdatePost)

		}
	}

	err := r.Run()
	if err != nil {
		log.Fatal("Run application is failure")
	}
}
