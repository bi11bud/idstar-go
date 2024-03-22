package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dtos "idstar.com/app/dtos/karyawan"
	resp "idstar.com/app/dtos/response"
	"idstar.com/app/services"
)

type KaryawanController struct {
	service *services.KaryawanService
}

func NewKaryawanController(service *services.KaryawanService) *KaryawanController {
	return &KaryawanController{
		service: service,
	}
}

// GetKaryawan godoc
//
//	@Summary		Get Karyawan Data
//	@Description	Get Karyawan
//	@Tags			Karyawan
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	false	"ID"
//	@Router			/karyawan/{id} [get]
func (ctrl *KaryawanController) GetKaryawan(ctx *gin.Context) {
	response := resp.FailedResponse{
		Code: 400,
	}

	result, err := ctrl.service.GetById(ctx.Param("id"))
	if err != nil {
		response.Status = err.Error()
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	resSukses := resp.Response{
		Code:   200,
		Data:   result,
		Status: "sukses",
	}

	ctx.JSON(http.StatusOK, resSukses)
}

// PostKaryawan godoc
//
//	@Summary		Post Karyawan Data
//	@Description	Add new Karyawan
//	@Tags			Karyawan
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dtos.SaveKaryawanRequest	true	"Karyawan"
//	@Router			/karyawan [post]
func (ctrl *KaryawanController) SaveKaryawan(ctx *gin.Context) {
	response := resp.FailedResponse{
		Code: 400,
	}

	req := dtos.SaveKaryawanRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Status = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// if err := req.Validate(); err != nil {
	// 	response.Status = err.Error()
	// 	ctx.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	result, err := ctrl.service.SaveKaryawan(&req)
	if err != nil {
		response.Status = err.Error()
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	resSukses := resp.Response{
		Code:   200,
		Data:   result,
		Status: "sukses",
	}

	ctx.JSON(http.StatusCreated, resSukses)
}

// GetAllKaryawan godoc
//
//	@Summary		Get All Karyawan Data
//	@Description	Get All Karyawan
//	@Tags			Karyawan
//	@Accept			json
//	@Produce		json
//	@Router			/karyawan/list [get]
func (ctrl *KaryawanController) GetAllKaryawan(ctx *gin.Context) {
	response := resp.FailedResponse{
		Code: 400,
	}

	result, err := ctrl.service.FindAll()
	if err != nil {
		response.Status = err.Error()
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	resSukses := resp.Response{
		Code:   200,
		Data:   result,
		Status: "sukses",
	}

	ctx.JSON(http.StatusOK, resSukses)
}

// PutKaryawan godoc
//
//	@Summary		Put Karyawan Data
//	@Description	Update Karyawan
//	@Tags			Karyawan
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dtos.UpdateKaryawanRequest	true	"Karyawan"
//	@Router			/karyawan [put]
func (ctrl *KaryawanController) UpdateKaryawan(ctx *gin.Context) {
	response := resp.FailedResponse{
		Code: 400,
	}

	req := dtos.UpdateKaryawanRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Status = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if req.Id == "" {
		response.Status = "Invalid Detail Karyawan ID"
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if err := req.Validate(); err != nil {
		response.Status = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	result, err := ctrl.service.UpdateKaryawan(req)
	if err != nil {
		response.Status = err.Error()
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	resSukses := resp.Response{
		Code:   200,
		Data:   result,
		Status: "sukses",
	}

	ctx.JSON(http.StatusOK, resSukses)
}

// DeleteKaryawan godoc
//
//	@Summary		Delete Karyawan Data
//	@Description	Delete Karyawan
//	@Tags			Karyawan
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"Karyawan ID"
//	@Router			/karyawan/{id} [delete]
func (ctrl *KaryawanController) DeleteKaryawan(ctx *gin.Context) {
	response := resp.FailedResponse{
		Code: 400,
	}

	var id string = ctx.Param("id")
	if id == "" {
		response.Status = "Invalid Karyawan ID"
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := ctrl.service.DeleteById(id)
	if err != nil {
		response.Status = err.Error()
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	resSukses := resp.Response{
		Code:   200,
		Data:   "Sukses",
		Status: "sukses",
	}

	ctx.JSON(http.StatusOK, resSukses)
}
