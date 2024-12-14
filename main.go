package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve static files from the templates directory
	fs := http.FileServer(http.Dir("templates"))
	http.Handle("/", fs)

	// Start server
	log.Println("Server starting on http://0.0.0.0:3000")
	if err := http.ListenAndServe("0.0.0.0:3000", nil); err != nil {
		log.Fatal(err)
	}
}