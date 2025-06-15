# Приложение "Список дел" (Todo App) - REST API Backend

## Описание

REST API сервер для управления задачами с функциями регистрации, авторизации и полного CRUD для задач. Предназначен для работы с фронтенд приложениями (веб, мобильные приложения).

## Технологии

- **Go** - основной язык программирования
- **Gin** - веб-фреймворк для создания REST API
- **GORM** - ORM для работы с базой данных
- **SQLite** - легковесная база данных
- **JWT** - токены для авторизации
- **bcrypt** - хеширование паролей

## Функционал

- ✅ Регистрация и авторизация пользователей
- ✅ Создание задачи (название, дата начала, дата окончания, статус)
- ✅ Просмотр списка задач
- ✅ Редактирование задачи
- ✅ Удаление задачи
- ✅ Поиск задач
- ✅ Сортировка задач по статусу, дате начала, дате окончания
- ✅ Фильтрация задач по пользователю

## Структура проекта

```
golang_server/
├── cmd/
│   └── server/
│       └── main.go              # Точка входа в приложение
├── internal/
│   ├── config/
│   │   └── config.go            # Конфигурация приложения
│   ├── handlers/
│   │   ├── auth.go              # Обработчики авторизации
│   │   └── tasks.go             # Обработчики задач
│   ├── middleware/
│   │   ├── auth.go              # Middleware для авторизации
│   │   └── cors.go              # Middleware для CORS
│   ├── models/
│   │   ├── user.go              # Модель пользователя
│   │   └── task.go              # Модель задачи
│   ├── repository/
│   │   ├── user.go              # Репозиторий пользователей
│   │   └── task.go              # Репозиторий задач
│   ├── services/
│   │   ├── auth.go              # Сервис авторизации
│   │   └── task.go              # Сервис задач
│   └── database/
│       └── database.go          # Подключение к БД
├── pkg/
│   ├── utils/
│   │   └── jwt.go               # Утилиты для JWT
│   └── validator/
│       └── validator.go         # Валидация данных
├── docs/
│   └── README.md                # Документация проекта
├── .env                         # Переменные окружения
├── .gitignore                   # Исключения для Git
├── go.mod                       # Модуль Go
├── go.sum                       # Контрольные суммы зависимостей
└── database.db                  # SQLite база данных
```

## Установка и запуск

### Предварительные требования

- Go 1.21 или выше
- Git

### Шаги установки

1. **Клонирование репозитория:**
```bash
git clone <repository-url>
cd golang_server
```

2. **Установка зависимостей:**
```bash
go mod tidy
```

3. **Создание файла .env:**
```env
PORT=8080
DB_PATH=./database.db
JWT_SECRET=your-secret-key-here
GIN_MODE=debug
```

4. **Запуск приложения:**
```bash
go run cmd/server/main.go
```

Сервер будет доступен по адресу: `http://localhost:8080`

### Альтернативный запуск

```bash
# Сборка приложения
go build -o app cmd/server/main.go

# Запуск собранного приложения
./app
```

## API Endpoints

### Авторизация

- `POST /api/auth/register` - Регистрация пользователя
- `POST /api/auth/login` - Вход в систему
- `POST /api/auth/logout` - Выход из системы

### Задачи (требуют авторизации)

- `GET /api/tasks` - Получить список задач
- `POST /api/tasks` - Создать новую задачу
- `GET /api/tasks/:id` - Получить задачу по ID
- `PUT /api/tasks/:id` - Обновить задачу
- `DELETE /api/tasks/:id` - Удалить задачу

### Параметры запросов

#### GET /api/tasks
- `status` - фильтр по статусу (pending, in_progress, completed)
- `sort` - сортировка (created_at, start_date, end_date, status)
- `order` - порядок сортировки (asc, desc)
- `search` - поиск по названию
- `page` - номер страницы
- `limit` - количество элементов на странице

## Модели данных

### User (Пользователь)
```go
type User struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Username  string    `json:"username" gorm:"unique;not null"`
    Email     string    `json:"email" gorm:"unique;not null"`
    Password  string    `json:"-" gorm:"not null"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

### Task (Задача)
```go
type Task struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Title       string    `json:"title" gorm:"not null"`
    Description string    `json:"description"`
    Status      string    `json:"status" gorm:"default:'pending'"`
    StartDate   time.Time `json:"start_date"`
    EndDate     time.Time `json:"end_date"`
    UserID      uint      `json:"user_id" gorm:"not null"`
    User        User      `json:"user" gorm:"foreignKey:UserID"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

## Безопасность

- Пароли хешируются с использованием bcrypt
- JWT токены для аутентификации
- Middleware для проверки авторизации
- Валидация входных данных
- CORS настройки

## Разработка

### Линтинг
```bash
go fmt ./...
go vet ./...
```

### Тестирование
```bash
go test ./...
```

### Сборка для production
```bash
GIN_MODE=release go build -o app cmd/server/main.go
```

## Дополнительные возможности

- Логирование запросов
- Пагинация результатов
- Валидация данных
- Обработка ошибок
- Graceful shutdown
- Документация API (Swagger)
- CORS поддержка для фронтенд приложений
- JSON API responses
- Готов для интеграции с любыми фронтенд фреймворками

## Лицензия

MIT License

## Авторы

Проект разработан для изучения современного стека технологий Go.

---
