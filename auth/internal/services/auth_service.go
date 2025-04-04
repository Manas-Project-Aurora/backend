package services

import (
	"errors"

	"github.com/Manas-Project-Aurora/backend/auth/internal/models"
	"github.com/Manas-Project-Aurora/backend/auth/internal/repository"
	"github.com/Manas-Project-Aurora/backend/auth/internal/utils"
	internalmodels "github.com/Manas-Project-Aurora/backend/internal/models"
)

type AuthService struct {
	repo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

// Register регистрирует нового пользователя
func (s *AuthService) Register(username, password string) (*internalmodels.User, error) {
	// Валидация входных данных
	if username == "" || password == "" {
		return nil, errors.New("username and password are required")
	}

	if len(password) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}

	// Создаем пользователя
	return s.repo.CreateUser(username, password)
}

// Login аутентифицирует пользователя и генерирует токены
func (s *AuthService) Login(username, password string) (*models.TokenPair, error) {
	// Получаем пользователя по username
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Проверяем пароль
	if !s.repo.VerifyPassword(user, password) {
		return nil, errors.New("invalid credentials")
	}

	// Проверяем активность пользователя
	if !user.IsActive {
		return nil, errors.New("account is not active")
	}

	// Генерируем токены
	return s.generateTokenPair(user)
}

// RefreshToken обновляет токены доступа
func (s *AuthService) RefreshToken(refreshToken string) (*models.TokenPair, error) {
	// Получаем информацию о refresh токене
	token, err := s.repo.GetRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	// Получаем пользователя
	user, err := s.repo.GetUserByID(token.UserID)
	if err != nil {
		return nil, err
	}

	// Отзываем старый токен
	if err := s.repo.RevokeRefreshToken(refreshToken); err != nil {
		return nil, err
	}

	// Генерируем новую пару токенов
	return s.generateTokenPair(user)
}

// Logout завершает сессию пользователя
func (s *AuthService) Logout(refreshToken string) error {
	if refreshToken != "" {
		return s.repo.RevokeRefreshToken(refreshToken)
	}
	return nil
}

// generateTokenPair создает пару access и refresh токенов
func (s *AuthService) generateTokenPair(user *internalmodels.User) (*models.TokenPair, error) {
	// Генерируем access token
	accessToken, err := utils.GenerateAccessToken(user.ID, user.Username, user.IsAdmin)
	if err != nil {
		return nil, err
	}

	// Генерируем refresh token
	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	// Сохраняем refresh token в базе
	if err := s.repo.StoreRefreshToken(user.ID, refreshToken); err != nil {
		return nil, err
	}

	return &models.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(utils.AccessTokenExpiry.Seconds()),
	}, nil
}
