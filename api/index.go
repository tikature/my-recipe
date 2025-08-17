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
	{
		ID:        3,
		Nama:      "Gado-Gado",
		Deskripsi: "Salad tradisional Indonesia yang sehat dan bergizi, terdiri dari berbagai sayuran segar yang disiram dengan bumbu kacang yang kaya rasa.",
		Gambar:    "https://ilovepeanutbutter.com/cdn/shop/articles/GadoGado_33.jpg?v=1651069093",
		Bahan: []string{
			"100g tahu, goreng",
			"100g tempe, goreng",
			"2 butir telur rebus",
			"100g tauge",
			"100g kangkung",
			"2 buah kentang rebus",
			"Kerupuk",
		},
		Instruksi: []string{
			"Rebus sayuran hingga matang",
			"Potong-potong tahu, tempe, dan telur",
			"Tata semua bahan dalam piring",
			"Siram dengan bumbu kacang",
			"Taburi kerupuk dan sajikan",
		},
		Tag: []string{"sayur", "sehat", "indonesia"},
	},
	{
		ID:        4,
		Nama:      "Rendang Daging",
		Deskripsi: "Masakan khas Minang yang terkenal di dunia. Daging sapi yang dimasak dengan santan dan rempah-rempah hingga empuk dan bercita rasa kaya.",
		Gambar:    "https://indonesiakaya.com/wp-content/uploads/2023/04/ren_Artboard_4.jpg",
		Bahan: []string{
			"1 kg daging sapi, potong kotak",
			"1 liter santan kental",
			"8 siung bawang merah",
			"6 siung bawang putih",
			"10 cabai merah",
			"3 cm jahe",
			"2 batang serai",
			"5 lembar daun jeruk",
		},
		Instruksi: []string{
			"Haluskan bumbu: bawang merah, bawang putih, cabai, jahe",
			"Tumis bumbu halus hingga harum",
			"Masukkan daging, masak hingga berubah warna",
			"Tuang santan, masak dengan api kecil",
			"Tambahkan serai dan daun jeruk",
			"Masak hingga santan mengering dan daging empuk (2-3 jam)",
		},
		Tag: []string{"daging", "rendang", "minang", "pedas"},
	},
	{
		ID:        5,
		Nama:      "Soto Ayam",
		Deskripsi: "Sup ayam tradisional Indonesia dengan kuah bening yang segar, dilengkapi dengan tauge, telur rebus, dan kerupuk. Cocok untuk segala cuaca.",
		Gambar:    "https://www.unileverfoodsolutions.co.id/dam/global-ufs/mcos/SEA/calcmenu/recipes/ID-recipes/chicken-&-other-poultry-dishes/soto-ayam/main-header.jpg",
		Bahan: []string{
			"1 ekor ayam kampung",
			"2 liter air",
			"200g tauge",
			"2 butir telur rebus",
			"100g mie soun",
			"4 siung bawang putih",
			"1 sdt garam",
			"Kerupuk dan sambal",
		},
		Instruksi: []string{
			"Rebus ayam hingga empuk, suwir-suwir",
			"Saring kaldu ayam",
			"Tumis bawang putih hingga harum",
			"Masukkan kaldu, didihkan",
			"Siapkan mangkok dengan tauge dan mie soun",
			"Tuang kuah panas, tambahkan ayam suwir dan telur",
		},
		Tag: []string{"soto", "ayam", "sup", "indonesia"},
	},
	{
	ID: 6,
	Nama: "Es Kopi Susu",
	Deskripsi: "Minuman kekinian yang menyegarkan: kopi robusta, susu segar, dan gula aren, disajikan dengan es.",
	Gambar: "https://jasindo.co.id/uploads/media/otwgs4vrqfbvvy30ymyeelsug-kopijpg",
	Bahan: []string{
		"100 ml kopi hitam kental",
		"120 ml susu segar",
		"2 sdm gula aren cair",
		"Es batu secukupnya",
	},
	Instruksi: []string{
		"Seduh kopi hitam kental lalu biarkan hangat.",
		"Isi gelas dengan es batu.",
		"Tuang kopi, susu, dan gula aren.",
		"Aduk rata dan sajikan dingin.",
	},
	Tag: []string{"minuman", "kopi", "gula-aren"},
	},
	{
		ID: 7,
		Nama: "Mie Goreng Jawa",
		Deskripsi: "Mie goreng manis-gurih khas Jawa dengan ayam, telur, dan sayuran.",
		Gambar: "https://allofresh.id/blog/wp-content/uploads/2023/09/cara-membuat-mie-goreng-3-1.jpg",
		Bahan: []string{
			"200 gr mie telur",
			"100 gr ayam suwir",
			"2 butir telur",
			"3 siung bawang putih, cincang",
			"2 sdm kecap manis",
			"1 sdt garam",
			"Sayuran (kol, sawi) secukupnya",
			"Minyak untuk menumis",
		},
		Instruksi: []string{
			"Rebus mie hingga al dente, tiriskan.",
			"Tumis bawang putih hingga harum.",
			"Masukkan ayam dan telur, orak-arik.",
			"Tambahkan sayuran, aduk layu.",
			"Masukkan mie, kecap, dan garam; aduk hingga merata.",
		},
		Tag: []string{"mie", "goreng", "jawa"},
	},
	{
		ID: 8,
		Nama: "Rawon Sapi",
		Deskripsi: "Sup daging khas Jawa Timur dengan kuah hitam dari kluwek, kaya rempah.",
		Gambar: "https://indonesiakaya.com/wp-content/uploads/2023/04/ra_Artboard_16.jpg",
		Bahan: []string{
			"500 gr daging sapi (sandung lamur), potong",
			"3 butir kluwek, ambil isinya",
			"5 siung bawang merah",
			"3 siung bawang putih",
			"1 sdt ketumbar sangrai",
			"1 batang serai, memarkan",
			"2 lembar daun jeruk",
			"Garam dan gula secukupnya",
			"Air untuk kuah",
		},
		Instruksi: []string{
			"Rebus daging hingga setengah empuk, sisihkan kaldu.",
			"Haluskan bawang, ketumbar, dan kluwek; tumis hingga harum.",
			"Masukkan bumbu tumis, serai, dan daun jeruk ke kaldu.",
			"Tambahkan daging; bumbui garam dan gula, masak hingga empuk.",
			"Sajikan panas dengan nasi, tauge rebus, dan bawang goreng.",
		},
		Tag: []string{"rawon", "sapi", "jawa-timur"},
	},
	{
		ID: 9,
		Nama:      "Ayam Rica-Rica",
		Deskripsi: "Masakan khas Manado dengan ayam pedas berbumbu cabai, bawang, dan rempah.",
		Gambar:    "https://o-cdf.oramiland.com/unsafe/cnc-magazine.oramiland.com/parenting/original_images/Resep_Ayam_Rica-rica_Sederhana.jpeg",
		Bahan: []string{
			"1 ekor ayam, potong kecil",
			"10 buah cabai merah keriting",
			"5 buah cabai rawit merah",
			"7 siung bawang merah",
			"5 siung bawang putih",
			"2 batang serai, memarkan",
			"5 lembar daun jeruk",
			"2 lembar daun salam",
			"1 ruas jahe, memarkan",
			"Garam, gula, dan minyak secukupnya",
		},
		Instruksi: []string{
			"Haluskan cabai merah, cabai rawit, bawang merah, dan bawang putih.",
			"Tumis bumbu halus dengan serai, daun jeruk, daun salam, dan jahe hingga harum.",
			"Masukkan potongan ayam, aduk hingga berubah warna.",
			"Tambahkan garam dan gula, masak dengan api kecil hingga ayam matang dan bumbu meresap.",
		},
		Tag: []string{"ayam", "rica-rica", "pedas"},
	},
	{
		ID: 10,
		Nama: "Es Cendol",
		Deskripsi: "Cendol hijau dengan santan dan gula merah cair—segar untuk cuaca panas.",
		Gambar: "https://safarfriendly.com/wp-content/uploads/2025/02/640d6eac51567.jpg",
		Bahan: []string{
			"100 gr cendol siap saji",
			"200 ml santan matang",
			"100 gr gula merah, larutkan",
			"Sejumput garam",
			"Es batu secukupnya",
		},
		Instruksi: []string{
			"Larutkan gula merah dengan sedikit air, saring.",
			"Campur santan dengan sejumput garam (sudah dimasak).",
			"Isi gelas dengan cendol dan es batu.",
			"Tuang gula merah dan santan; sajikan.",
		},
		Tag: []string{"minuman", "cendol", "dessert"},
	},
	{
		ID: 11,
		Nama: "Gulai Kambing",
		Deskripsi: "Kuah santan kental dengan daging kambing empuk, rempah khas Minang.",
		Gambar: "https://www.masakapahariini.com/wp-content/uploads/2023/06/shutterstock_2182650345-500x300.jpg",
		Bahan: []string{
			"500 gr daging kambing",
			"600 ml santan",
			"5 siung bawang merah",
			"3 siung bawang putih",
			"2 cm jahe, 2 cm kunyit, 2 cm lengkuas",
			"2 batang serai, daun salam",
			"Garam dan gula secukupnya",
		},
		Instruksi: []string{
			"Tumis bumbu halus hingga wangi.",
			"Masukkan daging, aduk hingga berubah warna.",
			"Tuang santan, tambahkan serai dan daun salam.",
			"Masak kecil hingga daging empuk dan bumbu meresap.",
		},
		Tag: []string{"gulai", "kambing", "minang"},
	},
	{
		ID: 12,
		Nama: "Pecel Madiun",
		Deskripsi: "Sayuran rebus segar disiram bumbu kacang pedas-manis khas Madiun.",
		Gambar: "https://cnc-magazine.oramiland.com/parenting/images/resep-bumbu-pecel-madiun.width-800.format-webp.webp",
		Bahan: []string{
			"100 gr kangkung",
			"100 gr tauge",
			"1 ikat bayam",
			"2 butir telur rebus",
			"Sambal kacang pecel",
			"Kerupuk untuk pelengkap",
		},
		Instruksi: []string{
			"Rebus aneka sayur hingga matang, tiriskan.",
			"Tata sayur dan telur rebus di piring.",
			"Siram dengan sambal kacang.",
			"Sajikan dengan kerupuk.",
		},
		Tag: []string{"pecel", "sayuran", "madiun"},
	},
	{
		ID: 13,
		Nama: "Sate Lilit",
		Deskripsi: "Sate khas Bali: daging cincang berbumbu dililit pada batang serai lalu dibakar.",
		Gambar: "https://jadilaper.com/wp-content/uploads/2023/12/resep-sate-lilit-ayam.jpg",
		Bahan: []string{
			"500 gr daging ayam cincang (atau ikan)",
			"2 batang serai besar (untuk tusuk)",
			"5 siung bawang merah",
			"3 siung bawang putih",
			"2 sdm kelapa parut sangrai",
			"1 sdt garam",
			"1/2 sdt gula",
		},
		Instruksi: []string{
			"Haluskan bumbu; campur dengan daging dan kelapa sangrai.",
			"Ambil adonan, lilitkan pada batang serai.",
			"Bakar hingga matang dan harum, balik sekali-sekali.",
			"Sajikan hangat dengan sambal matah/pencocol.",
		},
		Tag: []string{"sate", "bali", "panggangan"},
	},
	{
		ID: 14,
		Nama: "Klepon",
		Deskripsi: "Kue tradisional berisi gula merah cair, berbalut kelapa parut—kenyal dan manis.",
		Gambar: "https://www.unileverfoodsolutions.co.id/dam/global-ufs/mcos/SEA/calcmenu/recipes/ID-recipes/appetisers/klepon/main-header.jpg",
		Bahan: []string{
			"200 gr tepung ketan",
			"100 gr gula merah, sisir",
			"50 gr kelapa parut",
			"1 sdt pasta pandan",
			"Air hangat secukupnya",
			"Sejumput garam",
		},
		Instruksi: []string{
			"Campur tepung ketan, pasta pandan, garam; tuangi air sedikit-sedikit hingga bisa dipulung.",
			"Ambil adonan, pipihkan, isi gula merah; bulatkan.",
			"Rebus hingga mengapung; angkat dan tiriskan.",
			"Gulingkan di kelapa parut.",
		},
		Tag: []string{"jajanan", "tradisional", "pandan"},
	},
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