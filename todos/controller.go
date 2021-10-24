package todos

import "github.com/gin-gonic/gin"

var allTodos []Todo

func GetTodos(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "success",
		"todos":  allTodos,
	})
}

func AddTodos(c *gin.Context) {
	inputTodos := Todo{}

	// get ID from last Todos
	lengthTodos := len(allTodos)

	if lengthTodos == 0 {
		inputTodos.ID = 1
	} else {
		lastTodos := allTodos[lengthTodos-1]
		inputTodos.ID = lastTodos.ID + 1
	}

	err := c.ShouldBindJSON(&inputTodos)

	if err != nil {
		c.JSON(200, gin.H{
			"status": "error",
			"error":  err,
		})
		return
	}

	allTodos = append(allTodos, inputTodos)

	c.JSON(200, gin.H{
		"status":   "success",
		"messages": "Success add new todos",
		"data":     inputTodos,
	})
}
