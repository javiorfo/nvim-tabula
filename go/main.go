package main

import (
	"flag"
	"log"

	"github.com/javiorfo/nvim-tabula/go/database/engine/model"
	"github.com/javiorfo/nvim-tabula/go/database/factory"
	"github.com/javiorfo/nvim-tabula/go/logger"
)

func main() {
	engine := flag.String("engine", "", "Database engine")
	connStr := flag.String("conn-str", "", "Database string connection")
	dbName := flag.String("dbname", "", "Database name")
	queries := flag.String("queries", "", "Database queries semicolon-separated")
	borderStyle := flag.Int("border-style", 1, "Table border style")
	destFolder := flag.String("dest-folder", "/tmp", "Destinated folder for tabula files")
	tabulaLogFile := flag.String("tabula-log-file", "", "Neovim Tabula log file")
    option := flag.Int("option", 1, "Options to execute: 1:run/2:tables/3:table-info/4:ping")
	headerStyleLink := flag.String("header-style-link", "Type", "hi link header type")

	flag.Parse()

	logger.Initialize(*tabulaLogFile)

	if err := factory.Context(model.Option(*option), model.ProtoSQL{
		Engine:          *engine,
		ConnStr:         *connStr,
		DbName:          *dbName,
		Queries:         *queries,
		BorderStyle:     *borderStyle,
		DestFolder:      *destFolder,
		HeaderStyleLink: *headerStyleLink,
	}); err != nil {
		log.Fatal(err)
	}
}
