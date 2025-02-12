package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Email    string `bson:"email" json:"email"`
	Mobile   string `bson:"mobile" json:"mobile"`
	Password string `bson:"password" json:"password"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db_instance").(*mongo.Client)
	ctx := r.Context().Value("db_ctx").(context.Context)

	var data User
	var abc io.Reader= io.ReadCloser(r.Body)
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	customers := db.Database("test_db").Collection("customers")

	_, err := customers.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User created")

}
