package dtos

import (
	"github.com/go-playground/validator/v10"
)

type SaveRekeningRequest struct {
	Nama     string               `json:"nama"`
	Jenis    string               `json:"jenis"`
	Rekening string               `json:"rekening"`
	Karyawan SaveRekeningKaryawan `json:"karyawan"`
}

type SaveRekeningKaryawan struct {
	Id string `json:"id" validate:"required"`
}

func (c *SaveRekeningRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)

	if err != nil {
		return err
	}

	return nil
}
