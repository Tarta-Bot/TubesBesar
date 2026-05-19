package service

import (
	"autocare/models"
	"strings"
)

// SequentialSearch mencari kendaraan berdasarkan plat nomor secara sequential
// Time Complexity: O(n)
func SequentialSearch(kendaraanList []models.Kendaraan, platNomor string) *models.Kendaraan {
	platNomor = strings.ToUpper(strings.TrimSpace(platNomor))
	for i := range kendaraanList {
		if strings.ToUpper(kendaraanList[i].PlatNomor) == platNomor {
			return &kendaraanList[i]
		}
	}
	return nil
}

// BinarySearch mencari kendaraan berdasarkan plat nomor secara binary search
// SYARAT: Data harus sudah diurutkan berdasarkan PlatNomor (A-Z)
// Time Complexity: O(log n)
func BinarySearch(kendaraanList []models.Kendaraan, platNomor string) *models.Kendaraan {
	platNomor = strings.ToUpper(strings.TrimSpace(platNomor))
	kiri := 0
	kanan := len(kendaraanList) - 1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		current := strings.ToUpper(kendaraanList[tengah].PlatNomor)

		if current == platNomor {
			return &kendaraanList[tengah]
		} else if current < platNomor {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return nil
}

// UrutkanByPlatNomor mengurutkan slice kendaraan berdasarkan plat nomor (A-Z)
// diperlukan sebelum binary search
func UrutkanByPlatNomor(list []models.Kendaraan) []models.Kendaraan {
	sorted := make([]models.Kendaraan, len(list))
	copy(sorted, list)
	// Gunakan insertion sort untuk menjaga konsistensi dengan service sorting
	for i := 1; i < len(sorted); i++ {
		key := sorted[i]
		j := i - 1
		for j >= 0 && strings.ToUpper(sorted[j].PlatNomor) > strings.ToUpper(key.PlatNomor) {
			sorted[j+1] = sorted[j]
			j--
		}
		sorted[j+1] = key
	}
	return sorted
}
