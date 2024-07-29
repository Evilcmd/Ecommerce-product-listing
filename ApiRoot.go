package main

import "net/http"

func rootEndpoint(res http.ResponseWriter, req *http.Request) {
	payload := struct {
		Message string `json:"message"`
	}{
		"Welcome to the website",
	}
	respondWithJson(res, 200, payload)
}
