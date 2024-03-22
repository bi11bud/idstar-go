package dtos

import (
	"github.com/go-playground/validator/v10"
)

type SaveKaryawanRequest struct {
	Nama           string             `json:"nama"`
	Dob            string             `json:"dob"`
	Status         string             `json:"status"`
	Alamat         string             `json:"alamat"`
	DetailKaryawan SaveDetailKaryawan `json:"detailKaryawan"`
}

type SaveDetailKaryawan struct {
	Nik  string `json:"nik"`
	Npwp string `json:"npwp"`
}

func (c *SaveKaryawanRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)

	// dateString := c.Dob
	// var rxPat = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	// if !rxPat.MatchString(dateString) {

	// }

	if err != nil {
		return err
	}

	return nil
}

func (c *SaveDetailKaryawan) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)

	if err != nil {
		return err
	}

	return nil
}
