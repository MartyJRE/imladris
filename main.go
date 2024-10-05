package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"imladris/handlers"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mux := http.NewServeMux()
	port := os.Getenv("APP_PORT")
	address := fmt.Sprintf(":%s", port)

	log.Println("Starting server on port", port)

	mux.Handle("/login", handlers.LoginHandler{})
	mux.Handle("/", AuthMiddleware(handlers.HomeHandler{}))
	mux.Handle("/settings", AuthMiddleware(handlers.SettingsHandler{}))

	err = http.ListenAndServe(address, mux)
	if err != nil {
		log.Fatal("Error starting server", err)
	}
}
