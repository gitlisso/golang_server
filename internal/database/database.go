package database

import (
	"golang_server/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "modernc.org/sqlite"
)

// Init инициализирует подключение к базе данных
func Init(databasePath string) (*gorm.DB, error) {
	// Настройки GORM
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// Подключаемся к SQLite с modernc.org/sqlite драйвером
	db, err := gorm.Open(sqlite.Dialector{
		DriverName: "sqlite",
		DSN:        databasePath,
	}, config)
	if err != nil {
		return nil, err
	}

	// Выполняем миграции
	err = db.AutoMigrate(
		&models.User{},
		&models.Task{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// GetDB возвращает экземпляр базы данных
func GetDB(db *gorm.DB) *gorm.DB {
	return db
}
