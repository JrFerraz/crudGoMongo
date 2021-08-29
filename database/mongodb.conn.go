package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()
var database = "crud"
var client *mongo.Client
var err error

func init() {

	err = enableConnection()
	checkDatabaseConnection(err)
}

func GetCollection(collectionName string) *mongo.Collection {
	return client.Database(database).Collection(collectionName)
}

func enableConnection() error {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	client, err = mongo.Connect(ctx, clientOptions)
	return err
}

func checkDatabaseConnection(err error) {
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Conexion establecida")
	}
}

func CheckReturnError(err error) error {
	if err != nil {
		return err
	}
	return nil
}
