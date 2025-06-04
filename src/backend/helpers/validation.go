package helpers

import (
	"net/http"
)

func ValidateMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method " + method + " not allowed"))
		return false
	}
	return true
}
