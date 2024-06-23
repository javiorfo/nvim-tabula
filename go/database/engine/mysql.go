package engine

type MySql struct{}

const MYSQL = "mysql"

func (MySql) Execute(queries string, connStr string) {}
