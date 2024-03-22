package dtos

import "github.com/go-playground/validator/v10"

type UpdateKaryawanRequest struct {
	Id             string               `json:"id" validate:"required"`
	Nama           string               `json:"nama"`
	Dob            string               `json:"dob"`
	Status         string               `json:"status"`
	Alamat         string               `json:"alamat"`
	DetailKaryawan UpdateDetailKaryawan `json:"detailKaryawan"`
}

type UpdateDetailKaryawan struct {
	Id   string `json:"id" validate:"required"`
	Nik  string `json:"nik"`
	Npwp string `json:"npwp"`
}

type DeleteRequest struct {
	Id string `json:"id" validate:"required"`
}

func (c *UpdateKaryawanRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)

	if err != nil {
		return err
	}

	return nil
}

func (c *UpdateDetailKaryawan) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)

	if err != nil {
		return err
	}

	return nil
}
