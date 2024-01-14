package ports

import "time"

type AuthService interface {
	GenerateToken(identification string, expiresAt *time.Time, secret string) (string, error)
	ValidateToken(token string, secret string) (string, error)
}
