package main

import (
	"fmt"
	"log"
	"net/http"
	"project/controller"
	"project/model"
)

func main() {
	db, err := model.Connect()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	mux := controller.Register()

	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
