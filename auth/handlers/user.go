package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/Manas-Project-Aurora/gavna/auth/repository"
	"github.com/Manas-Project-Aurora/gavna/auth/services"
)

// AuthHandler обрабатывает HTTP-запросы, связанные с аутентификацией.
type AuthHandler struct {
	userService services.UserService
}

// NewAuthHandler создает новый экземпляр AuthHandler,
// инициализируя репозиторий и сервис для работы с пользователями.
func NewAuthHandler(db *gorm.DB) *AuthHandler {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	return &AuthHandler{
		userService: userService,
	}
}

// RegisterAuthRoutes регистрирует маршруты, связанные с аутентификацией.
// Реализованы эндпоинты: /register, /login, /token (refresh) и /logout.
func RegisterAuthRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	handler := NewAuthHandler(db)

	rg.POST("/register", handler.Register) // Регистрация нового пользователя
	rg.POST("/login", handler.Login)       // Вход пользователя (аутентификация)
	rg.POST("/token", handler.Refresh)     // Обновление access токена по refresh токену
	rg.POST("/logout", handler.Logout)     // Выход из профиля
}

// Register обрабатывает регистрацию нового пользователя.
// Принимает JSON с данными: username, password и telegram.
func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Telegram string `json:"telegram" binding:"required"`
	}
	// Привязываем JSON тело запроса к структуре
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Вызываем сервис для регистрации пользователя
	user, err := h.userService.RegisterUser(req.Username, req.Password, req.Telegram)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Возвращаем данные созданного пользователя (без пароля)
	c.JSON(http.StatusCreated, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"telegram": user.Telegram,
	})
}

// Login обрабатывает вход пользователя и возвращает пару токенов.
// В теле запроса ожидается username и password.
func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	// Привязываем входящие данные
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Аутентифицируем пользователя и получаем access и refresh токены
	accessToken, refreshToken, err := h.userService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	// Возвращаем токены в ответе
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// Refresh обрабатывает запрос на обновление access токена.
// Принимается JSON с полем "refresh_token".
func (h *AuthHandler) Refresh(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	// Проверяем наличие refresh токена в теле запроса
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Генерируем новый access токен, используя предоставленный refresh токен
	newAccessToken, err := h.userService.RefreshToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	// Возвращаем новый access токен
	c.JSON(http.StatusOK, gin.H{"access_token": newAccessToken})
}

// Logout обрабатывает выход пользователя.
// Для stateless JWT, как правило, дополнительная логика не требуется – просто возвращаем успешный ответ.
// При желании можно реализовать механизм отзыва токенов (blacklisting).
func (h *AuthHandler) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
