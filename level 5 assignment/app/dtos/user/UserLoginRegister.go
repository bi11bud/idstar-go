package dtos

import "github.com/go-playground/validator/v10"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ForgetPasswordRequest struct {
	Username string `json:"username"`
}

type UserResetPasswordRequest struct {
	Email              string `json:"email"`
	NewPassword        string `json:"newPassword"`
	ConfirmNewPassword string `json:"confirmNewPassword"`
	Otp                string `json:"otp"`
}

type CreateUserRequest struct {
	Name            string `json:"name" validate:"required"`
	Username        string `json:"username" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=5"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=5,eqfield=Password"`
}

func (c *CreateUserRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)

	if err != nil {
		return err
	}
	return nil
}
