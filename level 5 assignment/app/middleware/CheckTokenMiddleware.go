package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"idstar.com/app/configs"
	resp "idstar.com/app/dtos/response"
	dtos "idstar.com/app/dtos/token"
)

func CheckToken() gin.HandlerFunc {

	return func(c *gin.Context) {
		url := c.Request.URL
		// hadling for unnecessary token
		if url.Path == "/api/v1/user-login/login" ||
			url.Path == "/api/v1/forget-password/change-password" ||
			url.Path == "/api/v1/forget-password/send" ||
			url.Path == "/api/v1/user-register" ||
			strings.HasPrefix(url.Path, "/api/v1/user-register/register-confirm-otp/") ||
			url.Path == "/api/v1/user-register/send-otp" ||
			url.Path == "/api/v1/upload" ||
			strings.HasPrefix(url.Path, "/api/v1/showFile/") ||
			strings.Contains(url.Path, "/swagger/") {
			c.Next()
			return
		}

		// Get token from header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			panic(resp.ErrorResponse{
				Code:    401,
				Status:  "failed",
				Message: "Missing Header 'Authorization'",
			})
		}

		jwtKey := []byte(configs.GetJWTConfigurationInstance().Key)
		// Verify token
		tokenString := authHeader[len("Bearer "):]
		token, err := jwt.ParseWithClaims(tokenString, &dtos.Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrInvalidKey
			}
			return jwtKey, nil
		})

		if err != nil {
			panic(resp.ErrorResponse{
				Code:    401,
				Status:  "failed",
				Message: err.Error(),
			})
		}

		claims, ok := token.Claims.(*dtos.Claims)
		if !ok || !token.Valid {
			panic(resp.ErrorResponse{
				Code:    401,
				Status:  "failed",
				Message: "Not Authorize",
			})
		}

		if !claims.Approved {
			panic(resp.ErrorResponse{
				Code:    401,
				Status:  "failed",
				Message: "Account not active",
			})
		}

		c.Next()
	}
}
