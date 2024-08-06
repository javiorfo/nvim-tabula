package engine

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/javiorfo/nvim-tabula/go/database/engine/model"
	"github.com/javiorfo/nvim-tabula/go/database/query"
	_ "github.com/lib/pq"
)

type Postgres struct {
	model.Data
}

const POSTGRES = "postgres"

func (p Postgres) getDB() (*sql.DB, func()) {
	db, err := sql.Open(POSTGRES, p.ConnStr)
	if err != nil {
		panic(err)
	}
	return db, func() { db.Close() }
}

func (p Postgres) Run() {
	db, closer := p.getDB()
	defer closer()

	rows, err := db.Query(p.Queries)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	lenColumns := len(columns)

	tabula := query.Tabula{
		DestFolder: p.DestFolder,
		Headers:    make(map[int]query.Header, lenColumns),
		Rows:       make([][]string, 0),
	}

	for i, value := range columns {
		tabula.Headers[i+1] = query.Header{
			Name:   " " + strings.ToUpper(value),
			Length: len(value) + 2,
		}
	}

	values := make([]any, lenColumns)
	for i := range values {
		var value any
		values[i] = &value
	}

	for rows.Next() {
		err := rows.Scan(values...)
		if err != nil {
			log.Fatal(err)
		}

		results := make([]string, lenColumns)
		for i, value := range values {
			value := strings.Replace(fmt.Sprintf("%v", *value.(*any)), " +0000 +0000", "", -1)
			if value == "<nil>" {
				value = "NULL"
			}

			valueLength := len(value) + 2
			results[i] = " " + value
			index := i + 1

			if tabula.Headers[index].Length < valueLength {
				tabula.Headers[index] = query.Header{
					Name:   tabula.Headers[index].Name,
					Length: valueLength,
				}
			}
		}
		tabula.Rows = append(tabula.Rows, results)
	}

	tabula.Generate()
}

func (p Postgres) GetTables() {
    db, closer := p.getDB()
	defer closer()

	rows, err := db.Query("select table_name from information_schema.tables where table_schema = 'public'")
	if err != nil {
		log.Fatal("Error executing query:", err)
	}
	defer rows.Close()

    values := make([]string, 0)
	values = append(values, "return {\n")
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			log.Fatal("Error scanning row:", err)
		}
		values = append(values, fmt.Sprintf("    \"%s\",\n", table))
	}
	values = append(values, "}")

	if err := rows.Err(); err != nil {
		log.Fatal("Error iterating over rows:", err)
	}

    query.WriteTable(values, p.LuaTabulaPath, "tables.lua")
}

func (p Postgres) GetTableInfo() {
}
