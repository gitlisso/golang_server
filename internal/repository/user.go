package repository

import (
	"golang_server/internal/models"
	
	"gorm.io/gorm"
)

// UserRepository интерфейс для работы с пользователями
type UserRepository interface {
	Create(user *models.User) error
	GetByEmail(email string) (*models.User, error)
	GetByID(id uint) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
}

// userRepository реализация репозитория пользователей
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository создает новый репозиторий пользователей
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// Create создает нового пользователя
func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// GetByEmail получает пользователя по email
func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByID получает пользователя по ID
func (r *userRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update обновляет пользователя
func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

// Delete удаляет пользователя
func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
} 