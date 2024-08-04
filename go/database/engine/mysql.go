package engine

type MySql struct {
	ConnStr string
	Queries string
}

const MYSQL = "mysql"

func (my MySql) Execute() {}
