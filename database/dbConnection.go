package database

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/afifialaa/USER-AUTH/models"
	"github.com/afifialaa/USER-AUTH/secrets"

	"context"
	"fmt"
	"log"
)

var userCollection *mongo.Collection

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
	userCollection = client.Database("private").Collection("users")

	// Create index
	_, err = userCollection.Indexes().CreateOne(
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

// Insert new user
func SaveUser(user *models.User) bool {
	insertResult, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		if IsDup(err) {
			fmt.Println("Duplicate index")
			return false
		}
		fmt.Println("mongodb error ", err.Error())
		return false
	}

	fmt.Println("#saveuser -> user was created: ", insertResult.InsertedID)
	return true
}

// Helper function, duplicate keys
// Looping over the WriteErrors to find code 11000
func IsDup(err error) bool {
	var e mongo.WriteException
	if errors.As(err, &e) {
		for _, we := range e.WriteErrors {
			if we.Code == 11000 {
				return true
			}
		}
	}
	return false
}

func FindUser(user *models.User) bool {
	var result models.User

	filter := bson.D{{"email", user.Email}}
	err := userCollection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		fmt.Println("user was not found")
		return false
	}

	return true
}
