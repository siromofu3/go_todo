package infra

import (
	"fmt"
	"go_todo/domain/model"
	"go_todo/domain/repository"
)

type TodoRepository struct {
	SqlHandler
}

func NewTodoRepository(sqlHandler SqlHandler) repository.TodoRepository {
	todoRepository := TodoRepository{sqlHandler}
	return &todoRepository
}

func (todoRepo *TodoRepository) FindAll() ([]*model.Todo, error) {
	todos := []*model.Todo{}

	rows, err := todoRepo.SqlHandler.Conn.Query("SELECT * FROM todos")
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return todos, err
	}

	for rows.Next() {
		todo := model.Todo{}
		rows.Scan(&todo.ID, &todo.Task, &todo.LimitDate, &todo.Status)
		todos = append(todos, &todo)
	}
	return todos, err
}

func (todoRepo *TodoRepository) Find(word string) ([]*model.Todo, error) {
	todos := []*model.Todo{}

	rows, err := todoRepo.SqlHandler.Conn.Query("SELECT * FROM todos WHERE task LIKE ?", "%"+word+"%")
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return todos, err
	}

	for rows.Next() {
		todo := model.Todo{}
		rows.Scan(&todo.ID, &todo.Task, &todo.LimitDate, &todo.Status)
		todos = append(todos, &todo)
	}
	return todos, err
}

func (todoRepo *TodoRepository) Create(todo *model.Todo) (*model.Todo, error) {
	_, err := todoRepo.SqlHandler.Conn.Exec("INSERT INTO todos (task,limit_date,status) VALUES (?, ?, ?)", todo.Task, todo.LimitDate, todo.Status)
	return todo, err
}

func (todoRepo *TodoRepository) Update(todo *model.Todo) (*model.Todo, error) {
	_, err := todoRepo.SqlHandler.Conn.Exec("UPDATE todos SET task = ?,limit_date = ?,status = ? WHERE id = ?", todo.Task, todo.LimitDate, todo.Status, todo.ID)
	return todo, err
}