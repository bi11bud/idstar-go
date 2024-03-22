package models

import (
	"time"
)

// type RekeningEntity struct {
// 	Id          int    `gorm:"column:id;primaryKey;type:autoIncrement"`
// 	Jenis       string `gorm:"column:jenis"`
// 	Rekening    string `gorm:"column:rekening;"`
// 	Nama        string `gorm:"column:nama"`
// 	Karyawan    KaryawanEntity
// 	KaryawanID  int            `gorm:"not null"`
// 	CreatedDate time.Time      `gorm:"column:created_date;autoCreateTime:true;not null"`
// 	UpdatedDate *time.Time     `gorm:"column:updated_date;autoUpdateTime:true"`
// 	DeletedDate gorm.DeletedAt `gorm:"index;column:deleted_date;softDelete:true"`
// }

type RekeningEntity struct {
	RekeningID uint `gorm:"primaryKey"`
	Jenis      string
	Rekening   string
	Nama       string
	// Karyawan    KaryawanEntity
	KaryawanID  uint       `gorm:"foreignkey:KaryawanID"`
	CreatedDate time.Time  `gorm:"autoCreateTime:true;not null"`
	UpdatedDate time.Time  `gorm:"autoUpdateTime:true"`
	DeletedDate *time.Time `gorm:"softDelete:true"`
}

func (c *RekeningEntity) TableName() string {
	return "rekening"
}
