package repository

import (
	"autocare/database"
	"autocare/models"
	"fmt"
)

// TambahPemilik menyimpan data pemilik baru ke database
func TambahPemilik(p *models.Pemilik) error {
	query := `INSERT INTO pemilik (nama, telepon, alamat, email) VALUES (?, ?, ?, ?)`
	result, err := database.DB.Exec(query, p.Nama, p.Telepon, p.Alamat, p.Email)
	if err != nil {
		return fmt.Errorf("gagal menambah pemilik: %w", err)
	}
	id, _ := result.LastInsertId()
	p.ID = int(id)
	return nil
}

// GetSemuaPemilik mengambil semua data pemilik dari database
func GetSemuaPemilik() ([]models.Pemilik, error) {
	query := `SELECT id, nama, telepon, alamat, email FROM pemilik ORDER BY nama`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Pemilik
	for rows.Next() {
		var p models.Pemilik
		if err := rows.Scan(&p.ID, &p.Nama, &p.Telepon, &p.Alamat, &p.Email); err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, nil
}

// GetPemilikByID mengambil pemilik berdasarkan ID
func GetPemilikByID(id int) (*models.Pemilik, error) {
	query := `SELECT id, nama, telepon, alamat, email FROM pemilik WHERE id = ?`
	row := database.DB.QueryRow(query, id)
	var p models.Pemilik
	err := row.Scan(&p.ID, &p.Nama, &p.Telepon, &p.Alamat, &p.Email)
	if err != nil {
		return nil, fmt.Errorf("pemilik tidak ditemukan")
	}
	return &p, nil
}

// UpdatePemilik memperbarui data pemilik di database
func UpdatePemilik(p *models.Pemilik) error {
	query := `UPDATE pemilik SET nama=?, telepon=?, alamat=?, email=? WHERE id=?`
	res, err := database.DB.Exec(query, p.Nama, p.Telepon, p.Alamat, p.Email, p.ID)
	if err != nil {
		return fmt.Errorf("gagal update pemilik: %w", err)
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("pemilik dengan ID %d tidak ditemukan", p.ID)
	}
	return nil
}

// HapusPemilik menghapus data pemilik dari database
func HapusPemilik(id int) error {
	query := `DELETE FROM pemilik WHERE id = ?`
	res, err := database.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("gagal hapus pemilik: %w", err)
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("pemilik dengan ID %d tidak ditemukan", id)
	}
	return nil
}
