package repositories

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	config "idstar.com/app/configs"
	dtos "idstar.com/app/dtos/karyawan"
	"idstar.com/app/models"
)

type KaryawanRepository struct {
	DB *gorm.DB
}

func NewKaryawanRepository() *KaryawanRepository {
	return &KaryawanRepository{
		DB: config.GetDB(),
	}
}

func (repo *KaryawanRepository) Create(karyawanEntity *models.KaryawanEntity) (*models.KaryawanEntity, error) {
	result := repo.DB.Create(karyawanEntity)
	if result.Error != nil {
		repo.DB.Rollback()
		return nil, errors.New("failed to create Karyawan")
	}
	repo.DB.Commit()
	return karyawanEntity, nil
}

func (repo *KaryawanRepository) FindByID(id string) (*models.KaryawanEntity, error) {
	var karyawan models.KaryawanEntity
	result := repo.DB.Where("karyawan_id = ? and deleted_date is null", id).Find(&karyawan)
	if result.Error != nil {
		return nil, errors.New("failed to find Karyawan")
	}

	if karyawan.KaryawanID == 0 {
		return nil, errors.New("karyawan not found")
	}

	return &karyawan, nil
}

func (repo *KaryawanRepository) FindByIDIgnoreDelete(id string) (*models.KaryawanEntity, error) {
	var karyawan models.KaryawanEntity
	result := repo.DB.Where("karyawan_id = ?", id).Find(&karyawan)
	if result.Error != nil {
		return nil, errors.New("failed to find Karyawan")
	}

	if karyawan.KaryawanID == 0 {
		return nil, errors.New("karyawan not found")
	}

	return &karyawan, nil
}

func (repo *KaryawanRepository) FindAll() ([]models.KaryawanEntity, error) {
	var karyawan []models.KaryawanEntity
	result := repo.DB.Find(&karyawan)
	if result.Error != nil {
		return nil, errors.New("failed to find Karyawan")
	}

	return karyawan, nil
}

func (repo *KaryawanRepository) Update(req dtos.UpdateKaryawanRequest) (*models.KaryawanEntity, error) {
	tNow := time.Now()
	layout := "2006-01-02"
	parsedTime, err := time.Parse(layout, req.Dob)
	if err != nil {
		return nil, errors.New("invalid format dob (yyy-mm-dd)")
	}

	currentDate := time.Now().Truncate(24 * time.Hour)
	if parsedTime.After(currentDate) {
		return nil, errors.New("dob cannot be greater than current date")
	}

	if err := repo.DB.Model(models.KaryawanEntity{}).Where("karyawan_id = ?", req.Id).Updates(models.KaryawanEntity{
		Nama:        req.Nama,
		Alamat:      req.Alamat,
		Dob:         parsedTime,
		Status:      req.Status,
		UpdatedDate: tNow,
	}).Error; err != nil {
		repo.DB.Rollback()
		return nil, errors.New("failed to update Karyawan")
	}
	var updatedTraining models.KaryawanEntity
	if err := repo.DB.Where("karyawan_id = ?", req.Id).First(&updatedTraining).Error; err != nil {
		return nil, errors.New("failed to fetch updated Karyawan: " + err.Error())
	}
	repo.DB.Commit()
	return &updatedTraining, nil
}

func (repo *KaryawanRepository) Delete(id string) error {
	karyawan := models.KaryawanEntity{}

	if err := repo.DB.Clauses(clause.Returning{}).Delete(&karyawan, "karyawan_id", id).Error; err != nil {
		return errors.New("failed to delete Karyawan")
	}

	if karyawan.KaryawanID == 0 {
		return errors.New("karyawan not found")
	}

	return nil
}

func (repo *KaryawanRepository) DeleteTemp(id string) error {
	tNow := time.Now()
	if err := repo.DB.Model(models.KaryawanEntity{}).Where("karyawan_id = ?", id).Updates(models.KaryawanEntity{
		Status:      "deleted",
		DeletedDate: &tNow,
	}).Error; err != nil {
		repo.DB.Rollback()
		return errors.New("failed to update delete Karyawan")
	}

	repo.DB.Commit()
	return nil
}
