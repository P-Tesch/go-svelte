package main

import (
	"net/http"

	"github.com/P-Tesch/go-svelte/routes"
)

func main() {
	routes.RegisterFrontRoutes()
	routes.RegisterAPIRoutes()

	http.ListenAndServe(":8888", nil)
}
