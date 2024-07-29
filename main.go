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
	router.HandleFunc("GET /products", getAllProducts)
	router.HandleFunc("GET /products/{id}", getOneProduct)

	// Admin Routes
	adminRouter := http.NewServeMux()
	adminRouter.HandleFunc("POST /products/{id}", addProduct)
	adminRouter.HandleFunc("PATCH /products/{id}", updateProduct)
	adminRouter.HandleFunc("DELETE /products/{id}", deleteProduct)
	adminHandler := authenticate(adminRouter) // wrap the admin router in authentication middlerware

	// make the router handle the admin routes
	router.Handle("/", adminHandler)

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: router,
	}
	fmt.Printf("Starting server on port: %v\n", port)
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error in starting server: %v\n", err.Error())
	}
}
