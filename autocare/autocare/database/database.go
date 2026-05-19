package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

// Inisialisasi membuat koneksi ke database SQLite dan membuat tabel jika belum ada
func Inisialisasi(dbPath string) error {
	var err error
	db, err := sql.Open("sqlite", "autocare.db")
	if err != nil {
		return fmt.Errorf("gagal membuka database: %w", err)
	}
	DB = db

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("gagal terhubung ke database: %w", err)
	}

	if err = buatTabel(); err != nil {
		return fmt.Errorf("gagal membuat tabel: %w", err)
	}

	log.Println("Database berhasil diinisialisasi:", dbPath)
	return nil
}

// buatTabel membuat semua tabel yang diperlukan
func buatTabel() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS pemilik (
			id       INTEGER PRIMARY KEY AUTOINCREMENT,
			nama     TEXT NOT NULL,
			telepon  TEXT,
			alamat   TEXT,
			email    TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS kendaraan (
			id             INTEGER PRIMARY KEY AUTOINCREMENT,
			plat_nomor     TEXT NOT NULL UNIQUE,
			merek          TEXT NOT NULL,
			model          TEXT NOT NULL,
			tahun_produksi INTEGER NOT NULL,
			tipe_kendaraan TEXT NOT NULL,
			pemilik_id     INTEGER,
			FOREIGN KEY (pemilik_id) REFERENCES pemilik(id)
		)`,
		`CREATE TABLE IF NOT EXISTS riwayat_servis (
			id               INTEGER PRIMARY KEY AUTOINCREMENT,
			kendaraan_id     INTEGER NOT NULL,
			tanggal_servis   TEXT NOT NULL,
			jenis_kerusakan  TEXT NOT NULL,
			detail_perbaikan TEXT,
			biaya            REAL DEFAULT 0,
			teknisi          TEXT,
			status           TEXT DEFAULT 'selesai',
			FOREIGN KEY (kendaraan_id) REFERENCES kendaraan(id)
		)`,
	}

	for _, q := range queries {
		if _, err := DB.Exec(q); err != nil {
			return err
		}
	}
	return nil
}

// Tutup menutup koneksi database
func Tutup() {
	if DB != nil {
		DB.Close()
	}
}
