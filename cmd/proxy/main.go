package main

import (
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	_ "halyk-life_proxy_server/docs"
	"log"
	"net/http"
)

var app App

// @title Proxy Server API
// @version 1.0
// @description This is a proxy server API.
// @host localhost:8080
// @BasePath /
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/proxy", app.handleProxy).Methods("POST")
	r.HandleFunc("/history", app.retrieveHistory).Methods("GET")
	r.HandleFunc("/health", app.healthCheck).Methods("GET")
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
