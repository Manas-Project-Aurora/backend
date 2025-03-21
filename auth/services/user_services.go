package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/Manas-Project-Aurora/gavna/auth/repository"
	"github.com/Manas-Project-Aurora/gavna/internal/models"
)

// UserService определяет интерфейс бизнес-логики для работы с пользователями.
type UserService interface {
	// RegisterUser регистрирует нового пользователя.
	RegisterUser(username, password, telegram string) (*models.User, error)
	// Login аутентифицирует пользователя и возвращает access и refresh токены.
	Login(username, password string) (string, string, error)
	// RefreshToken проверяет refresh токен и генерирует новый access токен.
	RefreshToken(refreshToken string) (string, error)
}

// userService — конкретная реализация UserService.
type userService struct {
	repo repository.UserRepository
}

// jwtSecret используется для подписи JWT токенов.
// В реальном проекте рекомендуется хранить секрет в переменных окружения.
var jwtSecret = []byte("my-secret-key")

// Claims определяет структуру данных внутри JWT токена.
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// NewUserService создает новый экземпляр UserService.
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// RegisterUser регистрирует нового пользователя: хэширует пароль и сохраняет данные в БД.
func (s *userService) RegisterUser(username, password, telegram string) (*models.User, error) {
	// Хэшируем пароль с помощью bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Формируем объект пользователя
	user := &models.User{
		Username:     username,
		PasswordHash: string(hashedPassword),
		Telegram:     telegram,
		IsActive:     true, // Можно настроить по необходимости
	}

	// Сохраняем пользователя через репозиторий
	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

// generateAccessToken создает access токен с коротким временем жизни (например, 15 минут).
func generateAccessToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	// Создаем токен с алгоритмом HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// generateRefreshToken создает refresh токен с более длительным временем жизни (например, 7 дней).
func generateRefreshToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Login аутентифицирует пользователя: проверяет пароль и возвращает пару токенов (access и refresh).
func (s *userService) Login(username, password string) (string, string, error) {
	// Получаем пользователя по имени из БД
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return "", "", err
	}
	// Сравниваем хэш пароля с введенным паролем
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", "", errors.New("invalid credentials")
	}

	// Генерируем access токен
	accessToken, err := generateAccessToken(user.ID)
	if err != nil {
		return "", "", err
	}
	// Генерируем refresh токен
	refreshToken, err := generateRefreshToken(user.ID)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// RefreshToken проверяет refresh токен и, если он валиден, генерирует новый access токен.
func (s *userService) RefreshToken(refreshToken string) (string, error) {
	claims := &Claims{}
	// Разбираем и проверяем refresh токен
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return "", errors.New("invalid refresh token")
	}
	// Генерируем новый access токен на основе данных refresh токена
	newAccessToken, err := generateAccessToken(claims.UserID)
	if err != nil {
		return "", err
	}
	return newAccessToken, nil
}
