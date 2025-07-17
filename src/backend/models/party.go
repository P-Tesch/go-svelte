package models

import (
	"context"

	"github.com/P-Tesch/go-svelte/backend/database"
	"github.com/P-Tesch/go-svelte/backend/helpers"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var partyColl = database.GetDatabase().Collection("parties")

type Party struct {
	Id    *bson.ObjectID  `bson:"_id,omitempty"`
	Owner User            `bson:"owner"`
	Name  string          `bson:"name"`
	Type  TransactionType `bson:"type"`
}

type PartyDTO struct {
	Name string
	Type uint8
}

func (party *Party) FindById() bool {
	result := partyColl.FindOne(context.TODO(), bson.D{{Key: "id", Value: party.Id}})
	err := result.Decode(party)
	helpers.HandleError(err)
	return true
}

func (party *Party) Save() interface{} {
	if party.Id == nil {
		result, err := partyColl.InsertOne(context.TODO(), party)
		helpers.HandleError(err)

		return result.InsertedID
	}

	_, err := partyColl.UpdateByID(context.TODO(), party.Id, bson.M{"$set": party})
	helpers.HandleError(err)

	return party.Id
}

func (party Party) Delete() bool {
	// TODO: Delete from DB
	return false
}

func (party Party) CreateDTO() InstallmentDTO {
	return InstallmentDTO{
		Type: uint8(party.Type),
		Name: party.Name,
	}
}

func FindPartiesByOwner(owner User) []Installment {
	result, err := partyColl.Find(context.TODO(), bson.D{{Key: "owner", Value: owner}})
	helpers.HandleError(err)

	var parties []Installment
	result.All(context.TODO(), parties)
	return parties
}
