package todos

import (
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
)

var allTodos []Todo

// Get Todos
// @Summary Get All Data Todos
// @Description Get All Todos
// @Accept json
// @Produce json
// @Success 200 {object} Todos
// @Router /todos [get]
func GetTodos(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "success",
		"todos":  allTodos,
	})
}

// Add Todos
// @Summary Post New Data Todos
// @Description Post New Data Todos
// @Accept json
// @Produce json
// @Success 200 {object} Todo
// @Params todo body Todo true "Create Todo"
// @Router /todos [post]
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

// Delete Todos
// @Summary Delete Data Todos
// @Description Delete Data Todos
// @Accept json
// @Produce json
// @Success 200
// @Router /todos [delete]
func DeleteTodos(c *gin.Context) {
	stringId := c.Param("id")
	Id, _ := strconv.Atoi(stringId)
	for i, todo := range allTodos {
		if todo.ID == Id {
			allTodos = append(allTodos[:i], allTodos[i+1:]...)
			return
		}
	}
	c.JSON(200, gin.H{
		"status":   "success",
		"messages": "Success delete todos",
		"todos":    allTodos,
	})
}

// Update Todos
// @Summary Put Edit Data Todos
// @Description Edit Data Todos
// @Accept json
// @Produce json
// @Success 200
// @Router /todos [put]
func UpdateTodos(c *gin.Context) {
	stringId := c.Param("id")
	Id, _ := strconv.Atoi(stringId)
	for i, todo := range allTodos {
		if todo.ID == Id {
			allTodos = append(allTodos[:i], allTodos[i+1:]...)
			var updatedTodo Todo
			json.NewDecoder(c.Request.Body).Decode(&updatedTodo)
			allTodos = append(allTodos, updatedTodo)
			json.NewEncoder(c.Writer).Encode(updatedTodo)
			return
		}
	}
}
