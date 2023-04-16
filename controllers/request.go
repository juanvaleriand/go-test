package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/joho/godotenv"
	"github.com/juanvaleriand/go-test/models"
)

func MakeRequest(method, url string, body interface{}, isAuthenticatedRequest bool) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))

	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if isAuthenticatedRequest {
		request.Header.Add("Authorization", os.Getenv("TOKEN"))
	}

	models.ConnectDatabase()

	writer := httptest.NewRecorder()
	New().ServeHTTP(writer, request)

	return writer
}
