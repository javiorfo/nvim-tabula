package engine

import (
	"context"
	"fmt"

	"github.com/javiorfo/nvim-tabula/go/database/engine/model"
	"github.com/javiorfo/nvim-tabula/go/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	model.ProtoSQL
}

func (m *Mongo) GetDB() (*mongo.Database, func(), error) {
	clientOptions := options.Client().ApplyURI(m.ConnStr)

    // TODO ojo todo context
	client, err := mongo.Connect(context.TODO(), clientOptions)
	db := client.Database(m.DbName)
	if err != nil {
		logger.Errorf("Error initializing %s, connStr: %s", m.Engine, m.ConnStr)
		return nil, nil, fmt.Errorf("[ERROR] %v", err)
	}
	closer := func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			logger.Errorf("Error disconnecting from MongoDB: %v", err)
			return
		}
	}
	return db, closer, nil
}

func (m *Mongo) Run() {
	db, closer, err := m.GetDB()
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer closer()

	collection := db.Collection("dummies")

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		logger.Errorf("Error executing find:", err)
		fmt.Printf("[ERROR] %v", err)
		return
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			logger.Errorf("Error decoding collection:", err)
			fmt.Printf("[ERROR] %v", err)
			return
		}
		fmt.Println(result)
	}
}

func (m *Mongo) GetTables() {
	db, closer, err := m.GetDB()
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer closer()

	collections, err := db.ListCollections(context.TODO(), bson.D{})
	if err != nil {
		logger.Errorf("Error listing collection:", err)
		fmt.Printf("[ERROR] %v", err)
		return
	}

	values := make([]string, 0)
	for collections.Next(context.TODO()) {
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

}
