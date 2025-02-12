package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Data struct {
	CustomerId string `bson:"_id" json:"customerid"`
	ProductId  string `bson:"_id" json:"productid"`
}

func VerifyOrder(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db_instance").(*mongo.Client)
	ctx := r.Context().Value("db_ctx").(context.Context)

	var data Data
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}
	fmt.Println(data.CustomerId)

	products := db.Database("test_db").Collection("products")
	users := db.Database("test_db").Collection("customers")

	var res bson.M
	objID, _ := primitive.ObjectIDFromHex(data.ProductId)
	if err := products.FindOne(ctx, bson.M{"_id": objID}).Decode(&res); err != nil {
		w.WriteHeader(http.StatusNonAuthoritativeInfo)
		json.NewEncoder(w).Encode(map[string]bool{"success": false})
		return
	}
	objID, _ = primitive.ObjectIDFromHex(data.CustomerId)
	if err := users.FindOne(ctx, bson.M{"_id": objID}).Decode(&res); err != nil {
		w.WriteHeader(http.StatusNonAuthoritativeInfo)
		json.NewEncoder(w).Encode(map[string]bool{"success": false})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}
