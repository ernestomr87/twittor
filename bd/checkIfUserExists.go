package bd

import (
	"context"
	"time"

	"github.com/ernestomr87/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckIfUserExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("users")

	conditions := bson.M{"email": email}

	var results models.User

	err := col.FindOne(ctx, conditions).Decode(&results)

	ID := results.ID.Hex()
	if err != nil {
		return results, false, ID
	}
	return results, true, ID

}
