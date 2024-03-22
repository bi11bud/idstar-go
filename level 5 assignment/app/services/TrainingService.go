package services

import (
	"strconv"

	resp "idstar.com/app/dtos/response"
	dtos "idstar.com/app/dtos/training"
	"idstar.com/app/models"
	"idstar.com/app/repositories"
)

type TrainingService struct {
	trainingRepository         repositories.TrainingRepository
	karyawanTrainingRepository repositories.KaryawanTrainingRepository
}

func NewTrainingService(trainingRepository repositories.TrainingRepository,
	karyawanTrainingRepository repositories.KaryawanTrainingRepository) *TrainingService {
	return &TrainingService{trainingRepository, karyawanTrainingRepository}
}

func (u *TrainingService) SaveTraining(detail *models.TrainingEntity) (*models.TrainingEntity, error) {
	return u.trainingRepository.Create(detail)
}

func (u *TrainingService) GetById(id string) (*models.TrainingEntity, error) {
	return u.trainingRepository.FindByID(id)
}

func (u *TrainingService) FindAll() (any, error) {

	trainingList, err := u.trainingRepository.FindAll()
	if err != nil {
		return nil, err
	}

	result := resp.Content{
		Data:             trainingList,
		Pageable:         resp.Pageable{Offset: 0, PageNumber: 0, PageSize: len(trainingList), Unpaged: true, Paged: false},
		TotalElements:    len(trainingList),
		TotalPages:       1,
		Size:             len(trainingList),
		Number:           1,
		Sort:             resp.Sort{Empty: false, Sorted: false, Unsorted: false},
		First:            true,
		NumberOfElements: len(trainingList),
		Empty:            len(trainingList) == 0,
	}

	return result, nil
}

func (u *TrainingService) UpdateTraining(req dtos.UpdateTrainingRequest) (*models.TrainingEntity, error) {
	return u.trainingRepository.Update(req)
}

func (u *TrainingService) DeleteById(id string) error {
	training, err := u.trainingRepository.FindByID(id)
	if training == nil {
		return err
	}

	err = u.karyawanTrainingRepository.DeleteTempByTrainingId(strconv.Itoa(int(training.TrainingID)))
	if err != nil {
		return err
	}

	return u.trainingRepository.DeleteTemp(id)
}
