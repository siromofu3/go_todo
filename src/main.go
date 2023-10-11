package main

import (
	"fmt"
	"go_todo/injector"
	"go_todo/handler"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("server start")
	todoHandler := injector.InjectTodoHandler()
	e := echo.New()
	handler.InitRouting(e, todoHandler)
	e.Logger.Fatal(e.Start(":80"))
}