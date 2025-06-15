# Todo App - REST API Backend

REST API сервер для управления задачами на Go с авторизацией через JWT.

## 🚀 Быстрый старт

### Предварительные требования
- Go 1.21 или выше
- Git

### Установка и запуск

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

### Альтернативные способы запуска

**Сборка и запуск:**
```bash
go build -o app cmd/server/main.go
./app
```

**Production сборка:**
```bash
GIN_MODE=release go build -o app cmd/server/main.go
./app
```

## 📖 Полная документация

Подробная документация, включая описание API endpoints, модели данных, примеры использования и Postman коллекцию для тестирования доступна в: **[docs/Readmy.md](docs/Readmy.md)**

## 🔗 API Endpoints

- `POST /api/auth/register` - Регистрация
- `POST /api/auth/login` - Авторизация  
- `GET /api/tasks` - Список задач
- `POST /api/tasks` - Создать задачу
- `PUT /api/tasks/:id` - Обновить задачу
- `DELETE /api/tasks/:id` - Удалить задачу

## 🧪 Тестирование

Для тестирования API используйте Postman коллекцию из документации или выполните:

```bash
# Регистрация пользователя
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"test","email":"test@example.com","password":"password123"}'

# Авторизация
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'
```

## 📋 Технологии

- **Go** - основной язык
- **Gin** - веб-фреймворк  
- **GORM** - ORM для работы с БД
- **SQLite** - база данных
- **JWT** - авторизация
- **bcrypt** - хеширование паролей

## 📄 Лицензия

MIT License
