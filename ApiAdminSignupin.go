package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Evilcmd/Ecommerce-product-listing/internal/database/postgres"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserStruct struct {
	Username *string `json:"username"`
	Passwd   *string `json:"passwd"`
}

func (apiCfg *APIconfig) adminSignup(res http.ResponseWriter, req *http.Request) {
	userRes := UserStruct{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&userRes)
	if err != nil {
		respondWithError(res, 406, fmt.Sprintf("error decoding username and password: %v", err.Error()))
		return
	}

	if userRes.Username == nil || userRes.Passwd == nil {
		respondWithError(res, 406, "username or password not mentioned")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(*userRes.Passwd), 12)
	if err != nil {
		respondWithError(res, 406, fmt.Sprintf("error generating hash: %v", err.Error()))
		return
	}

	params := postgres.AddAdminParams{
		ID:       uuid.New(),
		Username: *userRes.Username,
		Passwd:   string(hash),
	}

	admin, err := apiCfg.DbQueries.AddAdmin(context.Background(), params)
	if err != nil {
		respondWithError(res, 406, fmt.Sprintf("error adding admin to database: %v", err.Error()))
		return
	}

	respondWithJson(res, 200, admin)
}

func (apiCfg *APIconfig) adminSignin(res http.ResponseWriter, req *http.Request) {
	userRes := UserStruct{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&userRes)
	if err != nil {
		respondWithError(res, 406, fmt.Sprintf("error decoding username and password: %v", err.Error()))
		return
	}

	if userRes.Username == nil || userRes.Passwd == nil {
		respondWithError(res, 406, "username or password not mentioned")
		return
	}

	adminRetreived, err := apiCfg.DbQueries.GetAdmin(context.Background(), *userRes.Username)
	if err != nil {
		respondWithError(res, 406, fmt.Sprintf("error fetching admin: %v", err.Error()))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(adminRetreived.Passwd), []byte(*userRes.Passwd))
	if err != nil {
		respondWithError(res, 406, "wrong password")
		return
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "ecomPL",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60)),
		Subject:   adminRetreived.ID.String(),
	})

	signedJwtTokenString, err := jwtToken.SignedString([]byte(apiCfg.jwtSecret))
	if err != nil {
		respondWithError(res, 406, fmt.Sprintf("error signing the token: %v", err.Error()))
		return
	}

	respondWithJson(res, 200, signedJwtTokenString)
}
