package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dtos "idstar.com/app/dtos/karyawan-training"
	resp "idstar.com/app/dtos/response"
	"idstar.com/app/services"
)

type KaryawanTrainingController struct {
	service *services.KaryawanTrainingService
}

func NewKaryawanTrainingController(service *services.KaryawanTrainingService) *KaryawanTrainingController {
	return &KaryawanTrainingController{
		service: service,
	}
}

// GetKaryawanTraining godoc
//
//	@Summary		Get KaryawanTraining Data
//	@Description	Get KaryawanTraining
//	@Tags			KaryawanTraining
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	false	"ID"
//	@Router			/karyawanTraining/{id} [get]
func (ctrl *KaryawanTrainingController) GetKaryawanTraining(ctx *gin.Context) {
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

// PostKaryawanTraining godoc
//
//	@Summary		Post KaryawanTraining Data
//	@Description	Add new KaryawanTraining
//	@Tags			KaryawanTraining
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dtos.SaveKaryawanTrainingRequest	true	"KaryawanTraining"
//	@Router			/karyawanTraining [post]
func (ctrl *KaryawanTrainingController) SaveKaryawanTraining(ctx *gin.Context) {
	response := resp.FailedResponse{
		Code: 400,
	}

	req := dtos.SaveKaryawanTrainingRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Status = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	result, err := ctrl.service.SaveKaryawanTraining(&req)
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

// GetAllKaryawanTraining godoc
//
//	@Summary		Get All KaryawanTraining Data
//	@Description	Get All KaryawanTraining
//	@Tags			KaryawanTraining
//	@Accept			json
//	@Produce		json
//	@Router			/karyawanTraining/list [get]
func (ctrl *KaryawanTrainingController) GetAllKaryawanTraining(ctx *gin.Context) {
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

// PutKaryawanTraining godoc
//
//	@Summary		Put KaryawanTraining Data
//	@Description	Update KaryawanTraining
//	@Tags			KaryawanTraining
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dtos.UpdateKaryawanTrainingRequest	true	"KaryawanTraining"
//	@Router			/karyawanTraining [put]
func (ctrl *KaryawanTrainingController) UpdateKaryawanTraining(ctx *gin.Context) {
	response := resp.FailedResponse{
		Code: 400,
	}

	req := dtos.UpdateKaryawanTrainingRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Status = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if err := req.Validate(); err != nil {
		response.Status = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	result, err := ctrl.service.UpdateKaryawanTraining(req)
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

// DeleteKaryawanTraining godoc
//
//	@Summary		Delete KaryawanTraining Data
//	@Description	Delete KaryawanTraining
//	@Tags			KaryawanTraining
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"KaryawanTraining ID"
//	@Router			/karyawanTraining/{id} [delete]
func (ctrl *KaryawanTrainingController) DeleteKaryawanTraining(ctx *gin.Context) {
	response := resp.FailedResponse{
		Code: 400,
	}

	var id string = ctx.Param("id")
	if id == "" {
		response.Status = "Invalid KaryawanTraining ID"
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
