# PROSES

## 1. Persona dan masalahnya
Persona Admin Himpunan butuh untuk menyelesaikan masalah pemimjaman ruang himpunan. Peminjaman ini harus disetujui oleh admin terlebih dahulu dan tidak boleh langsung dipakai. Jika permintaan ditolak, harus tertera alasannya kenapa. 

## 2. Rancangan struktur data
struct SumberDaya untuk menyimpan informasi ruang yang dipinjam. isinya ada ID, Nama, Jenis, dan Stok. 
struct Peminjaman untuk menyimpan informasi terkait dengan permintaan peminjaman. isinya ada ID, SumberDayaID, Peminjam, Tanggal, SlotMulai, SlotSelesai, Jumlah, Status, Alasan.
kedua struct dipilih untuk memenuhi fitur kelola sumber daya dan peminjaman. field SumberDayaID menghubungkan data peminjaman dengan sumber daya yang dipinjam. field Status digunakan untuk mencatat status permintaan, sedangkan Alasan disediakan untuk menyimpan alasan penolakan. Data disimpan menggunakan slice di dalam struct App.

## 3. Satu keputusan desain
saya memisahkan handler berdasarkan kelompok kebutuhan, yaitu B1handler.go, B2handler.go, dan B3handler.go, dan seterusnya agar lebih mudah untuk tracing per fiturnya.

## 4. Bug atau kendala (minimal dua)
### Bug 1
- Gejala: Request GET /peminjaman?sumber_daya_id=0 mengembalikan 200 OK 
- Cara menemukan: kirim request GET /peminjaman?sumber_daya_id=0 di powershell
- Penyebab dan solusi: di daftarPeminjaman di B3handler.go, parameter hanya diperiksa menggunakan strconv.Atoi, sehingga angka nol dan negatif tetap lolos. tambahkan validasi id < 1 agar ID harus berupa angka positif sebelum filter dijalankan.

### Bug 2
- Gejala:
- Cara menemukan:
- Penyebab dan solusi:

## 5. Change request
(Apa yang diminta, bagian kode mana yang kamu ubah, dan satu keputusan desain yang kamu ambil.)
