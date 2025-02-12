package main

import (
	"log"
	"net/http"

	"main.go/routes"
	"main.go/services/dbconn"
)

func main() {
	r := routes.GetRouter()

	dbconn.Connect("mongodb://localhost:27017")

	err := http.ListenAndServe(":1800", r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
