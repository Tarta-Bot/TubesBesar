# AutoCare - Aplikasi Manajemen dan Riwayat Servis Kendaraan

Tugas Besar Mata Kuliah Algoritma Pemrograman 2

## Struktur Proyek

```
autocare/
├── main.go                      # Entry point utama program
├── go.mod                       # Dependensi Go module
├── models/
│   └── models.go                # Struct data: Kendaraan, Pemilik, RiwayatServis
├── database/
│   └── database.go              # Koneksi & inisialisasi SQLite
├── repository/
│   ├── kendaraan_repo.go        # CRUD operasi kendaraan
│   ├── pemilik_repo.go          # CRUD operasi pemilik
│   └── servis_repo.go           # CRUD riwayat servis & statistik
├── service/
│   ├── search_service.go        # Sequential Search & Binary Search
│   └── sort_service.go          # Selection Sort & Insertion Sort
└── ui/
    ├── display.go               # Menu & helper tampilan
    ├── kendaraan_ui.go          # Handler UI kendaraan
    ├── pemilik_ui.go            # Handler UI pemilik
    ├── servis_ui.go             # Handler UI riwayat servis
    ├── search_sort_ui.go        # Handler UI pencarian & pengurutan
    └── statistik_ui.go          # Handler UI statistik & laporan
```

## Fitur Program

| No | Fitur | Keterangan |
|----|-------|------------|
| a | CRUD Kendaraan & Pemilik | Tambah, ubah, hapus, lihat data |
| b | Riwayat Servis | Catat tanggal, kerusakan, biaya, teknisi |
| c | Pencarian Kendaraan | Sequential Search O(n) & Binary Search O(log n) |
| d | Pengurutan Kendaraan | Selection Sort (tahun) & Insertion Sort (tanggal servis) |
| e | Statistik | Servis per bulan & kerusakan terbanyak |

## Cara Instalasi & Menjalankan

### 1. Prasyarat
- Go 1.21 atau lebih baru: https://go.dev/dl/
- GCC (diperlukan oleh go-sqlite3):
  - **Windows**: Install [TDM-GCC](https://jmeubank.github.io/tdm-gcc/) atau [MinGW-w64](https://www.mingw-w64.org/)
  - **Linux**: `sudo apt install gcc`
  - **macOS**: `xcode-select --install`

### 2. Clone / Download Proyek
```bash
# Masuk ke folder proyek
cd autocare
```

### 3. Install Dependensi
```bash
go mod tidy
```

### 4. Jalankan Program
```bash
go run main.go
```

### 5. Build Executable (opsional)
```bash
# Linux/macOS
go build -o autocare main.go

# Windows
go build -o autocare.exe main.go
```

## Database

Program menggunakan **SQLite** (`autocare.db`) yang otomatis dibuat saat pertama kali dijalankan.

### Tabel Database

| Tabel | Kolom Utama |
|-------|-------------|
| `pemilik` | id, nama, telepon, alamat, email |
| `kendaraan` | id, plat_nomor, merek, model, tahun_produksi, tipe_kendaraan, pemilik_id |
| `riwayat_servis` | id, kendaraan_id, tanggal_servis, jenis_kerusakan, detail_perbaikan, biaya, teknisi, status |

## Algoritma yang Diimplementasikan

### Pencarian
| Algoritma | File | Kompleksitas |
|-----------|------|-------------|
| Sequential Search | `service/search_service.go` | O(n) |
| Binary Search | `service/search_service.go` | O(log n) |

> **Catatan**: Binary Search membutuhkan data terurut. Program otomatis mengurutkan data sebelum pencarian.

### Pengurutan
| Algoritma | File | Kunci Pengurutan | Kompleksitas |
|-----------|------|-----------------|-------------|
| Selection Sort | `service/sort_service.go` | Tahun Produksi | O(n²) |
| Insertion Sort | `service/sort_service.go` | Tanggal Servis Terakhir | O(n²) worst, O(n) best |

## Pengembang
- Mata Kuliah: Algoritma Pemrograman 2
- Tugas Besar No. 4: Aplikasi Manajemen dan Riwayat Servis Kendaraan (AutoCare)
