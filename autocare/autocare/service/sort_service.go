package service

import (
	"autocare/models"
	"strings"
)

// SelectionSortByTahun mengurutkan kendaraan berdasarkan tahun produksi
// menggunakan algoritma Selection Sort
// Time Complexity: O(n^2)
func SelectionSortByTahun(list []models.Kendaraan, ascending bool) []models.Kendaraan {
	sorted := make([]models.Kendaraan, len(list))
	copy(sorted, list)

	n := len(sorted)
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if ascending {
				if sorted[j].TahunProduksi < sorted[idx].TahunProduksi {
					idx = j
				}
			} else {
				if sorted[j].TahunProduksi > sorted[idx].TahunProduksi {
					idx = j
				}
			}
		}
		// Tukar elemen
		sorted[i], sorted[idx] = sorted[idx], sorted[i]
	}
	return sorted
}

// InsertionSortByTanggalServis mengurutkan kendaraan berdasarkan tanggal servis terakhir
// menggunakan algoritma Insertion Sort
// Time Complexity: O(n^2) worst case, O(n) best case (sudah terurut)
func InsertionSortByTanggalServis(list []models.Kendaraan, servisMap map[int]string, ascending bool) []models.Kendaraan {
	sorted := make([]models.Kendaraan, len(list))
	copy(sorted, list)

	for i := 1; i < len(sorted); i++ {
		key := sorted[i]
		keyTgl := servisMap[key.ID]
		j := i - 1

		for j >= 0 {
			tglJ := servisMap[sorted[j].ID]
			var harus bool
			if ascending {
				harus = strings.Compare(tglJ, keyTgl) > 0
			} else {
				harus = strings.Compare(tglJ, keyTgl) < 0
			}

			if !harus {
				break
			}
			sorted[j+1] = sorted[j]
			j--
		}
		sorted[j+1] = key
	}
	return sorted
}
