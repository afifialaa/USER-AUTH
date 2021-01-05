package database

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/afifialaa/USER-AUTH/secrets"

	"context"
	"fmt"
	"log"
)

var UserCollection *mongo.Collection

func Connect() {
	clientOptions := options.Client().ApplyURI(secrets.MongoCloud())

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Set database and collection
	UserCollection = client.Database("private").Collection("users")

	// Create index
	_, err = UserCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.M{
				"email": 1,
			},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		fmt.Println("Email field is not unique")
		log.Fatal(err)
	}
}
