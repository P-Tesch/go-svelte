package middleware

import (
	"net/http"
)

func Default(w http.ResponseWriter, r *http.Request) bool {
	cors(w)
	return authenticate(w, r)
}
