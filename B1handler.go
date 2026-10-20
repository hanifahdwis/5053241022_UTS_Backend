package main

import (
	"net/http"
	"sort"
	"strings"
)

func (a *App) buatSumberDaya(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Nama  *string `json:"nama"`
		Jenis *string `json:"jenis"`
		Stok  *int    `json:"stok"`
	}

	if err := bacaJSON(r, &input); err != nil {
		kirimError(w, http.StatusBadRequest, "Body JSON tidak valid")
		return
	}

	if input.Nama == nil || input.Jenis == nil || input.Stok == nil {
		kirimError(w, http.StatusBadRequest, "Field nama, jenis, dan stok wajib diisi")
		return
	}

	nama := strings.TrimSpace(*input.Nama)
	jenis := strings.TrimSpace(*input.Jenis)
	stok := *input.Stok

	if !teksValid(nama, 50) {
		kirimError(w, http.StatusBadRequest, "Nama wajib diisi dan maksimal 50 karakter")
		return
	}

	if jenis != "ruang" && jenis != "alat" {
		kirimError(w, http.StatusBadRequest, "Jenis harus ruang atau alat")
		return
	}

	if stok < 1 {
		kirimError(w, http.StatusBadRequest, "Stok minimal 1")
		return
	}

	data := SumberDaya{
		ID:    a.nextSumberDayaID,
		Nama:  nama,
		Jenis: jenis,
		Stok:  stok,
	}

	a.nextSumberDayaID++
	a.sumberDaya = append(a.sumberDaya, data)

	kirimJSON(w, http.StatusCreated, data)
}

func (a *App) daftarSumberDaya(w http.ResponseWriter, r *http.Request) {
	data := append([]SumberDaya{}, a.sumberDaya...)

	sort.Slice(data, func(i, j int) bool {
		return data[i].ID < data[j].ID
	})

	kirimJSON(w, http.StatusOK, data)
}

func (a *App) detailSumberDaya(w http.ResponseWriter, r *http.Request) {
	id, err := ambilID(r)
	if err != nil {
		kirimError(w, http.StatusBadRequest, "ID harus berupa angka positif")
		return
	}

	index := a.cariSumberDaya(id)
	if index == -1 {
		kirimError(w, http.StatusNotFound, "Sumber daya tidak ditemukan")
		return
	}

	kirimJSON(w, http.StatusOK, a.sumberDaya[index])
}