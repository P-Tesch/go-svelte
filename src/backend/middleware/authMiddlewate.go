package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/P-Tesch/go-svelte/backend/domain/user"
	"github.com/P-Tesch/go-svelte/backend/global"
)

func authenticate(w http.ResponseWriter, r *http.Request) bool {
	token, err := r.Cookie("AuthToken")
	if err != nil {
		return false
	}

	user := user.User{Token: token.Value}
	if !user.FindByToken() {
		w.WriteHeader(http.StatusUnauthorized)
		resp, _ := json.Marshal(map[string]string{"error": "Invalid auth token"})
		w.Write([]byte(resp))
		return false
	}
	global.User = &user

	return true
}
