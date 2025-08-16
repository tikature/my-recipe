package main

import (
	"html/template"
	"net/http"
	"strings"

	"my-recipe/internal/data"
	"my-recipe/internal/models"
)

type PageData struct {
	Recipes []models.Resep
	Query   string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("q")))
	var filteredResep []models.Resep

	if query == "" {
		filteredResep = data.ResepData
	} else {
		for _, resep := range data.ResepData {
			if strings.Contains(strings.ToLower(resep.Nama), query) {
				filteredResep = append(filteredResep, resep)
				continue
			}
			for _, tag := range resep.Tag {
				if strings.Contains(strings.ToLower(tag), query) {
					filteredResep = append(filteredResep, resep)
					break
				}
			}
		}
	}

	data := PageData{
		Recipes: filteredResep,
		Query:   r.URL.Query().Get("q"),
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, data)
}
