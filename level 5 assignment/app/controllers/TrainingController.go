package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	resp "idstar.com/app/dtos/response"
	dtos "idstar.com/app/dtos/training"
	"idstar.com/app/models"
	"idstar.com/app/services"
)

type TrainingController struct {
	service *services.TrainingService
}

func NewTrainingController(service *services.TrainingService) *TrainingController {
	return &TrainingController{
		service: service,
	}
}

// GetTraining godoc
//
//	@Summary		Get Training Data
//	@Description	Get Training
//	@Tags			Training
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			id	path	string	false	"ID"
//	@Router			/training/{id} [get]
func (ctrl *TrainingController) GetTraining(ctx *gin.Context) {
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

// GetAllTraining godoc
//
//	@Summary		Get All Training Data
//	@Description	Get All Training
//	@Tags			Training
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Router			/training/list [get]
func (ctrl *TrainingController) GetAllTraining(ctx *gin.Context) {
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

// PostTraining godoc
//
//	@Summary		Post Training Data
//	@Description	Add new Training
//	@Tags			Training
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			request	body	dtos.SaveTrainingRequest	true	"Training"
//	@Router			/training [post]
func (ctrl *TrainingController) SaveTraining(ctx *gin.Context) {
	response := resp.FailedResponse{
		Code: 400,
	}

	req := dtos.SaveTrainingRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Status = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	m := models.TrainingEntity{
		Tema:     req.Tema,
		Pengajar: req.Pengajar,
	}

	result, err := ctrl.service.SaveTraining(&m)
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

// PutTraining godoc
//
//	@Summary		Put Training Data
//	@Description	Update Training
//	@Tags			Training
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			request	body	dtos.UpdateTrainingRequest	true	"Training"
//	@Router			/training [put]
func (ctrl *TrainingController) UpdateTraining(ctx *gin.Context) {
	response := resp.FailedResponse{
		Code: 400,
	}

	req := dtos.UpdateTrainingRequest{}
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

	result, err := ctrl.service.UpdateTraining(req)
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

// DeleteTraining godoc
//
//	@Summary		Delete Training Data
//	@Description	Delete Training
//	@Tags			Training
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			id	path	string	true	"Training ID"
//	@Router			/training/{id} [delete]
func (ctrl *TrainingController) DeleteTraining(ctx *gin.Context) {
	response := resp.FailedResponse{
		Code: 400,
	}

	var id string = ctx.Param("id")
	if id == "" {
		response.Status = "Invalid Training ID"
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
