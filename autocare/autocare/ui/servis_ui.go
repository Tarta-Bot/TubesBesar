package ui

import (
	"autocare/models"
	"autocare/repository"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// TampilkanServis mencetak daftar riwayat servis dalam format tabel
func TampilkanServis(list []models.RiwayatServis) {
	if len(list) == 0 {
		InfoTidakAda("Belum ada riwayat servis.")
		return
	}
	fmt.Println()
	fmt.Printf("  %-4s %-12s %-12s %-20s %-15s %s\n",
		"ID", "Plat", "Tanggal", "Kerusakan", "Teknisi", "Biaya")
	GarisTipis(90)
	for _, rs := range list {
		fmt.Printf("  %-4d %-12s %-12s %-20s %-15s Rp%.0f\n",
			rs.ID, rs.PlatNomor, rs.TanggalServis,
			rs.JenisKerusakan, rs.Teknisi, rs.Biaya)
	}
	GarisTipis(90)
	fmt.Printf("  Total: %d riwayat servis\n", len(list))
}

// HandlerCatatServis menangani input untuk mencatat servis baru
func HandlerCatatServis() {
	fmt.Println()
	SubHeader("CATAT SERVIS BARU")

	// Tampilkan daftar kendaraan
	kendaraanList, _ := repository.GetSemuaKendaraan()
	if len(kendaraanList) == 0 {
		InfoTidakAda("Belum ada data kendaraan. Tambah kendaraan terlebih dahulu.")
		TungguEnter()
		return
	}

	fmt.Println("  Daftar Kendaraan:")
	for _, k := range kendaraanList {
		fmt.Printf("    [%d] %s - %s %s (%d)\n", k.ID, k.PlatNomor, k.Merek, k.Model, k.TahunProduksi)
	}

	rs := models.RiwayatServis{}
	kendaraanIDStr := bacaInput("\n  ID Kendaraan    : ")
	kendaraanID, err := strconv.Atoi(kendaraanIDStr)
	if err != nil {
		InfoGagal("ID kendaraan tidak valid.")
		TungguEnter()
		return
	}
	rs.KendaraanID = kendaraanID

	// Verifikasi kendaraan ada
	_, err = repository.GetKendaraanByID(kendaraanID)
	if err != nil {
		InfoGagal("Kendaraan tidak ditemukan.")
		TungguEnter()
		return
	}

	// Default tanggal hari ini
	hariIni := time.Now().Format("2006-01-02")
	tglInput := bacaInput(fmt.Sprintf("  Tanggal Servis  (default %s): ", hariIni))
	if tglInput == "" {
		rs.TanggalServis = hariIni
	} else {
		rs.TanggalServis = tglInput
	}

	fmt.Println("\n  Jenis Kerusakan (contoh: Mesin, Rem, Oli, Ban, Kelistrikan, dll)")
	rs.JenisKerusakan = bacaInput("  Jenis Kerusakan : ")
	rs.DetailPerbaikan = bacaInput("  Detail Perbaikan: ")
	rs.Teknisi = bacaInput("  Nama Teknisi    : ")

	biayaStr := bacaInput("  Biaya (Rp)      : ")
	if biayaStr != "" {
		// Hapus titik dan koma pemisah ribuan jika ada
		biayaStr = strings.ReplaceAll(biayaStr, ".", "")
		biayaStr = strings.ReplaceAll(biayaStr, ",", "")
		biaya, err := strconv.ParseFloat(biayaStr, 64)
		if err == nil {
			rs.Biaya = biaya
		}
	}

	fmt.Println("  Status: [1] Selesai  [2] Proses  [3] Antri")
	statusPilihan := bacaInput("  Pilih Status    : ")
	switch statusPilihan {
	case "2":
		rs.Status = "proses"
	case "3":
		rs.Status = "antri"
	default:
		rs.Status = "selesai"
	}

	if rs.JenisKerusakan == "" {
		InfoGagal("Jenis kerusakan tidak boleh kosong.")
		TungguEnter()
		return
	}

	if err := repository.TambahServis(&rs); err != nil {
		InfoGagal(err.Error())
	} else {
		InfoSukses(fmt.Sprintf("Riwayat servis berhasil dicatat (ID: %d)", rs.ID))
	}
	TungguEnter()
}

// HandlerLihatSemuaServis menampilkan semua riwayat servis
func HandlerLihatSemuaServis() {
	fmt.Println()
	SubHeader("SEMUA RIWAYAT SERVIS")
	list, err := repository.GetSemuaServis()
	if err != nil {
		InfoGagal(err.Error())
	} else {
		TampilkanServis(list)
	}
	TungguEnter()
}

// HandlerLihatServisByKendaraan menampilkan riwayat servis untuk kendaraan tertentu
func HandlerLihatServisByKendaraan() {
	fmt.Println()
	SubHeader("SERVIS PER KENDARAAN")

	kendaraanList, _ := repository.GetSemuaKendaraan()
	if len(kendaraanList) == 0 {
		InfoTidakAda("Belum ada data kendaraan.")
		TungguEnter()
		return
	}

	fmt.Println("  Daftar Kendaraan:")
	for _, k := range kendaraanList {
		fmt.Printf("    [%d] %s - %s %s\n", k.ID, k.PlatNomor, k.Merek, k.Model)
	}

	idStr := bacaInput("\n  Masukkan ID Kendaraan: ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		InfoGagal("ID tidak valid.")
		TungguEnter()
		return
	}

	k, err := repository.GetKendaraanByID(id)
	if err != nil {
		InfoGagal("Kendaraan tidak ditemukan.")
		TungguEnter()
		return
	}

	fmt.Printf("\n  Riwayat Servis: %s - %s %s\n", k.PlatNomor, k.Merek, k.Model)
	list, err := repository.GetServisByKendaraan(id)
	if err != nil {
		InfoGagal(err.Error())
	} else {
		TampilkanServis(list)
		// Hitung total biaya
		var total float64
		for _, rs := range list {
			total += rs.Biaya
		}
		fmt.Printf("  Total Biaya Servis: Rp%.0f\n", total)
	}
	TungguEnter()
}

// HandlerHapusServis menangani penghapusan riwayat servis
func HandlerHapusServis() {
	fmt.Println()
	SubHeader("HAPUS RIWAYAT SERVIS")
	list, _ := repository.GetSemuaServis()
	TampilkanServis(list)
	if len(list) == 0 {
		TungguEnter()
		return
	}

	idStr := bacaInput("\n  Masukkan ID riwayat servis yang akan dihapus: ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		InfoGagal("ID tidak valid.")
		TungguEnter()
		return
	}

	konfirmasi := bacaInput(fmt.Sprintf("  Hapus riwayat servis ID %d? (y/n): ", id))
	if strings.ToLower(konfirmasi) != "y" {
		InfoTidakAda("Penghapusan dibatalkan.")
		TungguEnter()
		return
	}

	if err := repository.HapusServis(id); err != nil {
		InfoGagal(err.Error())
	} else {
		InfoSukses("Riwayat servis berhasil dihapus.")
	}
	TungguEnter()
}
