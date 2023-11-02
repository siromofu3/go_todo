package infra

import (
	"fmt"
	"database/sql"
	"go_todo/config"

	"github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler(conf *config.Config) *SqlHandler {
	c := mysql.Config{
		DBName: conf.DB.DBName,
		User:   conf.DB.User,
		Passwd: conf.DB.Password,
		Addr:   fmt.Sprintf("%s:%d", conf.DB.Host, conf.DB.Port),
		Net:    "tcp",
	}

	conn, err := sql.Open("mysql", c.FormatDSN())

	if err != nil {
		panic(err.Error)
	}
	sqlHandler := SqlHandler{conn}
	return &sqlHandler
}