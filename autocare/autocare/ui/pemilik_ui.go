package ui

import (
	"autocare/models"
	"autocare/repository"
	"fmt"
	"strconv"
	"strings"
)

// TampilkanPemilik mencetak daftar pemilik dalam format tabel
func TampilkanPemilik(list []models.Pemilik) {
	if len(list) == 0 {
		InfoTidakAda("Belum ada data pemilik.")
		return
	}
	fmt.Println()
	fmt.Printf("  %-4s %-20s %-15s %-25s %-20s\n",
		"ID", "Nama", "Telepon", "Email", "Alamat")
	GarisTipis(90)
	for _, p := range list {
		fmt.Printf("  %-4d %-20s %-15s %-25s %-20s\n",
			p.ID, p.Nama, p.Telepon, p.Email, p.Alamat)
	}
	GarisTipis(90)
	fmt.Printf("  Total: %d pemilik\n", len(list))
}

// HandlerTambahPemilik menangani input untuk menambah pemilik baru
func HandlerTambahPemilik() {
	fmt.Println()
	SubHeader("TAMBAH PEMILIK BARU")

	p := models.Pemilik{}
	p.Nama = bacaInput("  Nama    : ")
	p.Telepon = bacaInput("  Telepon : ")
	p.Email = bacaInput("  Email   : ")
	p.Alamat = bacaInput("  Alamat  : ")

	if p.Nama == "" {
		InfoGagal("Nama tidak boleh kosong.")
		TungguEnter()
		return
	}

	if err := repository.TambahPemilik(&p); err != nil {
		InfoGagal(err.Error())
	} else {
		InfoSukses(fmt.Sprintf("Pemilik %s berhasil ditambahkan (ID: %d)", p.Nama, p.ID))
	}
	TungguEnter()
}

// HandlerLihatPemilik menampilkan semua pemilik
func HandlerLihatPemilik() {
	fmt.Println()
	SubHeader("DAFTAR SEMUA PEMILIK")
	list, err := repository.GetSemuaPemilik()
	if err != nil {
		InfoGagal(err.Error())
	} else {
		TampilkanPemilik(list)
	}
	TungguEnter()
}

// HandlerUbahPemilik menangani input untuk mengubah data pemilik
func HandlerUbahPemilik() {
	fmt.Println()
	SubHeader("UBAH DATA PEMILIK")
	list, _ := repository.GetSemuaPemilik()
	TampilkanPemilik(list)
	if len(list) == 0 {
		TungguEnter()
		return
	}

	idStr := bacaInput("\n  Masukkan ID pemilik yang akan diubah: ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		InfoGagal("ID tidak valid.")
		TungguEnter()
		return
	}

	p, err := repository.GetPemilikByID(id)
	if err != nil {
		InfoGagal(err.Error())
		TungguEnter()
		return
	}

	fmt.Printf("\n  Data saat ini: %s | %s | %s\n", p.Nama, p.Telepon, p.Email)
	fmt.Println("  (Tekan Enter untuk mempertahankan data lama)")

	input := bacaInput("  Nama Baru    : ")
	if input != "" {
		p.Nama = input
	}
	input = bacaInput("  Telepon Baru : ")
	if input != "" {
		p.Telepon = input
	}
	input = bacaInput("  Email Baru   : ")
	if input != "" {
		p.Email = input
	}
	input = bacaInput("  Alamat Baru  : ")
	if input != "" {
		p.Alamat = input
	}

	if err := repository.UpdatePemilik(p); err != nil {
		InfoGagal(err.Error())
	} else {
		InfoSukses("Data pemilik berhasil diperbarui.")
	}
	TungguEnter()
}

// HandlerHapusPemilik menangani penghapusan pemilik
func HandlerHapusPemilik() {
	fmt.Println()
	SubHeader("HAPUS PEMILIK")
	list, _ := repository.GetSemuaPemilik()
	TampilkanPemilik(list)
	if len(list) == 0 {
		TungguEnter()
		return
	}

	idStr := bacaInput("\n  Masukkan ID pemilik yang akan dihapus: ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		InfoGagal("ID tidak valid.")
		TungguEnter()
		return
	}

	p, err := repository.GetPemilikByID(id)
	if err != nil {
		InfoGagal(err.Error())
		TungguEnter()
		return
	}

	konfirmasi := bacaInput(fmt.Sprintf("  Hapus pemilik %s? (y/n): ", p.Nama))
	if strings.ToLower(konfirmasi) != "y" {
		InfoTidakAda("Penghapusan dibatalkan.")
		TungguEnter()
		return
	}

	if err := repository.HapusPemilik(id); err != nil {
		InfoGagal(err.Error())
	} else {
		InfoSukses("Pemilik berhasil dihapus.")
	}
	TungguEnter()
}
