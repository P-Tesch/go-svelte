package middleware

import (
	"net/http"

	"github.com/P-Tesch/go-svelte/backend/models"
)

func authenticate(w http.ResponseWriter, r *http.Request) bool {
	token, err := r.Cookie("auth")
	if err != nil {
		return false
	}

	user := models.User{Token: token.Value}
	if !user.FindByToken() {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("User not found"))
		return false
	}

	return true
}
