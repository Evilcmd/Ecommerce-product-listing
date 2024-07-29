package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type productStructure struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Price       *int    `json:"price"`
}

func addProduct(res http.ResponseWriter, req *http.Request) {
	productReceived := productStructure{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&productReceived)
	if err != nil {
		respondWithError(res, 405, fmt.Sprintf("error while decoding request body: %v", err.Error()))
		return
	}

	// below is for dummy api
	fmt.Println(req.PathValue("id"))
	respondWithJson(res, 200, productReceived)
}

func updateProduct(res http.ResponseWriter, req *http.Request) {
	productReceived := productStructure{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&productReceived)
	if err != nil {
		respondWithError(res, 405, fmt.Sprintf("error while decoding request body: %v", err.Error()))
		return
	}

	// below is for dummy api
	fmt.Println(req.PathValue("id"))
	temp := productStructure{}
	if productReceived.Name != nil {
		temp.Name = productReceived.Name
	}
	if productReceived.Description != nil {
		temp.Description = productReceived.Description
	}
	if productReceived.Price != nil {
		temp.Price = productReceived.Price
	}
	respondWithJson(res, 200, temp)

}

func deleteProduct(res http.ResponseWriter, req *http.Request) {

	// below is for dummy api
	fmt.Println(req.PathValue("id"))
}
