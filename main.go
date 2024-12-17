package main

import (
	"log"
	"net/http"
	"os"
	"server/config"
	"server/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
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

	config.ConnectDB()
	port := os.Getenv("PORT")

	if port == "" {
		port = ":8080"
	}
	// fmt.println("hello world")
	r := mux.NewRouter()
	routes.AuthRoutes(r)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	corsRouter := enableCORS(r)
	log.Println("Server running in ", port)
	log.Fatal(http.ListenAndServe(":"+port, corsRouter))
}
