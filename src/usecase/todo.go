package usecase

import (
	"log"
	"go_todo/domain/model"
	"go_todo/domain/repository"
)

type TodoUsecase interface {
	Search(string) (todo []*model.Todo, err error)
	View() (todo[]*model.Todo, err error)
	Add(*model.Todo) (err error)
	Edit(*model.Todo) (err error)
}

type todoUsecase struct {
	todoRepo repository.TodoRepository
}

func NewTodoUsecase(todoRepo repository.TodoRepository) TodoUsecase {
	todoUsecase := todoUsecase{todoRepo}
	return &todoUsecase
}

func (tu *todoUsecase) Search(word string) (todo []*model.Todo, err error) {
	todo, err = tu.todoRepo.Find(word)
	return
}

func (tu *todoUsecase) View() (todo []*model.Todo, err error) {
	log.Print("TodoUsecase View")
	todo, err = tu.todoRepo.FindAll()
	return
}

func (tu *todoUsecase) Add(todo *model.Todo) (err error) {
	_, err = tu.todoRepo.Create(todo)
	log.Print(err)
	return
}

func (tu *todoUsecase) Edit(todo *model.Todo) (err error) {
	_, err = tu.todoRepo.Update(todo)
	return
}