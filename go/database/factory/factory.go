package factory

import (
	"errors"

	"github.com/javiorfo/nvim-tabula/go/database/engine"
	"github.com/javiorfo/nvim-tabula/go/database/engine/model"
	"github.com/javiorfo/nvim-tabula/go/database/engine/mongo"
)

const MYSQL = "mysql"
const MSSQL = "mssql"
const POSTGRES = "postgres"
const MONGO = "mongo"

type Executor interface {
	Run()
	GetTables()
	GetTableInfo()
    Ping()
}

func Context(option model.Option, proto model.ProtoSQL) error {
	switch proto.Engine {
	case POSTGRES:
		return run(&engine.Postgres{ProtoSQL: proto}, option)
	case MONGO:
		return run(&mongo.Mongo{ProtoSQL: proto}, option)
	case MYSQL:
		return run(&engine.MySql{ProtoSQL: proto}, option)
	case MSSQL:
		return run(&engine.MSSql{ProtoSQL: proto}, option)
	default:
		return errors.New("engine does not exist: " + proto.Engine)
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
	case model.PING:
		executor.Ping()
		return nil
	default:
		return errors.New("Option does not exist")
	}
}
