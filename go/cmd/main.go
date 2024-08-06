package main

import (
	"flag"
	"log"

	"github.com/javiorfo/nvim-tabula/go/database"
	"github.com/javiorfo/nvim-tabula/go/database/engine/model"
)

func main() {
	engine := *flag.String("engine", "postgres", "Database engine")
	connStr := *flag.String("conn-str", "user=admin dbname=db_dummy password=admin host=localhost sslmode=disable", "Database string connection")
	queries := *flag.String("queries", "select * from dummies;", "Database queries semicolon-separated")
	destFolder := *flag.String("dest-folder", "/tmp", "Destinated folder for tabula files")
	luaTabulaPath := *flag.String("lua-tabula-path", "/home/javier/.local/share/nvim/lazy/nvim-tabula/lua/tabula/", "Folder where Lua files are stored in tabula")
	tabulaLogFile := *flag.String("tabula-log-file", "", "Neovim Tabula log file")
	option := *flag.Int("option", 2, "Options to execute: 1:run/2:tables/3:table-info")

	flag.Parse()

	// 	queries := "select * from dummies;"
	//     queries := "select CAST(table_name as varchar), table_type, CAST(table_catalog as varchar), cast(table_schema as varchar) from information_schema.tables where table_schema = 'public';"
	//     queries := "select cast(column_name as varchar), data_type, is_nullable from information_schema.columns where table_name = 'dummies';"

	/*     engine := "mongo"
	       connStr := "mongodb://admin:admin@localhost:27017/db_dummy"
	   	queries := "select * from dummies;" */

	// 	err := database.Context(engine, connStr, queries, destFolder, luaTabulaPath, tabulaLogFile, option)
	err := database.Context(engine, model.Option(option), model.Data{
		ConnStr:       connStr,
		Queries:       queries,
		DestFolder:    destFolder,
		LuaTabulaPath: luaTabulaPath,
		TabulaLogFile: tabulaLogFile,
	})
	if err != nil {
		log.Fatal(err)
	}
}
