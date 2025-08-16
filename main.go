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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // lokal
	}

	log.Printf("Server berjalan di http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
