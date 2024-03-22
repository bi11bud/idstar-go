package dtos

import (
	"github.com/go-playground/validator/v10"
)

type UpdateRekeningRequest struct {
	Id       string           `json:"id" validate:"required"`
	Nama     string           `json:"nama"`
	Jenis    string           `json:"jenis"`
	Rekening string           `json:"rekening"`
	Alamat   string           `json:"alamat"`
	Karyawan RekeningKaryawan `json:"karyawan"`
}

type RekeningKaryawan struct {
	Id string `json:"id" validate:"required"`
}

func (c *UpdateRekeningRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)

	if err != nil {
		return err
	}

	return nil
}
