package service

import (
	"fmt"
	"time"
	"users-service/pkg"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

//go:generate go run go.uber.org/mock/mockgen -destination mock_service/jwt_service.go . JwtService
type JwtService interface {
	GenerateToken(id uuid.UUID) (string, error)
	ValidateToken(token string) bool
}

type jwtService struct {
	secretKey string
}

func NewJwtService(secretKey string) JwtService {
	return &jwtService{
		secretKey: secretKey,
	}
}

func (j *jwtService) GenerateToken(id uuid.UUID) (string, error) {
	claim := jwt.MapClaims{
		"issuer":     "github.com/daniellsantiago/realworld-microservices-example-app",
		"sum":        id.String(),
		"expired_at": time.Now().Add(time.Hour * 2).Unix(),
		"issued_at":  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		pkg.Logger.Error("Failed to generate JWT token", zap.Error(err))
		return "", err
	}

	return signedToken, nil
}

func (j *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			pkg.Logger.Error("Invalid token")
			return nil, fmt.Errorf("invalid token: %v. error_type: %w", token, pkg.ErrValidation)
		}

		return []byte(j.secretKey), nil
	})

	return err == nil
}
