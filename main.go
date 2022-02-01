package main

import (
	"backend-inventory-app/injector"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {

	//Load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	server := injector.InitializedServer()

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
