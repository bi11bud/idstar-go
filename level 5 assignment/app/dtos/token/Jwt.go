package dtos

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	Value     string    `json:"value"`
	IssuedOn  time.Time `json:"issuedOn"`
	ExpiresOn time.Time `json:"expiresOn"`
}

type Claims struct {
	Username string `json:"username"`
	Approved bool
	jwt.RegisteredClaims
}
