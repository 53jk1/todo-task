package dao

import (
	models "github.com/53jk1/todo-task/dao/models"
)

type Todo interface {
	Create(models.Todo) (*models.Todo, error)
}
