package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db_instance").(*mongo.Client)
	ctx := r.Context().Value("db_ctx").(context.Context)

	var data []bson.M
	// if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	customers := db.Database("test_db").Collection("customers")

	res, err := customers.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer res.Close(ctx)

	for res.Next(ctx) {
		var result bson.M
		if err := res.Decode(&result); err != nil {
			log.Fatal(err)
		}
		data = append(data, result)
	}

	if err := res .Err(); err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

	fmt.Println("Users fetched: ", res)

}
