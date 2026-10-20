package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func kirimJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func kirimError(w http.ResponseWriter, status int, pesan string) {
	kirimJSON(w, status, map[string]string{
		"error": pesan,
	})
}

func bacaJSON(r *http.Request, tujuan any) error {
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(tujuan); err != nil {
		return err
	}

	var tambahan any
	if err := decoder.Decode(&tambahan); err != io.EOF {
		return strconv.ErrSyntax
	}

	return nil
}

func ambilID(r *http.Request) (int, error) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		return 0, strconv.ErrSyntax
	}

	return id, nil
}

func tanggalValid(tanggal string) bool {
	t, err := time.Parse("2006-01-02", tanggal)
	return err == nil && t.Format("2006-01-02") == tanggal
}

func teksValid(teks string, maksimal int) bool {
	teks = strings.TrimSpace(teks)
	return teks != "" && utf8.RuneCountInString(teks) <= maksimal
}

func (a *App) cariSumberDaya(id int) int {
	for i := range a.sumberDaya {
		if a.sumberDaya[i].ID == id {
			return i
		}
	}
	return -1
}

func (a *App) cariPeminjaman(id int) int {
	for i := range a.peminjaman {
		if a.peminjaman[i].ID == id {
			return i
		}
	}
	return -1
}

func (a *App) validasiPeminjaman(
	sumberDayaID int,
	peminjam string,
	tanggal string,
	slotMulai int,
	slotSelesai int,
	jumlah int,
) string {
	peminjam = strings.TrimSpace(peminjam)
	tanggal = strings.TrimSpace(tanggal)

	if !teksValid(peminjam, 50) {
		return "Peminjam wajib diisi dan maksimal 50 karakter"
	}

	if !tanggalValid(tanggal) {
		return "Tanggal harus valid dengan format YYYY-MM-DD"
	}

	if slotMulai < 1 || slotMulai > 8 ||
		slotSelesai < 1 || slotSelesai > 8 ||
		slotMulai > slotSelesai {
		return "Slot harus 1 sampai 8 dan slot mulai tidak boleh melebihi slot selesai"
	}

	index := a.cariSumberDaya(sumberDayaID)
	if index == -1 {
		return "Sumber daya tidak ditemukan"
	}

	if jumlah < 1 || jumlah > a.sumberDaya[index].Stok {
		return "Jumlah harus antara 1 dan stok sumber daya"
	}

	return ""
}

func (a *App) adaBentrok(
	sumberDayaID int,
	tanggal string,
	slotMulai int,
	slotSelesai int,
	jumlah int,
	abaikanID int,
) bool {
	index := a.cariSumberDaya(sumberDayaID)
	if index == -1 {
		return true
	}

	stok := a.sumberDaya[index].Stok

	for slot := slotMulai; slot <= slotSelesai; slot++ {
		totalTerpakai := 0

		for _, p := range a.peminjaman {
			if p.ID == abaikanID {
				continue
			}

			if p.SumberDayaID != sumberDayaID ||
				p.Tanggal != tanggal {
				continue
			}

			if p.Status != "diajukan" &&
				p.Status != "disetujui" {
				continue
			}

			if slot >= p.SlotMulai && slot <= p.SlotSelesai {
				totalTerpakai += p.Jumlah
			}
		}

		if totalTerpakai+jumlah > stok {
			return true
		}
	}

	return false
}