package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	router := http.NewServeMux()

	// Test routes
	router.HandleFunc("GET /health", checkHealth)
	router.HandleFunc("GET /error", errCheck)

	// General Routes
	router.HandleFunc("GET /", rootEndpoint)

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: router,
	}
	fmt.Printf("Starting server on port: %v\n", port)
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error in starting server: %v\n", err.Error())
	}
}
