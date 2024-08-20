package engine

import (
	"errors"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/javiorfo/nvim-tabula/go/database/engine/model"
	"github.com/javiorfo/nvim-tabula/go/logger"
)

type MySql struct {
	model.ProtoSQL
}

func (ms *MySql) GetTables() {
	dbName, err := ms.getDBName()
	if err != nil {
		logger.Errorf("Error executing query %v", err)
		fmt.Printf("[ERROR] %v", err)
		return
	}
	ms.Queries = fmt.Sprintf("select table_name from information_schema.tables where table_schema = '%s'", *dbName)
	ms.ProtoSQL.GetTables()
}

func (ms *MySql) getDBName() (*string, error) {
	parts := strings.Split(ms.ConnStr, "/")
	if len(parts) > 1 {
		return &parts[1], nil
	}
	return nil, errors.New("DB name does not exist in connection string.")
}
