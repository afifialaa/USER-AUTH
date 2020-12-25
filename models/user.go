package models

import (
	"context"

	"github.com/afifialaa/user-auth/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u)

// Create user
func (u User) Create() (*mongo.InsertOneResult, error) {
	result, err := database.UserCollection.InsertOne(context.TODO(), u)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Delete user
func (u User) Delete() (*mongo.DeleteResult, error) {
	filter := bson.D{{"email", u.Email}}

	result, err := database.User.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Update user
func (u User) Update() (*mongo.UpdateResult, error) {
	result, err := database.UserCollection.UpdateOne(context.TODO(), bson.M{"email": u.Email},
		bson.D{
			{"$set", u},
		},
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// Find user
func (u User) Find() (*mongo.InsertOneResult, error){
	result, err := database.UserCollection.InsertOne(context.TODO(), u)
	if err != nil {
		return nil, err
	}

	return result, nil	
}
