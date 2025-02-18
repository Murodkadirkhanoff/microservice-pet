package db

import (
	"fmt"
	"log"

	entity "github.com/Murodkadirkhanoff/pet-microservice-golang-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// DSN (Data Source Name)
	dsn := "host=postgres_user_service port=5432 user=postgres password=secret dbname=user_service sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = db
}

func MigrateDB() {
	// Миграция моделей
	if err := DB.AutoMigrate(&entity.User{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	fmt.Println("Migration completed")
}
