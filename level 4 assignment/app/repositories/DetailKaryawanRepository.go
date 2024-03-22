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

type DetailKaryawanRepository struct {
	DB *gorm.DB
}

func NewDetailKaryawanRepository() *DetailKaryawanRepository {
	return &DetailKaryawanRepository{
		DB: config.GetDB(),
	}
}

func (repo *DetailKaryawanRepository) Create(detailKaryawan *models.DetailKaryawanEntity) (*models.DetailKaryawanEntity, error) {
	result := repo.DB.Create(detailKaryawan)
	if result.Error != nil {
		return nil, errors.New("failed to create detail karyawan")
	}
	return detailKaryawan, nil
}

func (repo *DetailKaryawanRepository) FindByID(id string) (*models.DetailKaryawanEntity, error) {
	var detail models.DetailKaryawanEntity
	result := repo.DB.Where("detail_karyawan_id = ? and deleted_date is null", id).Find(&detail)
	if result.Error != nil {
		return nil, errors.New("failed to find detail karyawan")
	}

	if detail.DetailKaryawanID == 0 {
		return nil, errors.New("detail karyawan not found")
	}

	return &detail, nil
}

func (repo *DetailKaryawanRepository) FindByIDIgnoreDelete(id string) (*models.DetailKaryawanEntity, error) {
	var detail models.DetailKaryawanEntity
	result := repo.DB.Where("detail_karyawan_id = ?", id).Find(&detail)
	if result.Error != nil {
		return nil, errors.New("failed to find detail karyawan")
	}

	if detail.DetailKaryawanID == 0 {
		return nil, errors.New("detail karyawan not found")
	}

	return &detail, nil
}

func (repo *DetailKaryawanRepository) Update(detailKaryawan dtos.UpdateDetailKaryawan) (*models.DetailKaryawanEntity, error) {
	tNow := time.Now()
	if err := repo.DB.Model(models.DetailKaryawanEntity{}).Where("detail_karyawan_id = ?", detailKaryawan.Id).Updates(models.DetailKaryawanEntity{
		NIK:         detailKaryawan.Nik,
		NPWP:        detailKaryawan.Npwp,
		UpdatedDate: tNow,
	}).Error; err != nil {
		return nil, errors.New("failed to update detail karyawan")
	}
	var updatedDetail models.DetailKaryawanEntity
	if err := repo.DB.Where("detail_karyawan_id = ?", detailKaryawan.Id).First(&updatedDetail).Error; err != nil {
		return nil, errors.New("failed to fetch updated detail karyawan: " + err.Error())
	}

	return &updatedDetail, nil
}

func (repo *DetailKaryawanRepository) Delete(id string) error {
	detail := models.DetailKaryawanEntity{}

	if err := repo.DB.Clauses(clause.Returning{}).Delete(&detail, "detail_karyawan_id", id).Error; err != nil {
		return errors.New("failed to delete detail karyawan")
	}

	if detail.DetailKaryawanID == 0 {
		return errors.New("detail karyawan not found")
	}

	return nil
}

func (repo *DetailKaryawanRepository) DeleteTemp(id string) error {
	tNow := time.Now()
	if err := repo.DB.Model(models.DetailKaryawanEntity{}).Where("detail_karyawan_id = ?", id).Updates(models.DetailKaryawanEntity{
		DeletedDate: &tNow,
	}).Error; err != nil {
		repo.DB.Rollback()
		return errors.New("failed to update delete detail karyawan")
	}

	return nil
}
