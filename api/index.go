package handler

import (
	"embed"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

//go:embed public/*.html
var templatesFS embed.FS

//go:embed public/*
var staticFS embed.FS


// Struct untuk resep
type Resep struct {
	ID        int
	Nama      string
	Deskripsi string
	Gambar    string
	Bahan     []string
	Instruksi []string
	Tag       []string
}

// Data resep sample
var resepData = []Resep{
	{
		ID:        1,
		Nama:      "Nasi Goreng",
		Deskripsi: "Nasi goreng khas Indonesia yang gurih dan lezat, dimasak dengan bumbu tradisional dan telur. Cocok untuk sarapan atau makan malam.",
		Gambar:    "https://images.unsplash.com/photo-1512058564366-18510be2db19?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1000&q=80",
		Bahan: []string{
			"2 piring nasi putih",
			"2 butir telur",
			"3 siung bawang putih",
			"2 siung bawang merah",
			"2 sdm kecap manis",
			"1 sdt garam",
			"Minyak goreng",
		},
		Instruksi: []string{
			"Panaskan minyak dalam wajan",
			"Tumis bawang putih dan bawang merah hingga harum",
			"Masukkan telur, orak-arik",
			"Masukkan nasi putih, aduk rata",
			"Tambahkan kecap manis dan garam",
			"Aduk hingga bumbu merata, sajikan",
		},
		Tag: []string{"nasi", "goreng", "indonesia"},
	},
	{
		ID:        2,
		Nama:      "Ayam Geprek",
		Deskripsi: "Ayam goreng crispy yang digeprek dengan sambal pedas segar. Menu favorit anak muda yang suka makanan pedas dan gurih.",
		Gambar:    "https://radarkaur.disway.id/upload/d2c6b9d19d66be8f71843509d9f44359.jpg",
		Bahan: []string{
			"1 ekor ayam, potong bagian dada",
			"5 siung bawang putih",
			"1 sdt garam",
			"10 buah cabai rawit",
			"2 buah cabai merah",
			"1 buah tomat",
			"Minyak untuk menggoreng",
		},
		Instruksi: []string{
			"Bersihkan ayam, lumuri dengan garam",
			"Goreng ayam hingga matang dan kering",
			"Haluskan cabai, bawang putih, dan tomat",
			"Geprek ayam dengan sambal hingga bumbu meresap",
			"Sajikan dengan nasi hangat",
		},
		Tag: []string{"ayam", "pedas", "geprek"},
	},
	// ... tambahkan semua resep lainnya sama seperti di main.go
	{
		ID: 15,
		Nama: "Wedang Jahe",
		Deskripsi: "Minuman hangat dari jahe segar, gula merah, dan serai—menghangatkan tubuh.",
		Gambar: "https://www.masakapahariini.com/wp-content/uploads/2021/04/shutterstock_1725419560-500x300.jpg",
		Bahan: []string{
			"2 ruas jahe, memarkan",
			"1 batang serai, memarkan",
			"500 ml air",
			"50 gr gula merah",
		},
		Instruksi: []string{
			"Rebus air bersama jahe dan serai hingga harum.",
			"Tambahkan gula merah, aduk hingga larut.",
			"Didihkan kembali 2–3 menit.",
			"Sajikan hangat.",
		},
		Tag: []string{"minuman", "hangat", "jahe"},
	},
}

// Struct untuk data template
type PageData struct {
	Recipes []Resep
	Query   string
}

// Main handler untuk Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	// Handle static files
	if strings.HasPrefix(r.URL.Path, "/css/") || 
	   strings.HasPrefix(r.URL.Path, "/js/") || 
	   strings.HasPrefix(r.URL.Path, "/images/") ||
	   strings.HasSuffix(r.URL.Path, ".png") ||
	   strings.HasSuffix(r.URL.Path, ".jpg") ||
	   strings.HasSuffix(r.URL.Path, ".ico") {
		serveStatic(w, r)
		return
	}

	// Route ke handler yang sesuai berdasarkan query parameter atau path
	if strings.Contains(r.URL.RawQuery, "resep=") || strings.HasPrefix(r.URL.Path, "/resep/") {
		detailHandler(w, r)
	} else {
		homeHandler(w, r)
	}
}

// Handler untuk static files
func serveStatic(w http.ResponseWriter, r *http.Request) {
	// Remove leading slash
	path := strings.TrimPrefix(r.URL.Path, "/")
	
	// Add public prefix
	filePath := "public/" + path
	
	content, err := staticFS.ReadFile(filePath)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	
	// Set content type based on file extension
	if strings.HasSuffix(path, ".css") {
		w.Header().Set("Content-Type", "text/css")
	} else if strings.HasSuffix(path, ".js") {
		w.Header().Set("Content-Type", "application/javascript")
	} else if strings.HasSuffix(path, ".png") {
		w.Header().Set("Content-Type", "image/png")
	} else if strings.HasSuffix(path, ".jpg") || strings.HasSuffix(path, ".jpeg") {
		w.Header().Set("Content-Type", "image/jpeg")
	}
	
	w.Write(content)
}

// Handler untuk halaman utama
func homeHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("q")))

	var filteredResep []Resep

	if query == "" {
		filteredResep = resepData
	} else {
		for _, resep := range resepData {
			// Cek apakah query cocok dengan nama resep atau tag
			if strings.Contains(strings.ToLower(resep.Nama), query) {
				filteredResep = append(filteredResep, resep)
				continue
			}

			// Cek tag
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

	// Parse template dari embedded filesystem
	tmpl, err := template.ParseFS(templatesFS, "public/index.html")
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		return
	}

}

// Handler untuk detail resep
func detailHandler(w http.ResponseWriter, r *http.Request) {
	// Extract ID dari URL atau query parameter
	var idStr string
	
	// Cek dari query parameter dulu (untuk Vercel routing)
	if resepParam := r.URL.Query().Get("resep"); resepParam != "" {
		idStr = resepParam
	} else {
		// Fallback ke path-based routing
		idStr = strings.TrimPrefix(r.URL.Path, "/resep/")
	}
	
	id, err := strconv.Atoi(idStr)
	if err != nil {
		notFoundHandler(w, r)
		return
	}

	var selectedResep *Resep
	for _, resep := range resepData {
		if resep.ID == id {
			selectedResep = &resep
			break
		}
	}

	if selectedResep == nil {
		notFoundHandler(w, r)
		return
	}

	// Parse template dari embedded filesystem
	tmpl, err := template.ParseFS(templatesFS, "public/detail.html")
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, selectedResep); err != nil {
		http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// Handler untuk halaman 404
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	
	// Parse template dari embedded filesystem
	tmpl, err := template.ParseFS(templatesFS, "public/404.html")
	if err != nil {
		http.Error(w, "404 - Page not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, nil)
}