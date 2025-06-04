package controllers

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/P-Tesch/go-svelte/backend/helpers"
)

type User struct {
	Username string
	Password string
}

func Register(w http.ResponseWriter, r *http.Request) {
	if !helpers.ValidateMethod(w, r, "POST") {
		return
	}
	defer r.Body.Close()
	jsonBody, err := io.ReadAll(r.Body)
	helpers.HandleError(err)

	var body User
	err = json.Unmarshal(jsonBody, &body)
	helpers.HandleError(err)

	// TODO: Register user in DB

	//

	responseBody := make(map[string]string)
	token := fmt.Sprintf("%x", sha256.Sum256([]byte(body.Username+strconv.FormatInt(time.Now().Unix(), 10))))
	responseBody["token"] = token

	rBodyJson, err := json.Marshal(responseBody)
	helpers.HandleError(err)

	w.WriteHeader(http.StatusCreated)
	w.Write(rBodyJson)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if !helpers.ValidateMethod(w, r, "GET") {
		return
	}

	user, pass, _ := r.BasicAuth()
	body := User{user, pass}

	// TODO: Check user in DB

	responseBody := make(map[string]string)
	token := fmt.Sprintf("%x", sha256.Sum256([]byte(body.Username+strconv.FormatInt(time.Now().Unix(), 10))))
	responseBody["token"] = token

	rBodyJson, err := json.Marshal(responseBody)
	helpers.HandleError(err)

	w.Write(rBodyJson)
}
