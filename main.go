package main

import (
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"

	"EventBackend/database"
	"EventBackend/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	// Initialize database
	database.InitDB()
	defer database.DB.Close()

	// Set up router
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/api/names", handlers.GetNames).Methods("GET")
	router.HandleFunc("/api/names", handlers.AddName).Methods("POST")
	router.HandleFunc("/api/names/{id}", handlers.UpdateArrivedStatus).Methods("PUT")
	router.HandleFunc("/api/names/{id}", handlers.DeleteName).Methods("DELETE") // Add

	// Start server
	log.Println("Server running on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
