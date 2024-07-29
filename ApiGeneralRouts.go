package main

import (
	"fmt"
	"net/http"
)

func rootEndpoint(res http.ResponseWriter, req *http.Request) {
	payload := struct {
		Message string `json:"message"`
	}{
		"Welcome to the website",
	}
	respondWithJson(res, 200, payload)
}

func getAllProducts(res http.ResponseWriter, req *http.Request) {
	// dummy
	temp := make([]productStructure, 3)
	name := "name"
	desc := "desc"
	price := 1
	temp[0] = productStructure{&name, &desc, &price}
	temp[1] = productStructure{&name, &desc, &price}
	temp[2] = productStructure{&name, &desc, &price}
	respondWithJson(res, 200, temp)
}

func getOneProduct(res http.ResponseWriter, req *http.Request) {
	name := "name"
	desc := "desc"
	price := 1
	fmt.Println(req.PathValue("id"))
	respondWithJson(res, 200, productStructure{&name, &desc, &price})
}
