package handler

import (
	"net/http"
	"go_todo/domain/model"
	"go_todo/usecase"
	"github.com/labstack/echo"
)

type TodoHandler struct {
	todoUsecase usecase.TodoUsecase
}

func NewTodoHandler(todoUsecase usecase.TodoUsecase) TodoHandler {
	todoHandler := TodoHandler{todoUsecase}
	return todoHandler
}

func (th *TodoHandler) View() echo.HandlerFunc {
	return func(c echo.Context) error {
		models, err := th.todoUsecase.View()
		if err != nil {
			return c.JSON(http.StatusBadRequest, models)
		}
		return c.JSON(http.StatusOK, models)
	}
}

func (th *TodoHandler) Search() echo.HandlerFunc {
	return func(c echo.Context) error {
		word := c.QueryParam("word")
		models, err := th.todoUsecase.Search(word)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models)
		}
		return c.JSON(http.StatusOK, models)
	}
}

func (th *TodoHandler) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var todo model.Todo
		c.Bind(&todo)
		err := th.todoUsecase.Add(&todo)
		return c.JSON(http.StatusOK, err)
	}
}

func (th *TodoHandler) Edit() echo.HandlerFunc {
	return func(c echo.Context) error {
		var todo model.Todo
		c.Bind(&todo)
		err := th.todoUsecase.Edit(&todo)
		return c.JSON(http.StatusOK, err)
	}
}