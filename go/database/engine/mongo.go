package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/javiorfo/nvim-tabula/go/database/engine/model"
	"github.com/javiorfo/nvim-tabula/go/database/table"
	"github.com/javiorfo/nvim-tabula/go/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	model.ProtoSQL
}

func (m *Mongo) getDB(c context.Context) (*mongo.Database, func(), error) {
	clientOptions := options.Client().ApplyURI(m.ConnStr)

	client, err := mongo.Connect(c, clientOptions)
	db := client.Database(m.DbName)
	if err != nil {
		logger.Errorf("Error initializing %s, connStr: %s", m.Engine, m.ConnStr)
		return nil, nil, fmt.Errorf("[ERROR] %v", err)
	}
	closer := func() {
		if err = client.Disconnect(c); err != nil {
			logger.Errorf("Error disconnecting from MongoDB: %v", err)
			return
		}
	}
	return db, closer, nil
}

func (m *Mongo) Run() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	db, closer, err := m.getDB(ctx)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer closer()

	collection := db.Collection("dummies")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		logger.Errorf("Error finding collection:", err)
		fmt.Printf("[ERROR] %v", err)
		return
	}
	defer cursor.Close(ctx)

	values := make([]string, 0)
	for cursor.Next(ctx) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			logger.Errorf("Error decoding collection:", err)
			fmt.Printf("[ERROR] %v", err)
			return
		}
		prettyJSON, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			logger.Errorf("Error prettifying:", err)
			fmt.Printf("[ERROR] %v", err)
			return
		}
		values = append(values, string(prettyJSON))
	}

	if err := cursor.Err(); err != nil {
		logger.Errorf("Error in collection cursor:", err)
		fmt.Printf("[ERROR] %v", err)
		return
	}

	filePath := table.CreateTabulaMongoFileFormat(m.DestFolder)
	fmt.Println("syn match tabulaStmtErr ' ' | hi link tabulaStmtErr ErrorMsg")
	fmt.Println(filePath)

	table.WriteToFile(filePath, values...)
}

func (m *Mongo) GetTables() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	db, closer, err := m.getDB(ctx)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer closer()

	collections, err := db.ListCollections(ctx, bson.D{})
	if err != nil {
		logger.Errorf("Error listing collection:", err)
		fmt.Printf("[ERROR] %v", err)
		return
	}

	values := make([]string, 0)
	for collections.Next(ctx) {
		var collection bson.M
		err := collections.Decode(&collection)
		if err != nil {
			logger.Errorf("Error decoding collection:", err)
			fmt.Printf("[ERROR] %v", err)
			return
		}
		values = append(values, collection["name"].(string))
	}

	if err := collections.Err(); err != nil {
		logger.Errorf("Error iterating over rows:", err)
		fmt.Printf("[ERROR] %v", err)
		return
	}

	fmt.Print(values)
}

func (m *Mongo) GetTableInfo() {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	db, closer, err := m.getDB(ctx)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer closer()

	collection := db.Collection(m.Queries)

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		logger.Errorf("Error listing collection:", err)
		fmt.Printf("[ERROR] %v", err)
		return
	}
	defer cursor.Close(ctx)

	var maxKeysDoc bson.M
	maxKeysCount := 0

	for cursor.Next(ctx) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			logger.Errorf("Error decoding:", err)
			fmt.Printf("[ERROR] %v", err)
			return
		}

		keysCount := len(result)
		if keysCount > maxKeysCount {
			maxKeysCount = keysCount
			maxKeysDoc = result
		}
	}

	if err := cursor.Err(); err != nil {
		logger.Errorf("Error in collection cursor:", err)
		fmt.Printf("[ERROR] %v", err)
		return
	}

	tabula := table.Tabula{
		DestFolder:      m.DestFolder,
		BorderStyle:     m.BorderStyle,
		HeaderStyleLink: m.HeaderStyleLink,
		Headers: map[int]table.Header{
			1: {Name: " 󰠵 KEY", Length: 7},
			2: {Name: " 󰠵 DATA_TYPE", Length: 13},
		},
		Rows: make([][]string, len(maxKeysDoc)),
	}

	index := 0
	for key, value := range maxKeysDoc {
		valueKey := " " + strings.ToUpper(key)
		valueType := " " + reflect.TypeOf(value).String()
		tabula.Rows[index] = []string{valueKey, valueType}

		valueKeyLength := utf8.RuneCountInString(valueKey) + 2
		if tabula.Headers[1].Length < valueKeyLength {
			tabula.Headers[1] = table.Header{
				Name:   tabula.Headers[1].Name,
				Length: valueKeyLength,
			}
		}
		valueTypeLength := utf8.RuneCountInString(valueType) + 2
		if tabula.Headers[2].Length < valueTypeLength {
			tabula.Headers[2] = table.Header{
				Name:   tabula.Headers[2].Name,
				Length: valueTypeLength,
			}
		}
		index++
	}

	tabula.Generate()
}
