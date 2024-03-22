package models

import (
	"time"
)

// type KaryawanEntity struct {
// 	Id               int       `gorm:"column:id;primaryKey;type:autoIncrement"`
// 	Alamat           string    `gorm:"column:alamat"`
// 	Dob              time.Time `gorm:"column:dob;"`
// 	Nama             string    `gorm:"column:nama"`
// 	Status           string    `gorm:"column:status"`
// 	DetailKaryawan   DetailKaryawanEntity
// 	DetailKaryawanId int            `gorm:"not null"`
// 	CreatedDate      time.Time      `gorm:"column:created_date;autoCreateTime:true;not null"`
// 	UpdatedDate      *time.Time     `gorm:"column:updated_date;autoUpdateTime:true"`
// 	DeletedDate      gorm.DeletedAt `gorm:"index;column:deleted_date;softDelete:true"`
// }

type KaryawanEntity struct {
	KaryawanID uint `gorm:"primaryKey"`
	Alamat     string
	Dob        time.Time `gorm:"type:date;format:yyyy-mm-dd"`
	Nama       string
	Status     string
	// DetailKaryawan   DetailKaryawanEntity
	DetailKaryawanId uint       `gorm:"foreignkey:DetailKaryawanID"`
	CreatedDate      time.Time  `gorm:"autoCreateTime:true;not null"`
	UpdatedDate      time.Time  `gorm:"autoUpdateTime:true"`
	DeletedDate      *time.Time `gorm:"softDelete:true"`
}

func (c *KaryawanEntity) TableName() string {
	return "karyawan"
}
