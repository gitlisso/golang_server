package repository

import (
	"golang_server/internal/models"
	
	"gorm.io/gorm"
)

// TaskRepository интерфейс для работы с задачами
type TaskRepository interface {
	Create(task *models.Task) error
	GetByID(id uint) (*models.Task, error)
	GetByUserID(userID uint, params models.TaskQueryParams) ([]models.Task, int64, error)
	Update(task *models.Task) error
	Delete(id uint) error
}

// taskRepository реализация репозитория задач
type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepository создает новый репозиторий задач
func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{
		db: db,
	}
}

// Create создает новую задачу
func (r *taskRepository) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

// GetByID получает задачу по ID
func (r *taskRepository) GetByID(id uint) (*models.Task, error) {
	var task models.Task
	err := r.db.Preload("User").First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// GetByUserID получает задачи пользователя с фильтрацией и пагинацией
func (r *taskRepository) GetByUserID(userID uint, params models.TaskQueryParams) ([]models.Task, int64, error) {
	var tasks []models.Task
	var total int64

	query := r.db.Where("user_id = ?", userID)

	// Фильтрация по статусу
	if params.Status != "" {
		query = query.Where("status = ?", params.Status)
	}

	// Поиск по названию
	if params.Search != "" {
		query = query.Where("title ILIKE ?", "%"+params.Search+"%")
	}

	// Подсчет общего количества
	query.Model(&models.Task{}).Count(&total)

	// Сортировка
	orderBy := "created_at"
	if params.Sort != "" {
		switch params.Sort {
		case "created_at", "start_date", "end_date", "status", "title":
			orderBy = params.Sort
		}
	}

	order := "DESC"
	if params.Order == "asc" {
		order = "ASC"
	}

	query = query.Order(orderBy + " " + order)

	// Пагинация
	if params.Page > 0 && params.Limit > 0 {
		offset := (params.Page - 1) * params.Limit
		query = query.Offset(offset).Limit(params.Limit)
	}

	err := query.Find(&tasks).Error
	return tasks, total, err
}

// Update обновляет задачу
func (r *taskRepository) Update(task *models.Task) error {
	return r.db.Save(task).Error
}

// Delete удаляет задачу
func (r *taskRepository) Delete(id uint) error {
	return r.db.Delete(&models.Task{}, id).Error
} 