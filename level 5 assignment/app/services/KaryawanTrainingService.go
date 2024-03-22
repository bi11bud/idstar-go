package services

import (
	"errors"
	"strconv"
	"time"

	dtosKaryawan "idstar.com/app/dtos/karyawan"
	dtos "idstar.com/app/dtos/karyawan-training"
	resp "idstar.com/app/dtos/response"
	dtosTraining "idstar.com/app/dtos/training"
	"idstar.com/app/models"
	"idstar.com/app/repositories"
)

type KaryawanTrainingService struct {
	karyawanTrainingRepository repositories.KaryawanTrainingRepository
	karyawanRepository         repositories.KaryawanRepository
	trainingRepository         repositories.TrainingRepository
	detailRepository           repositories.DetailKaryawanRepository
}

func NewKaryawanTrainingService(karyawanTrainingRepository repositories.KaryawanTrainingRepository,
	karyawanRepository repositories.KaryawanRepository,
	trainingRepository repositories.TrainingRepository,
	detailRepository repositories.DetailKaryawanRepository,
) *KaryawanTrainingService {
	return &KaryawanTrainingService{
		karyawanTrainingRepository,
		karyawanRepository,
		trainingRepository,
		detailRepository,
	}
}

func (u *KaryawanTrainingService) SaveKaryawanTraining(req *dtos.SaveKaryawanTrainingRequest) (*dtos.KaryawanTrainingResponse, error) {
	tanggal, err := time.Parse(tanggalFormat, req.Tanggal)
	if err != nil {
		return nil, errors.New("invalid format tanggal (yyy-mm-dd hh:mm:ss)")
	}

	currentDate := time.Now().Truncate(24 * time.Hour)
	if tanggal.Before(currentDate) {
		return nil, errors.New("tanggal cannot be lower than current date")
	}

	karyawan, err := u.karyawanRepository.FindByID(req.Karyawan.Id)
	if karyawan == nil {
		return nil, err
	}

	training, err := u.trainingRepository.FindByID(req.Training.Id)
	if training == nil {
		return nil, err
	}

	karyawanTrainingReq := models.KaryawanTrainingEntity{
		KaryawanID: karyawan.KaryawanID,
		TrainingID: training.TrainingID,
		Tanggal:    tanggal,
	}

	karyawanTraining, err := u.karyawanTrainingRepository.Create(&karyawanTrainingReq)
	if karyawanTraining == nil {
		return nil, err
	}

	detail, err := u.detailRepository.FindByID(strconv.Itoa(int(karyawan.DetailKaryawanId)))
	if detail == nil {
		return nil, err
	}

	result := dtos.KaryawanTrainingResponse{
		ID:          int(karyawanTraining.KaryawanTrainingID),
		Tanggal:     karyawanTraining.Tanggal,
		CreatedDate: karyawanTraining.CreatedDate,
		UpdatedDate: karyawanTraining.UpdatedDate,
		DeletedDate: karyawanTraining.DeletedDate,
		Karyawan: dtosKaryawan.KaryawanResponse{
			CreatedDate: karyawan.CreatedDate,
			UpdatedDate: karyawan.UpdatedDate,
			DeletedDate: karyawan.DeletedDate,
			ID:          int(karyawan.KaryawanID),
			Nama:        karyawan.Nama,
			Alamat:      karyawan.Alamat,
			Dob:         karyawan.Dob.Format(dobFormat),
			Status:      karyawan.Status,
			DetailKaryawan: dtosKaryawan.DetailKaryawanResponse{
				ID:          int(detail.DetailKaryawanID),
				Nik:         detail.NIK,
				Npwp:        detail.NPWP,
				CreatedDate: detail.CreatedDate,
				UpdatedDate: detail.UpdatedDate,
				DeletedDate: detail.DeletedDate,
			},
		},

		Training: dtosTraining.TrainingResponse{
			CreatedDate: training.CreatedDate,
			UpdatedDate: training.UpdatedDate,
			DeletedDate: training.DeletedDate,
			ID:          int(training.TrainingID),
			Tema:        training.Tema,
			Pengajar:    training.Pengajar,
		},
	}

	return &result, nil
}

