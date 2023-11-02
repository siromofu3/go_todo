package repository

import (
	"go_todo/domain/model"
)

type TodoRepository interface {
	FindAll() ([]*model.Todo, error)
	Find(string) ([]*model.Todo, error)
	Create(*model.Todo) (*model.Todo, error)
	Update(*model.Todo) (*model.Todo, error)
}