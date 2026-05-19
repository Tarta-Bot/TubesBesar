package main

import (
	"autocare/database"
	"autocare/ui"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Inisialisasi database SQLite
	if err := database.Inisialisasi("autocare.db"); err != nil {
		log.Fatalf("Gagal inisialisasi database: %v", err)
	}
	defer database.Tutup()

	reader := bufio.NewReader(os.Stdin)

	for {
		ui.MenuUtama()
		input, _ := reader.ReadString('\n')
		pilihan := strings.TrimSpace(input)

		switch pilihan {
		case "1":
			menuKendaraan(reader)
		case "2":
			menuPemilik(reader)
		case "3":
			menuServis(reader)
		case "4":
			menuPencarian(reader)
		case "5":
			menuPengurutan(reader)
		case "6":
			ui.HandlerStatistik()
		case "0":
			fmt.Println("\n  Terima kasih telah menggunakan AutoCare. Sampai jumpa!")
			return
		default:
			fmt.Println("\n  Pilihan tidak valid. Coba lagi.")
		}
	}
}

func menuKendaraan(reader *bufio.Reader) {
	for {
		ui.MenuKendaraan()
		input, _ := reader.ReadString('\n')
		pilihan := strings.TrimSpace(input)

		switch pilihan {
		case "1":
			ui.HandlerTambahKendaraan()
		case "2":
			ui.HandlerLihatKendaraan()
		case "3":
			ui.HandlerUbahKendaraan()
		case "4":
			ui.HandlerHapusKendaraan()
		case "0":
			return
		default:
			fmt.Println("\n  Pilihan tidak valid.")
		}
	}
}

func menuPemilik(reader *bufio.Reader) {
	for {
		ui.MenuPemilik()
		input, _ := reader.ReadString('\n')
		pilihan := strings.TrimSpace(input)

		switch pilihan {
		case "1":
			ui.HandlerTambahPemilik()
		case "2":
			ui.HandlerLihatPemilik()
		case "3":
			ui.HandlerUbahPemilik()
		case "4":
			ui.HandlerHapusPemilik()
		case "0":
			return
		default:
			fmt.Println("\n  Pilihan tidak valid.")
		}
	}
}

func menuServis(reader *bufio.Reader) {
	for {
		ui.MenuServis()
		input, _ := reader.ReadString('\n')
		pilihan := strings.TrimSpace(input)

		switch pilihan {
		case "1":
			ui.HandlerCatatServis()
		case "2":
			ui.HandlerLihatSemuaServis()
		case "3":
			ui.HandlerLihatServisByKendaraan()
		case "4":
			ui.HandlerHapusServis()
		case "0":
			return
		default:
			fmt.Println("\n  Pilihan tidak valid.")
		}
	}
}

func menuPencarian(reader *bufio.Reader) {
	for {
		ui.MenuPencarian()
		input, _ := reader.ReadString('\n')
		pilihan := strings.TrimSpace(input)

		switch pilihan {
		case "1":
			ui.HandlerSequentialSearch()
		case "2":
			ui.HandlerBinarySearch()
		case "0":
			return
		default:
			fmt.Println("\n  Pilihan tidak valid.")
		}
	}
}

func menuPengurutan(reader *bufio.Reader) {
	for {
		ui.MenuPengurutan()
		input, _ := reader.ReadString('\n')
		pilihan := strings.TrimSpace(input)

		switch pilihan {
		case "1":
			ui.HandlerSelectionSortTahun()
		case "2":
			ui.HandlerInsertionSortServis()
		case "0":
			return
		default:
			fmt.Println("\n  Pilihan tidak valid.")
		}
	}
}
