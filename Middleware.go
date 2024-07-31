package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func (apiCfg *APIconfig) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		signedJwtTokenString := req.Header.Get("Authorization")
		if len(signedJwtTokenString) == 0 {
			respondWithError(res, 406, "expected authoriztion")
			return
		}
		signedJwtTokenString = strings.Split(signedJwtTokenString, " ")[1]

		_, err := jwt.ParseWithClaims(signedJwtTokenString, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(apiCfg.jwtSecret), nil
		})
		if err != nil {
			respondWithError(res, http.StatusUnauthorized, fmt.Sprintf("Unauthorized: %v", err.Error()))
			return
		}

		// skipping refresh tokens for now

		next.ServeHTTP(res, req)
	})
}
