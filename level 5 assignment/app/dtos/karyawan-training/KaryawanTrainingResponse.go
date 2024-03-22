package dtos

import (
	"time"

	karyawan "idstar.com/app/dtos/karyawan"
	training "idstar.com/app/dtos/training"
)

type KaryawanTrainingResponse struct {
	CreatedDate time.Time                 `json:"created_date"`
	UpdatedDate time.Time                 `json:"updated_date"`
	DeletedDate *time.Time                `json:"deleted_date"`
	ID          int                       `json:"id"`
	Karyawan    karyawan.KaryawanResponse `json:"karyawan"`
	Training    training.TrainingResponse `json:"training"`
	Tanggal     time.Time                 `json:"tanggal"`
}
