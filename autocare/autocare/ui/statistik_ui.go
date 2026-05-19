package ui

import (
	"autocare/repository"
	"fmt"
	"strings"
	"time"
)

// HandlerStatistik menampilkan statistik dan laporan
func HandlerStatistik() {
	fmt.Println()
	SubHeader("STATISTIK & LAPORAN")

	// Gunakan tahun saat ini sebagai default
	tahunDefault := time.Now().Format("2006")
	tahunInput := bacaInput(fmt.Sprintf("  Tahun (default %s): ", tahunDefault))
	if tahunInput == "" {
		tahunInput = tahunDefault
	}

	fmt.Printf("\n  === STATISTIK SERVIS TAHUN %s ===\n\n", tahunInput)

	// Statistik per bulan
	statBulanan, err := repository.GetStatistikBulanan(tahunInput)
	if err != nil {
		InfoGagal(err.Error())
	} else if len(statBulanan) == 0 {
		InfoTidakAda(fmt.Sprintf("Tidak ada data servis untuk tahun %s.", tahunInput))
	} else {
		fmt.Println("  Jumlah Kendaraan Diservis per Bulan:")
		GarisTipis(50)

		maxJumlah := 0
		for _, s := range statBulanan {
			if s.JumlahServis > maxJumlah {
				maxJumlah = s.JumlahServis
			}
		}

		totalServis := 0
		for _, s := range statBulanan {
			totalServis += s.JumlahServis
			barLen := 0
			if maxJumlah > 0 {
				barLen = s.JumlahServis * 25 / maxJumlah
			}
			bar := strings.Repeat("█", barLen)
			fmt.Printf("  %-10s | %-25s %d\n", s.Bulan, bar, s.JumlahServis)
		}
		GarisTipis(50)
		fmt.Printf("  Total servis: %d\n", totalServis)
	}

	fmt.Println()

	// Statistik kerusakan terbanyak
	statKerusakan, err := repository.GetStatistikKerusakan()
	if err != nil {
		InfoGagal(err.Error())
	} else if len(statKerusakan) == 0 {
		InfoTidakAda("Belum ada data kerusakan.")
	} else {
		fmt.Println("  Kategori Kerusakan yang Paling Sering Muncul:")
		GarisTipis(50)

		maxJml := statKerusakan[0].Jumlah
		for i, s := range statKerusakan {
			barLen := 0
			if maxJml > 0 {
				barLen = s.Jumlah * 20 / maxJml
			}
			bar := strings.Repeat("█", barLen)
			ranking := fmt.Sprintf("#%d", i+1)
			fmt.Printf("  %-3s %-18s | %-20s %d kali\n", ranking, s.JenisKerusakan, bar, s.Jumlah)
		}
		GarisTipis(50)
		if len(statKerusakan) > 0 {
			fmt.Printf("  Kerusakan paling sering: %s (%d kali)\n",
				statKerusakan[0].JenisKerusakan, statKerusakan[0].Jumlah)
		}
	}

	// Ringkasan total kendaraan dan pemilik
	fmt.Println()
	GarisTipis(50)
	kendaraanList, _ := repository.GetSemuaKendaraan()
	pemilikList, _ := repository.GetSemuaPemilik()
	semuaServis, _ := repository.GetSemuaServis()

	fmt.Printf("  Total Kendaraan Terdaftar : %d\n", len(kendaraanList))
	fmt.Printf("  Total Pemilik Terdaftar   : %d\n", len(pemilikList))
	fmt.Printf("  Total Riwayat Servis      : %d\n", len(semuaServis))

	var totalBiaya float64
	for _, rs := range semuaServis {
		totalBiaya += rs.Biaya
	}
	fmt.Printf("  Total Pendapatan Servis   : Rp%.0f\n", totalBiaya)
	GarisTipis(50)

	TungguEnter()
}
