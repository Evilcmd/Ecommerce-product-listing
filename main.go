package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/Evilcmd/Ecommerce-product-listing/internal/database/postgres"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type APIconfig struct {
	DbQueries *postgres.Queries
	jwtSecret string
}

func main() {
	godotenv.Load()
	dbUrl := os.Getenv("DBURL")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	port := os.Getenv("PORT")
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		fmt.Printf("error starting databse connection: %v", err.Error())
		return
	}

	dbQeries := postgres.New(db)
	apiConfig := APIconfig{
		DbQueries: dbQeries,
		jwtSecret: JWT_SECRET,
	}

	router := http.NewServeMux()

	// Test routes
	router.HandleFunc("GET /health", checkHealth)
	router.HandleFunc("GET /error", errCheck)

	// General Routes
	router.HandleFunc("GET /", rootEndpoint)
	router.HandleFunc("GET /products", apiConfig.getAllProducts)
	router.HandleFunc("GET /products/{id}", apiConfig.getOneProduct)

	// Sign up admin
	router.HandleFunc("POST /admin/signup", apiConfig.adminSignup)
	router.HandleFunc("POST /admin/signin", apiConfig.adminSignin)

	// Admin Routes
	adminRouter := http.NewServeMux()
	adminRouter.HandleFunc("POST /products", apiConfig.addProduct)
	adminRouter.HandleFunc("PATCH /products/{id}", apiConfig.updateProduct)
	adminRouter.HandleFunc("DELETE /products/{id}", apiConfig.deleteProduct)
	adminHandler := apiConfig.authenticate(adminRouter) // wrap the admin router in authentication middleware

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
