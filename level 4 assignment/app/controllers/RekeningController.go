package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	dtos "idstar.com/app/dtos/rekening"
	resp "idstar.com/app/dtos/response"
	"idstar.com/app/models"
	"idstar.com/app/services"
)

type RekeningController struct {
	service *services.RekeningService
}

func NewRekeningController(service *services.RekeningService) *RekeningController {
	return &RekeningController{
		service: service,
	}
}

// GetRekening godoc
//
//	@Summary		Get Rekening Data
//	@Description	Get Rekening
//	@Tags			Rekening
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	false	"ID"
//	@Router			/rekening/{id} [get]
func (ctrl *RekeningController) GetRekening(ctx *gin.Context) {
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

// GetAllRekening godoc
//
//	@Summary		Get All Rekening Data
//	@Description	Get All Rekening
//	@Tags			Rekening
//	@Accept			json
//	@Produce		json
//	@Router			/rekening/list [get]
func (ctrl *RekeningController) GetAllRekening(ctx *gin.Context) {
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

// PostRekening godoc
//
//	@Summary		Post Rekening Data
//	@Description	Add new Rekening
//	@Tags			Rekening
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dtos.SaveRekeningRequest	true	"Rekening"
//	@Router			/rekening [post]
func (ctrl *RekeningController) SaveRekening(ctx *gin.Context) {
	response := resp.FailedResponse{
		Code: 400,
	}

	req := dtos.SaveRekeningRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Status = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	id, err := strconv.ParseUint(req.Karyawan.Id, 10, 64)
	if err != nil {
		response.Status = "Invalid Karyawan ID"
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	m := models.RekeningEntity{
		Nama:       req.Nama,
		Jenis:      req.Jenis,
		Rekening:   req.Rekening,
		KaryawanID: uint(id),
	}

	result, err := ctrl.service.SaveRekening(&m)
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

// PutRekening godoc
//
//	@Summary		Put Rekening Data
//	@Description	Update Rekening
//	@Tags			Rekening
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dtos.UpdateRekeningRequest	true	"Rekening"
//	@Router			/rekening [put]
func (ctrl *RekeningController) UpdateRekening(ctx *gin.Context) {
	response := resp.FailedResponse{
		Code: 400,
	}

	req := dtos.UpdateRekeningRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Status = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if req.Id == "" {
		response.Status = "Invalid Rekening ID"
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if err := req.Validate(); err != nil {
		response.Status = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	result, err := ctrl.service.UpdateRekening(req)
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

// DeleteRekening godoc
//
//	@Summary		Delete Rekening Data
//	@Description	Delete Rekening
//	@Tags			Rekening
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"Rekening ID"
//	@Router			/rekening/{id} [delete]
func (ctrl *RekeningController) DeleteRekenig(ctx *gin.Context) {
	response := resp.FailedResponse{
		Code: 400,
	}

	var id string = ctx.Param("id")
	if id == "" {
		response.Status = "Invalid Rekening ID"
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
