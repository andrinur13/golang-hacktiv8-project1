package main

import (
	"golang-hacktiv8-project1/todos"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/todos", todos.GetTodos)
	router.Run()
}
