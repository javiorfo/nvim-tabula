package main

import (
	"log"

	"github.com/javiorfo/nvim-tabula/go/database"
)

func main() {
    db := "postgres"
    connStr := "user=admin dbname=db_dummy password=admin host=localhost sslmode=disable"
    query := "select * from dummies;"
    engine, err := database.Context(db)
    if err == nil {
        engine.Execute(query, connStr)
    } else {
        log.Fatal(err)
    }
}
