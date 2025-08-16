package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"my-recipe/internal/data"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Redirect(w, r, "/api/notfound", http.StatusTemporaryRedirect)
		return
	}

	var selectedResep *data.Resep
	for _, resep := range data.ResepData {
		if resep.ID == id {
			selectedResep = &resep
			break
		}
	}

	if selectedResep == nil {
		http.Redirect(w, r, "/api/notfound", http.StatusTemporaryRedirect)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/detail.html"))
	tmpl.Execute(w, selectedResep)
}
