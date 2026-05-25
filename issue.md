# Issue: Milestone 1 - Fondasi Backend & Dokumentasi API

## 🎯 Tujuan
Membangun fondasi awal untuk server backend menggunakan Golang, menghubungkannya dengan database PostgreSQL, dan menyiapkan kerangka dokumentasi API menggunakan Swagger.

## 📝 Tugas yang Perlu Dikerjakan (Task List)

### 1. Inisialisasi Proyek Go
- [ ] Jalankan perintah `go mod init <nama-project>` di terminal untuk membuat file `go.mod`.
- [ ] Buat file `main.go` sebagai titik awal (entry point) berjalannya aplikasi.

### 2. Persiapan Struktur Folder (Clean Architecture)
Buat folder-folder berikut untuk merapikan kode:
- [ ] `domain`: Tempat menyimpan struktur data (*struct*) inti.
- [ ] `repository`: Tempat kode untuk interaksi langsung ke database (PostgreSQL).
- [ ] `usecase`: Tempat meletakkan logika bisnis aplikasi.
- [ ] `delivery`: Tempat mengatur rute API (HTTP Request/Response).

### 3. Setup Koneksi Database PostgreSQL
- [ ] Install library database (bisa menggunakan **GORM** atau **SQLX**) dan driver PostgreSQL.
- [ ] Buat kode untuk melakukan koneksi ke database.
- [ ] Panggil fungsi koneksi tersebut di `main.go` dan pastikan tampil pesan sukses (atau tidak ada error) saat server dinyalakan.

### 4. Setup HTTP Server & Routing
- [ ] Install framework router (misalnya: **Gin Gonic** atau **Go-Chi**).
- [ ] Buat satu endpoint percobaan (contoh: `GET /ping` yang membalas `{"message": "pong"}`) di `main.go` untuk mengecek apakah server merespon dengan baik.

### 5. Setup Dokumentasi Swagger API
- [ ] Install library `swaggo/swag` untuk Golang.
- [ ] Tambahkan komentar dasar Swagger di atas fungsi `main()` pada `main.go` (berisi info judul API, deskripsi, versi).
- [ ] Jalankan perintah `swag init` di terminal untuk menghasilkan file dokumentasi.
- [ ] Tambahkan *route* khusus agar halaman dokumentasi (Swagger UI) bisa diakses langsung via browser.

---
## 💡 Catatan Penting
- **Jangan memikirkan fitur spesifik dulu** (seperti Login atau Tambah Tugas). Fokus kita sekarang murni persiapan alat dan kerangka kerja!
- Pastikan di akhir *issue* ini: server Go bisa dijalankan, koneksi ke PostgreSQL berhasil, dan halaman Swagger UI bisa terbuka di browser tanpa kendala.
- Kerjakan langkah demi langkah secara berurutan. Semangat! 🚀
