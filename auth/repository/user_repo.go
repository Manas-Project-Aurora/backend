package repository

import (
	"gorm.io/gorm"

	"github.com/Manas-Project-Aurora/gavna/internal/models"
)

// UserRepository определяет интерфейс для операций с данными пользователя.
type UserRepository interface {
	// CreateUser создает нового пользователя в БД.
	CreateUser(user *models.User) error
	// GetUserByUsername возвращает пользователя по имени.
	GetUserByUsername(username string) (*models.User, error)
}

// userRepo — конкретная реализация интерфейса UserRepository с использованием Gorm.
type userRepo struct {
	db *gorm.DB
}

// NewUserRepository создает новый экземпляр UserRepository.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

// CreateUser создает нового пользователя в базе данных.
func (r *userRepo) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

// GetUserByUsername ищет пользователя по его имени.
func (r *userRepo) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
