package main

import (
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func (a *App) daftarPeminjaman(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	data := []Peminjaman{}

	var sumberDayaID *int

	if nilai := query.Get("sumber_daya_id"); nilai != "" {
		id, err := strconv.Atoi(nilai)
		if err != nil || id < 1 {
			kirimError(
				w,
				http.StatusBadRequest,
				"sumber_daya_id harus berupa angka positif",
			)
			return
		}

		sumberDayaID = &id
	}

	tanggal := strings.TrimSpace(query.Get("tanggal"))
	status := strings.TrimSpace(query.Get("status"))

	if tanggal != "" && !tanggalValid(tanggal) {
		kirimError(
			w,
			http.StatusBadRequest,
			"Tanggal filter harus valid dengan format YYYY-MM-DD",
		)
		return
	}

	for _, p := range a.peminjaman {
		if sumberDayaID != nil && p.SumberDayaID != *sumberDayaID {
			continue
		}

		if tanggal != "" && p.Tanggal != tanggal {
			continue
		}

		if status != "" && p.Status != status {
			continue
		}

		data = append(data, p)
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].ID < data[j].ID
	})

	kirimJSON(w, http.StatusOK, data)
}

func (a *App) detailPeminjaman(w http.ResponseWriter, r *http.Request) {
	id, err := ambilID(r)
	if err != nil {
		kirimError(
			w,
			http.StatusBadRequest,
			"ID harus berupa angka positif",
		)
		return
	}

	index := a.cariPeminjaman(id)
	if index == -1 {
		kirimError(
			w,
			http.StatusNotFound,
			"Peminjaman tidak ditemukan",
		)
		return
	}

	kirimJSON(w, http.StatusOK, a.peminjaman[index])
}