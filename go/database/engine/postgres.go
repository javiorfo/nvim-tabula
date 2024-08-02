package engine

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/javiorfo/nvim-tabula/go/database/query"
	_ "github.com/lib/pq"
)

type Postgres struct{}

const POSTGRES = "postgres"

func getDB(connStr string) (*sql.DB, func()) {
	db, err := sql.Open(POSTGRES, connStr)
	if err != nil {
		panic(err)
	}
	return db, func() { db.Close() }
}

func (Postgres) Execute(queries string, connStr string) {
	db, closer := getDB(connStr)
	defer closer()

	rows, err := db.Query(queries)
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
		Headers: make(map[int]query.Header, lenColumns),
		Rows:   make(map[int][]string),
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

	rowNr := 1
	for rows.Next() {
		err := rows.Scan(values...)
		if err != nil {
			log.Fatal(err)
		}

		tabula.Rows[rowNr] = make([]string, lenColumns)
		for i, value := range values {
			value := strings.Replace(fmt.Sprintf("%v", *value.(*any)), " +0000 +0000", "", -1)
			valueLength := len(value) + 2

			if value == "<nil>" {
				value = "NULL"
			}

			tabula.Rows[rowNr][i] = " " + value
			index := i + 1

			if tabula.Headers[index].Length < valueLength {
				tabula.Headers[index] = query.Header{
					Name:   tabula.Headers[index].Name,
					Length: valueLength,
				}
			}
		}
		rowNr++
	}

    tabula.Generate()
}
