package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	dtos "idstar.com/app/dtos/token"
	user "idstar.com/app/dtos/user"
	"idstar.com/app/models"
	"idstar.com/app/repositories"
	"idstar.com/app/tools"
)

const key = "abcdefghij1234567890"

type AuthService struct {
	userRepository repositories.UserRepository
}

func NewAuthenticationService(userRepository repositories.UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (u *AuthService) Login(dto user.LoginRequest) (int, *dtos.Token, error) {
	user, err := u.userRepository.FindByUsernameOrEmail(dto.Username)
	if err != nil {
		return 500, nil, err
	}

	aes128 := tools.Aes128{}
	decryptedPassword, err := aes128.Decrypt(user.Password)
	if err != nil {
		return 500, nil, err
	}

	if dto.Password != *decryptedPassword {
		return 401, nil, errors.New("invalid password")
	}

	if !user.Approved {
		return 401, nil, errors.New("your account is not active, please activate it first")
	}

	token, err := GenerateToken(user)
	if err != nil {
		return 500, nil, err
	}

	return 200, token, nil
}

func GenerateToken(user *models.UserEntity) (*dtos.Token, error) {
	iat := jwt.NewNumericDate(time.Now())
	exp := jwt.NewNumericDate(time.Now().Add(time.Millisecond * 3600000))

	claims := &dtos.Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "issuer",
			Audience:  []string{"audience01", "audience02"},
			IssuedAt:  iat,
			ExpiresAt: exp,
		},
		Approved: user.Approved,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return nil, err
	}

	return &dtos.Token{
		Value:     tokenString,
		IssuedOn:  time.Unix(iat.Unix(), 0),
		ExpiresOn: time.Unix(exp.Unix(), 0),
	}, nil
}
