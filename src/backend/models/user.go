package models

import (
	"context"

	"github.com/P-Tesch/go-svelte/backend/database"
	"github.com/P-Tesch/go-svelte/backend/helpers"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var coll = database.GetDatabase().Collection("users")

type User struct {
	_id      *bson.ObjectID
	Username string
	Password string
	Token    string
}

func (user User) FindById() bool {
	result := coll.FindOne(context.TODO(), bson.D{{Key: "_id", Value: user._id}})
	err := result.Decode(user)
	helpers.HandleError(err)
	return true
}

func (user User) FindByUsername() bool {
	result, err := coll.Find(context.TODO(), bson.D{{Key: "Username", Value: user.Username}})
	helpers.HandleError(err)

	var users []User
	err = result.All(context.TODO(), &users)
	helpers.HandleError(err)

	if len(users) == 0 {
		return false
	}

	user = users[0]
	return true
}

func (user User) FindByToken() bool {
	result, err := coll.Find(context.TODO(), bson.D{{Key: "Token", Value: user.Token}})
	helpers.HandleError(err)

	var users []User
	err = result.All(context.TODO(), &users)
	helpers.HandleError(err)

	if len(users) == 0 {
		return false
	}

	user = users[0]
	return true
}

func (user User) Save() interface{} {
	if &user._id != nil {
		result, err := coll.InsertOne(context.TODO(), user)
		helpers.HandleError(err)

		return result.InsertedID
	}

	_, err := coll.UpdateByID(context.TODO(), user._id, user)
	helpers.HandleError(err)

	return user._id
}

func (user User) Delete() bool {
	// TODO: Delete from DB
	return false
}
