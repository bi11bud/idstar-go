package services

import (
	"errors"
	"strconv"
	"time"

	dtos "idstar.com/app/dtos/karyawan"
	resp "idstar.com/app/dtos/response"
	"idstar.com/app/models"
	"idstar.com/app/repositories"
)

var dobFormat string = "2006-01-02"
var tanggalFormat string = "2006-01-02 15:04:05"

type KaryawanService struct {
	karyawanRepository         repositories.KaryawanRepository
	detailKaryawanRepository   repositories.DetailKaryawanRepository
	rekeningRepository         repositories.RekeningRepository
	karyawanTrainingRepository repositories.KaryawanTrainingRepository
}

func NewKaryawanService(karyawanRepository repositories.KaryawanRepository,
	detailKaryawanRepository repositories.DetailKaryawanRepository,
	rekeningRepository repositories.RekeningRepository,
	karyawanTrainingRepository repositories.KaryawanTrainingRepository) *KaryawanService {
	return &KaryawanService{
		karyawanRepository,
		detailKaryawanRepository,
		rekeningRepository,
		karyawanTrainingRepository,
	}
}

func (u *KaryawanService) SaveKaryawan(req *dtos.SaveKaryawanRequest) (*dtos.KaryawanResponse, error) {
	dob, err := time.Parse(dobFormat, req.Dob)
	if err != nil {
		return nil, errors.New("invalid format dob (yyy-mm-dd)")
	}

	currentDate := time.Now().Truncate(24 * time.Hour)
	if dob.After(currentDate) {
		return nil, errors.New("dob cannot be greater than current date")
	}

	detailKaryawan := models.DetailKaryawanEntity{
		NIK:  req.DetailKaryawan.Nik,
		NPWP: req.DetailKaryawan.Npwp,
	}

	detail, err := u.detailKaryawanRepository.Create(&detailKaryawan)
	if detail == nil {
		return nil, err
	}

	karyawan := models.KaryawanEntity{
		Nama:             req.Nama,
		Alamat:           req.Alamat,
		Dob:              dob,
		Status:           req.Status,
		DetailKaryawanId: detail.DetailKaryawanID,
	}

	respKaryawan, err := u.karyawanRepository.Create(&karyawan)
	if respKaryawan == nil {
		return nil, err
	}

	result := dtos.KaryawanResponse{
		ID:          int(respKaryawan.KaryawanID),
		Nama:        respKaryawan.Nama,
		Alamat:      respKaryawan.Alamat,
		Dob:         respKaryawan.Dob.Format(dobFormat),
		Status:      respKaryawan.Status,
		CreatedDate: respKaryawan.CreatedDate,
		UpdatedDate: respKaryawan.UpdatedDate,
		DeletedDate: respKaryawan.DeletedDate,
		DetailKaryawan: dtos.DetailKaryawanResponse{
			ID:          int(detail.DetailKaryawanID),
			Nik:         detail.NIK,
			Npwp:        detail.NPWP,
			CreatedDate: detail.CreatedDate,
			UpdatedDate: detail.UpdatedDate,
			DeletedDate: detail.DeletedDate,
		},
	}

	return &result, nil
}

func (u *KaryawanService) GetById(id string) (*dtos.KaryawanResponse, error) {
	respKaryawan, err := u.karyawanRepository.FindByID(id)
	if respKaryawan == nil {
		return nil, err
	}

	detailkaryawanId := strconv.FormatUint(uint64(respKaryawan.DetailKaryawanId), 10)
	detail, err := u.detailKaryawanRepository.FindByID(detailkaryawanId)
	if detail == nil {
		return nil, err
	}

	result := dtos.KaryawanResponse{
		ID:          int(respKaryawan.KaryawanID),
		Nama:        respKaryawan.Nama,
		Alamat:      respKaryawan.Alamat,
		Dob:         respKaryawan.Dob.Format(dobFormat),
		Status:      respKaryawan.Status,
		CreatedDate: respKaryawan.CreatedDate,
		UpdatedDate: respKaryawan.UpdatedDate,
		DeletedDate: respKaryawan.DeletedDate,
		DetailKaryawan: dtos.DetailKaryawanResponse{
			ID:          int(detail.DetailKaryawanID),
			Nik:         detail.NIK,
			Npwp:        detail.NPWP,
			CreatedDate: detail.CreatedDate,
			UpdatedDate: detail.UpdatedDate,
			DeletedDate: detail.DeletedDate,
		},
	}

	return &result, nil
}

