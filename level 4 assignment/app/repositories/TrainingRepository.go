package repositories

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	config "idstar.com/app/configs"
	dtos "idstar.com/app/dtos/training"
	"idstar.com/app/models"
)

type TrainingRepository struct {
	DB *gorm.DB
}

func NewTrainingRepository() *TrainingRepository {
	return &TrainingRepository{
		DB: config.GetDB(),
	}
}

func (repo *TrainingRepository) Create(trainingEntity *models.TrainingEntity) (*models.TrainingEntity, error) {
	result := repo.DB.Create(trainingEntity)
	if result.Error != nil {
		return nil, errors.New("failed to create Training")
	}
	return trainingEntity, nil
}

func (repo *TrainingRepository) FindByID(id string) (*models.TrainingEntity, error) {
	var training models.TrainingEntity
	result := repo.DB.Where("training_id = ? and deleted_date is null", id).Find(&training)
	if result.Error != nil {
		return nil, errors.New("failed to find Training")
	}

	if training.TrainingID == 0 {
		return nil, errors.New("training not found")
	}

	return &training, nil
}

func (repo *TrainingRepository) FindByIDIgnoreDelete(id string) (*models.TrainingEntity, error) {
	var training models.TrainingEntity
	result := repo.DB.Where("training_id = ?", id).Find(&training)
	if result.Error != nil {
		return nil, errors.New("failed to find Training")
	}

	if training.TrainingID == 0 {
		return nil, errors.New("training not found")
	}

	return &training, nil
}

func (repo *TrainingRepository) FindAll() ([]models.TrainingEntity, error) {
	var training []models.TrainingEntity
	result := repo.DB.Find(&training)
	if result.Error != nil {
		return nil, errors.New("failed to find Training")
	}

	return training, nil
}

func (repo *TrainingRepository) Update(req dtos.UpdateTrainingRequest) (*models.TrainingEntity, error) {
	tNow := time.Now()
	if err := repo.DB.Model(models.TrainingEntity{}).Where("training_id = ?", req.Id).Updates(models.TrainingEntity{
		Pengajar:    req.Pengajar,
		Tema:        req.Tema,
		UpdatedDate: tNow,
	}).Error; err != nil {
		return nil, errors.New("failed to update Training")
	}
	var updatedTraining models.TrainingEntity
	if err := repo.DB.Where("training_id = ?", req.Id).First(&updatedTraining).Error; err != nil {
		return nil, errors.New("failed to fetch updated training: " + err.Error())
	}
	return &updatedTraining, nil
}

func (repo *TrainingRepository) Delete(id string) error {
	training := models.TrainingEntity{}

	if err := repo.DB.Clauses(clause.Returning{}).Delete(&training, "training_id", id).Error; err != nil {
		return errors.New("failed to delete Training")
	}

	if training.TrainingID == 0 {
		return errors.New("training not found")
	}

	return nil
}

func (repo *TrainingRepository) DeleteTemp(id string) error {
	tNow := time.Now()
	if err := repo.DB.Model(models.TrainingEntity{}).Where("training_id = ?", id).Updates(models.TrainingEntity{
		DeletedDate: &tNow,
	}).Error; err != nil {
		repo.DB.Rollback()
		return errors.New("failed to update delete training")
	}

	repo.DB.Commit()
	return nil
}
