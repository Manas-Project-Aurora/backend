package handlers

import (
	"net/http"

	"github.com/Manas-Project-Aurora/gavna/auth/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// RegisterRequest представляет запрос на регистрацию
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginRequest представляет запрос на вход
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// TokenRequest представляет запрос на обновление токена
type TokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// Register обрабатывает запрос на регистрацию
func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Используем email как username
	username := req.Email

	user, err := h.service.Register(username, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user_id": user.ID,
	})
}

// Login обрабатывает запрос на вход
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := req.Email
	tokens, err := h.service.Login(username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Устанавливаем access token и refresh token в cookie
	c.SetCookie("access_token", tokens.AccessToken, 3600, "/", "", true, true)
	c.SetCookie("refresh_token", tokens.RefreshToken, 604800, "/", "", true, true) // 7 дней

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

// RefreshToken обрабатывает запрос на обновление токена
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No refresh token provided"})
		return
	}

	tokens, err := h.service.RefreshToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	// Обновляем токены в cookie
	c.SetCookie("access_token", tokens.AccessToken, 3600, "/", "", true, true)
	c.SetCookie("refresh_token", tokens.RefreshToken, 604800, "/", "", true, true)

	c.JSON(http.StatusOK, gin.H{"message": "Token refreshed"})
}

// Logout обрабатывает запрос на выход
func (h *AuthHandler) Logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No refresh token provided"})
		return
	}

	if err := h.service.Logout(refreshToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Очищаем куки
	c.SetCookie("access_token", "", -1, "/", "", true, true)
	c.SetCookie("refresh_token", "", -1, "/", "", true, true)

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
