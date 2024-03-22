package services

import (
	"errors"
	"log"
	"strconv"
	"time"

	dtos "idstar.com/app/dtos/user"
	"idstar.com/app/models"
	"idstar.com/app/repositories"
	"idstar.com/app/tools"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (u *UserService) CreateUser(user *models.UserEntity) (*models.UserEntity, error) {
	// Enkripsi password
	aes128 := tools.Aes128{}
	encryptedPassword, err := aes128.Encrypt(user.Password)
	if err != nil {
		return nil, err
	}

	user.OtpExpiredDate = time.Now().Add(30 * time.Minute)
	generateOtp := tools.GenerateOtp{}
	user.Otp = generateOtp.GenerateOTP()
	user.Password = *encryptedPassword

	userResult, err := u.userRepository.Create(user)
	if err == nil {
		log.Println("Send Activation Email")
		//kirim email
		sendEmail := tools.SendEmail{}
		err = sendEmail.ActivationOtp(userResult)
		if err != nil {
			log.Println("Send Activation Email Failed")
			log.Println(err.Error())
		}
	}

	return userResult, err
}

func (u *UserService) GetUserByID(userID string) (*models.UserEntity, error) {
	return u.userRepository.FindByID(userID)
}

func (u *UserService) UpdatePasswordUser(user dtos.UserResetPasswordRequest) error {
	if user.NewPassword != user.ConfirmNewPassword {
		return errors.New("password is not match")
	}

	aes128 := tools.Aes128{}
	// encrypt before save to db
	encryptedPassword, err := aes128.Encrypt(user.ConfirmNewPassword)
	if err != nil {
		return err
	}

	userEnt, err := u.userRepository.FindByUsernameOrEmail(user.Email)
	if err != nil {
		return err
	}

	if userEnt.Otp != user.Otp {
		return errors.New("otp is not match")
	}

	tNow := time.Now()
	if !userEnt.OtpExpiredDate.After(tNow) {
		return errors.New("otp is expired")
	}

	user.ConfirmNewPassword = *encryptedPassword

	if userEnt.Password == user.ConfirmNewPassword {
		return errors.New("new password cannot be same like old password")
	}

	id := strconv.FormatUint(uint64(userEnt.Id), 10)
	return u.userRepository.UpdatePassword(id, user.ConfirmNewPassword)
}

func (u *UserService) UpdateOtpUser(email string) error {
	generateOtp := tools.GenerateOtp{}
	otp := generateOtp.GenerateOTP()

	user, err := u.userRepository.FindByUsernameOrEmail(email)
	if err != nil {
		return err
	}

	updatedUser, err := u.userRepository.UpdateOtp(user.Email, otp)
	if err == nil {
		log.Println("Send Otp Email")
		//kirim email
		sendEmail := tools.SendEmail{}
		err = sendEmail.ActivationOtp(updatedUser)
		if err != nil {
			log.Println("Send Otp Email Failed")
			log.Println(err.Error())
		}
	}

	return nil
}

func (u *UserService) UpdateOtpForgetPasswordUser(userId string) error {
	generateOtp := tools.GenerateOtp{}
	otp := generateOtp.GenerateOTP()

	user, err := u.userRepository.FindByUsernameOrEmail(userId)
	if err != nil {
		return err
	}

	updatedUser, err := u.userRepository.UpdateOtp(user.Email, otp)
	if err == nil {
		log.Println("Send Otp Forget Password Email")
		//kirim email
		sendEmail := tools.SendEmail{}
		err = sendEmail.ResetPasswordOtp(updatedUser)
		if err != nil {
			log.Println("Send Otp Forget Password Email Failed")
			log.Println(err.Error())
		}
	}

	return nil
}

func (u *UserService) UpdateApprovedUser(userId string, otpReq string) error {
	userEnt, err := u.userRepository.FindByUsernameOrEmail(userId)
	if err != nil {
		return err
	}

	if userEnt.Otp != otpReq {
		return errors.New("otp is not match")
	}

	tNow := time.Now()
	if !userEnt.OtpExpiredDate.After(tNow) {
		return errors.New("otp is expired")
	}

	id := strconv.FormatUint(uint64(userEnt.Id), 10)
	return u.userRepository.UpdateStatus(id)
}
