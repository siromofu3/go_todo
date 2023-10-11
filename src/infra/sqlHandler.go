package infra

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	log.Print("NewSqlHandler")
	conn, err := sql.Open("mysql", "root:rootpass@tcp(db:3306)/todo")
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := SqlHandler{conn}
	return &sqlHandler
}