func (u *KaryawanService) FindAll() (any, error) {
	karyawanList, err := u.karyawanRepository.FindAll()
	if err != nil {
		return nil, err
	}

	karyawanListResponse := []dtos.KaryawanResponse{}

	for _, karyawan := range karyawanList {
		detailkaryawanId := strconv.FormatUint(uint64(karyawan.DetailKaryawanId), 10)
		detail, err := u.detailKaryawanRepository.FindByIDIgnoreDelete(detailkaryawanId)
		if detail == nil {
			return nil, err
		}

		karyawanResponse := dtos.KaryawanResponse{
			ID:          int(karyawan.KaryawanID),
			Nama:        karyawan.Nama,
			Alamat:      karyawan.Alamat,
			Dob:         karyawan.Dob.Format(dobFormat),
			Status:      karyawan.Status,
			CreatedDate: karyawan.CreatedDate,
			UpdatedDate: karyawan.UpdatedDate,
			DeletedDate: karyawan.DeletedDate,
			DetailKaryawan: dtos.DetailKaryawanResponse{
				ID:          int(detail.DetailKaryawanID),
				Nik:         detail.NIK,
				Npwp:        detail.NPWP,
				CreatedDate: detail.CreatedDate,
				UpdatedDate: detail.UpdatedDate,
				DeletedDate: detail.DeletedDate,
			},
		}

		karyawanListResponse = append(karyawanListResponse, karyawanResponse)

	}

	result := resp.Content{
		Data:             karyawanListResponse,
		Pageable:         resp.Pageable{Offset: 0, PageNumber: 0, PageSize: len(karyawanList), Unpaged: true, Paged: false},
		TotalElements:    len(karyawanList),
		TotalPages:       1,
		Size:             len(karyawanList),
		Number:           1,
		Sort:             resp.Sort{Empty: false, Sorted: false, Unsorted: false},
		First:            true,
		NumberOfElements: len(karyawanList),
		Empty:            len(karyawanList) == 0,
	}

	return result, nil
}

func (u *KaryawanService) UpdateKaryawan(req dtos.UpdateKaryawanRequest) (*dtos.KaryawanResponse, error) {
	karyawan, err := u.karyawanRepository.FindByID(req.Id)
	if karyawan == nil {
		return nil, err
	}

	if strconv.Itoa(int(karyawan.DetailKaryawanId)) != req.DetailKaryawan.Id {
		return nil, errors.New("detail karyawan Id not match")
	}

	detailReq := dtos.UpdateDetailKaryawan{
		Id:   req.DetailKaryawan.Id,
		Nik:  req.DetailKaryawan.Nik,
		Npwp: req.DetailKaryawan.Npwp,
	}

	detail, err := u.detailKaryawanRepository.Update(detailReq)
	if err != nil {
		return nil, err
	}

	respKaryawan, err := u.karyawanRepository.Update(req)
	if err != nil {
		return nil, err
	}

	result := dtos.KaryawanResponse{
		ID:          int(respKaryawan.KaryawanID),
		Nama:        respKaryawan.Nama,
		Alamat:      respKaryawan.Alamat,
		Dob:         respKaryawan.Dob.Format(dobFormat),
		Status:      respKaryawan.Status,
		CreatedDate: respKaryawan.CreatedDate,
		UpdatedDate: respKaryawan.UpdatedDate,
		DeletedDate: respKaryawan.DeletedDate,
		DetailKaryawan: dtos.DetailKaryawanResponse{
			ID:          int(detail.DetailKaryawanID),
			Nik:         detail.NIK,
			Npwp:        detail.NPWP,
			CreatedDate: detail.CreatedDate,
			UpdatedDate: detail.UpdatedDate,
			DeletedDate: detail.DeletedDate,
		},
	}

	return &result, nil
}

func (u *KaryawanService) DeleteById(id string) error {
	karyawan, err := u.karyawanRepository.FindByID(id)
	if karyawan == nil {
		return err
	}

	err = u.detailKaryawanRepository.DeleteTemp(strconv.Itoa(int(karyawan.DetailKaryawanId)))
	if err != nil {
		return err
	}

	err = u.rekeningRepository.DeleteTempByKaryawanId(strconv.Itoa(int(karyawan.KaryawanID)))
	if err != nil {
		return err
	}

	err = u.karyawanTrainingRepository.DeleteTempByKaryawanId(strconv.Itoa(int(karyawan.KaryawanID)))
	if err != nil {
		return err
	}

	return u.karyawanRepository.DeleteTemp(id)
}
