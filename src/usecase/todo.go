package usecase

import (
	"go_todo/domain/model"
	"go_todo/domain/repository"
)

type TodoUsecase interface {
	Search(string) ([]*model.Todo, error)
	View() ([]*model.Todo, error)
	Add(*model.Todo) (error)
	Edit(*model.Todo) (error)
}

type todoUsecase struct {
	todoRepo repository.TodoRepository
}

func NewTodoUsecase(todoRepo repository.TodoRepository) TodoUsecase {
	todoUsecase := todoUsecase{todoRepo}
	return &todoUsecase
}

func (tu *todoUsecase) Search(word string) ([]*model.Todo, error) {
	todo, err := tu.todoRepo.Find(word)
	return todo, err
}

func (tu *todoUsecase) View() ([]*model.Todo, error) {
	todo, err := tu.todoRepo.FindAll()
	return todo, err
}

func (tu *todoUsecase) Add(todo *model.Todo) (error) {
	_, err := tu.todoRepo.Create(todo)
	return err
}

func (tu *todoUsecase) Edit(todo *model.Todo) (error) {
	_, err := tu.todoRepo.Update(todo)
	return err
}