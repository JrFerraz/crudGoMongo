package user_repository

import (
	"context"
	"time"

	db "github.com/JrFerraz/GolangMongodbCRUD/database"
	m "github.com/JrFerraz/GolangMongodbCRUD/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection = db.GetCollection("cruduserv2")
var ctx = context.Background()

//lo de retornar el error es simplemente buena practica
func Create(user m.User) error {

	var err error
	_, err = collection.InsertOne(ctx, user)

	return db.CheckReturnError(err)
}

func Read() (m.Users, error) {
	filter := bson.D{}
	//cur es un cursor, que pa que se entienda, es como un contenedor de todos los documentos que hicieron match de una collection con un filter
	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	return collectUsers(cur, err)
}

func Update(user m.User, userId string) error {

	objectId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"_id": objectId}

	update := bson.M{
		"$set": bson.M{
			"name":      user.Name,
			"email":     user.Email,
			"updatedat": time.Now(),
		},
	}
	_, err := collection.UpdateOne(ctx, filter, update)

	return db.CheckReturnError(err)
}

func Delete(userId string) error {

	objectId, err := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"_id": objectId}

	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(ctx, filter)

	return db.CheckReturnError(err)
}

//el cursor contiene los docu, los recorreremos uno a uno. Con decode los pasamos a struct
func collectUsers(cur *mongo.Cursor, err error) (m.Users, error) {
	var users m.Users
	for cur.Next(ctx) {

		var user m.User
		err = cur.Decode(&user)

		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
