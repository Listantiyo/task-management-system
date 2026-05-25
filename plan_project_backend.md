# 🗺️ Master Plan Proyek: Sisi Backend (Golang API)

Dokumen ini berfokus pada arsitektur tingkat tinggi untuk komponen Backend. Server ini bertindak sebagai penyedia data tunggal (*Single Source of Truth*) yang akan dikonsumsi oleh aplikasi klien (Frontend React). Fokus utama fase awal adalah menyelesaikan seluruh kontrak API backend sebelum frontend mulai dibangun.

---

## 🔗 1. Keterkaitan dengan Frontend (React)
Backend ini dirancang khusus untuk mendukung aplikasi Frontend React yang bersifat *Stateless*. Artinya:
*   **Komunikasi**: Menggunakan HTTP REST API dengan format pertukaran data JSON.
*   **Autentikasi**: Menggunakan JWT (JSON Web Token). Backend tidak menyimpan sesi user; Frontend wajib mengirimkan Token JWT ini pada HTTP Header `Authorization: Bearer <token>` di setiap *request* yang membutuhkan hak akses.
*   **Dokumentasi**: Wajib mengintegrasikan Swagger/OpenAPI sebagai jembatan agar tim Frontend tahu persis struktur data, URL, dan *status code* yang dihasilkan oleh Backend.

---

## 🛠️ 2. Teknologi & Pustaka Utama
*   **Bahasa Pemrograman**: Golang (Go)
*   **Database**: PostgreSQL (Relasional, digunakan untuk relasi data User dan Task).
*   **HTTP Framework**: Go-Chi atau Gin Gonic (Untuk manajemen routing).
*   **ORM / Database Toolkit**: GORM atau SQLX (Pemetaan data objek Go ke PostgreSQL).
*   **Dokumentasi**: Swagger Go (Mengotomatiskan pembuatan file `swagger.json`).

---

## 📐 3. Struktur Kode (Clean Architecture)
Kode Backend dipisahkan menjadi 4 lapisan utama untuk mempermudah AI Agent melakukan *Code Review* dan *Unit Testing*:
1.  **Layer Domain**: Definisi struktur data (*struct*) inti dan kontrak interface.
2.  **Layer Repository**: Operasi CRUD langsung ke database PostgreSQL menggunakan ORM.
3.  **Layer Usecase**: Tempat validasi aturan sistem dan seluruh logika bisnis.
4.  **Layer Delivery/HTTP**: Membaca JSON request dari Frontend, meneruskan ke Usecase, dan mengembalikan HTTP Response JSON beserta status code (200, 400, 401, 500, dll).

---

## 🏁 4. Roadmap Pengembangan Backend
*   **Milestone 1: Fondasi & Dokumentasi API** (Setup folder, koneksi PostgreSQL, dan Swagger UI).
*   **Milestone 2: Akses & Keamanan** (Fitur Register, Login, Enkripsi Password, dan Auth Middleware JWT).
*   **Milestone 3: Fitur Utama (Manajemen Tugas)** (CRUD Task dengan validasi kepemilikan data user).
