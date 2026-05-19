package ui

import (
	"autocare/models"
	"autocare/repository"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func bacaInput(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// TampilkanKendaraan mencetak daftar kendaraan dalam format tabel
func TampilkanKendaraan(list []models.Kendaraan) {
	if len(list) == 0 {
		InfoTidakAda("Belum ada data kendaraan.")
		return
	}
	fmt.Println()
	fmt.Printf("  %-4s %-12s %-12s %-15s %-6s %-10s %-10s\n",
		"No", "Plat Nomor", "Merek", "Model", "Tahun", "Tipe", "Servis Terakhir")
	GarisTipis(80)
	for i, k := range list {
		servisTerakhir := repository.GetServisTermakhir(k.ID)
		fmt.Printf("  %-4d %-12s %-12s %-15s %-6d %-10s %-10s\n",
			i+1, k.PlatNomor, k.Merek, k.Model, k.TahunProduksi, k.TipeKendaraan, servisTerakhir)
	}
	GarisTipis(80)
	fmt.Printf("  Total: %d kendaraan\n", len(list))
}

// HandlerTambahKendaraan menangani input untuk menambah kendaraan baru
func HandlerTambahKendaraan() {
	fmt.Println()
	SubHeader("TAMBAH KENDARAAN BARU")

	// Tampilkan daftar pemilik
	pemilikList, _ := repository.GetSemuaPemilik()
	if len(pemilikList) == 0 {
		InfoTidakAda("Belum ada data pemilik. Tambah pemilik terlebih dahulu.")
		TungguEnter()
		return
	}

	fmt.Println("  Daftar Pemilik:")
	for _, p := range pemilikList {
		fmt.Printf("    [%d] %s - %s\n", p.ID, p.Nama, p.Telepon)
	}

	k := models.Kendaraan{}
	k.PlatNomor = strings.ToUpper(bacaInput("\n  Plat Nomor     : "))
	k.Merek = bacaInput("  Merek          : ")
	k.Model = bacaInput("  Model          : ")
	tahunStr := bacaInput("  Tahun Produksi : ")
	k.TipeKendaraan = bacaInput("  Tipe Kendaraan : ")
	pemilikIDStr := bacaInput("  ID Pemilik     : ")

	tahun, err := strconv.Atoi(tahunStr)
	if err != nil || tahun < 1900 || tahun > 2100 {
		InfoGagal("Tahun produksi tidak valid.")
		TungguEnter()
		return
	}
	k.TahunProduksi = tahun

	pemilikID, err := strconv.Atoi(pemilikIDStr)
	if err != nil {
		InfoGagal("ID pemilik tidak valid.")
		TungguEnter()
		return
	}
	k.PemilikID = pemilikID

	if k.PlatNomor == "" || k.Merek == "" || k.Model == "" {
		InfoGagal("Plat nomor, merek, dan model tidak boleh kosong.")
		TungguEnter()
		return
	}

	if err := repository.TambahKendaraan(&k); err != nil {
		InfoGagal(err.Error())
	} else {
		InfoSukses(fmt.Sprintf("Kendaraan %s berhasil ditambahkan (ID: %d)", k.PlatNomor, k.ID))
	}
	TungguEnter()
}

// HandlerLihatKendaraan menampilkan semua kendaraan
func HandlerLihatKendaraan() {
	fmt.Println()
	SubHeader("DAFTAR SEMUA KENDARAAN")
	list, err := repository.GetSemuaKendaraan()
	if err != nil {
		InfoGagal(err.Error())
	} else {
		TampilkanKendaraan(list)
	}
	TungguEnter()
}

// HandlerUbahKendaraan menangani input untuk mengubah data kendaraan
func HandlerUbahKendaraan() {
	fmt.Println()
	SubHeader("UBAH DATA KENDARAAN")
	list, _ := repository.GetSemuaKendaraan()
	TampilkanKendaraan(list)
	if len(list) == 0 {
		TungguEnter()
		return
	}

	idStr := bacaInput("\n  Masukkan ID kendaraan yang akan diubah: ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		InfoGagal("ID tidak valid.")
		TungguEnter()
		return
	}

	k, err := repository.GetKendaraanByID(id)
	if err != nil {
		InfoGagal(err.Error())
		TungguEnter()
		return
	}

	fmt.Printf("\n  Data saat ini: %s | %s %s | %d\n", k.PlatNomor, k.Merek, k.Model, k.TahunProduksi)
	fmt.Println("  (Tekan Enter untuk mempertahankan data lama)")

	input := bacaInput("  Plat Nomor Baru     : ")
	if input != "" {
		k.PlatNomor = strings.ToUpper(input)
	}
	input = bacaInput("  Merek Baru          : ")
	if input != "" {
		k.Merek = input
	}
	input = bacaInput("  Model Baru          : ")
	if input != "" {
		k.Model = input
	}
	input = bacaInput("  Tahun Produksi Baru : ")
	if input != "" {
		tahun, err := strconv.Atoi(input)
		if err == nil {
			k.TahunProduksi = tahun
		}
	}
	input = bacaInput("  Tipe Kendaraan Baru : ")
	if input != "" {
		k.TipeKendaraan = input
	}

	if err := repository.UpdateKendaraan(k); err != nil {
		InfoGagal(err.Error())
	} else {
		InfoSukses("Data kendaraan berhasil diperbarui.")
	}
	TungguEnter()
}

// HandlerHapusKendaraan menangani penghapusan kendaraan
func HandlerHapusKendaraan() {
	fmt.Println()
	SubHeader("HAPUS KENDARAAN")
	list, _ := repository.GetSemuaKendaraan()
	TampilkanKendaraan(list)
	if len(list) == 0 {
		TungguEnter()
		return
	}

	idStr := bacaInput("\n  Masukkan ID kendaraan yang akan dihapus: ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		InfoGagal("ID tidak valid.")
		TungguEnter()
		return
	}

	k, err := repository.GetKendaraanByID(id)
	if err != nil {
		InfoGagal(err.Error())
		TungguEnter()
		return
	}

	konfirmasi := bacaInput(fmt.Sprintf("  Hapus kendaraan %s (%s %s)? (y/n): ", k.PlatNomor, k.Merek, k.Model))
	if strings.ToLower(konfirmasi) != "y" {
		InfoTidakAda("Penghapusan dibatalkan.")
		TungguEnter()
		return
	}

	if err := repository.HapusKendaraan(id); err != nil {
		InfoGagal(err.Error())
	} else {
		InfoSukses("Kendaraan berhasil dihapus.")
	}
	TungguEnter()
}
