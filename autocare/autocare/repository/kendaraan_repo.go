package repository

import (
	"autocare/database"
	"autocare/models"
	"fmt"
)

// TambahKendaraan menyimpan data kendaraan baru ke database
func TambahKendaraan(k *models.Kendaraan) error {
	query := `INSERT INTO kendaraan (plat_nomor, merek, model, tahun_produksi, tipe_kendaraan, pemilik_id)
	          VALUES (?, ?, ?, ?, ?, ?)`
	result, err := database.DB.Exec(query, k.PlatNomor, k.Merek, k.Model,
		k.TahunProduksi, k.TipeKendaraan, k.PemilikID)
	if err != nil {
		return fmt.Errorf("gagal menambah kendaraan: %w", err)
	}
	id, _ := result.LastInsertId()
	k.ID = int(id)
	return nil
}

// GetSemuaKendaraan mengambil semua data kendaraan dari database
func GetSemuaKendaraan() ([]models.Kendaraan, error) {
	query := `SELECT id, plat_nomor, merek, model, tahun_produksi, tipe_kendaraan, pemilik_id
	          FROM kendaraan ORDER BY plat_nomor`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Kendaraan
	for rows.Next() {
		var k models.Kendaraan
		if err := rows.Scan(&k.ID, &k.PlatNomor, &k.Merek, &k.Model,
			&k.TahunProduksi, &k.TipeKendaraan, &k.PemilikID); err != nil {
			return nil, err
		}
		list = append(list, k)
	}
	return list, nil
}

// GetKendaraanByID mengambil kendaraan berdasarkan ID
func GetKendaraanByID(id int) (*models.Kendaraan, error) {
	query := `SELECT id, plat_nomor, merek, model, tahun_produksi, tipe_kendaraan, pemilik_id
	          FROM kendaraan WHERE id = ?`
	row := database.DB.QueryRow(query, id)
	var k models.Kendaraan
	err := row.Scan(&k.ID, &k.PlatNomor, &k.Merek, &k.Model,
		&k.TahunProduksi, &k.TipeKendaraan, &k.PemilikID)
	if err != nil {
		return nil, fmt.Errorf("kendaraan tidak ditemukan")
	}
	return &k, nil
}

// UpdateKendaraan memperbarui data kendaraan di database
func UpdateKendaraan(k *models.Kendaraan) error {
	query := `UPDATE kendaraan SET plat_nomor=?, merek=?, model=?, tahun_produksi=?, tipe_kendaraan=?, pemilik_id=?
	          WHERE id=?`
	res, err := database.DB.Exec(query, k.PlatNomor, k.Merek, k.Model,
		k.TahunProduksi, k.TipeKendaraan, k.PemilikID, k.ID)
	if err != nil {
		return fmt.Errorf("gagal update kendaraan: %w", err)
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("kendaraan dengan ID %d tidak ditemukan", k.ID)
	}
	return nil
}

// HapusKendaraan menghapus data kendaraan dari database
func HapusKendaraan(id int) error {
	// Hapus riwayat servis terkait terlebih dahulu
	database.DB.Exec(`DELETE FROM riwayat_servis WHERE kendaraan_id = ?`, id)

	query := `DELETE FROM kendaraan WHERE id = ?`
	res, err := database.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("gagal hapus kendaraan: %w", err)
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("kendaraan dengan ID %d tidak ditemukan", id)
	}
	return nil
}

// GetServisTermakhir mengambil tanggal servis terakhir tiap kendaraan
func GetServisTermakhir(kendaraanID int) string {
	query := `SELECT tanggal_servis FROM riwayat_servis
	          WHERE kendaraan_id = ? ORDER BY tanggal_servis DESC LIMIT 1`
	row := database.DB.QueryRow(query, kendaraanID)
	var tgl string
	row.Scan(&tgl)
	if tgl == "" {
		return "-"
	}
	return tgl
}
