package main

import (
	"log"
	"net/http"

	"EventBackend/database"
	"EventBackend/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize database
	database.InitDB()
	defer database.DB.Close()

	// Set up router
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/api/names", handlers.GetNames).Methods("GET")
	router.HandleFunc("/api/names", handlers.AddName).Methods("POST")
	router.HandleFunc("/api/names/{id}", handlers.UpdateArrivedStatus).Methods("PUT")

	// Start server
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
