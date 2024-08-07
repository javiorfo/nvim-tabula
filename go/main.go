package main

import (
	"flag"
	"log"

	"github.com/javiorfo/nvim-tabula/go/database/engine/model"
	"github.com/javiorfo/nvim-tabula/go/database/factory"
)

func main() {
	engine := *flag.String("engine", "postgres", "Database engine")
	connStr := *flag.String("conn-str", "user=admin dbname=db_dummy password=admin host=localhost sslmode=disable", "Database string connection")
	queries := *flag.String("queries", "select * from dummies;", "Database queries semicolon-separated")
	borderStyle := *flag.Int("border-style", 1, "Table border style")
	destFolder := *flag.String("dest-folder", "/tmp", "Destinated folder for tabula files")
	luaTabulaPath := *flag.String("lua-tabula-path", "/home/javier/.local/share/nvim/lazy/nvim-tabula/lua/tabula/", "Folder where Lua files are stored in tabula")
	tabulaLogFile := *flag.String("tabula-log-file", "", "Neovim Tabula log file")
	option := *flag.Int("option", 1, "Options to execute: 1:run/2:tables")

	flag.Parse()

	//     queries := "select cast(column_name as varchar), data_type, is_nullable from information_schema.columns where table_name = 'dummies';"

	/*     engine := "mongo"
	       connStr := "mongodb://admin:admin@localhost:27017/db_dummy"
	   	queries := "select * from dummies;" */

	// 	err := database.Context(engine, connStr, queries, destFolder, luaTabulaPath, tabulaLogFile, option)
	err := factory.Context(model.Option(option), model.Data{
		Engine:        engine,
		ConnStr:       connStr,
		Queries:       queries,
        BorderStyle:   borderStyle,
		DestFolder:    destFolder,
		LuaTabulaPath: luaTabulaPath,
		TabulaLogFile: tabulaLogFile,
	})
	if err != nil {
		log.Fatal(err)
	}
}
