package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

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

	redisProd := apiCfg.redisClient.JSONGet(context.Background(), id.String(), "$")
	x, err := redisProd.Result()
	if err == nil && x != "" {
		fmt.Println("cache hit", x)
		// Too lazy to unmarshal into proper format
		respondWithJson(res, 200, x)
		return
	}

	prod, err := apiCfg.DbQueries.GetProduct(context.Background(), id)
	if err != nil {
		respondWithError(res, 406, fmt.Sprintf("error fetching from database: %v", err.Error()))
		return
	}

	apiCfg.redisClient.JSONSet(context.Background(), id.String(), "$", prod)
	apiCfg.redisClient.Expire(context.Background(), id.String(), time.Second*20)

	respondWithJson(res, 200, prod)
}
