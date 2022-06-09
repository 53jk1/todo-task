package db

import (
	models "github.com/53jk1/todo-task/dao/models"
	"gorm.io/gorm"
)

type Todo struct {
	Db *gorm.DB
}

func (t *Todo) Create(todo models.Todo) (*models.Todo, error) {
	if err := t.Db.Create(&todo).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (t *Todo) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	if err := t.Db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}