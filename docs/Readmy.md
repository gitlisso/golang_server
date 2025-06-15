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
git clone https://github.com/gitlisso/golang_server.git
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

## Тестирование API с Postman

### Настройка переменных окружения

Создайте в Postman новое окружение со следующими переменными:

```json
{
  "base_url": "http://localhost:8080",
  "jwt_token": "",
  "user_id": ""
}
```

### Postman коллекция

#### 1. Авторизация

**Регистрация пользователя**
```
POST {{base_url}}/api/auth/register
Content-Type: application/json

{
  "username": "testuser",
  "email": "test@example.com",
  "password": "password123"
}
```

**Вход в систему**
```
POST {{base_url}}/api/auth/login
Content-Type: application/json

{
  "email": "test@example.com",
  "password": "password123"
}

// В Tests добавьте скрипт для сохранения токена:
pm.test("Save JWT token", function () {
    var jsonData = pm.response.json();
    if (jsonData.token) {
        pm.environment.set("jwt_token", jsonData.token);
    }
});
```

**Выход из системы**
```
POST {{base_url}}/api/auth/logout
Authorization: Bearer {{jwt_token}}
```

#### 2. Управление задачами

**Получить список задач**
```
GET {{base_url}}/api/tasks
Authorization: Bearer {{jwt_token}}

// С параметрами:
GET {{base_url}}/api/tasks?status=pending&sort=created_at&order=desc&page=1&limit=10&search=важная
```

**Создать новую задачу**
```
POST {{base_url}}/api/tasks
Authorization: Bearer {{jwt_token}}
Content-Type: application/json

{
  "title": "Важная задача",
  "description": "Описание важной задачи",
  "status": "pending",
  "start_date": "2024-01-15T09:00:00Z",
  "end_date": "2024-01-20T18:00:00Z"
}

// В Tests добавьте скрипт для сохранения ID задачи:
pm.test("Save task ID", function () {
    var jsonData = pm.response.json();
    if (jsonData.task && jsonData.task.id) {
        pm.environment.set("task_id", jsonData.task.id);
    }
});
```

**Получить задачу по ID**
```
GET {{base_url}}/api/tasks/{{task_id}}
Authorization: Bearer {{jwt_token}}
```

**Обновить задачу**
```
PUT {{base_url}}/api/tasks/{{task_id}}
Authorization: Bearer {{jwt_token}}
Content-Type: application/json

{
  "title": "Обновленная задача",
  "description": "Обновленное описание",
  "status": "in_progress",
  "start_date": "2024-01-15T09:00:00Z",
  "end_date": "2024-01-22T18:00:00Z"
}
```

**Удалить задачу**
```
DELETE {{base_url}}/api/tasks/{{task_id}}
Authorization: Bearer {{jwt_token}}
```

### Примеры ответов API

**Успешная регистрация (201)**
```json
{
  "message": "Пользователь успешно зарегистрирован",
  "user": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "created_at": "2024-01-15T12:00:00Z"
  }
}
```

**Успешный вход (200)**
```json
{
  "message": "Успешный вход в систему",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com"
  }
}
```

**Список задач (200)**
```json
{
  "tasks": [
    {
      "id": 1,
      "title": "Важная задача",
      "description": "Описание важной задачи",
      "status": "pending",
      "start_date": "2024-01-15T09:00:00Z",
      "end_date": "2024-01-20T18:00:00Z",
      "user_id": 1,
      "created_at": "2024-01-15T12:00:00Z",
      "updated_at": "2024-01-15T12:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 1
  }
}
```

**Ошибка авторизации (401)**
```json
{
  "error": "Необходима авторизация"
}
```

**Ошибка валидации (400)**
```json
{
  "error": "Ошибка валидации",
  "details": [
    "Поле 'title' обязательно для заполнения"
  ]
}
```

### Коллекция Postman (JSON)

Скопируйте и импортируйте следующую коллекцию:

