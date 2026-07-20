package main

import (
	"log"
	"net/http"
)

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Create a file server for the current directory
	fileServer := http.FileServer(http.Dir("."))

	// Serve files from the root path
	mux.Handle("/", fileServer)

	// Create the HTTP server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Server running at http://localhost:8080")

	// Start the server
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}