package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	resp "idstar.com/app/dtos/response"
	dtos "idstar.com/app/dtos/user"
	"idstar.com/app/models"
	"idstar.com/app/services"
	"idstar.com/app/tools"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

// GetUser godoc
//
//	@Summary		Get User Data
//	@Description	Get User Account
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			id	path	string	false	"ID"
//	@Router			/user/{id} [get]
func (ctrl *UserController) GetUser(ctx *gin.Context) {
	result, err := ctrl.service.GetUserByID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	response := resp.Response{
		Code:   200,
		Data:   result,
		Status: "Successfully",
	}

	ctx.JSON(http.StatusOK, response)
}

// PostUser godoc
//
//	@Summary		Create New User
//	@Description	Register New User Account
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dtos.CreateUserRequest	true	"User"
//	@Router			/user-register [post]
func (ctrl *UserController) RegisterUser(ctx *gin.Context) {
	req := dtos.CreateUserRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		response := resp.Response{
			Code:   400,
			Data:   errMsg,
			Status: "Failed On Validation",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if err := req.Validate(); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		response := resp.Response{
			Code:   400,
			Data:   errMsg,
			Status: "Failed On Validation",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	m := models.UserEntity{
		Username: req.Username,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	result, err := ctrl.service.CreateUser(&m)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := resp.Response{
		Code:   200,
		Data:   result,
		Status: "Successfully, Please check your email to activate your account",
	}

	ctx.JSON(http.StatusCreated, response)
}

// PostUser godoc
//
//	@Summary		Post User Data
//	@Description	Send Otp Registration User Account
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dtos.ForgetPasswordRequest	true	"User"
//	@Router			/user-register/send-otp [post]
func (ctrl *UserController) SendOtpRegister(ctx *gin.Context) {
	req := dtos.ForgetPasswordRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		response := resp.Response{
			Code:   400,
			Data:   errMsg,
			Status: "Failed On Validation",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := ctrl.service.UpdateOtpUser(req.Username)
	if err != nil {
		response := resp.FailedResponse{
			Code:   500,
			Status: err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := resp.ErrorResponse{
		Code:    200,
		Status:  "Success",
		Message: "Thanks, please check your email for activation",
	}

	ctx.JSON(http.StatusOK, response)
}

// PostUser godoc
//
//	@Summary		Post User Data
//	@Description	Approved User Account
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			otp	path	string	false	"otp"
//	@Param			request	body	dtos.ForgetPasswordRequest	true	"User"
//	@Router			/user-register/register-confirm-otp/{otp} [post]
func (ctrl *UserController) ApprovedUser(ctx *gin.Context) {
	req := dtos.ForgetPasswordRequest{}
	otp := ctx.Param("otp")
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		response := resp.Response{
			Code:   400,
			Data:   errMsg,
			Status: "Failed On Validation",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := ctrl.service.UpdateApprovedUser(req.Username, otp)
	if err != nil {
		response := resp.FailedResponse{
			Code:   500,
			Status: err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := resp.ErrorResponse{
		Code:    200,
		Status:  "Success",
		Message: "Your Account is active now, please try to log in",
	}

	ctx.JSON(http.StatusOK, response)
}

// PostUser godoc
//
//	@Summary		Post User Data
//	@Description	Forget Password User Account
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dtos.ForgetPasswordRequest	true	"User"
//	@Router			/forget-password/send [post]
func (ctrl *UserController) SendOtpForgetPassword(ctx *gin.Context) {
	req := dtos.ForgetPasswordRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		response := resp.Response{
			Code:   400,
			Data:   errMsg,
			Status: "Failed On Validation",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := ctrl.service.UpdateOtpForgetPasswordUser(req.Username)
	if err != nil {
		response := resp.FailedResponse{
			Code:   500,
			Status: err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := resp.ErrorResponse{
		Code:    200,
		Status:  "Success",
		Message: "Please check your email to reset your password",
	}

	ctx.JSON(http.StatusOK, response)
}

// PostUser godoc
//
//	@Summary		Post User Data
//	@Description	Forget Password User Account
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dtos.UserResetPasswordRequest	true	"User"
//	@Router			/forget-password/change-password [post]
func (ctrl *UserController) ResetPassword(ctx *gin.Context) {
	req := dtos.UserResetPasswordRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		response := resp.Response{
			Code:   400,
			Data:   errMsg,
			Status: "Failed On Validation",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := ctrl.service.UpdatePasswordUser(req)
	if err != nil {
		response := resp.FailedResponse{
			Code:   500,
			Status: err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := resp.ErrorResponse{
		Code:    200,
		Status:  "Success",
		Message: "Reset Password Success",
	}

	ctx.JSON(http.StatusOK, response)
}
