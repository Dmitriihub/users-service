package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/Dmitriihub/users-service/internal/user"
)

var DB *gorm.DB

func InitDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("не удалось подключиться к базе данных: %v", err)
	}

	err = DB.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}
}
