package engine

import "github.com/javiorfo/nvim-tabula/go/database/engine/model"

type MySql struct {
	model.Data
}

const MYSQL = "mysql"

func (my MySql) Run() {}

func (m MySql) GetTables() {
}
