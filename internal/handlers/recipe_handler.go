package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"my-recipe/internal/data"
	"my-recipe/internal/models"
)

type PageData struct {
	Recipes []models.Resep
	Query   string
}

// Handler home
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		NotFoundHandler(w, r)
		return
	}

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

// Handler detail
func DetailHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/resep/")
	id, err := strconv.Atoi(path)
	if err != nil {
		NotFoundHandler(w, r)
		return
	}

	var selectedResep *models.Resep
	for _, resep := range data.ResepData {
		if resep.ID == id {
			selectedResep = &resep
			break
		}
	}

	if selectedResep == nil {
		NotFoundHandler(w, r)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/detail.html"))
	tmpl.Execute(w, selectedResep)
}

// Handler not found
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tmpl := template.Must(template.ParseFiles("templates/404.html"))
	tmpl.Execute(w, nil)
}
