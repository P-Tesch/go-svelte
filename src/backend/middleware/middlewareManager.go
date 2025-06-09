package middleware

import (
	"net/http"
)

func Default(w http.ResponseWriter, r *http.Request) bool {
	return authenticate(w, r)
}
