package models

// Kendaraan menyimpan data kendaraan
type Kendaraan struct {
	ID           int
	PlatNomor    string
	Merek        string
	Model        string
	TahunProduksi int
	TipeKendaraan string // mobil, motor, truk, dll
	PemilikID    int
}

// Pemilik menyimpan data pemilik kendaraan
type Pemilik struct {
	ID       int
	Nama     string
	Telepon  string
	Alamat   string
	Email    string
}

// RiwayatServis menyimpan riwayat servis kendaraan
type RiwayatServis struct {
	ID              int
	KendaraanID     int
	PlatNomor       string // untuk tampilan
	TanggalServis   string
	JenisKerusakan  string
	DetailPerbaikan string
	Biaya           float64
	Teknisi         string
	Status          string // selesai, proses, antri
}

// Statistik untuk laporan bulanan
type StatistikBulanan struct {
	Bulan         string
	JumlahServis  int
}

// StatistikKerusakan untuk laporan kerusakan terbanyak
type StatistikKerusakan struct {
	JenisKerusakan string
	Jumlah         int
}