func (u *KaryawanTrainingService) GetById(id string) (*dtos.KaryawanTrainingResponse, error) {
	karyawanTraining, err := u.karyawanTrainingRepository.FindByID(id)
	if karyawanTraining == nil {
		return nil, err
	}

	karyawanId := strconv.FormatUint(uint64(karyawanTraining.KaryawanID), 10)
	karyawan, err := u.karyawanRepository.FindByID(karyawanId)
	if karyawan == nil {
		return nil, err
	}

	detailId := strconv.FormatUint(uint64(karyawan.DetailKaryawanId), 10)
	detail, err := u.detailRepository.FindByID(detailId)
	if detail == nil {
		return nil, err
	}

	trainingId := strconv.FormatUint(uint64(karyawanTraining.TrainingID), 10)
	training, err := u.trainingRepository.FindByID(trainingId)
	if training == nil {
		return nil, err
	}

	result := dtos.KaryawanTrainingResponse{
		ID:          int(karyawanTraining.KaryawanTrainingID),
		Tanggal:     karyawanTraining.Tanggal,
		CreatedDate: karyawanTraining.CreatedDate,
		UpdatedDate: karyawanTraining.UpdatedDate,
		DeletedDate: karyawanTraining.DeletedDate,
		Karyawan: dtosKaryawan.KaryawanResponse{
			CreatedDate: karyawan.CreatedDate,
			UpdatedDate: karyawan.UpdatedDate,
			DeletedDate: karyawan.DeletedDate,
			ID:          int(karyawan.KaryawanID),
			Nama:        karyawan.Nama,
			Alamat:      karyawan.Alamat,
			Dob:         karyawan.Dob.Format(dobFormat),
			Status:      karyawan.Status,
			DetailKaryawan: dtosKaryawan.DetailKaryawanResponse{
				ID:          int(detail.DetailKaryawanID),
				Nik:         detail.NIK,
				Npwp:        detail.NPWP,
				CreatedDate: detail.CreatedDate,
				UpdatedDate: detail.UpdatedDate,
				DeletedDate: detail.DeletedDate,
			},
		},
		Training: dtosTraining.TrainingResponse{
			CreatedDate: training.CreatedDate,
			UpdatedDate: training.UpdatedDate,
			DeletedDate: training.DeletedDate,
			ID:          int(training.TrainingID),
			Tema:        training.Tema,
			Pengajar:    training.Pengajar,
		},
	}

	return &result, nil
}

func (u *KaryawanTrainingService) FindAll() (any, error) {
	karyawanTrainingList, err := u.karyawanTrainingRepository.FindAll()
	if err != nil {
		return nil, err
	}

	karyawanTrainingListResponse := []dtos.KaryawanTrainingResponse{}

	for _, karyawanTraining := range karyawanTrainingList {
		karyawanId := strconv.FormatUint(uint64(karyawanTraining.KaryawanID), 10)
		karyawan, err := u.karyawanRepository.FindByIDIgnoreDelete(karyawanId)
		if karyawan == nil {
			return nil, err
		}

		detailkaryawanId := strconv.FormatUint(uint64(karyawan.DetailKaryawanId), 10)
		detail, err := u.detailRepository.FindByIDIgnoreDelete(detailkaryawanId)
		if detail == nil {
			return nil, err
		}

		trainingId := strconv.FormatUint(uint64(karyawanTraining.TrainingID), 10)
		training, err := u.trainingRepository.FindByIDIgnoreDelete(trainingId)
		if training == nil {
			return nil, err
		}

		karyawanTrainingResponse := dtos.KaryawanTrainingResponse{
			ID:          int(karyawanTraining.KaryawanTrainingID),
			Tanggal:     karyawanTraining.Tanggal,
			CreatedDate: karyawanTraining.CreatedDate,
			UpdatedDate: karyawanTraining.UpdatedDate,
			DeletedDate: karyawanTraining.DeletedDate,
			Karyawan: dtosKaryawan.KaryawanResponse{
				CreatedDate: karyawan.CreatedDate,
				UpdatedDate: karyawan.UpdatedDate,
				DeletedDate: karyawan.DeletedDate,
				ID:          int(karyawan.KaryawanID),
				Nama:        karyawan.Nama,
				Alamat:      karyawan.Alamat,
				Dob:         karyawan.Dob.Format(dobFormat),
				Status:      karyawan.Status,
				DetailKaryawan: dtosKaryawan.DetailKaryawanResponse{
					ID:          int(detail.DetailKaryawanID),
					Nik:         detail.NIK,
					Npwp:        detail.NPWP,
					CreatedDate: detail.CreatedDate,
					UpdatedDate: detail.UpdatedDate,
					DeletedDate: detail.DeletedDate,
				},
			},
			Training: dtosTraining.TrainingResponse{
				CreatedDate: training.CreatedDate,
				UpdatedDate: training.UpdatedDate,
				DeletedDate: training.DeletedDate,
				ID:          int(training.TrainingID),
				Tema:        training.Tema,
				Pengajar:    training.Pengajar,
			},
		}

		karyawanTrainingListResponse = append(karyawanTrainingListResponse, karyawanTrainingResponse)

	}

	result := resp.Content{
		Data:             karyawanTrainingListResponse,
		Pageable:         resp.Pageable{Offset: 0, PageNumber: 0, PageSize: len(karyawanTrainingList), Unpaged: true, Paged: false},
		TotalElements:    len(karyawanTrainingList),
		TotalPages:       1,
		Size:             len(karyawanTrainingList),
		Number:           1,
		Sort:             resp.Sort{Empty: false, Sorted: false, Unsorted: false},
		First:            true,
		NumberOfElements: len(karyawanTrainingList),
		Empty:            len(karyawanTrainingList) == 0,
	}

	return result, nil
}

