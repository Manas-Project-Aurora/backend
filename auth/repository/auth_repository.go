package repository

import (
	"errors"
	"time"

	"github.com/Manas-Project-Aurora/gavna/auth/models"
	"github.com/Manas-Project-Aurora/gavna/auth/utils"
	internalmodels "github.com/Manas-Project-Aurora/gavna/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

// CreateUser создает нового пользователя
func (r *AuthRepository) CreateUser(username, password string) (*internalmodels.User, error) {
	// Проверяем, существует ли пользователь с таким username
	var existingUser internalmodels.User
	result := r.DB.Where("username = ?", username).First(&existingUser)
	if result.Error == nil {
		return nil, errors.New("user with this username already exists")
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	// Хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Создаем нового пользователя
	user := internalmodels.User{
		Username:     username,
		PasswordHash: string(hashedPassword),
		IsActive:     true,
	}

	result = r.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// GetUserByUsername находит пользователя по имени пользователя
func (r *AuthRepository) GetUserByUsername(username string) (*internalmodels.User, error) {
	var user internalmodels.User
	result := r.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// VerifyPassword проверяет правильность пароля для пользователя
func (r *AuthRepository) VerifyPassword(user *internalmodels.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	return err == nil
}

// StoreRefreshToken сохраняет токен обновления
func (r *AuthRepository) StoreRefreshToken(userID uint, token string) error {
	refreshToken := models.RefreshToken{
		Token:     token,
		UserID:    userID,
		ExpiresAt: time.Now().Add(utils.RefreshTokenExpiry),
		IssuedAt:  time.Now(),
	}

	return r.DB.Create(&refreshToken).Error
}

// GetRefreshToken получает токен обновления
func (r *AuthRepository) GetRefreshToken(token string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	result := r.DB.Where("token = ? AND is_revoked = ?", token, false).First(&refreshToken)
	if result.Error != nil {
		return nil, result.Error
	}

	// Проверяем, не истек ли токен
	if refreshToken.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("refresh token expired")
	}

	return &refreshToken, nil
}

// GetUserByID получает пользователя по ID
func (r *AuthRepository) GetUserByID(userID uint) (*internalmodels.User, error) {
	var user internalmodels.User
	result := r.DB.First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// RevokeRefreshToken помечает токен как отозванный
func (r *AuthRepository) RevokeRefreshToken(token string) error {
	result := r.DB.Model(&models.RefreshToken{}).
		Where("token = ?", token).
		Update("is_revoked", true)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("token not found")
	}

	return nil
}

// RevokeAllUserTokens отзывает все токены пользователя
func (r *AuthRepository) RevokeAllUserTokens(userID uint) error {
	return r.DB.Model(&models.RefreshToken{}).
		Where("user_id = ?", userID).
		Update("is_revoked", true).Error
}
