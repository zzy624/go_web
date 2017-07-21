package model

import (
	"database/sql"

	_ "github.com/bmizerany/pq"
	"github.com/labstack/gommon/log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=postgres password=123456 dbname=goweb sslmode=disable")
	if err != nil {
		log.Error("Open DB failed", err)
	}
}
