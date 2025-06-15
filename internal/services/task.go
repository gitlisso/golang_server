package services

import (
	"errors"

	"golang_server/internal/models"
	"golang_server/internal/repository"

	"gorm.io/gorm"
)

// TaskService интерфейс для сервиса задач
type TaskService interface {
	CreateTask(userID uint, req models.CreateTaskRequest) (*models.TaskResponse, error)
	GetTasks(userID uint, params models.TaskQueryParams) ([]models.TaskResponse, int64, error)
	GetTaskByID(userID, taskID uint) (*models.TaskResponse, error)
	UpdateTask(userID, taskID uint, req models.UpdateTaskRequest) (*models.TaskResponse, error)
	DeleteTask(userID, taskID uint) error
}

// taskService реализация сервиса задач
type taskService struct {
	taskRepo repository.TaskRepository
}

// NewTaskService создает новый сервис задач
func NewTaskService(taskRepo repository.TaskRepository) TaskService {
	return &taskService{
		taskRepo: taskRepo,
	}
}

// CreateTask создает новую задачу
func (s *taskService) CreateTask(userID uint, req models.CreateTaskRequest) (*models.TaskResponse, error) {
	// Проверяем, что дата окончания не раньше даты начала
	if req.EndDate.Before(req.StartDate) {
		return nil, errors.New("end date cannot be before start date")
	}

	task := &models.Task{
		Title:       req.Title,
		Description: req.Description,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		Status:      models.TaskStatusPending,
		UserID:      userID,
	}

	if err := s.taskRepo.Create(task); err != nil {
		return nil, err
	}

	taskResponse := task.ToResponse()
	return &taskResponse, nil
}

// GetTasks получает список задач пользователя
func (s *taskService) GetTasks(userID uint, params models.TaskQueryParams) ([]models.TaskResponse, int64, error) {
	// Устанавливаем значения по умолчанию для пагинации
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.Limit <= 0 {
		params.Limit = 10
	}
	if params.Limit > 100 {
		params.Limit = 100
	}

	tasks, total, err := s.taskRepo.GetByUserID(userID, params)
	if err != nil {
		return nil, 0, err
	}

	taskResponses := make([]models.TaskResponse, len(tasks))
	for i, task := range tasks {
		taskResponses[i] = task.ToResponse()
	}

	return taskResponses, total, nil
}

// GetTaskByID получает задачу по ID
func (s *taskService) GetTaskByID(userID, taskID uint) (*models.TaskResponse, error) {
	task, err := s.taskRepo.GetByID(taskID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("task not found")
		}
		return nil, err
	}

	// Проверяем, что задача принадлежит пользователю
	if task.UserID != userID {
		return nil, errors.New("access denied")
	}

	taskResponse := task.ToResponse()
	return &taskResponse, nil
}

// UpdateTask обновляет задачу
func (s *taskService) UpdateTask(userID, taskID uint, req models.UpdateTaskRequest) (*models.TaskResponse, error) {
	task, err := s.taskRepo.GetByID(taskID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("task not found")
		}
		return nil, err
	}

	// Проверяем, что задача принадлежит пользователю
	if task.UserID != userID {
		return nil, errors.New("access denied")
	}

	// Обновляем поля, если они предоставлены
	if req.Title != nil {
		task.Title = *req.Title
	}
	if req.Description != nil {
		task.Description = *req.Description
	}
	if req.Status != nil {
		if !req.Status.IsValid() {
			return nil, errors.New("invalid status")
		}
		task.Status = *req.Status
	}
	if req.StartDate != nil {
		task.StartDate = *req.StartDate
	}
	if req.EndDate != nil {
		task.EndDate = *req.EndDate
	}

	// Проверяем, что дата окончания не раньше даты начала
	if task.EndDate.Before(task.StartDate) {
		return nil, errors.New("end date cannot be before start date")
	}

	if err := s.taskRepo.Update(task); err != nil {
		return nil, err
	}

	taskResponse := task.ToResponse()
	return &taskResponse, nil
}

// DeleteTask удаляет задачу
func (s *taskService) DeleteTask(userID, taskID uint) error {
	task, err := s.taskRepo.GetByID(taskID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("task not found")
		}
		return err
	}

	// Проверяем, что задача принадлежит пользователю
	if task.UserID != userID {
		return errors.New("access denied")
	}

	return s.taskRepo.Delete(taskID)
}