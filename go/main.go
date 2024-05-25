package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

func main() {
    // Establish a connection to the PostgreSQL database
    connStr := "user=admin dbname=db_dummy password=admin host=localhost sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Perform a select query
    rows, err := db.Query("SELECT id as hola, info FROM dummies")
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    // Get column names
	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	// Create a slice of interface{} to store the values
	values := make([]any, len(columns))
	for i := range values {
		var value any
		values[i] = &value
	}

	// Iterate over the rows
	for rows.Next() {
		err := rows.Scan(values...)
		if err != nil {
			log.Fatal(err)
		}

		// Process the values dynamically based on their types
		for i, value := range values {
			fmt.Printf("%s: %v\n", columns[i], *value.(*any))
		}
	}
}

