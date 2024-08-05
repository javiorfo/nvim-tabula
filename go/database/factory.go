package database

import (
	"errors"

	"github.com/javiorfo/nvim-tabula/go/database/engine"
)

type Executor interface {
	Execute() // TODO return error
}

func Context(db_engine, connStr, queries string) error {
	switch db_engine {
	case engine.POSTGRES:
		engine.Postgres{
			ConnStr: connStr,
			Queries: queries,
		}.Execute()
		return nil
	case engine.MONGO:
		engine.Mongo{
			ConnStr: connStr,
			Queries: queries,
		}.Execute()
		return nil
	case engine.MYSQL:
		engine.MySql{
			ConnStr: connStr,
			Queries: queries,
		}.Execute()
		return nil
	default:
		return errors.New("engine does not exist")
	}
}
