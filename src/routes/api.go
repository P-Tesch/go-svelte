package routes

import (
	"net/http"

	"github.com/P-Tesch/go-svelte/backend/controllers"
)

func RegisterAPIRoutes() {
	http.HandleFunc("/api/register", controllers.Register)
	http.HandleFunc("/api/login", controllers.Login)
}
