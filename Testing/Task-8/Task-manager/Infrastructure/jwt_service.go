package infrastructure

import (
	"errors"
	"os"
	domain "tskmgr/Domain"

	"github.com/golang-jwt/jwt"
)

type JWTService struct {
}

func NewJWTService() *JWTService {
	return &JWTService{}
}

func (j *JWTService) GetToken(claim *domain.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	key := os.Getenv("JWT_SECRET")
	jwtSecret := []byte(key)

	if jwtSecret == nil {
		return "", errors.New("JWT_SECRET environment variable not set")
	}

	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
