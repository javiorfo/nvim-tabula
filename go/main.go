package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Column struct {
	Name   string
	Length int
}

type QueryResult struct {
	header map[int]Column
	rows   map[int][]string
}

func main() {
	connStr := "user=admin dbname=db_dummy password=admin host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM dummies")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	queryResult := QueryResult{
		make(map[int]Column),
		make(map[int][]string),
	}

	for i, value := range columns {
		queryResult.header[i+1] = Column{value, len(value)}
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

			queryResult.rows[rowNr] = append(queryResult.rows[rowNr], value)
			index := len(queryResult.rows[rowNr])

			if queryResult.header[index].Length < valueLength {
				queryResult.header[index] = Column{queryResult.header[index].Name, valueLength}
			}
		}
		rowNr++
	}
	fmt.Println(queryResult.header)
	fmt.Println(queryResult.rows)
}
