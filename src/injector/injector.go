package injector

import (
	"go_todo/infra"
	"go_todo/config"
	"go_todo/domain/repository"
	"go_todo/handler"
	"go_todo/usecase"
)

func injectDB(conf *config.Config) infra.SqlHandler {
	sqlHandler := infra.NewSqlHandler(conf)
	return *sqlHandler
}

func InjectTodoRepository(conf *config.Config) repository.TodoRepository {
	return infra.NewTodoRepository(injectDB(conf))
}

func InjectTodoUsecase(conf *config.Config) usecase.TodoUsecase {
	return usecase.NewTodoUsecase(InjectTodoRepository(conf))
}

func InjectTodoHandler(conf *config.Config) handler.TodoHandler {
	return handler.NewTodoHandler(InjectTodoUsecase(conf))
}