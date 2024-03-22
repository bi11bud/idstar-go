package repositories

import (
	"errors"
	"strconv"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	config "idstar.com/app/configs"
	dtos "idstar.com/app/dtos/rekening"
	"idstar.com/app/models"
)

type RekeningRepository struct {
	DB *gorm.DB
}

func NewRekeningRepository() *RekeningRepository {
	return &RekeningRepository{
		DB: config.GetDB(),
	}
}

func (repo *RekeningRepository) Create(rekeningEntity *models.RekeningEntity) (*models.RekeningEntity, error) {
	result := repo.DB.Create(rekeningEntity)
	if result.Error != nil {
		return nil, errors.New("failed to create Rekening")
	}
	return rekeningEntity, nil
}

func (repo *RekeningRepository) FindByID(id string) (*models.RekeningEntity, error) {
	var rekening models.RekeningEntity
	result := repo.DB.Where("rekening_id = ? and deleted_date is null", id).Find(&rekening)
	if result.Error != nil {
		return nil, errors.New("failed to find Rekening")
	}

	if rekening.RekeningID == 0 {
		return nil, errors.New("rekening not found")
	}

	return &rekening, nil
}

func (repo *RekeningRepository) FindAll() ([]models.RekeningEntity, error) {
	var rekening []models.RekeningEntity
	result := repo.DB.Find(&rekening)
	if result.Error != nil {
		return nil, errors.New("failed to find Rekening")
	}

	return rekening, nil
}

func (repo *RekeningRepository) Update(req dtos.UpdateRekeningRequest) (*models.RekeningEntity, error) {
	tNow := time.Now()
	karyawanId, err := strconv.ParseUint(req.Karyawan.Id, 10, 64)
	if err != nil {
		return nil, errors.New("failed to parse karyawan id")
	}

	if err := repo.DB.Model(models.RekeningEntity{}).Where("rekening_id = ?", req.Id).Updates(models.RekeningEntity{
		Jenis:       req.Jenis,
		Rekening:    req.Rekening,
		Nama:        req.Nama,
		KaryawanID:  uint(karyawanId),
		UpdatedDate: tNow,
	}).Error; err != nil {
		return nil, errors.New("failed to update Rekening")
	}
	var rekening models.RekeningEntity
	if err := repo.DB.Where("rekening_id = ?", req.Id).First(&rekening).Error; err != nil {
		return nil, errors.New("failed to fetch updated rekening: " + err.Error())
	}
	return &rekening, nil
}

func (repo *RekeningRepository) Delete(id string) error {
	rekening := models.RekeningEntity{}

	if err := repo.DB.Clauses(clause.Returning{}).Delete(&rekening, "rekening_id", id).Error; err != nil {
		return errors.New("failed to delete Rekening")
	}

	if rekening.RekeningID == 0 {
		return errors.New("rekening not found")
	}

	return nil
}

func (repo *RekeningRepository) DeleteTemp(id string) error {
	tNow := time.Now()
	if err := repo.DB.Model(models.RekeningEntity{}).Where("rekening_id = ?", id).Updates(models.RekeningEntity{
		DeletedDate: &tNow,
	}).Error; err != nil {
		return errors.New("failed to update delete rekening")
	}

	return nil
}

func (repo *RekeningRepository) DeleteTempByKaryawanId(id string) error {
	tNow := time.Now()
	if err := repo.DB.Model(models.RekeningEntity{}).Where("karyawan_id = ?", id).Updates(models.RekeningEntity{
		DeletedDate: &tNow,
	}).Error; err != nil {
		repo.DB.Rollback()
		return errors.New("failed to update delete rekening")
	}

	return nil
}
