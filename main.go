package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/juanvaleriand/go-test/controllers"
	"github.com/juanvaleriand/go-test/models"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	handler := controllers.New()

	fmt.Println(os.Getenv("BASE_URL"))

	server := &http.Server{
		Addr:    os.Getenv("BASE_URL"),
		Handler: handler,
	}

	models.ConnectDatabase()

	server.ListenAndServe()
}
