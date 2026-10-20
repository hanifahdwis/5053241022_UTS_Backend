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

	if !teksValid(peminjam, 50) {
		kirimError(w, http.StatusBadRequest, "Peminjam wajib diisi dan maksimal 50 karakter")
		return
	}

	if !tanggalValid(tanggal) {
		kirimError(w, http.StatusBadRequest, "Tanggal harus valid dengan format YYYY-MM-DD")
		return
	}

	if slotMulai < 1 || slotMulai > 8 ||
		slotSelesai < 1 || slotSelesai > 8 ||
		slotMulai > slotSelesai {
		kirimError(w, http.StatusBadRequest, "Slot harus 1 sampai 8 dan slot mulai tidak boleh melebihi slot selesai")
		return
	}

	index := a.cariSumberDaya(sumberDayaID)
	if index == -1 {
		kirimError(w, http.StatusBadRequest, "Sumber daya tidak ditemukan")
		return
	}

	if jumlah < 1 || jumlah > a.sumberDaya[index].Stok {
		kirimError(w, http.StatusBadRequest, "Jumlah harus antara 1 dan stok sumber daya")
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