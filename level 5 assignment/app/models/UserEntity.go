package models

import (
	"time"

	"gorm.io/gorm"
)

type UserEntity struct {
	Id             uint           `gorm:"primaryKey"`
	Username       string         `gorm:"index;unique;column:username;required"`
	Name           string         `gorm:"column:name;required"`
	Email          string         `gorm:"index;unique;column:email;required"`
	Password       string         `gorm:"column:password;required"`
	CreatedDate    time.Time      `gorm:"column:created_date;autoCreateTime:true;not null"`
	UpdatedDate    *time.Time     `gorm:"column:updated_date;autoUpdateTime:true"`
	DeletedDate    gorm.DeletedAt `gorm:"index;column:deleted_date;softDelete:true"`
	Otp            string         `gorm:"column:otp"`
	OtpExpiredDate time.Time      `gorm:"column:otp_expired_date"`
	Approved       bool           `gorm:"column:approved;default:false"`
}

func (c *UserEntity) TableName() string {
	return "useraccount"
}
