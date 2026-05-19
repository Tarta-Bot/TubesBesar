# Autocare

Aplikasi manajemen bengkel sederhana berbasis **Go** + **SQLite**.

## Fitur

- Manajemen data pelanggan
- Manajemen data kendaraan
- Manajemen layanan/servis
- Penyimpanan data lokal dengan SQLite
- UI terminal/CLI sederhana

## Teknologi

- Go
- SQLite
- Driver SQLite: `modernc.org/sqlite`

## Struktur Project

```text
autocare/
├── main.go
├── go.mod
├── go.sum
├── database/
│   └── database.go
├── models/
├── repository/
├── service/
└── ui/
