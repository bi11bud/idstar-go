package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dtos "idstar.com/app/dtos/user"
	"idstar.com/app/services"
)

type AuthController struct {
	service *services.AuthService
}

func NewAuthenticationController(service *services.AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

// RequestToken godoc
//
//	@Summary		Login Request Token user
//	@Description	Login Request Token for Authorization
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dtos.LoginRequest	true	"body"
//	@Router			/user-login/login [post]
func (ctrl *AuthController) Login(ctx *gin.Context) {
	req := dtos.LoginRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	statusCode, token, err := ctrl.service.Login(req)
	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(statusCode, gin.H{"data": token})
}
