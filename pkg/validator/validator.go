package validator

import (
	"regexp"
	"strings"
	"time"
)

// IsValidEmail проверяет валидность email
func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// IsValidPassword проверяет валидность пароля
func IsValidPassword(password string) bool {
	// Минимум 6 символов
	if len(password) < 6 {
		return false
	}
	return true
}

// IsValidUsername проверяет валидность имени пользователя
func IsValidUsername(username string) bool {
	// Минимум 3 символа, максимум 50
	if len(username) < 3 || len(username) > 50 {
		return false
	}
	
	// Только буквы, цифры и подчеркивания
	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	return usernameRegex.MatchString(username)
}

// IsValidTaskTitle проверяет валидность названия задачи
func IsValidTaskTitle(title string) bool {
	title = strings.TrimSpace(title)
	return len(title) >= 1 && len(title) <= 255
}

// IsValidDateRange проверяет валидность диапазона дат
func IsValidDateRange(startDate, endDate time.Time) bool {
	return !endDate.Before(startDate)
}

// IsValidTaskStatus проверяет валидность статуса задачи
func IsValidTaskStatus(status string) bool {
	validStatuses := []string{"pending", "in_progress", "completed"}
	for _, validStatus := range validStatuses {
		if status == validStatus {
			return true
		}
	}
	return false
}

// SanitizeString очищает строку от лишних пробелов
func SanitizeString(str string) string {
	return strings.TrimSpace(str)
} 