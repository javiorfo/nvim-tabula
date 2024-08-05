package main

import (
	"log"
	"os"

	"github.com/javiorfo/nvim-tabula/go/database"
)

func main() {
    _ = os.Args
    // 0 prog name
    // 1 engine
    // 2 connStr
    // 3 queries
    // 4 opt "gen_tables", "execute_query", "get_table_info"
    
	engine := "postgres"
	connStr := "user=admin dbname=db_dummy password=admin host=localhost sslmode=disable"
// 	queries := "select * from dummies;"
//     queries := "select CAST(table_name as varchar), table_type, CAST(table_catalog as varchar), cast(table_schema as varchar) from information_schema.tables where table_schema = 'public';"
    queries := "select cast(column_name as varchar), data_type, is_nullable from information_schema.columns where table_name = 'dummies';"

/*     engine := "mongo"
    connStr := "mongodb://admin:admin@localhost:27017/db_dummy"
	queries := "select * from dummies;" */

	err := database.Context(engine, connStr, queries)
	if err != nil {
		log.Fatal(err)
	}
}
