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
	engine_str = strings.ToLower(engine_str)
	switch engine_str {
	case engine.POSTGRES:
		return &engine.Postgres{}, nil
	default:
		return nil, errors.New("engine does not exist")
	}
}
