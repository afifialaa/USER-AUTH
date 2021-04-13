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

	result, err := database.UserCollection.DeleteOne(context.TODO(), filter)
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
func (u User) Find() (*User, error) {
	var user User
	filter := bson.M{"email": u.Email, "password": u.Password}

	err := database.UserCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
