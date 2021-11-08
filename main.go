package main

import (
	"golang-hacktiv8-project1/docs"
	"golang-hacktiv8-project1/todos"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Project 1 Kampus Merdeka X Hacktiv8 API Docs
// @version 1.0
// @description This is a sample server celler server.
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @query.collection.format multi
func main() {
	router := gin.Default()

	router.GET("/todos", todos.GetTodos)
	router.POST("/todos", todos.AddTodos)
	router.PUT("/todos", todos.UpdateTodos)
	router.DELETE("/todos", todos.DeleteTodos)

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
