package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juanvaleriand/go-test/middlewares"
)

func New() http.Handler {
	router := mux.NewRouter()

	// Public endpoints
	public := router.PathPrefix("/api/v1").Subrouter()
	public.HandleFunc("/users", GetAllUsers).Methods("GET")
	public.HandleFunc("/user/{id}", GetUser).Methods("GET")
	public.HandleFunc("/user", CreateUser).Methods("POST")
	public.HandleFunc("/user/{id}", UpdateUser).Methods("PUT")

	// Protected delete endpoint
	protected := router.PathPrefix("/api/v1").Subrouter()
	protected.Use(middlewares.AuthMiddleware)
	protected.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")

	return router
}
