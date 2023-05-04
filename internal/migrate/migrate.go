package main

// Импортируем библиотеки
import (
	"game_go/internal/database"
	"game_go/internal/models"
	"log"
)

func main() {
	// Подключаемся к базе данных
	if err := db.Connect(); err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer db.Close()

	// Создаем таблицы
	db.DB.AutoMigrate(
		&models.Category{},
		&models.Post{},
		&models.Review{},
		&models.Comment{},
	)

	// Теперь таблицы должны быть созданы в базе данных
}
