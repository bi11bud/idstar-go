package models

import (
	"time"
)

// type DetailKaryawanEntity struct {
// 	Id          int            `gorm:"column:id;primaryKey;type:autoIncrement"`
// 	NIK         string         `gorm:"column:nik"`
// 	NPWP        string         `gorm:"column:npwp"`
// 	CreatedDate time.Time      `gorm:"column:created_date;autoCreateTime:true;not null"`
// 	UpdatedDate *time.Time     `gorm:"column:updated_date;autoUpdateTime:true"`
// 	DeletedDate gorm.DeletedAt `gorm:"index;column:deleted_date;softDelete:true"`
// }

type DetailKaryawanEntity struct {
	DetailKaryawanID uint `gorm:"column:detail_karyawan_id;primaryKey"`
	NIK              string
	NPWP             string
	CreatedDate      time.Time  `gorm:"autoCreateTime:true;not null"`
	UpdatedDate      time.Time  `gorm:"autoUpdateTime:true"`
	DeletedDate      *time.Time `gorm:"softDelete:true"`
}

func (c *DetailKaryawanEntity) TableName() string {
	return "detail_karyawan"
}
