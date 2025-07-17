package routes

import (
	"net/http"

	"github.com/P-Tesch/go-svelte/backend/domain/party"
	"github.com/P-Tesch/go-svelte/backend/domain/transaction"
	"github.com/P-Tesch/go-svelte/backend/domain/user"
)

func RegisterAPIRoutes() {
	http.HandleFunc("/api/register", user.Register)
	http.HandleFunc("/api/login", user.Login)
	http.HandleFunc("/api/transactions", transaction.Index)
	http.HandleFunc("GET /api/parties", party.Index)
	http.HandleFunc("POST /api/parties", party.Add)
}
