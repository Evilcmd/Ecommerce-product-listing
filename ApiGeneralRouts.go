package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func rootEndpoint(res http.ResponseWriter, req *http.Request) {
	payload := struct {
		Message string `json:"message"`
	}{
		"Welcome to the website",
	}
	respondWithJson(res, 200, payload)
}

func (apiCfg *APIconfig) getAllProducts(res http.ResponseWriter, req *http.Request) {
	prods, err := apiCfg.DbQueries.GetAllProducts(context.Background())
	if err != nil {
		respondWithError(res, 406, fmt.Sprintf("error fetching from database: %v", err.Error()))
	}
	respondWithJson(res, 200, prods)
}

func (apiCfg *APIconfig) getOneProduct(res http.ResponseWriter, req *http.Request) {
	pathid := req.PathValue("id")
	id, err := uuid.Parse(pathid)
	if err != nil {
		respondWithError(res, 406, fmt.Sprintf("error parsing uuid: %v, error message: %v", pathid, err.Error()))
		return
	}
	prod, err := apiCfg.DbQueries.GetProduct(context.Background(), id)
	if err != nil {
		respondWithError(res, 406, fmt.Sprintf("error fetching from database: %v", err.Error()))
		return
	}
	respondWithJson(res, 200, prod)
}
