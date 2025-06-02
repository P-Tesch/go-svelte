package controllers

import (
	"net/http"
)

func TestEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TESTE"))
}
