package routes

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/P-Tesch/go-svelte/backend/helpers"
	"github.com/P-Tesch/go-svelte/backend/middleware"
)

func RegisterFrontRoutes() {
	env := helpers.GetEnvVar("ENVIROMENT")

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
		if strings.HasPrefix(r.URL.Path, "/@") || strings.HasPrefix(r.URL.Path, "/.") || strings.HasPrefix(r.URL.Path, "/node_modules") || strings.HasPrefix(r.URL.Path, "/src") || strings.EqualFold(r.Header.Get("Upgrade"), "websocket") {
			proxy.ServeHTTP(w, r)
			return
		}

		if r.URL.Path != "/login" && !middleware.Default(w, r) {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}

		proxy.ServeHTTP(w, r)
	})
}

func prodFrontend() {
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !middleware.Default(w, r) {
			return
		}

		http.ServeFile(w, r, "./dist/index.html")
	})
}
