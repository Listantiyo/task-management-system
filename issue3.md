# Issue: Milestone 3 - Fitur Utama (Manajemen Tugas / Task CRUD)

## 🎯 Tujuan
Membangun fitur inti dari aplikasi, yaitu manajemen tugas (Task). Pengguna yang telah berhasil *login* dapat membuat, melihat, mengubah, dan menghapus daftar tugas miliknya sendiri. Data tugas ini akan terkait (berelasi) dengan data pengguna (*User*).

## 📝 Tugas yang Perlu Dikerjakan (Task List)

### 1. Definisi Entitas Task (Layer `domain`)
- [ ] Buat struct `Task` (misal: ID, UserID, Title, Description, Status, CreatedAt). *Status* bisa berupa tipe data *string* atau *enum* ("pending", "in-progress", "completed").
- [ ] Buat struktur request/response (DTO) untuk Task (misal: `CreateTaskRequest`, `UpdateTaskRequest`, `TaskResponse`).
- [ ] Definisikan *interface* `TaskRepository` dan `TaskUsecase` yang berisi operasi CRUD (Create, Read, Update, Delete).

### 2. Relasi & Query Database (Layer `repository`)
- [ ] Buat `TaskModel` di layer `models` dengan relasi yang mengarah ke `UserModel` (seorang *User* memiliki banyak *Task* / *One-to-Many*).
- [ ] Auto-migrate tabel `tasks` ke PostgreSQL.
- [ ] Implementasikan `TaskRepository` dengan GORM untuk:
  - `CreateTask`: Menyimpan tugas baru.
  - `GetTasksByUserID`: Mengambil semua tugas milik *User* tertentu.
  - `GetTaskByID`: Mengambil detail satu tugas.
  - `UpdateTask`: Memperbarui data tugas.
  - `DeleteTask`: Menghapus data tugas.

### 3. Logika Bisnis & Otorisasi (Layer `usecase`)
- [ ] Buat implementasi `TaskUsecase`.
- [ ] **Penting**: Pada fungsi Update dan Delete, validasi bahwa `Task` yang akan diubah/dihapus benar-benar milik `UserID` yang sedang *login* (berdasarkan ID dari JWT). Jangan sampai *user* lain bisa mengubah tugas yang bukan miliknya (mencegah *IDOR vulnerability*).

### 4. Controller HTTP & Routing (Layer `delivery`)
Buat *handler* dan *router* dengan framework `gin` untuk rute berikut, yang semuanya **wajib** dilindungi oleh `AuthMiddleware`:
- [ ] `POST /api/tasks` : Membuat tugas baru.
- [ ] `GET /api/tasks` : Menampilkan seluruh tugas milik user yang sedang *login*.
- [ ] `GET /api/tasks/:id` : Menampilkan detail satu tugas.
- [ ] `PUT /api/tasks/:id` : Memperbarui tugas.
- [ ] `DELETE /api/tasks/:id` : Menghapus tugas.
- [ ] Jangan lupa tambahkan anotasi **Swagger** untuk masing-masing fungsi *handler* di atas (wajib ada `@Security BearerAuth` agar JWT bisa di-*test* lewat Swagger UI).

---
## 💡 Catatan Penting
- Anda dapat mengambil `UserID` (*user* yang sedang *login*) dari `gin.Context` yang sebelumnya di-set oleh `AuthMiddleware`.
- Fitur ini adalah inti dari aplikasi. Pastikan *error handling*-nya tertangani dengan baik (contoh: *return 404* jika *Task* tidak ditemukan).
- Selamat mengerjakan Milestone penutup untuk Backend ini! 🚀
