package mongo

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/javiorfo/nvim-tabula/go/database/table"
	"github.com/javiorfo/nvim-tabula/go/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func find(ctx context.Context, mongoCommand *mongoCommand, db *mongo.Database, destFolder string) {
	filter, err := getBsonParsed(mongoCommand.FuncParam.Params)
	if err != nil {
		fmt.Printf("[ERROR] %v", err)
		return
	}

	findOptions := options.Find()
	if mongoCommand.FuncParam2 != nil {
		switch mongoCommand.FuncParam2.Func {
		case Sort:
			filter, err := getBsonParsed(mongoCommand.FuncParam2.Params)
			if err != nil {
				fmt.Printf("[ERROR] %v", err)
				return
			}
			findOptions.SetSort(*filter)
		case Limit:
			n, err := strconv.ParseInt(mongoCommand.FuncParam2.Params, 10, 64)
			if err != nil {
				fmt.Printf("[ERROR] %v", err)
				return
			}
			findOptions.SetLimit(n)
		case Skip:
			n, err := strconv.ParseInt(mongoCommand.FuncParam2.Params, 10, 64)
			if err != nil {
				fmt.Printf("[ERROR] %v", err)
				return
			}
			findOptions.SetSkip(n)
		}
	}

	cursor, err := db.Collection(mongoCommand.Collection).Find(ctx, *filter, findOptions)
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

	if len(values) == 0 {
		fmt.Print("  Query has returned 0 results.")
		return
	}

	filePath := table.CreateTabulaMongoFileFormat(destFolder)
	fmt.Println("syn match tabulaStmtErr ' ' | hi link tabulaStmtErr ErrorMsg")
	fmt.Println(filePath)

	table.WriteToFile(filePath, values...)
}

func findOne(ctx context.Context, mongoCommand *mongoCommand, db *mongo.Database, destFolder string) {
	filter, err := getBsonParsed(mongoCommand.FuncParam.Params)
	if err != nil {
		fmt.Printf("[ERROR] %v", err)
		return
	}

	var result bson.M
	err = db.Collection(mongoCommand.Collection).FindOne(ctx, *filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Print("  Query has returned 0 results.")
            return
		} else {
			logger.Errorf("Error finding document:", err)
			fmt.Printf("[ERROR] %v", err)
			return
		}
	}
	prettyJSON, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		logger.Errorf("Error prettifying:", err)
		fmt.Printf("[ERROR] %v", err)
		return
	}

	filePath := table.CreateTabulaMongoFileFormat(destFolder)
	fmt.Println("syn match tabulaStmtErr ' ' | hi link tabulaStmtErr ErrorMsg")
	fmt.Println(filePath)

	table.WriteToFile(filePath, string(prettyJSON))
}

func countDocuments(ctx context.Context, mongoCommand *mongoCommand, db *mongo.Database) {
	filter, err := getBsonParsed(mongoCommand.FuncParam.Params)
	if err != nil {
		fmt.Printf("[ERROR] %v", err)
		return
	}

	total, err := db.Collection(mongoCommand.Collection).CountDocuments(ctx, *filter)
	if err != nil {
		logger.Errorf("Error counting collection:", err)
		fmt.Printf("[ERROR] %v", err)
		return
	}

	fmt.Printf("  Collection %s count: %d results.", mongoCommand.Collection, total)
}

func insertOne(ctx context.Context, mongoCommand *mongoCommand, db *mongo.Database) {
	obj, err := getBsonParsed(mongoCommand.FuncParam.Params)
	if err != nil {
		fmt.Printf("[ERROR] %v", err)
		return
	}
    
    result, err := db.Collection(mongoCommand.Collection).InsertOne(ctx, *obj)
	if err != nil {
		logger.Errorf("Error inserting collection:", err)
		fmt.Printf("[ERROR] %v", err)
		return
	}

	fmt.Printf("  Collection %s, document inserted with ID: %v", mongoCommand.Collection, result.InsertedID)
}

func insertMany(ctx context.Context, mongoCommand *mongoCommand, db *mongo.Database) {
	array, err := getArrayParsed(mongoCommand.FuncParam.Params)
	if err != nil {
		fmt.Printf("[ERROR] %v", err)
		return
	}
    
    result, err := db.Collection(mongoCommand.Collection).InsertMany(ctx, array)
	if err != nil {
		logger.Errorf("Error inserting collection:", err)
		fmt.Printf("[ERROR] %v", err)
		return
	}

	fmt.Printf("  Collection %s, documents inserted with ID/s: %v", mongoCommand.Collection, result.InsertedIDs)
}

func deleteOne(ctx context.Context, mongoCommand *mongoCommand, db *mongo.Database) {
	obj, err := getBsonParsed(mongoCommand.FuncParam.Params)
	if err != nil {
		fmt.Printf("[ERROR] %v", err)
		return
	}
    
    result, err := db.Collection(mongoCommand.Collection).DeleteOne(ctx, *obj)
	if err != nil {
		logger.Errorf("Error deleting collection:", err)
		fmt.Printf("[ERROR] %v", err)
		return
	}

	fmt.Printf("  Collection %s, deleted: %d document.", mongoCommand.Collection, result.DeletedCount)
}

func deleteMany(ctx context.Context, mongoCommand *mongoCommand, db *mongo.Database) {
	obj, err := getBsonParsed(mongoCommand.FuncParam.Params)
	if err != nil {
		fmt.Printf("[ERROR] %v", err)
		return
	}
    
    result, err := db.Collection(mongoCommand.Collection).DeleteMany(ctx, *obj)
	if err != nil {
		logger.Errorf("Error deleting collection:", err)
		fmt.Printf("[ERROR] %v", err)
		return
	}

	fmt.Printf("  Collection %s, deleted: %d document/s.", mongoCommand.Collection, result.DeletedCount)
}

func updateOne(ctx context.Context, mongoCommand *mongoCommand, db *mongo.Database) {
	obj, err := getTwoBsonParsed(mongoCommand.FuncParam.Params)
	if err != nil {
		fmt.Printf("[ERROR] %v", err)
		return
	}
    
    result, err := db.Collection(mongoCommand.Collection).UpdateOne(ctx, obj.First, obj.Second)
	if err != nil {
		logger.Errorf("Error updating collection:", err)
		fmt.Printf("[ERROR] %v", err)
		return
	}

	fmt.Printf("  Collection %s, document updated with ID: %v", mongoCommand.Collection, result.ModifiedCount)
}

func updateMany(ctx context.Context, mongoCommand *mongoCommand, db *mongo.Database) {
	obj, err := getTwoBsonParsed(mongoCommand.FuncParam.Params)
	if err != nil {
		fmt.Printf("[ERROR] %v", err)
		return
	}
    
    result, err := db.Collection(mongoCommand.Collection).UpdateMany(ctx, obj.First, obj.Second)
	if err != nil {
		logger.Errorf("Error updating collection:", err)
		fmt.Printf("[ERROR] %v", err)
		return
	}

	fmt.Printf("  Collection %s, document updated with ID/s: %v", mongoCommand.Collection, result.ModifiedCount)
}

func dropCollection(ctx context.Context, mongoCommand *mongoCommand, db *mongo.Database) {
    err := db.Collection(mongoCommand.Collection).Drop(ctx)
	if err != nil {
		logger.Errorf("Error dropping collection:", err)
		fmt.Printf("[ERROR] %v", err)
		return
	}

	fmt.Printf("  Collection %s dropped succesfully.", mongoCommand.Collection)
}
