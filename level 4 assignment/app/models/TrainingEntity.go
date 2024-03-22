package models

import (
	"time"
)

// type TrainingEntity struct {
// 	Id       int    `gorm:"column:id;primaryKey;type:autoIncrement"`
// 	Pengajar string `gorm:"column:pengajar"`
// 	Tema     string `gorm:"column:tema"`
// 	// KaryawanTrainingId KaryawanTrainingEntity
// 	CreatedDate time.Time      `gorm:"column:created_date;autoCreateTime:true;not null"`
// 	UpdatedDate *time.Time     `gorm:"column:updated_date;autoUpdateTime:true"`
// 	DeletedDate gorm.DeletedAt `gorm:"index;column:deleted_date;softDelete:true"`
// }

type TrainingEntity struct {
	TrainingID  uint `gorm:"primaryKey"`
	Pengajar    string
	Tema        string
	CreatedDate time.Time  `gorm:"autoCreateTime:true;not null"`
	UpdatedDate time.Time  `gorm:"autoUpdateTime:true"`
	DeletedDate *time.Time `gorm:"softDelete:true"`
}

func (c *TrainingEntity) TableName() string {
	return "training"
}
