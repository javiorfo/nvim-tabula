package engine

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/javiorfo/nvim-tabula/go/database/engine/model"
	"github.com/javiorfo/nvim-tabula/go/database/query"
	"github.com/javiorfo/nvim-tabula/go/database/table"
	"github.com/javiorfo/nvim-tabula/go/logger"
	_ "github.com/lib/pq"
)

type Postgres struct {
	model.Data
}

func (p Postgres) getDB() (*sql.DB, func(), error) {
	db, err := sql.Open(p.Engine, p.ConnStr)
	if err != nil {
		logger.Errorf("Error initializing %s, connStr: %s", p.Engine, p.ConnStr)
		return nil, nil, fmt.Errorf("[ERROR] %v", err)
	}
	return db, func() { db.Close() }, nil
}

func (p Postgres) Run() {
	db, closer, err := p.getDB()
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer closer()

    if query.IsSelectQuery(p.Queries) {
        p.executeSelect(db)
    } else {
        p.execute(db)
    }
}

func (p Postgres) execute(db *sql.DB) {
    res, err := db.Exec(p.Queries)
	if err != nil {
		logger.Errorf("Error executing query %v", err)
		fmt.Printf("[ERROR] %v", err)
        return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logger.Errorf("Error executing query %v", err)
		fmt.Printf("[ERROR] %v", err)
        return
	}
    // TODO change message if there is no rows affected (drop, create)
    table.WriteToFile(p.DestFolder, "tabula", fmt.Sprintf("Row(s) affected: %d", rowsAffected))
}

func (p Postgres) executeSelect(db *sql.DB) {
    rows, err := db.Query(p.Queries)
	if err != nil {
		logger.Errorf("Error executing query %v", err)
		fmt.Printf("[ERROR] %v", err)
		return
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		logger.Errorf("Could not get columns %v", err)
		fmt.Printf("[ERROR] %v", err)
		return
	}
	lenColumns := len(columns)

	tabula := table.Tabula{
		DestFolder:      p.DestFolder,
		BorderStyle:     p.BorderStyle,
		HeaderStyleLink: p.HeaderStyleLink,
		Headers:         make(map[int]table.Header, lenColumns),
		Rows:            make([][]string, 0),
	}

	for i, value := range columns {
		tabula.Headers[i+1] = table.Header{
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
			logger.Errorf("Error getting rows %v", err)
			fmt.Printf("[ERROR] %v", err)
			return
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
				tabula.Headers[index] = table.Header{
					Name:   tabula.Headers[index].Name,
					Length: valueLength,
				}
			}
		}
		tabula.Rows = append(tabula.Rows, results)
	}

	if len(tabula.Rows) > 0 {
		tabula.Generate()
	} else {
		table.WriteToFile(tabula.DestFolder, "tabula", "Query has returned 0 results.")
	}
}

func (p Postgres) GetTables() {
	db, closer, err := p.getDB()
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer closer()

	rows, err := db.Query("select table_name from information_schema.tables where table_schema = 'public'")
	if err != nil {
		logger.Errorf("Error executing query:", err)
		fmt.Printf("[ERROR] %v", err)
		return
	}
	defer rows.Close()

	values := make([]string, 0)
	values = append(values, "return { ")
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			logger.Errorf("Error scanning row:", err)
			fmt.Printf("[ERROR] %v", err)
			return
		}
		values = append(values, fmt.Sprintf(" \"%s\", ", table))
	}
	values = append(values, "}")

	if err := rows.Err(); err != nil {
		logger.Errorf("Error iterating over rows:", err)
		fmt.Printf("[ERROR] %v", err)
		return
	}

	table.WriteToFile(p.LuaTabulaPath, "tables.lua", values...)
}
