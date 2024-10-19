package engine

import (
	"fmt"

    _ "github.com/alexbrainman/odbc"
	"github.com/javiorfo/nvim-dbeer/go/database/engine/model"
)

type Db2 struct {
	model.ProtoSQL
}

func (d *Db2) GetTables() {
	d.Queries = "select tabname as table_name from syscat.tables where tabschema = 'PUBLIC' order by tabname;"
	d.ProtoSQL.GetTables()
}

func (d *Db2) GetTableInfo() {
    fmt.Print("Not supported in DB2")	
}
