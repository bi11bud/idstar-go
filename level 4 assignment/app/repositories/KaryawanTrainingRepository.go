package repositories

import (
	"errors"
	"strconv"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	config "idstar.com/app/configs"
	dtos "idstar.com/app/dtos/karyawan-training"
	"idstar.com/app/models"
)

type KaryawanTrainingRepository struct {
	DB *gorm.DB
}

func NewKaryawanTrainingRepository() *KaryawanTrainingRepository {
	return &KaryawanTrainingRepository{
		DB: config.GetDB(),
	}
}

func (repo *KaryawanTrainingRepository) Create(karyawanTrainingEntity *models.KaryawanTrainingEntity) (*models.KaryawanTrainingEntity, error) {
	result := repo.DB.Create(karyawanTrainingEntity)
	if result.Error != nil {
		return nil, errors.New("failed to create karyawan training")
	}
	return karyawanTrainingEntity, nil
}

func (repo *KaryawanTrainingRepository) FindByID(id string) (*models.KaryawanTrainingEntity, error) {
	var karyawan models.KaryawanTrainingEntity
	result := repo.DB.Where("karyawan_training_id = ? and deleted_date is null", id).Find(&karyawan)
	if result.Error != nil {
		return nil, errors.New("failed to find Karyawan Training")
	}

	if karyawan.KaryawanID == 0 {
		return nil, errors.New("karyawan training not found")
	}

	return &karyawan, nil
}

func (repo *KaryawanTrainingRepository) Update(req dtos.UpdateKaryawanTrainingRequest) (*models.KaryawanTrainingEntity, error) {
	tNow := time.Now()
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, req.Tanggal)
	if err != nil {
		return nil, errors.New("invalid format tanggal (yyy-mm-dd hh:mm:ss)")
	}

	currentDate := time.Now().Truncate(24 * time.Hour)
	if parsedTime.Before(currentDate) {
		return nil, errors.New("tanggal cannot be lower than current date")
	}

	karyawanId, err := strconv.ParseUint(req.Karyawan.Id, 10, 64)
	if err != nil {
		return nil, err
	}
	trainingId, err := strconv.ParseUint(req.Training.Id, 10, 64)
	if err != nil {
		return nil, err
	}

	if err := repo.DB.Model(models.KaryawanTrainingEntity{}).Where("karyawan_training_id = ?", req.Id).Updates(models.KaryawanTrainingEntity{
		KaryawanID:  uint(karyawanId),
		TrainingID:  uint(trainingId),
		Tanggal:     parsedTime,
		UpdatedDate: tNow,
	}).Error; err != nil {
		return nil, errors.New("failed to update Karyawan")
	}
	var updatedTraining models.KaryawanTrainingEntity
	if err := repo.DB.Where("karyawan_id = ?", req.Id).First(&updatedTraining).Error; err != nil {
		return nil, errors.New("failed to fetch updated Karyawan: " + err.Error())
	}
	return &updatedTraining, nil
}

func (repo *KaryawanTrainingRepository) FindAll() ([]models.KaryawanTrainingEntity, error) {
	var karyawanTraining []models.KaryawanTrainingEntity
	result := repo.DB.Find(&karyawanTraining)
	if result.Error != nil {
		return nil, errors.New("failed to find Karyawan Training")
	}

	return karyawanTraining, nil
}

func (repo *KaryawanTrainingRepository) Delete(id string) error {
	karyawan := models.KaryawanTrainingEntity{}

	if err := repo.DB.Clauses(clause.Returning{}).Delete(&karyawan, "karyawan_training_id", id).Error; err != nil {
		return errors.New("failed to delete Karyawan Training")
	}

	if karyawan.KaryawanID == 0 {
		return errors.New("karyawan training not found")
	}

	return nil
}

func (repo *KaryawanTrainingRepository) DeleteTemp(id string) error {
	tNow := time.Now()
	if err := repo.DB.Model(models.KaryawanTrainingEntity{}).Where("karyawan_training_id = ?", id).Updates(models.KaryawanTrainingEntity{
		DeletedDate: &tNow,
	}).Error; err != nil {
		return errors.New("failed to update delete Karyawan Training")
	}

	return nil
}

func (repo *KaryawanTrainingRepository) DeleteTempByKaryawanId(id string) error {
	tNow := time.Now()
	if err := repo.DB.Model(models.KaryawanTrainingEntity{}).Where("karyawan_id = ?", id).Updates(models.KaryawanTrainingEntity{
		DeletedDate: &tNow,
	}).Error; err != nil {
		repo.DB.Rollback()
		return errors.New("failed to update delete Karyawan Training")
	}

	return nil
}

func (repo *KaryawanTrainingRepository) DeleteTempByTrainingId(id string) error {
	tNow := time.Now()
	if err := repo.DB.Model(models.KaryawanTrainingEntity{}).Where("training_id = ?", id).Updates(models.KaryawanTrainingEntity{
		DeletedDate: &tNow,
	}).Error; err != nil {
		repo.DB.Rollback()
		return errors.New("failed to update delete Karyawan Training")
	}

	return nil
}
