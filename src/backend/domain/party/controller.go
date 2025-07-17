package party

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/P-Tesch/go-svelte/backend/global"
	"github.com/P-Tesch/go-svelte/backend/helpers"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var result []PartyDTO

	parties := FindPartiesByOwner(*global.User)
	for _, parties := range parties {
		result = append(result, parties.CreateDTO())
	}

	response, err := json.Marshal(result)
	helpers.HandleError(err)

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if !helpers.ValidateMethod(w, r, "POST") {
		return
	}
	defer r.Body.Close()
	jsonBody, err := io.ReadAll(r.Body)
	helpers.HandleError(err)

	var party Party
	err = json.Unmarshal(jsonBody, &party)
	helpers.HandleError(err)
	party.Save()
	w.WriteHeader(http.StatusOK)
}
