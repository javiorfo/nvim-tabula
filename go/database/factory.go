package database

import (
	"errors"

	"github.com/javiorfo/nvim-tabula/go/database/engine"
)

type Executor interface {
	Execute(queries string, connStr string)
}

func Context(engine_str string) (Executor, error) {
	switch engine_str {
	case engine.POSTGRES:
		return engine.Postgres{}, nil
	case engine.MONGO:
		return engine.Mongo{}, nil
	case engine.MYSQL:
		return engine.MySql{}, nil
	default:
		return nil, errors.New("engine does not exist")
	}
}
