package dtos

import (
	"github.com/go-playground/validator/v10"
)

type SaveTrainingRequest struct {
	Tema     string `json:"tema"`
	Pengajar string `json:"pengajar"`
}

type UpdateTrainingRequest struct {
	Id       string `json:"id" validate:"required"`
	Tema     string `json:"tema"`
	Pengajar string `json:"pengajar"`
}

func (c *UpdateTrainingRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)

	if err != nil {
		return err
	}

	return nil
}
