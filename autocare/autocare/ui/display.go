package ui

import (
	"fmt"
	"strings"
)

// GarisHorizontal mencetak garis horizontal
func GarisHorizontal(panjang int) {
	fmt.Println(strings.Repeat("=", panjang))
}

// GarisTipis mencetak garis tipis
func GarisTipis(panjang int) {
	fmt.Println(strings.Repeat("-", panjang))
}

// Header mencetak header dengan judul
func Header(judul string) {
	lebar := 60
	GarisHorizontal(lebar)
	padding := (lebar - len(judul)) / 2
	if padding < 0 {
		padding = 0
	}
	fmt.Printf("%s%s\n", strings.Repeat(" ", padding), judul)
	GarisHorizontal(lebar)
}

// SubHeader mencetak sub header
func SubHeader(judul string) {
	GarisTipis(60)
	fmt.Printf("  [ %s ]\n", judul)
	GarisTipis(60)
}

// MenuUtama menampilkan menu utama aplikasi
func MenuUtama() {
	fmt.Println()
	Header("AUTOCARE - Manajemen Servis Kendaraan")
	fmt.Println()
	fmt.Println("  [1] Manajemen Kendaraan")
	fmt.Println("  [2] Manajemen Pemilik")
	fmt.Println("  [3] Riwayat Servis")
	fmt.Println("  [4] Pencarian Kendaraan")
	fmt.Println("  [5] Pengurutan Kendaraan")
	fmt.Println("  [6] Statistik & Laporan")
	fmt.Println("  [0] Keluar")
	fmt.Println()
	GarisTipis(60)
	fmt.Print("  Pilih menu: ")
}

// MenuKendaraan menampilkan submenu kendaraan
func MenuKendaraan() {
	fmt.Println()
	SubHeader("MANAJEMEN KENDARAAN")
	fmt.Println("  [1] Tambah Kendaraan")
	fmt.Println("  [2] Lihat Semua Kendaraan")
	fmt.Println("  [3] Ubah Data Kendaraan")
	fmt.Println("  [4] Hapus Kendaraan")
	fmt.Println("  [0] Kembali")
	fmt.Println()
	fmt.Print("  Pilih: ")
}

// MenuPemilik menampilkan submenu pemilik
func MenuPemilik() {
	fmt.Println()
	SubHeader("MANAJEMEN PEMILIK")
	fmt.Println("  [1] Tambah Pemilik")
	fmt.Println("  [2] Lihat Semua Pemilik")
	fmt.Println("  [3] Ubah Data Pemilik")
	fmt.Println("  [4] Hapus Pemilik")
	fmt.Println("  [0] Kembali")
	fmt.Println()
	fmt.Print("  Pilih: ")
}

// MenuServis menampilkan submenu riwayat servis
func MenuServis() {
	fmt.Println()
	SubHeader("RIWAYAT SERVIS")
	fmt.Println("  [1] Catat Servis Baru")
	fmt.Println("  [2] Lihat Semua Riwayat Servis")
	fmt.Println("  [3] Lihat Servis per Kendaraan")
	fmt.Println("  [4] Hapus Riwayat Servis")
	fmt.Println("  [0] Kembali")
	fmt.Println()
	fmt.Print("  Pilih: ")
}

// MenuPencarian menampilkan submenu pencarian
func MenuPencarian() {
	fmt.Println()
	SubHeader("PENCARIAN KENDARAAN")
	fmt.Println("  [1] Sequential Search")
	fmt.Println("  [2] Binary Search")
	fmt.Println("  [0] Kembali")
	fmt.Println()
	fmt.Print("  Pilih: ")
}

// MenuPengurutan menampilkan submenu pengurutan
func MenuPengurutan() {
	fmt.Println()
	SubHeader("PENGURUTAN KENDARAAN")
	fmt.Println("  [1] Urutkan berdasarkan Tahun Produksi (Selection Sort)")
	fmt.Println("  [2] Urutkan berdasarkan Tanggal Servis (Insertion Sort)")
	fmt.Println("  [0] Kembali")
	fmt.Println()
	fmt.Print("  Pilih: ")
}

// MenuUrutan menampilkan pilihan arah urutan
func MenuUrutan() string {
	fmt.Println("  Urutan:")
	fmt.Println("  [1] Ascending (Terkecil/Terlama ke Terbesar/Terbaru)")
	fmt.Println("  [2] Descending (Terbesar/Terbaru ke Terkecil/Terlama)")
	fmt.Print("  Pilih: ")
	return ""
}

// InfoSukses menampilkan pesan sukses
func InfoSukses(pesan string) {
	fmt.Printf("\n  ✔ %s\n", pesan)
}

// InfoGagal menampilkan pesan gagal
func InfoGagal(pesan string) {
	fmt.Printf("\n  ✘ ERROR: %s\n", pesan)
}

// InfoTidakAda menampilkan pesan data tidak ditemukan
func InfoTidakAda(pesan string) {
	fmt.Printf("\n  ℹ %s\n", pesan)
}

// TungguEnter menunggu user menekan Enter
func TungguEnter() {
	fmt.Print("\n  Tekan Enter untuk melanjutkan...")
	fmt.Scanln()
}
