# Cinema Ticket API

Backend REST API untuk sistem pemesanan tiket bioskop online yang dirancang untuk menangani proses pengelolaan film, cinema, studio, kursi, jadwal tayang, autentikasi pengguna, serta pemesanan tiket.

Project ini dibuat menggunakan **Golang** dengan PostgreSQL sebagai database utama dan Redis untuk kebutuhan caching, temporary seat locking, serta mendukung proses dengan concurrent request.

## Features

- User registration & login
- JWT authentication
- Role-based access control
- Email verification
- Cinema management
- Studio management
- Automatic seat generation berdasarkan layout studio
- Movie management
- Showtime management
- PostgreSQL database
- Redis integration
- Swagger API documentation
- RESTful API
- Database migration
- Protection terhadap double booking melalui seat locking dan database transaction

## Tech Stack

- **Golang 1.25.5**
- **PostgreSQL**
- **Redis 7**
- **JWT**
- **Gorilla Mux**
- **SQLX**
- **PGX**
- **Swagger**
- **Docker**
- **golang-migrate**

## Project Structure

```text
cinema-ticket/
├── cmd/
│   └── main.go
├── config/
├── internal/
│   ├── dto/
│   ├── handler/
│   ├── middleware/
│   ├── model/
│   ├── repository/
│   └── service/
├── migrations/
├── pkg/
│   ├── database/
│   ├── jwt/
│   └── response/
├── docs/
├── .env
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## Requirements

Pastikan environment berikut sudah tersedia:

- Go `1.25.5`
- PostgreSQL
- Docker
- Git
- golang-migrate

Redis dapat dijalankan menggunakan Docker.

---

# Installation

## 1. Clone Repository

Clone repository menggunakan Git:

```bash
git clone https://github.com/Adhityaekp/cinema-ticket.git
```

Masuk ke directory project:

```bash
cd cinema-ticket
```

Kemudian download dependency:

```bash
go mod download
```

Atau:

```bash
go mod tidy
```

## 2. Setup PostgreSQL

Buat database PostgreSQL dengan nama:

```text
cinema_ticket
```

Contoh menggunakan PostgreSQL:

```sql
CREATE DATABASE cinema_ticket;
```

Pastikan PostgreSQL berjalan pada:

```text
Host     : localhost
Port     : 5432
Username : postgres
Password : postgres
Database : cinema_ticket
```

Jika konfigurasi PostgreSQL berbeda, sesuaikan dengan file `.env`.

## 3. Setup Redis

Redis digunakan untuk temporary seat locking, caching, dan kebutuhan lainnya.

Jalankan Redis menggunakan Docker:

```bash
docker run -d \
  --name cinema-redis \
  -p 6379:6379 \
  redis:7
```

Untuk mengecek container:

```bash
docker ps
```

Jika container sudah pernah dibuat dan hanya dalam kondisi stopped:

```bash
docker start cinema-redis
```

Test Redis:

```bash
docker exec -it cinema-redis redis-cli ping
```

Jika berhasil akan menghasilkan:

```text
PONG
```

## 4. Setup Environment

Buat file `.env` di root project:

```env
APP_NAME=cinema-ticket
APP_ENV=local
APP_PORT=8080
APP_BASE_URL=http://localhost:8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=cinema_ticket
DB_SSLMODE=disable

REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

JWT_SECRET=change-this-secret-key
JWT_EXPIRE_HOURS=24
JWT_ACCESS_EXPIRE_MINUTES=15
JWT_REFRESH_EXPIRE_DAYS=7

COOKIE_SECURE=false
COOKIE_DOMAIN=

SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=
SMTP_PASSWORD=
SMTP_FROM=
```

> Untuk production, jangan menggunakan `JWT_SECRET` default dan jangan menyimpan credential SMTP atau password sensitif di repository.

## 5. Run Database Migration

Pastikan PostgreSQL sudah berjalan.

Kemudian jalankan migration:

```bash
migrate -path migrations \
  -database "postgres://postgres:postgres@localhost:5432/cinema_ticket?sslmode=disable" \
  up
```

Migration akan membuat tabel yang dibutuhkan oleh aplikasi dan memasukkan data dummy/master yang telah disediakan.

Untuk mengecek migration:

```bash
migrate -path migrations \
  -database "postgres://postgres:postgres@localhost:5432/cinema_ticket?sslmode=disable" \
  version
