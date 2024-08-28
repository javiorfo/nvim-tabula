package mongo

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type mongoCommand struct {
	Collection string
	FuncParam  mongoFuncParam
	FuncParam2 *mongoFuncParam
}

type mongoFuncParam struct {
	Func   mongoFunc
	Params string
}

func newMongoFuncParam(s string) *mongoFuncParam {
	openParenIndexExtra := strings.Index(s, "(")
	if openParenIndexExtra != -1 {
		return &mongoFuncParam{
			Func:   mongoFunc(s[:openParenIndexExtra]),
			Params: s[openParenIndexExtra+1 : len(s)-1],
		}
	}
	return nil
}

type mongoFunc string

const (
	Find                   mongoFunc = "find"
	Sort                             = "sort"
	Skip                             = "skip"
	Limit                            = "limit"
	CountDocuments                   = "countDocuments"
	FindOne                          = "findOne"

    // TODO
	InsertOne                        = "insertOne"
	InsertMany                       = "insertMany"
	UpdateOne                        = "updateOne"
	UpdateMany                       = "updateMany"
	ReplaceOne                       = "replaceOne"
	DeleteOne                        = "deleteOne"
	DeleteMany                       = "deleteMany"
	CreateIndex                      = "createIndex"
	DropIndex                        = "dropIndex"
	ListIndexes                      = "listIndexes"
	Drop                             = "drop"
	Rename                           = "rename"
)

func getQuerySections(query string) (*mongoCommand, error) {
	if len(query) == 0 {
		return nil, errors.New("Query empty.")
	}

	parts := strings.Split(query, ".")

	var collection string
	var function string
	var extra string
	if len(parts) > 0 {
		var index int
		if parts[index] == "db" {
			index = 1
		}
		collection = parts[index]
		function = parts[index+1]
		if len(parts) > 2 {
			extra = parts[index+2]
		}

		funcParam := newMongoFuncParam(function)
		if funcParam != nil {
			mongoCommand := mongoCommand{
				Collection: collection,
				FuncParam:  *funcParam,
				FuncParam2: newMongoFuncParam(extra),
			}
			return &mongoCommand, nil
		} else {
			return nil, errors.New("Error format: " + function)
		}
	}
	return nil, errors.New("Error format: " + query)
}

func getBsonParsed(s string) (*primitive.M, error) {
	filter := bson.M{}
	if len(s) > 2 {
		if err := json.Unmarshal([]byte(s), &filter); err != nil {
			return nil, fmt.Errorf("parsing filter %v", err)
		}
	}
	return &filter, nil
}
