package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Evilcmd/Ecommerce-product-listing/internal/database/postgres"
	"github.com/google/uuid"
)

type productStructure struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Price       *int    `json:"price"`
}

func (apiCfg *APIconfig) addProduct(res http.ResponseWriter, req *http.Request) {
	productReceived := productStructure{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&productReceived)
	if err != nil {
		respondWithError(res, 405, fmt.Sprintf("error while decoding request body: %v", err.Error()))
		return
	}

	if productReceived.Name == nil || productReceived.Description == nil || productReceived.Price == nil {
		respondWithError(res, http.StatusBadRequest, "name, description, and price are required fields")
		return
	}

	id := uuid.New()

	params := postgres.AddProductParams{
		ID:          id,
		Name:        *productReceived.Name,
		Description: *productReceived.Description,
		Price:       int32(*productReceived.Price),
	}

	product, err := apiCfg.DbQueries.AddProduct(context.Background(), params)
	if err != nil {
		respondWithError(res, 406, fmt.Sprintf("error adding to databade: %v", err.Error()))
		return
	}

	respondWithJson(res, 200, product)
}

func (apiCfg *APIconfig) updateProduct(res http.ResponseWriter, req *http.Request) {
	productReceived := productStructure{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&productReceived)
	if err != nil {
		respondWithError(res, 405, fmt.Sprintf("error while decoding request body: %v", err.Error()))
		return
	}

	pathid := req.PathValue("id")
	id, err := uuid.Parse(pathid)
	if err != nil {
		respondWithError(res, 406, fmt.Sprintf("error parsing uuid: %v, error message: %v", pathid, err.Error()))
		return
	}

	if productReceived.Name != nil {
		params := postgres.UpdateProductNameParams{
			ID:   id,
			Name: *productReceived.Name,
		}
		err = apiCfg.DbQueries.UpdateProductName(context.Background(), params)
		if err != nil {
			respondWithError(res, 406, fmt.Sprintf("error updating product name in databse: %v", err.Error()))
			return
		}
	}
	if productReceived.Description != nil {
		params := postgres.UpdateProductDescriptionParams{
			ID:          id,
			Description: *productReceived.Description,
		}
		err = apiCfg.DbQueries.UpdateProductDescription(context.Background(), params)
		if err != nil {
			respondWithError(res, 406, fmt.Sprintf("error updating product description in databse: %v", err.Error()))
			return
		}
	}
	if productReceived.Price != nil {
		params := postgres.UpdateProductPriceParams{
			ID:    id,
			Price: int32(*productReceived.Price),
		}
		err = apiCfg.DbQueries.UpdateProductPrice(context.Background(), params)
		if err != nil {
			respondWithError(res, 406, fmt.Sprintf("error updating product name in databse: %v", err.Error()))
			return
		}
	}

	apiCfg.redisClient.Expire(context.Background(), id.String(), time.Millisecond)

	respondWithJson(res, 200, "OK")

}

func (apiCfg *APIconfig) deleteProduct(res http.ResponseWriter, req *http.Request) {
	pathid := req.PathValue("id")
	id, err := uuid.Parse(pathid)
	if err != nil {
		respondWithError(res, 406, fmt.Sprintf("error parsing uuid: %v, error message: %v", pathid, err.Error()))
		return
	}

	err = apiCfg.DbQueries.DeleteProduct(context.Background(), id)
	if err != nil {
		respondWithError(res, 406, fmt.Sprintf("error deleting product from database: %v", err.Error()))
		return
	}

	apiCfg.redisClient.Expire(context.Background(), id.String(), time.Millisecond)

	respondWithJson(res, 200, "ok")
}
