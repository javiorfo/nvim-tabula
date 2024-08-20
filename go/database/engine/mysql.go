package engine

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/javiorfo/nvim-tabula/go/database/engine/model"
)

type MySql struct {
	model.ProtoSQL
}

func (ms *MySql) GetTables() {
	ms.Queries = fmt.Sprintf("select table_name from information_schema.tables where table_schema = '%s'", ms.DbName)
	ms.ProtoSQL.GetTables()
}

