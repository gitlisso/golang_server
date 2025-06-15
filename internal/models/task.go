package models

import (
	"time"
)

// TaskStatus представляет статус задачи
type TaskStatus string

const (
	TaskStatusPending    TaskStatus = "pending"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusCompleted  TaskStatus = "completed"
)

// Task представляет модель задачи
type Task struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Title       string     `json:"title" gorm:"not null"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status" gorm:"default:'pending'"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     time.Time  `json:"end_date"`
	UserID      uint       `json:"user_id" gorm:"not null"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	
	// Связи
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// CreateTaskRequest представляет запрос на создание задачи
type CreateTaskRequest struct {
	Title       string    `json:"title" binding:"required,min=1,max=255"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     time.Time `json:"end_date" binding:"required"`
}

// UpdateTaskRequest представляет запрос на обновление задачи
type UpdateTaskRequest struct {
	Title       *string     `json:"title,omitempty" binding:"omitempty,min=1,max=255"`
	Description *string     `json:"description,omitempty"`
	Status      *TaskStatus `json:"status,omitempty" binding:"omitempty,oneof=pending in_progress completed"`
	StartDate   *time.Time  `json:"start_date,omitempty"`
	EndDate     *time.Time  `json:"end_date,omitempty"`
}

// TaskResponse представляет ответ с данными задачи
type TaskResponse struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     time.Time  `json:"end_date"`
	UserID      uint       `json:"user_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// TaskQueryParams представляет параметры запроса для получения задач
type TaskQueryParams struct {
	Status string `form:"status"`
	Sort   string `form:"sort"`
	Order  string `form:"order"`
	Search string `form:"search"`
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
}

// ToResponse конвертирует модель в ответ
func (t *Task) ToResponse() TaskResponse {
	return TaskResponse{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		StartDate:   t.StartDate,
		EndDate:     t.EndDate,
		UserID:      t.UserID,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

// IsValidStatus проверяет валидность статуса
func (t TaskStatus) IsValid() bool {
	switch t {
	case TaskStatusPending, TaskStatusInProgress, TaskStatusCompleted:
		return true
	}
	return false
} 