package factory

import (
	"errors"

	"github.com/javiorfo/nvim-tabula/go/database/engine"
	"github.com/javiorfo/nvim-tabula/go/database/engine/model"
)

const MYSQL = "mysql"
const POSTGRES = "postgres"
const MONGO = "mongo"

type Executor interface {
	Run()
	GetTables()
    GetTableInfo()
}

func Context(option model.Option, data model.Data) error {
	switch data.Engine {
	case POSTGRES:
		return run(&engine.Postgres{Data: data}, option)
	case MONGO:
		return run(&engine.Mongo{Data: data}, option)
	case MYSQL:
		return run(&engine.MySql{Data: data}, option)
	default:
        return errors.New("engine does not exist: " + data.Engine)
	}
}

func run(executor Executor, option model.Option) error {
	switch option {
	case model.RUN:
		executor.Run()
		return nil
	case model.TABLES:
		executor.GetTables()
		return nil
	case model.TABLE_INFO:
		executor.GetTableInfo()
		return nil
	default:
		return errors.New("Option does not exist")
	}
}
