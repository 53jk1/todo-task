package dao

import (
	models "github.com/53jk1/todo-task/models"
)

type Todo interface {
	Create (models.Todo)
}