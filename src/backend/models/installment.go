package models

import (
	"context"
	"time"

	"github.com/P-Tesch/go-svelte/backend/database"
	"github.com/P-Tesch/go-svelte/backend/helpers"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var installColl = database.GetDatabase().Collection("installments")

// TODO
type Installment struct {
	Id      *bson.ObjectID    `bson:"_id,omitempty"`
	Owner   User              `bson:"owner"`
	Name    string            `bson:"name"`
	Type    TransactionType   `bson:"type"`
	Value   uint32            `bson:"value"`
	DueDate bson.DateTime     `bson:"due_date"`
	Status  TransactionStatus `bson:"status"`
}

type InstallmentDTO struct {
	Name    string
	Type    uint8
	Value   uint32
	DueDate string
	Status  uint8
}

func (installment *Installment) FindById() bool {
	result := installColl.FindOne(context.TODO(), bson.D{{Key: "id", Value: installment.Id}})
	err := result.Decode(installment)
	helpers.HandleError(err)
	return true
}

func (installment *Installment) Save() interface{} {
	if installment.Id == nil {
		result, err := installColl.InsertOne(context.TODO(), installment)
		helpers.HandleError(err)

		return result.InsertedID
	}

	_, err := installColl.UpdateByID(context.TODO(), installment.Id, bson.M{"$set": installment})
	helpers.HandleError(err)

	return installment.Id
}

func (installment Installment) Delete() bool {
	// TODO: Delete from DB
	return false
}

func (installment Installment) CreateDTO() InstallmentDTO {
	return InstallmentDTO{
		Type:    uint8(installment.Type),
		Value:   installment.Value,
		DueDate: installment.DueDate.Time().Format(time.DateOnly),
		Status:  uint8(installment.Status),
	}
}

func FindInstallmentsByOwner(owner User) []Installment {
	result, err := installColl.Find(context.TODO(), bson.D{{Key: "owner", Value: owner}})
	helpers.HandleError(err)

	var installments []Installment
	result.All(context.TODO(), installments)
	return installments
}
