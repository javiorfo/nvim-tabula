package main

import (
	"log"
	"os"

	"github.com/javiorfo/nvim-tabula/go/database"
)

func main() {
    _ = os.Args

    
	db := "postgres"
	connStr := "user=admin dbname=db_dummy password=admin host=localhost sslmode=disable"
	queries := "select * from dummies;"

/*     db := "mongo"
    connStr := "mongodb://admin:admin@localhost:27017/db_dummy"
	queries := "select * from dummies;" */

	err := database.Context(db, connStr, queries)
	if err != nil {
		log.Fatal(err)
	}
}
