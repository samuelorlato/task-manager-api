package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/samuelorlato/task-manager-api/internal/core/ports"
)

type JWTService struct{}

func NewJWTService() ports.AuthService {
	return &JWTService{}
}

type JWTClaim struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func (j *JWTService) GenerateToken(email string, expiresAt *time.Time, secret string) (string, error) {
	claims := JWTClaim{
		Email: email,
	}

	if expiresAt != nil {
		claims.RegisteredClaims = jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(*expiresAt),
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JWTService) ValidateToken(tokenString string, secret string) (string, error) {
	claims := &JWTClaim{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			err := errors.New("Invalid JWT")
			return "", err
		}

		return "", err
	}

	if !token.Valid {
		err := errors.New("Invalid JWT")
		return "", err
	}

	return claims.Email, nil
}
