# AI-LOG

Catat setiap penggunaan AI yang menghasilkan kode. Jika tidak memakai AI sama sekali, tulis satu baris: **tidak memakai AI**.

Log yang jujur tidak mengurangi nilai, sebanyak apa pun AI dipakai. Yang mengurangi nilai adalah log yang tidak sesuai dengan kode, atau isi log yang tidak bisa dijelaskan.

## (Validasi tanggal, B2)
- Prompt: time.Parse saja cukup untuk validasi tanggal dengan format YYYY-MM-DD, termasuk tanggal yang tidak valid?
- Dipakai: fungsi tanggalValid pada helper.go.
- Diubah: menambahkan pemeriksaan hasil format ulang tanggal agar input harus sesuai format dan tanggal kalender yang valid.

## (Penerapan filter, B3)
- Prompt: bagaimana cara penerapan filter di daftarPeminjaman?
- Dipakai: fungsi daftarPeminjaman pada file B3handler.go
- Diubah: menerapkan filter sumber_daya_id, tanggal, dan status secara berurutan
