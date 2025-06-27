package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/P-Tesch/go-svelte/backend/global"
	"github.com/P-Tesch/go-svelte/backend/helpers"
	"github.com/P-Tesch/go-svelte/backend/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var result []models.TransactionDTO

	transactions := models.FindTransactionsByOwner(*global.User)
	for _, transaction := range transactions {
		result = append(result, transaction.CreateDTO())
	}

	response, err := json.Marshal(result)
	helpers.HandleError(err)

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
