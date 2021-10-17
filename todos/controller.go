package todos

import "github.com/gin-gonic/gin"

var allTodos []Todo

func GetTodos(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Bismillah",
		"todos":  allTodos,
	})
}