```

## 6. Run Application

Jalankan aplikasi:

```bash
go run ./cmd
```

Jika berhasil, API dapat diakses melalui:

```text
http://localhost:8080
```

## 7. Swagger Documentation

API documentation tersedia menggunakan Swagger.

Buka:

```text
http://localhost:8080/swagger/index.html
```

Swagger dapat digunakan untuk melihat endpoint, request body, response, serta mencoba API secara langsung.

## Authentication

API menggunakan JWT untuk authentication.

Access token disimpan menggunakan **HttpOnly Cookie** sehingga client tidak perlu mengirimkan token secara manual pada setiap request.

Beberapa endpoint membutuhkan role tertentu.

### Customer

Customer dapat mengakses endpoint yang bersifat public dan endpoint yang berkaitan dengan proses pemesanan tiket.

### Admin

Admin memiliki akses untuk melakukan pengelolaan:

- Cinema
- Studio
- Movie
- Showtime
- Data terkait sistem lainnya

## Dummy Account

Untuk kebutuhan testing tersedia akun dummy:

### Admin

```text
Email    : admin@cinematicket.com
Password : Admin123!
Role     : ADMIN
```

### Customer

```text
Email    : user@cinematicket.com
Password : User123!
Role     : CUSTOMER
```

## API Endpoint

Beberapa endpoint utama:

### Authentication

```text
POST /api/auth/register
POST /api/auth/login
POST /api/auth/refresh
POST /api/auth/logout
```

### Cinema

```text
GET    /api/cinemas
GET    /api/cinemas/{id}
POST   /api/cinemas
PUT    /api/cinemas/{id}
DELETE /api/cinemas/{id}
```

### Studio

```text
GET    /api/studios
GET    /api/studios/{id}
POST   /api/studios
PUT    /api/studios/{id}
DELETE /api/studios/{id}
```

Saat membuat studio, sistem akan otomatis membuat kursi berdasarkan jumlah baris dan jumlah kursi per baris.

Contoh:

```json
{
  "cinema_id": "11111111-1111-1111-1111-111111111111",
  "name": "Studio 3",
  "rows": 10,
  "seats_per_row": 12
}
```

Request tersebut akan menghasilkan 120 kursi dengan format:

```text
A1 - A12
B1 - B12
...
J1 - J12
```

### Seats

```text
GET /api/seats
GET /api/seats/{id}
GET /api/seats?studio_id={studio_id}
```

Seat dibuat otomatis ketika studio dibuat sehingga tidak perlu membuat kursi satu per satu.

### Movies

```text
GET    /api/movies
GET    /api/movies/{id}
POST   /api/movies
PUT    /api/movies/{id}
DELETE /api/movies/{id}
```

### Showtimes

```text
GET    /api/showtimes
GET    /api/showtimes/{id}
POST   /api/showtimes
PUT    /api/showtimes/{id}
DELETE /api/showtimes/{id}
```

Showtime menghubungkan:

```text
Movie
  ↓
Cinema
  ↓
Studio
  ↓
Showtime
  ↓
Showtime Seats
```

Setiap showtime memiliki daftar kursi sendiri sehingga status kursi untuk satu jadwal tidak memengaruhi jadwal lainnya.

## Seat Locking

Salah satu masalah utama dalam sistem pemesanan tiket bioskop adalah **double booking**, yaitu dua customer mencoba memilih kursi yang sama pada waktu yang hampir bersamaan.

Untuk mengatasi hal tersebut, sistem menggunakan Redis sebagai temporary seat lock.

Contoh key:

```text
seat:lock:showtime:123:A10
```

Seat akan di-lock sementara dengan TTL tertentu.

Konsepnya:

```text
Customer memilih kursi
        ↓
Redis SET NX
        ↓
Berhasil?
   ↓          ↓
  YES         NO
   ↓           ↓
Lock seat    Seat sudah
sementara    dipilih user lain
   ↓
Customer melakukan pembayaran
```

Database PostgreSQL tetap menjadi **source of truth** untuk data transaksi dan status kursi.

## Booking Flow

Secara umum proses pemesanan:

```text
Customer
   ↓
Pilih Movie
   ↓
Pilih Cinema
   ↓
Pilih Showtime
   ↓
Pilih Seat
   ↓
Redis Seat Lock
   ↓
Create Booking
   ↓
Payment
   ↓
Payment Success
   ↓
Booking PAID
   ↓
E-Ticket
```

Jika customer tidak menyelesaikan pembayaran sampai batas waktu:

```text
PENDING
   ↓
Expired
   ↓
Seat kembali AVAILABLE
```

## Showtime Cancellation & Refund

Jika cinema membatalkan suatu showtime, sistem akan memproses booking yang sudah dibayar.

Flow:

```text
Admin Cancel Showtime
        ↓
Showtime = CANCELLED
        ↓
Cari booking PAID
        ↓
Create Refund
        ↓
Refund berhasil?
     ↓       ↓
    YES      NO
     ↓        ↓
REFUNDED   REFUND_PENDING
     ↓
Release Seat
     ↓
Notify Customer
```

Data booking tidak dihapus sehingga riwayat transaksi tetap dapat disimpan dan dilacak.

## Database

Database utama menggunakan PostgreSQL.

Beberapa tabel utama:

```text
users
cinemas
studios
seats
movies
showtimes
showtime_seats
bookings
booking_items
payments
refunds
```

Relasi utama:

```text
Cinema
  └── Studio
        └── Seat

Movie
  └── Showtime
        ├── Cinema
        ├── Studio
        └── Showtime Seat

User
  └── Booking
        ├── Booking Item
        ├── Payment
        └── Refund
```

## Development

Untuk menjalankan aplikasi dalam mode development:

```bash
go run ./cmd
```

Untuk melakukan pengecekan build:

```bash
go build ./...
```

Untuk menjalankan test:

```bash
go test ./...
```

Untuk generate ulang Swagger:

```bash
swag init -g cmd/main.go
```

## Git Clone – Quick Start

Jika ingin menjalankan project dari awal:

```bash
git clone https://github.com/Adhityaekp/cinema-ticket.git

cd cinema-ticket

go mod download

# Buat database PostgreSQL
# Jalankan Redis

# Buat file .env

# Jalankan migration
migrate -path migrations \
  -database "postgres://postgres:postgres@localhost:5432/cinema_ticket?sslmode=disable" \
  up

# Jalankan aplikasi
go run ./cmd
```

Kemudian akses:

```text
API      : http://localhost:8080
Swagger  : http://localhost:8080/swagger/index.html
```

## Notes

Project ini dibuat sebagai implementasi backend untuk sistem pemesanan tiket bioskop dengan fokus pada **REST API, database transaction, authentication & authorization, concurrency handling, seat locking, serta maintainable backend architecture**.
