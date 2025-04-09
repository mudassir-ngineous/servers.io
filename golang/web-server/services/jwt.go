package services

import (
	"errors"
	"time"
	"web-server/config"
	"web-server/models"

	"github.com/golang-jwt/jwt/v5"
)

var JWTService *JWTServiceType

type JWTServiceType struct {
	secretKey string
	expiresIn int
}

func NewJWTService(cfg *config.Config) *JWTServiceType {
	return &JWTServiceType{
		secretKey: cfg.JWT.SecretKey,
		expiresIn: cfg.JWT.ExpiresIn,
	}
}

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func (s *JWTServiceType) GenerateToken(user *models.User) (string, error) {
	claims := Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(s.expiresIn))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *JWTServiceType) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
} 