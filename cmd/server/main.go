package main

import (
	"log"
	"os"

	"golang_server/internal/config"
	"golang_server/internal/database"
	"golang_server/internal/handlers"
	"golang_server/internal/middleware"
	"golang_server/internal/repository"
	"golang_server/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Загружаем переменные окружения
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Инициализируем конфигурацию
	cfg := config.New()

	// Подключаемся к базе данных
	db, err := database.Init(cfg.DatabasePath)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Создаем репозитории
	userRepo := repository.NewUserRepository(db)
	taskRepo := repository.NewTaskRepository(db)

	// Создаем сервисы
	authService := services.NewAuthService(userRepo, cfg.JWTSecret)
	taskService := services.NewTaskService(taskRepo)

	// Создаем обработчики
	authHandler := handlers.NewAuthHandler(authService)
	taskHandler := handlers.NewTaskHandler(taskService)

	// Настраиваем Gin
	r := gin.Default()

	// Применяем middleware
	r.Use(middleware.CORS())

	// Публичные маршруты
	auth := r.Group("/api/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	// Защищенные маршруты
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	{
		api.GET("/tasks", taskHandler.GetTasks)
		api.POST("/tasks", taskHandler.CreateTask)
		api.GET("/tasks/:id", taskHandler.GetTask)
		api.PUT("/tasks/:id", taskHandler.UpdateTask)
		api.DELETE("/tasks/:id", taskHandler.DeleteTask)
	}

	// Запускаем сервер
	log.Printf("Server starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
} 