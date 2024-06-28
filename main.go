package main

import (
	"log"
	"net/http"
)

func main() {
	// Port
	port := "8080"
	// Create a new ServerMux
	mux := http.NewServeMux()

	// Set the server struct
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// Start the server
	log.Println("Serving on port", port)
	log.Fatal(server.ListenAndServe())
}
