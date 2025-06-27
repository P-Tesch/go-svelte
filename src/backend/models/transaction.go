package models

import (
	"context"
	"time"

	"github.com/P-Tesch/go-svelte/backend/database"
	"github.com/P-Tesch/go-svelte/backend/helpers"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var transColl = database.GetDatabase().Collection("transactions")

type TransactionType int
type TransactionStatus int

const (
	Income TransactionType = iota
	Outgoing
)

const (
	Paid TransactionStatus = iota
	Pending
	Late
)

type Transaction struct {
	Id      *bson.ObjectID    `bson:"_id,omitempty"`
	Owner   User              `bson:"owner"`
	Type    TransactionType   `bson:"type"`
	Value   uint32            `bson:"value"`
	DueDate bson.DateTime     `bson:"due_date"`
	Status  TransactionStatus `bson:"status"`
}

type TransactionDTO struct {
	Type    uint8
	Value   uint32
	DueDate string
	Status  uint8
}

func (transaction *Transaction) FindById() bool {
	result := transColl.FindOne(context.TODO(), bson.D{{Key: "id", Value: transaction.Id}})
	err := result.Decode(transaction)
	helpers.HandleError(err)
	return true
}

func (transaction *Transaction) Save() interface{} {
	if transaction.Id == nil {
		result, err := transColl.InsertOne(context.TODO(), transaction)
		helpers.HandleError(err)

		return result.InsertedID
	}

	_, err := transColl.UpdateByID(context.TODO(), transaction.Id, bson.M{"$set": transaction})
	helpers.HandleError(err)

	return transaction.Id
}

func (transaction Transaction) Delete() bool {
	// TODO: Delete from DB
	return false
}

func (transaction Transaction) CreateDTO() TransactionDTO {
	return TransactionDTO{
		Type:    uint8(transaction.Type),
		Value:   transaction.Value,
		DueDate: transaction.DueDate.Time().Format(time.DateOnly),
		Status:  uint8(transaction.Status),
	}
}

func FindTransactionsByOwner(owner User) []Transaction {
	result, err := transColl.Find(context.TODO(), bson.D{{Key: "owner", Value: owner}})
	helpers.HandleError(err)

	var transactions []Transaction
	result.All(context.TODO(), transactions)
	return transactions
}
