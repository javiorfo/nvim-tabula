package engine

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/javiorfo/nvim-tabula/go/database/query"
	_ "github.com/lib/pq"
)

type Postgres struct{}

const POSTGRES = "postgres"

func (Postgres) Execute(queries string, connStr string) {
	db, err := sql.Open(POSTGRES, connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(queries)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
    
    columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	selectResult := query.SelectResult {
		Header: make(map[int]query.ColumnResult),
		Rows:   make(map[int][]string),
	}

	for i, value := range columns {
		selectResult.Header[i+1] = query.ColumnResult{
			Name:   value,
			Length: len(value),
		}
	}

	values := make([]any, len(columns))
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

		for _, value := range values {
			value := fmt.Sprintf("%v", *value.(*any))
			valueLength := len(value)

			if value == "<nil>" {
				value = "null"
			}

			selectResult.Rows[rowNr] = append(selectResult.Rows[rowNr], value)
			index := len(selectResult.Rows[rowNr])

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
