package dtos

import "time"

type TrainingResponse struct {
	CreatedDate time.Time  `json:"created_date"`
	UpdatedDate time.Time  `json:"updated_date"`
	DeletedDate *time.Time `json:"deleted_date"`
	ID          int        `json:"id"`
	Tema        string     `json:"tema"`
	Pengajar    string     `json:"pengajar"`
}
