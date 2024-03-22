package tools

import (
	"fmt"
	"net/smtp"
	"strings"

	"idstar.com/app/models"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "PT. Karyawan Makmur"
const CONFIG_AUTH_EMAIL = "ukires@gmail.com"
const CONFIG_AUTH_PASSWORD = "mcymymlfjecbvonl"

type SendEmail struct{}

func (c *SendEmail) ResetPasswordOtp(user *models.UserEntity) error {
	subject := "Forget Password"
	message := "If you requested to reset password for " + user.Email + ", use the confirmation code below to complete " +
		"the process. If you didn't make this request, just ignore the email. \n" +
		"code: " + user.Otp + "\n" +
		"Regards, \n" +
		CONFIG_SENDER_NAME

	cc := []string{}
	to := []string{user.Email}

	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}

func (c *SendEmail) ActivationOtp(user *models.UserEntity) error {
	subject := "Activation User"
	message := "Halo, " + user.Name + ". Selamat bergabung \n" +
		"Harap konfirmasikan email kamu dengan memasukkan kode dibawah ini: \n" +
		"kode: " + user.Otp + "\n" +
		"Jika kamu butuh bantuan atau pertanyaan, hubungi customer care kami yaa \n" +
		"Semoga harimu menyenangkan! \n" +
		CONFIG_SENDER_NAME

	cc := []string{}
	to := []string{user.Email}

	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}
