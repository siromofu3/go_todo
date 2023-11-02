package main

import (
	"fmt"
	"log"
	"go_todo/injector"
	"go_todo/handler"
	"go_todo/config"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("server start")

	conf, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	todoHandler := injector.InjectTodoHandler(conf)
	e := echo.New()
	handler.InitRouting(e, todoHandler)
	e.Logger.Fatal(e.Start(":80"))
}