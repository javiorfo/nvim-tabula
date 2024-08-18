package engine

import (
	"database/sql"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/javiorfo/nvim-tabula/go/database/engine/model"
	"github.com/javiorfo/nvim-tabula/go/database/query"
	"github.com/javiorfo/nvim-tabula/go/database/table"
	"github.com/javiorfo/nvim-tabula/go/logger"
	_ "github.com/lib/pq"
)

type Postgres struct {
	model.Data
}

func (p *Postgres) getDB() (*sql.DB, func(), error) {
	db, err := sql.Open(p.Engine, p.ConnStr)
	if err != nil {
		logger.Errorf("Error initializing %s, connStr: %s", p.Engine, p.ConnStr)
		return nil, nil, fmt.Errorf("[ERROR] %v", err)
	}
	return db, func() { db.Close() }, nil
}

func (p *Postgres) Run() {
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

func (p *Postgres) execute(db *sql.DB) {
	if !query.ContainsSemicolonInMiddle(p.Queries) {
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

		if query.IsInsertUpdateOrDelete(p.Queries) {
			fmt.Print(fmt.Sprintf("  Row(s) affected: %d", rowsAffected))
		} else {
			fmt.Print("  Statement executed correctly.")
		}
	} else {
		queries := query.SplitQueries(p.Queries)
		results := make([]string, len(queries))
		for i, q := range queries {
			if res, err := db.Exec(q); err != nil {
				logger.Errorf("Error executing query %v", err)
				results[i] = fmt.Sprintf("%d)   %v\n", i+1, err)
			} else {
                if rowsAffected, err := res.RowsAffected(); err != nil {
                    logger.Errorf("Error executing query %v", err)
                    results[i] = fmt.Sprintf("%d)   %v\n", i+1, err)
                } else {
                    if query.IsInsertUpdateOrDelete(q) {
                        results[i] = fmt.Sprintf("%d)   Row(s) affected: %d\n", i+1, rowsAffected)
                    } else {
                        results[i] = fmt.Sprintf("%d)   Statement executed correctly.\n", i+1)
                    }
                }
            }
		}
		filePath := table.CreateTabulaFileFormat(p.DestFolder)
		fmt.Println("syn match tabulaStmtErr ' ' | hi link tabulaStmtErr ErrorMsg")
		fmt.Println(filePath)

		table.WriteToFile(filePath, results...)
	}
}

func (p *Postgres) executeSelect(db *sql.DB) {
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
        name := " 󰠵 " + strings.ToUpper(value)
		tabula.Headers[i+1] = table.Header{
			Name:   name,
			Length: utf8.RuneCountInString(name) + 1,
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

			valueLength := utf8.RuneCountInString(value) + 2
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
		fmt.Print("  Query has returned 0 results.")
	}
}

func (p *Postgres) GetTables() {
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
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			logger.Errorf("Error scanning row:", err)
			fmt.Printf("[ERROR] %v", err)
			return
		}
		values = append(values, fmt.Sprintf("%s", table))
	}

	if err := rows.Err(); err != nil {
		logger.Errorf("Error iterating over rows:", err)
		fmt.Printf("[ERROR] %v", err)
		return
	}
    
    fmt.Print(values)
}

func (p *Postgres) GetTableInfo() {
    db, closer, err := p.getDB()
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer closer()

    p.Queries = p.GetTableInfoQuery(p.Queries)
    p.executeSelect(db)
}

