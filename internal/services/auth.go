package services

import (
	"errors"

	"golang_server/internal/models"
	"golang_server/internal/repository"
	"golang_server/pkg/utils"

	"gorm.io/gorm"
)

// AuthService интерфейс для сервиса авторизации
type AuthService interface {
	Register(req models.CreateUserRequest) (*models.UserResponse, error)
	Login(req models.LoginRequest) (string, *models.UserResponse, error)
}

// authService реализация сервиса авторизации
type authService struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

// NewAuthService создает новый сервис авторизации
func NewAuthService(userRepo repository.UserRepository, jwtSecret string) AuthService {
	return &authService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

// Register регистрирует нового пользователя
func (s *authService) Register(req models.CreateUserRequest) (*models.UserResponse, error) {
	// Проверяем, существует ли пользователь с таким email
	existingUser, err := s.userRepo.GetByEmail(req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Создаем нового пользователя
	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password, // Пароль будет хеширован в BeforeCreate хуке
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	userResponse := user.ToResponse()
	return &userResponse, nil
}

// Login авторизует пользователя
func (s *authService) Login(req models.LoginRequest) (string, *models.UserResponse, error) {
	// Находим пользователя по email
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil, errors.New("invalid email or password")
		}
		return "", nil, err
	}

	// Проверяем пароль
	if !user.CheckPassword(req.Password) {
		return "", nil, errors.New("invalid email or password")
	}

	// Генерируем JWT токен
	token, err := utils.GenerateToken(user.ID, user.Email, s.jwtSecret)
	if err != nil {
		return "", nil, err
	}

	userResponse := user.ToResponse()
	return token, &userResponse, nil
} 