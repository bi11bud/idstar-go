package dtos

import "time"

type RekeningResponse struct {
	CreatedDate time.Time  `json:"created_date"`
	UpdatedDate time.Time  `json:"updated_date"`
	DeletedDate *time.Time `json:"deleted_date"`
	ID          int        `json:"id"`
	Nama        string     `json:"nama"`
	Jenis       string     `json:"jenis"`
	Rekening    string     `json:"rekening"`
	Karyawan    Karyawan   `json:"karyawan"`
}

type Karyawan struct {
	ID   string `json:"id"`
	Nama string `json:"nama"`
}
