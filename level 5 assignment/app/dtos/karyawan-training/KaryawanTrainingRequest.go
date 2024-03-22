package dtos

import (
	"github.com/go-playground/validator/v10"
)

type SaveKaryawanTrainingRequest struct {
	Karyawan KaryawanRequest `json:"karyawan"`
	Training TrainingRequest `json:"training"`
	Tanggal  string          `json:"tanggal"`
}

type UpdateKaryawanTrainingRequest struct {
	Id       int             `json:"id" validate:"required"`
	Karyawan KaryawanRequest `json:"karyawan"`
	Training TrainingRequest `json:"training"`
	Tanggal  string          `json:"tanggal"`
}

type KaryawanRequest struct {
	Id string `json:"id" validate:"required"`
}

type TrainingRequest struct {
	Id string `json:"id" validate:"required"`
}

func (c *UpdateKaryawanTrainingRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)

	if err != nil {
		return err
	}

	return nil
}
