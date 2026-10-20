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