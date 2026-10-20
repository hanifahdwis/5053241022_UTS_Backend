package main

import "net/http"

func NewHandler() http.Handler {
	app := &App{
		sumberDaya:       []SumberDaya{},
		peminjaman:       []Peminjaman{},
		nextSumberDayaID: 1,
		nextPeminjamanID: 1,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /sumber-daya", app.buatSumberDaya)
	mux.HandleFunc("GET /sumber-daya", app.daftarSumberDaya)
	mux.HandleFunc("GET /sumber-daya/{id}", app.detailSumberDaya)

	mux.HandleFunc("POST /peminjaman", app.buatPeminjaman)
	mux.HandleFunc("GET /peminjaman", app.daftarPeminjaman)
	mux.HandleFunc("GET /peminjaman/{id}", app.detailPeminjaman)
	mux.HandleFunc("PUT /peminjaman/{id}", app.ubahPeminjaman)
	mux.HandleFunc("DELETE /peminjaman/{id}", app.hapusPeminjaman)

	return mux
}