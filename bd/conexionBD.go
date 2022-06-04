package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN is an Connection object
var MongoCN = ConnectionDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://root:epastor@cluster0.v6opscz.mongodb.net/?retryWrites=true&w=majority")

// ConnectionDB Connection to MongoDB
func ConnectionDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Successfully connected")
	return client
}

// CheckConnection is ping the connection
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
