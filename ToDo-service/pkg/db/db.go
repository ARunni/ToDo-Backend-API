package db

import (
	"log"

	"github.com/ARunni/ToDo-Backend-API/ToDo-service/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Todo{})
	return Handler{db}
}
