package main

import (
	"net/http"
	"strings"
)

func (a *App) buatPeminjaman(w http.ResponseWriter, r *http.Request) {
	var input struct {
		SumberDayaID *int    `json:"sumber_daya_id"`
		Peminjam     *string `json:"peminjam"`
		Tanggal      *string `json:"tanggal"`
		SlotMulai    *int    `json:"slot_mulai"`
		SlotSelesai  *int    `json:"slot_selesai"`
		Jumlah       *int    `json:"jumlah"`
	}

	if err := bacaJSON(r, &input); err != nil {
		kirimError(w, http.StatusBadRequest, "Body JSON tidak valid")
		return
	}

	if input.SumberDayaID == nil ||
		input.Peminjam == nil ||
		input.Tanggal == nil ||
		input.SlotMulai == nil ||
		input.SlotSelesai == nil ||
		input.Jumlah == nil {
		kirimError(w, http.StatusBadRequest, "Semua enam field peminjaman wajib diisi")
		return
	}

	sumberDayaID := *input.SumberDayaID
	peminjam := strings.TrimSpace(*input.Peminjam)
	tanggal := strings.TrimSpace(*input.Tanggal)
	slotMulai := *input.SlotMulai
	slotSelesai := *input.SlotSelesai
	jumlah := *input.Jumlah

	pesan := a.validasiPeminjaman(
		sumberDayaID,
		peminjam,
		tanggal,
		slotMulai,
		slotSelesai,
		jumlah,
	)

	if pesan != "" {
		kirimError(w, http.StatusBadRequest, pesan)
		return
	}

	if a.adaBentrok(
		sumberDayaID,
		tanggal,
		slotMulai,
		slotSelesai,
		jumlah,
		0,
	) {
		kirimError(
			w,
			http.StatusConflict,
			"Kapasitas sumber daya tidak mencukupi pada slot yang dipilih",
		)
		return
	}

	data := Peminjaman{
		ID:           a.nextPeminjamanID,
		SumberDayaID: sumberDayaID,
		Peminjam:     peminjam,
		Tanggal:      tanggal,
		SlotMulai:    slotMulai,
		SlotSelesai:  slotSelesai,
		Jumlah:       jumlah,
		Status:       "diajukan",
	}

	a.nextPeminjamanID++
	a.peminjaman = append(a.peminjaman, data)

	kirimJSON(w, http.StatusCreated, data)
}