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

	selectResult := query.SelectResult{
		Header: make(map[int]query.ColumnResult, lenColumns),
		Rows:   make(map[int][]string),
	}

	for i, value := range columns {
		selectResult.Header[i+1] = query.ColumnResult{
			Name:   value,
			Length: len(value),
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

		selectResult.Rows[rowNr] = make([]string, lenColumns)
		for i, value := range values {
			value := strings.Replace(fmt.Sprintf("%v", *value.(*any)), " +0000 +0000", "", -1)
			valueLength := len(value)

			if value == "<nil>" {
				value = "null"
			}

			selectResult.Rows[rowNr][i] = value
			index := i + 1

			if selectResult.Header[index].Length < valueLength {
				selectResult.Header[index] = query.ColumnResult{
					Name:   selectResult.Header[index].Name,
					Length: valueLength,
				}
			}
		}
		rowNr++
	}
	fmt.Println(selectResult.Header)
	fmt.Println(selectResult.Rows)
}
