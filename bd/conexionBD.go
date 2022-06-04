package bd

import (
	"context"
	"log"

	"github.com/ernestomr87/twittor/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN is an Connection object
var MongoCN = ConnectionDB()
var clientOptions = options.Client().ApplyURI(utils.GetEnvs("MONGODB_URL"))

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

// CheckDBConnection is ping the connection
func CheckDBConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
