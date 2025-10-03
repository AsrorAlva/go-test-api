# 🎬 Go REST API - Movies App

Project ini adalah contoh sederhana REST API menggunakan **Go**, **Gin**, **GORM**, dan **PostgreSQL** untuk melakukan operasi CRUD (Create, Read, Update, Delete) pada data `Movie`.

---

## 📚 Daftar Isi
- [🎬 Go REST API - Movies App](#-go-rest-api---movies-app)
  - [📚 Daftar Isi](#-daftar-isi)
  - [🚀 Fitur](#-fitur)
  - [📁 Struktur Folder](#-struktur-folder)
  - [🧪 Requirement](#-requirement)
  - [⚙️ Konfigurasi `.env`](#️-konfigurasi-env)
  - [🧰 Instalasi](#-instalasi)
  - [🌐 Endpoint API](#-endpoint-api)
    - [1. Get All Movies](#1-get-all-movies)
    - [2. Get Movie by ID](#2-get-movie-by-id)
    - [3. Create Movie](#3-create-movie)
    - [4. Update Movie](#4-update-movie)
    - [5. Delete Movie](#5-delete-movie)
  - [📝 Catatan](#-catatan)
  - [📄 Lisensi](#-lisensi)

---

## 🚀 Fitur

- ✅ Create Movie  
- ✅ Read All Movies  
- ✅ Read Movie by ID  
- ✅ Update Movie  
- ✅ Delete Movie  
- 🌐 Terhubung ke PostgreSQL  
- 📦 Menggunakan GORM untuk ORM  
- ⚡ Dibangun dengan Gin Web Framework

---

## 📁 Struktur Folder

```bash
.
├── db
│   └── db.go           # Koneksi DB & fungsi CRUD
├── router
│   └── router.go       # Routing & handler HTTP
├── .env               # Konfigurasi database
├── go.mod
├── go.sum
├── main.go            # Entry point aplikasi
└── README.md
🧪 Requirement

Go 1.20+

PostgreSQL

pgAdmin
 (opsional, untuk GUI)

curl / Postman (untuk testing API)

⚙️ Konfigurasi .env

Buat file .env di root project:

DB_HOST=127.0.0.1
DB_PORT=5432
DB_USER=postgres
DB_NAME=postgres
DB_PASSWORD=yourpassword


Ganti yourpassword dengan password PostgreSQL kamu.

🧰 Instalasi

Clone repository

git clone https://github.com/username/go-test-api.git
cd go-test-api


Inisialisasi module & install dependency

go mod tidy


Jalankan aplikasi

go run main.go


Server akan jalan di:
👉 http://localhost:8080

🌐 Endpoint API
1. Get All Movies
curl -X GET http://localhost:8080/movies

2. Get Movie by ID
curl -X GET http://localhost:8080/movies/{id}


Contoh:

curl -X GET http://localhost:8080/movies/89dc2468-4e65-4175-a94b-7431878428c1

3. Create Movie
curl -X POST http://localhost:8080/movies \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Inception",
    "description": "Dream within a dream"
  }'

4. Update Movie
curl -X PUT http://localhost:8080/movies/{id} \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Inception (Updated)",
    "description": "Amazing"
  }'


Contoh:

curl -X PUT http://localhost:8080/movies/89dc2468-4e65-4175-a94b-7431878428c1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Inception (Updated)",
    "description": "Amazing"
  }'

5. Delete Movie
curl -X DELETE http://localhost:8080/movies/{id}


Contoh:

curl -X DELETE http://localhost:8080/movies/89dc2468-4e65-4175-a94b-7431878428c1

📝 Catatan

ID yang digunakan adalah UUID, bukan angka. Pastikan saat GET/PUT/DELETE, gunakan ID UUID yang benar.

Pastikan PostgreSQL sudah jalan sebelum menjalankan aplikasi.

GORM otomatis membuat tabel movies jika belum ada.

Kamu bisa mengecek isi tabel lewat pgAdmin atau psql dengan query:

SELECT * FROM movies;

📄 Lisensi

MIT License © 2025
Created by [AsrorAlva]