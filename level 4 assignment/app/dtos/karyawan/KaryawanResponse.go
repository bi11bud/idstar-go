package dtos

import "time"

type KaryawanResponse struct {
	CreatedDate    time.Time              `json:"created_date"`
	UpdatedDate    time.Time              `json:"updated_date"`
	DeletedDate    *time.Time             `json:"deleted_date"`
	ID             int                    `json:"id"`
	Nama           string                 `json:"nama"`
	Dob            string                 `json:"dob"`
	Status         string                 `json:"status"`
	Alamat         string                 `json:"alamat"`
	DetailKaryawan DetailKaryawanResponse `json:"detailKaryawan"`
}

type DetailKaryawanResponse struct {
	CreatedDate time.Time  `json:"created_date"`
	UpdatedDate time.Time  `json:"updated_date"`
	DeletedDate *time.Time `json:"deleted_date"`
	ID          int        `json:"id"`
	Nik         string     `json:"nik"`
	Npwp        string     `json:"npwp"`
}
