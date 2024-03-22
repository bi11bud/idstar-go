package services

import (
	"strconv"

	dtos "idstar.com/app/dtos/rekening"
	resp "idstar.com/app/dtos/response"
	"idstar.com/app/models"
	"idstar.com/app/repositories"
)

type RekeningService struct {
	rekeningRepository repositories.RekeningRepository
	karyawanRepository repositories.KaryawanRepository
}

func NewRekeningService(rekeningRepository repositories.RekeningRepository, karyawanRepository repositories.KaryawanRepository) *RekeningService {
	return &RekeningService{
		rekeningRepository: rekeningRepository,
		karyawanRepository: karyawanRepository,
	}
}

func (u *RekeningService) SaveRekening(detail *models.RekeningEntity) (*dtos.RekeningResponse, error) {
	karyawanId := strconv.Itoa(int(detail.KaryawanID))
	var karyawan *models.KaryawanEntity
	karyawan, err := u.karyawanRepository.FindByID(karyawanId)
	if karyawan == nil {
		return nil, err
	}

	rekening, err := u.rekeningRepository.Create(detail)
	if rekening == nil {
		return nil, err
	}

	detailKaryawan := dtos.Karyawan{
		ID:   strconv.Itoa(int(karyawan.KaryawanID)),
		Nama: karyawan.Nama,
	}

	respRekening := dtos.RekeningResponse{
		ID:          int(rekening.RekeningID),
		Nama:        rekening.Nama,
		Rekening:    rekening.Rekening,
		Jenis:       rekening.Jenis,
		Karyawan:    detailKaryawan,
		CreatedDate: rekening.CreatedDate,
		UpdatedDate: rekening.UpdatedDate,
		DeletedDate: rekening.DeletedDate,
	}

	return &respRekening, nil
}

func (u *RekeningService) GetById(id string) (*dtos.RekeningResponse, error) {
	rekening, err := u.rekeningRepository.FindByID(id)
	if rekening == nil {
		return nil, err
	}

	karyawanId := strconv.FormatUint(uint64(rekening.KaryawanID), 10)
	karyawan, err := u.karyawanRepository.FindByID(karyawanId)
	if karyawan == nil {
		return nil, err
	}

	detailKaryawan := dtos.Karyawan{
		ID:   strconv.Itoa(int(karyawan.KaryawanID)),
		Nama: karyawan.Nama,
	}

	respRekening := dtos.RekeningResponse{
		ID:          int(rekening.RekeningID),
		Nama:        rekening.Nama,
		Rekening:    rekening.Rekening,
		Jenis:       rekening.Jenis,
		Karyawan:    detailKaryawan,
		CreatedDate: rekening.CreatedDate,
		UpdatedDate: rekening.UpdatedDate,
		DeletedDate: rekening.DeletedDate,
	}

	return &respRekening, nil
}

func (u *RekeningService) FindAll() (any, error) {
	rekeningList, err := u.rekeningRepository.FindAll()
	if err != nil {
		return nil, err
	}

	rekeningListResponse := []dtos.RekeningResponse{}

	for _, rekening := range rekeningList {
		karayawanId := strconv.FormatUint(uint64(rekening.KaryawanID), 10)
		karyawan, err := u.karyawanRepository.FindByIDIgnoreDelete(karayawanId)
		if karyawan == nil {
			return nil, err
		}

		rekeningResponse := dtos.RekeningResponse{
			ID:       int(rekening.RekeningID),
			Nama:     rekening.Nama,
			Rekening: rekening.Rekening,
			Jenis:    rekening.Jenis,
			Karyawan: dtos.Karyawan{
				ID:   strconv.Itoa(int(karyawan.KaryawanID)),
				Nama: karyawan.Nama,
			},
			CreatedDate: rekening.CreatedDate,
			UpdatedDate: rekening.UpdatedDate,
			DeletedDate: rekening.DeletedDate,
		}

		rekeningListResponse = append(rekeningListResponse, rekeningResponse)

	}

	result := resp.Content{
		Data:             rekeningListResponse,
		Pageable:         resp.Pageable{Offset: 0, PageNumber: 0, PageSize: len(rekeningList), Unpaged: true, Paged: false},
		TotalElements:    len(rekeningList),
		TotalPages:       1,
		Size:             len(rekeningList),
		Number:           1,
		Sort:             resp.Sort{Empty: false, Sorted: false, Unsorted: false},
		First:            true,
		NumberOfElements: len(rekeningList),
		Empty:            len(rekeningList) == 0,
	}

	return result, nil
}

func (u *RekeningService) UpdateRekening(req dtos.UpdateRekeningRequest) (*dtos.RekeningResponse, error) {
	karyawan, err := u.karyawanRepository.FindByID(req.Karyawan.Id)
	if karyawan == nil {
		return nil, err
	}

	rekening, err := u.rekeningRepository.Update(req)
	if rekening == nil {
		return nil, err
	}

	detailKaryawan := dtos.Karyawan{
		ID:   strconv.Itoa(int(karyawan.KaryawanID)),
		Nama: karyawan.Nama,
	}

	respRekening := dtos.RekeningResponse{
		ID:          int(rekening.RekeningID),
		Nama:        rekening.Nama,
		Rekening:    rekening.Rekening,
		Jenis:       rekening.Jenis,
		Karyawan:    detailKaryawan,
		CreatedDate: rekening.CreatedDate,
		UpdatedDate: rekening.UpdatedDate,
		DeletedDate: rekening.DeletedDate,
	}

	return &respRekening, nil
}

func (u *RekeningService) DeleteById(id string) error {
	return u.rekeningRepository.DeleteTemp(id)
}
