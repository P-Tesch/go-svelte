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
	"github.com/P-Tesch/go-svelte/backend/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if !helpers.ValidateMethod(w, r, "POST") {
		return
	}
	defer r.Body.Close()
	jsonBody, err := io.ReadAll(r.Body)
	helpers.HandleError(err)

	var user models.User
	err = json.Unmarshal(jsonBody, &user)
	helpers.HandleError(err)

	if user.FindByUsername() {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Username already taken"))
		return
	}

	responseBody := make(map[string]string)
	token := fmt.Sprintf("%x", sha256.Sum256([]byte(user.Username+strconv.FormatInt(time.Now().Unix(), 10))))
	responseBody["token"] = token
	user.Token = token

	user.Save()

	rBodyJson, err := json.Marshal(responseBody)
	helpers.HandleError(err)

	w.WriteHeader(http.StatusCreated)
	w.Write(rBodyJson)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if !helpers.ValidateMethod(w, r, "GET") {
		return
	}

	username, pass, _ := r.BasicAuth()
	user := models.User{Username: username}

	if !user.FindByUsername() {
		w.WriteHeader(http.StatusUnprocessableEntity)
		resp, _ := json.Marshal(map[string]string{"error": "Incorrect username or password"})
		w.Write([]byte(resp))
		return
	}

	if user.Password != pass {
		w.WriteHeader(http.StatusUnprocessableEntity)
		resp, _ := json.Marshal(map[string]string{"error": "Incorrect username or password"})
		w.Write([]byte(resp))
		return
	}

	token := fmt.Sprintf("%x", sha256.Sum256([]byte(user.Username+strconv.FormatInt(time.Now().Unix(), 10))))

	user.Token = token
	user.Save()

	cookie := http.Cookie{
		Name:     "AuthToken",
		Value:    token,
		MaxAge:   0,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteDefaultMode,
	}

	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}
