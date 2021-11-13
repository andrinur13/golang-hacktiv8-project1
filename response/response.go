package response

import "golang-hacktiv8-project1/todos"

type SuccessCreate struct {
	Status   string     `json:"status" example:"success"`
	Messages string     `json:"messages" example:"success add new todos"`
	Data     todos.Todo `json:"data"`
}

type SuccessDelete struct {
	Status   string `json:"status" example:"success"`
	Messages string `json:"messages" example:"success delete todos"`
}

type SuccessUpdate struct {
	Status   string `json:"status" example:"success"`
	Messages string `json:"messages" example:"success delete todos"`
}
