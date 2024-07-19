package main

import (
	"log"

	"github.com/javiorfo/nvim-tabula/go/database"
)

func main() {
/* 	db := "postgres"
	connStr := "user=admin dbname=db_dummy password=admin host=localhost sslmode=disable"
	queries := "select * from dummies;" */

    db := "mongo"
    connStr := "mongodb://admin:admin@localhost:27017"
	queries := "select * from dummies;"

	engine, err := database.Context(db)
	if err == nil {
		engine.Execute(queries, connStr)
	} else {
		log.Fatal(err)
	}
}
