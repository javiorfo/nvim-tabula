package engine

import (
	"github.com/javiorfo/nvim-tabula/go/database/engine/model"
	_ "github.com/lib/pq"
)

type Postgres struct {
	model.ProtoSQL
}

func (p *Postgres) GetTables() {
	p.Queries = "select table_name from information_schema.tables where table_schema = 'public' order by table_name"
	p.ProtoSQL.GetTables()
}
