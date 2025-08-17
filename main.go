package main

import (
	"log"
	"net/http"

	"my-recipe/internal/data"
	"my-recipe/internal/handlers"
)

func main() {
	log.Printf("Starting server with %d recipes loaded", len(data.ResepData))

	// static
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// routes
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/resep/", handlers.DetailHandler)

	log.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
