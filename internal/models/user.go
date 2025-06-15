package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User представляет модель пользователя
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"-" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	// Связи
	Tasks []Task `json:"tasks,omitempty" gorm:"foreignKey:UserID"`
}

// CreateUserRequest представляет запрос на создание пользователя
type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginRequest представляет запрос на вход в систему
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// UserResponse представляет ответ с данными пользователя
type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// BeforeCreate хук для хеширования пароля перед созданием
func (u *User) BeforeCreate(tx *gorm.DB) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

// CheckPassword проверяет пароль
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// ToResponse конвертирует модель в ответ
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}
} 