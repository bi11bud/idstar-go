package models

import (
	"time"
)

// type KaryawanTrainingEntity struct {
// 	Id          int    `gorm:"column:id;primaryKey;type:autoIncrement"`
// 	Tanggal     string `gorm:"column:tanggal"`
// 	Karyawan    KaryawanEntity
// 	KaryawanID  int `gorm:"not null"`
// 	Training    TrainingEntity
// 	TrainingID  int            `gorm:"not null"`
// 	CreatedDate time.Time      `gorm:"column:created_date;autoCreateTime:true;not null"`
// 	UpdatedDate *time.Time     `gorm:"column:updated_date;autoUpdateTime:true"`
// 	DeletedDate gorm.DeletedAt `gorm:"index;column:deleted_date;softDelete:true"`
// }

type KaryawanTrainingEntity struct {
	KaryawanTrainingID uint `gorm:"primaryKey"`
	Tanggal            time.Time
	// Karyawan           KaryawanEntity
	KaryawanID uint `gorm:"foreignkey:KaryawanID"`
	// Training           TrainingEntity
	TrainingID  uint       `gorm:"foreignkey:TrainingID"`
	CreatedDate time.Time  `gorm:"autoCreateTime:true;not null"`
	UpdatedDate time.Time  `gorm:"autoUpdateTime:true"`
	DeletedDate *time.Time `gorm:"softDelete:true"`
}

func (c *KaryawanTrainingEntity) TableName() string {
	return "karyawan_training"
}
