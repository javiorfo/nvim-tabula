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
	index := strings.Index(s, "(")
	if index != -1 {
		return &mongoFuncParam{
			Func:   mongoFunc(s[:index]),
			Params: s[index+1 : len(s)-1],
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
	InsertOne                        = "insertOne"
	InsertMany                       = "insertMany"
	DeleteOne                        = "deleteOne"
	DeleteMany                       = "deleteMany"
	Drop                             = "drop"

    // TODO
	UpdateOne                        = "updateOne"
	UpdateMany                       = "updateMany"
	ReplaceOne                       = "replaceOne"
	CreateIndex                      = "createIndex"
	DropIndex                        = "dropIndex"
	ListIndexes                      = "listIndexes"
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
		if len(parts) > index+2 {
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
	bsonObj := bson.M{}
	if len(s) > 2 {
		if err := json.Unmarshal([]byte(s), &bsonObj); err != nil {
			return nil, fmt.Errorf("parsing filter %v", err)
		}
	}
	return &bsonObj, nil
}

func getArrayParsed(s string) ([]any, error) {
	var array []any
	if len(s) > 2 {
		if err := json.Unmarshal([]byte(s), &array); err != nil {
			return nil, fmt.Errorf("parsing filter %v", err)
		}
	}
	return array, nil
}
