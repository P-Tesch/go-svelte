package routes

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/P-Tesch/go-svelte/backend/helpers/dotenv"
)

func RegisterFrontRoutes() {
	env := dotenv.GetVar("ENVIROMENT")

	if env == "dev" {
		fmt.Println("Serving dev server on port 8888")
		devFrontend()
	} else {
		fmt.Println("Serving files on port 8888")
		prodFrontend()
	}
}

func devFrontend() {
	target, _ := url.Parse("http://svelte:5173")

	proxy := httputil.NewSingleHostReverseProxy(target)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})
}

func prodFrontend() {
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./dist/index.html")
	})
}
