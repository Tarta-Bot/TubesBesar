package repository

import (
	"autocare/database"
	"autocare/models"
	"fmt"
)

// TambahServis menyimpan riwayat servis baru ke database
func TambahServis(rs *models.RiwayatServis) error {
	query := `INSERT INTO riwayat_servis (kendaraan_id, tanggal_servis, jenis_kerusakan, detail_perbaikan, biaya, teknisi, status)
	          VALUES (?, ?, ?, ?, ?, ?, ?)`
	result, err := database.DB.Exec(query, rs.KendaraanID, rs.TanggalServis,
		rs.JenisKerusakan, rs.DetailPerbaikan, rs.Biaya, rs.Teknisi, rs.Status)
	if err != nil {
		return fmt.Errorf("gagal menambah riwayat servis: %w", err)
	}
	id, _ := result.LastInsertId()
	rs.ID = int(id)
	return nil
}

// GetSemuaServis mengambil semua riwayat servis beserta plat nomor kendaraan
func GetSemuaServis() ([]models.RiwayatServis, error) {
	query := `SELECT rs.id, rs.kendaraan_id, k.plat_nomor, rs.tanggal_servis,
	                 rs.jenis_kerusakan, rs.detail_perbaikan, rs.biaya, rs.teknisi, rs.status
	          FROM riwayat_servis rs
	          JOIN kendaraan k ON rs.kendaraan_id = k.id
	          ORDER BY rs.tanggal_servis DESC`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.RiwayatServis
	for rows.Next() {
		var rs models.RiwayatServis
		if err := rows.Scan(&rs.ID, &rs.KendaraanID, &rs.PlatNomor, &rs.TanggalServis,
			&rs.JenisKerusakan, &rs.DetailPerbaikan, &rs.Biaya, &rs.Teknisi, &rs.Status); err != nil {
			return nil, err
		}
		list = append(list, rs)
	}
	return list, nil
}

// GetServisByKendaraan mengambil riwayat servis berdasarkan kendaraan ID
func GetServisByKendaraan(kendaraanID int) ([]models.RiwayatServis, error) {
	query := `SELECT rs.id, rs.kendaraan_id, k.plat_nomor, rs.tanggal_servis,
	                 rs.jenis_kerusakan, rs.detail_perbaikan, rs.biaya, rs.teknisi, rs.status
	          FROM riwayat_servis rs
	          JOIN kendaraan k ON rs.kendaraan_id = k.id
	          WHERE rs.kendaraan_id = ?
	          ORDER BY rs.tanggal_servis DESC`
	rows, err := database.DB.Query(query, kendaraanID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.RiwayatServis
	for rows.Next() {
		var rs models.RiwayatServis
		if err := rows.Scan(&rs.ID, &rs.KendaraanID, &rs.PlatNomor, &rs.TanggalServis,
			&rs.JenisKerusakan, &rs.DetailPerbaikan, &rs.Biaya, &rs.Teknisi, &rs.Status); err != nil {
			return nil, err
		}
		list = append(list, rs)
	}
	return list, nil
}

// HapusServis menghapus riwayat servis dari database
func HapusServis(id int) error {
	res, err := database.DB.Exec(`DELETE FROM riwayat_servis WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("gagal hapus servis: %w", err)
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("riwayat servis ID %d tidak ditemukan", id)
	}
	return nil
}

// GetStatistikBulanan mengambil jumlah servis per bulan dalam setahun
func GetStatistikBulanan(tahun string) ([]models.StatistikBulanan, error) {
	query := `SELECT strftime('%m', tanggal_servis) as bulan, COUNT(*) as jumlah
	          FROM riwayat_servis
	          WHERE strftime('%Y', tanggal_servis) = ?
	          GROUP BY bulan ORDER BY bulan`
	rows, err := database.DB.Query(query, tahun)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	namaBulan := map[string]string{
		"01": "Januari", "02": "Februari", "03": "Maret", "04": "April",
		"05": "Mei", "06": "Juni", "07": "Juli", "08": "Agustus",
		"09": "September", "10": "Oktober", "11": "November", "12": "Desember",
	}

	var list []models.StatistikBulanan
	for rows.Next() {
		var s models.StatistikBulanan
		var bulanNum string
		if err := rows.Scan(&bulanNum, &s.JumlahServis); err != nil {
			return nil, err
		}
		s.Bulan = namaBulan[bulanNum]
		list = append(list, s)
	}
	return list, nil
}

// GetStatistikKerusakan mengambil jenis kerusakan yang paling sering muncul
func GetStatistikKerusakan() ([]models.StatistikKerusakan, error) {
	query := `SELECT jenis_kerusakan, COUNT(*) as jumlah
	          FROM riwayat_servis
	          GROUP BY jenis_kerusakan
	          ORDER BY jumlah DESC LIMIT 10`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.StatistikKerusakan
	for rows.Next() {
		var s models.StatistikKerusakan
		if err := rows.Scan(&s.JenisKerusakan, &s.Jumlah); err != nil {
			return nil, err
		}
		list = append(list, s)
	}
	return list, nil
}