func (u *KaryawanTrainingService) UpdateKaryawanTraining(req dtos.UpdateKaryawanTrainingRequest) (*dtos.KaryawanTrainingResponse, error) {
	checkKaryawanTraining, err := u.karyawanTrainingRepository.FindByID(strconv.Itoa(req.Id))
	if checkKaryawanTraining == nil {
		return nil, err
	}

	if strconv.Itoa(int(checkKaryawanTraining.KaryawanID)) != req.Karyawan.Id {
		return nil, errors.New("karyawan Id not match")
	}

	if strconv.Itoa(int(checkKaryawanTraining.TrainingID)) != req.Training.Id {
		return nil, errors.New("training Id not match")
	}

	karyawanTraining, err := u.karyawanTrainingRepository.Update(req)
	if err != nil {
		return nil, err
	}

	karyawanId := strconv.FormatUint(uint64(karyawanTraining.KaryawanID), 10)
	karyawan, err := u.karyawanRepository.FindByID(karyawanId)
	if karyawan == nil {
		return nil, err
	}

	detailId := strconv.FormatUint(uint64(karyawan.DetailKaryawanId), 10)
	detail, err := u.detailRepository.FindByID(detailId)
	if detail == nil {
		return nil, err
	}

	trainingId := strconv.FormatUint(uint64(karyawanTraining.TrainingID), 10)
	training, err := u.trainingRepository.FindByID(trainingId)
	if training == nil {
		return nil, err
	}

	result := dtos.KaryawanTrainingResponse{
		ID:          int(karyawanTraining.KaryawanTrainingID),
		Tanggal:     karyawanTraining.Tanggal,
		CreatedDate: karyawanTraining.CreatedDate,
		UpdatedDate: karyawanTraining.UpdatedDate,
		DeletedDate: karyawanTraining.DeletedDate,
		Karyawan: dtosKaryawan.KaryawanResponse{
			CreatedDate: karyawan.CreatedDate,
			UpdatedDate: karyawan.UpdatedDate,
			DeletedDate: karyawan.DeletedDate,
			ID:          int(karyawan.KaryawanID),
			Nama:        karyawan.Nama,
			Alamat:      karyawan.Alamat,
			Dob:         karyawan.Dob.Format(dobFormat),
			Status:      karyawan.Status,
			DetailKaryawan: dtosKaryawan.DetailKaryawanResponse{
				ID:          int(detail.DetailKaryawanID),
				Nik:         detail.NIK,
				Npwp:        detail.NPWP,
				CreatedDate: detail.CreatedDate,
				UpdatedDate: detail.UpdatedDate,
				DeletedDate: detail.DeletedDate,
			},
		},
		Training: dtosTraining.TrainingResponse{
			CreatedDate: training.CreatedDate,
			UpdatedDate: training.UpdatedDate,
			DeletedDate: training.DeletedDate,
			ID:          int(training.TrainingID),
			Tema:        training.Tema,
			Pengajar:    training.Pengajar,
		},
	}

	return &result, nil
}

func (u *KaryawanTrainingService) DeleteById(id string) error {
	return u.karyawanTrainingRepository.DeleteTemp(id)
}
