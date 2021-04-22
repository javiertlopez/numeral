package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

const (
	writeTimeout = 15 * time.Second
	readTimeout  = 15 * time.Second
	idleTimeout  = 60 * time.Second
)

func main() {
	// Environment variables
	addr := os.Getenv("ADDR")
	mongoString := os.Getenv("MONGO_STRING")

	application := New(
		AppConfig{
			MongoURI: mongoString,
		},
	)

	// Create a Gorilla Mux router
	router := application.router

	// Create a Server instance with the router
	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
		IdleTimeout:  idleTimeout,
		Handler:      router,
	}

	// Start the server
	log.Fatal(srv.ListenAndServe())
}
