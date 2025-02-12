package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type Product struct {
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	Price       string `bson:"price" json:"price"`
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db_instance").(*mongo.Client)
	ctx := r.Context().Value("db_ctx").(context.Context)

	var data Product
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	products := db.Database("test_db").Collection("products")

	_, err := products.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Product created")

}
