package main

import (
	"log"
	"net/http"
)

func main() {
	// Create a new ServerMux
	mux := http.NewServeMux()

	// Set the server struct
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Start the server
	log.Fatal(server.ListenAndServe())
}
