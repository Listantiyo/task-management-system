# Issue: Milestone 2 - Akses & Keamanan (Register, Login, & JWT)

## 🎯 Tujuan
Membangun fitur otentikasi agar aplikasi menjadi aman. Sistem akan memiliki kapabilitas untuk mendaftarkan pengguna baru (Register), masuk (Login), mengenkripsi password dengan *hash*, serta mengamankan API menggunakan token JWT (JSON Web Token).

## 📝 Tugas yang Perlu Dikerjakan (Task List)

### 1. Definisi Entitas & Struktur Data (Layer `domain`)
- [ ] Buat struct `User` yang merepresentasikan tabel di database (misal: ID, Email, Password, CreatedAt).
- [ ] Buat struct untuk Payload/Request dari Frontend (misal: `RegisterRequest`, `LoginRequest`).
- [ ] Buat *interface* untuk UserRepository dan UserUsecase sebagai kontrak fungsi yang dibutuhkan.

### 2. Koneksi & Query Database (Layer `repository`)
- [ ] Lakukan auto-migrasi tabel `users` di PostgreSQL (bisa dilakukan saat *startup* aplikasi menggunakan fungsi AutoMigrate dari GORM).
- [ ] Buat fungsi `CreateUser` untuk menyimpan data pengguna baru ke database.
- [ ] Buat fungsi `GetUserByEmail` untuk mengambil data pengguna (diperlukan saat proses login dan pengecekan duplikasi email).

### 3. Logika Bisnis & Enkripsi (Layer `usecase`)
- [ ] Install library untuk enkripsi password (misalnya: `golang.org/x/crypto/bcrypt`).
- [ ] Buat fungsi *Register*: Validasi apakah email sudah terdaftar, lalu *hash* password, sebelum mengirimnya ke *repository*.
- [ ] Buat fungsi *Login*: Cek apakah email ada di database, bandingkan password *hash* dengan password yang dikirim (*compare bcrypt*).
- [ ] Install library JWT (misalnya: `github.com/golang-jwt/jwt/v5`).
- [ ] Jika *Login* sukses, *generate* token JWT dan kembalikan ke layer *delivery*.

### 4. Controller HTTP REST API (Layer `delivery`)
- [ ] Buat endpoint `POST /api/register` untuk menangani proses registrasi.
- [ ] Buat endpoint `POST /api/login` untuk menangani proses login dan membalas dengan JSON berisi token JWT.
- [ ] Pastikan kedua endpoint ini terdokumentasi dengan Swagger (tambahkan anotasi komentar di atas fungsinya).

### 5. Middleware JWT
- [ ] Buat *middleware* fungsi (misal: `AuthMiddleware`) di router.
- [ ] *Middleware* bertugas membaca HTTP Header `Authorization: Bearer <token>`, memvalidasi token JWT tersebut, dan meneruskan proses atau memblokirnya (Return 401 Unauthorized) jika token tidak valid/kedaluwarsa.
- [ ] (Opsional/Testing) Buat satu endpoint percobaan misalnya `GET /api/me` yang dilindungi *middleware* ini untuk memastikan *middleware* berfungsi dengan benar.

---
## 💡 Catatan Penting
- Jangan pernah menyimpan password dalam bentuk *plain text* di database! Pastikan selalu di-hash menggunakan `bcrypt` terlebih dahulu.
- Saat melakukan *testing* (baik lewat Postman atau Swagger), simpan Token JWT hasil *login* untuk digunakan mengakses *endpoint* yang dilindungi (*protected route*).
