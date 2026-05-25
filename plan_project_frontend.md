# 🗺️ Master Plan Proyek: Sisi Frontend (React Client)

Dokumen ini berfokus pada arsitektur tingkat tinggi untuk komponen Frontend. Aplikasi ini bertindak sebagai antarmuka pengguna (*User Interface*) berbasis Single Page Application (SPA) yang akan berinteraksi langsung dengan Server Backend berbasis Golang. Pengembangan frontend akan dimulai setelah fondasi API Backend selesai.

---

## 🔗 1. Keterkaitan dengan Backend (Golang)
Aplikasi React ini bergantung penuh pada ketersediaan API dari Backend. Aturan integrasinya adalah:
*   **Konsumsi Data**: Menggunakan pustaka HTTP Client (seperti Axios atau Fetch API) untuk menembak endpoint JSON yang disediakan oleh Backend.
*   **Manajemen Sesi**: Setelah pengguna sukses Login, React akan menerima Token JWT dari Backend. React bertanggung jawab menyimpan token ini dengan aman (misal di *local storage* atau *secure cookie*) dan menyisipkannya ke setiap HTTP Header saat meminta data sensitif.
*   **Panduan Integrasi**: Pengembang Frontend wajib merujuk pada dokumentasi Swagger UI yang disediakan oleh Backend untuk mengetahui skema data yang valid.

---

## 🛠️ 2. Teknologi & Pustaka Utama
*   **Library Utama**: React.js 
*   **State Management**: React Context API atau Redux Toolkit (Untuk menyimpan status login user dan daftar tugas secara global).
*   **Routing**: React Router (Untuk perpindahan halaman tanpa *reload*).
*   **HTTP Client**: Axios (Untuk komunikasi data dengan Backend).
*   **Styling**: Tailwind CSS (Untuk antarmuka yang cepat dan responsif).

---

## 📐 3. Komponen & Alur Halaman Utama
Aplikasi Frontend akan dibagi menjadi beberapa area komponen utama:
1.  **Halaman Publik**: Form Login dan Form Registrasi User.
2.  **Halaman Privat (Protected Route)**: Dashboard Utama, Komponen List Task, Form Tambah Task, dan Tombol Aksi (Edit/Delete). Halaman ini terkunci jika tidak ada Token JWT aktif.
3.  **Interseptor HTTP**: Fungsi otomatis yang bertugas menangkap *error 401 (Unauthorized)* dari Backend untuk otomatis me-lempar (*redirect*) user kembali ke halaman Login jika token kedaluwarsa.

---

## 🏁 4. Roadmap Pengembangan Frontend
*   **Milestone 1: Setup Proyek & UI Kit** (Inisialisasi React, install Tailwind, dan routing dasar).
*   **Milestone 2: Autentikasi & Sesi** (Pembuatan halaman Login/Register dan integrasi JWT dengan API Backend).
*   **Milestone 3: Manajemen Tugas (Dashboard)** (Integrasi CRUD Task secara visual dan sinkronisasi data dengan database melalui API Backend).
