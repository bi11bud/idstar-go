package repositories

import (
	"errors"
	"time"

	"gorm.io/gorm"
	config "idstar.com/app/configs"
	"idstar.com/app/models"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		DB: config.GetDB(),
	}
}

func (repo *UserRepository) FindByID(id string) (*models.UserEntity, error) {
	var user models.UserEntity
	result := repo.DB.Where("id = ?", id).Find(&user)
	if result.Error != nil {
		return nil, errors.New("failed to find user")
	}
	return &user, nil
}

func (repo *UserRepository) Create(user *models.UserEntity) (*models.UserEntity, error) {
	result := repo.DB.Create(user)
	if result.Error != nil {
		return nil, errors.New("failed to register user")
	}

	return user, nil
}

func (repo *UserRepository) UpdateOtp(userId string, otp string) (*models.UserEntity, error) {
	tNow := time.Now()
	tNowPlus30Min := tNow.Add(30 * time.Minute)
	if err := repo.DB.Model(models.UserEntity{}).Where("email = ? or username = ?", userId, userId).Updates(models.UserEntity{
		Otp:            otp,
		OtpExpiredDate: tNowPlus30Min,
		UpdatedDate:    &tNow,
	}).Error; err != nil {
		return nil, errors.New("failed to update otp user")
	}
	var user models.UserEntity
	if err := repo.DB.Where("email = ? or username = ?", userId, userId).First(&user).Error; err != nil {
		return nil, errors.New("failed to fetch updated otp user: " + err.Error())
	}

	return &user, nil
}

func (repo *UserRepository) UpdateStatus(userId string) error {
	tNow := time.Now()
	if err := repo.DB.Model(models.UserEntity{}).Where("id = ?", userId).Updates(models.UserEntity{
		Approved:    true,
		UpdatedDate: &tNow,
	}).Error; err != nil {
		return errors.New("failed to update user")
	}
	return nil
}

func (repo *UserRepository) UpdatePassword(userId string, password string) error {
	tNow := time.Now()
	if err := repo.DB.Model(models.UserEntity{}).Where("id = ?", userId).Updates(models.UserEntity{
		Password:    password,
		UpdatedDate: &tNow,
	}).Error; err != nil {
		return errors.New("failed to update password user")
	}
	return nil
}

func (repo *UserRepository) FindByUsernameOrEmail(username string) (*models.UserEntity, error) {
	var user models.UserEntity
	result := repo.DB.Where("username = ? OR email = ?", username, username).Find(&user)
	if result.Error != nil {
		return nil, errors.New("failed to find user")
	}
	return &user, nil
}
