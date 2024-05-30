package database

import (
	"errors"
	"strings"

	"github.com/javiorfo/nvim-tabula/go/database/engine"
)

type Executor interface {
	Execute(string, string)
}

func Context(engine_str string) (Executor, error) {
	switch engine_str = strings.ToLower(engine_str); engine_str {
	case engine.POSTGRES:
		return engine.Postgres{}, nil
	case engine.MYSQL:
		return engine.MySql{}, nil
	default:
		return nil, errors.New("engine does not exist")
	}
}