```json
{
  "info": {
    "name": "Todo App API",
    "description": "REST API для управления задачами",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
  },
  "variable": [
    {
      "key": "base_url",
      "value": "http://localhost:8080"
    }
  ],
  "item": [
    {
      "name": "Auth",
      "item": [
        {
          "name": "Register",
          "request": {
            "method": "POST",
            "header": [
              {
                "key": "Content-Type",
                "value": "application/json"
              }
            ],
            "body": {
              "mode": "raw",
              "raw": "{\n  \"username\": \"testuser\",\n  \"email\": \"test@example.com\",\n  \"password\": \"password123\"\n}"
            },
            "url": {
              "raw": "{{base_url}}/api/auth/register",
              "host": ["{{base_url}}"],
              "path": ["api", "auth", "register"]
            }
          }
        },
        {
          "name": "Login",
          "event": [
            {
              "listen": "test",
              "script": {
                "exec": [
                  "pm.test(\"Save JWT token\", function () {",
                  "    var jsonData = pm.response.json();",
                  "    if (jsonData.token) {",
                  "        pm.environment.set(\"jwt_token\", jsonData.token);",
                  "    }",
                  "});"
                ]
              }
            }
          ],
          "request": {
            "method": "POST",
            "header": [
              {
                "key": "Content-Type",
                "value": "application/json"
              }
            ],
            "body": {
              "mode": "raw",
              "raw": "{\n  \"email\": \"test@example.com\",\n  \"password\": \"password123\"\n}"
            },
            "url": {
              "raw": "{{base_url}}/api/auth/login",
              "host": ["{{base_url}}"],
              "path": ["api", "auth", "login"]
            }
          }
        }
      ]
    },
    {
      "name": "Tasks",
      "item": [
        {
          "name": "Get Tasks",
          "request": {
            "method": "GET",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{jwt_token}}"
              }
            ],
            "url": {
              "raw": "{{base_url}}/api/tasks",
              "host": ["{{base_url}}"],
              "path": ["api", "tasks"]
            }
          }
        },
        {
          "name": "Create Task",
          "event": [
            {
              "listen": "test",
              "script": {
                "exec": [
                  "pm.test(\"Save task ID\", function () {",
                  "    var jsonData = pm.response.json();",
                  "    if (jsonData.task && jsonData.task.id) {",
                  "        pm.environment.set(\"task_id\", jsonData.task.id);",
                  "    }",
                  "});"
                ]
              }
            }
          ],
          "request": {
            "method": "POST",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{jwt_token}}"
              },
              {
                "key": "Content-Type",
                "value": "application/json"
              }
            ],
            "body": {
              "mode": "raw",
              "raw": "{\n  \"title\": \"Важная задача\",\n  \"description\": \"Описание важной задачи\",\n  \"status\": \"pending\",\n  \"start_date\": \"2024-01-15T09:00:00Z\",\n  \"end_date\": \"2024-01-20T18:00:00Z\"\n}"
            },
            "url": {
              "raw": "{{base_url}}/api/tasks",
              "host": ["{{base_url}}"],
              "path": ["api", "tasks"]
            }
          }
        }
      ]
    }
  ]
}
```

### Инструкции по импорту

1. Откройте Postman
2. Нажмите "Import" в верхней части интерфейса
3. Вставьте JSON коллекции выше или создайте файл `.json`
4. Настройте переменные окружения
5. Запустите сервер: `go run cmd/server/main.go`
6. Начните тестирование с регистрации пользователя

### Автоматические тесты

Добавьте в каждый запрос следующие тесты:

```javascript
// Проверка статуса ответа
pm.test("Status code is 200", function () {
    pm.response.to.have.status(200);
});

// Проверка формата JSON
pm.test("Response is JSON", function () {
    pm.response.to.be.json;
});

// Проверка времени ответа
pm.test("Response time is less than 500ms", function () {
    pm.expect(pm.response.responseTime).to.be.below(500);
});
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
