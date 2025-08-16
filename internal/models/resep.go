package models

type Resep struct {
	ID        int
	Nama      string
	Deskripsi string
	Gambar    string
	Bahan     []string
	Instruksi []string
	Tag       []string
}
