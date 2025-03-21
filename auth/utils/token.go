package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	// 🛑 В реальном приложении эти значения должны быть в env-переменных 🛑
	// 🛑 Секретные данные пока что внутри кода 🛑
	AccessTokenSecret  = "access_secret_key_very_secure"
	RefreshTokenSecret = "refresh_secret_key_very_secure"
	AccessTokenExpiry  = time.Minute * 15   // 15 минут
	RefreshTokenExpiry = time.Hour * 24 * 7 // 7 дней
)

type JWTClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

// GenerateAccessToken создает новый JWT токен доступа
func GenerateAccessToken(userID uint, username string, isAdmin bool) (string, error) {
	claims := JWTClaims{
		UserID:   userID,
		Username: username,
		IsAdmin:  isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "gavna-auth-service",
			Subject:   username,
			ID:        uuid.New().String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(AccessTokenSecret))
}

// GenerateRefreshToken создает новый токен для обновления
func GenerateRefreshToken() (string, error) {
	return uuid.NewString(), nil
}

// ValidateAccessToken проверяет валидность access token
func ValidateAccessToken(tokenString string) (*JWTClaims, error) {
	claims := &JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(AccessTokenSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
