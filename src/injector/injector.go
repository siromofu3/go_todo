package injector

import (
	"go_todo/infra"
	"go_todo/domain/repository"
	"go_todo/handler"
	"go_todo/usecase"
)

func injectDB() infra.SqlHandler {
	sqlHandler := infra.NewSqlHandler()
	return *sqlHandler
}

func InjectTodoRepository() repository.TodoRepository {
	return infra.NewTodoRepository(injectDB())
}

func InjectTodoUsecase() usecase.TodoUsecase {
	return usecase.NewTodoUsecase(InjectTodoRepository())
}

func InjectTodoHandler() handler.TodoHandler {
	return handler.NewTodoHandler(InjectTodoUsecase())
}