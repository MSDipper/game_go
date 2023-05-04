package db

import (
	"game_go/internal/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"log"
)

var db *gorm.DB

func Init() {
	var err error
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=game_db password=11 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(
		&models.Category{},
		&models.Post{},
		&models.Review{},
		&models.Comment{},
	).Error; err != nil {
		log.Fatal(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
