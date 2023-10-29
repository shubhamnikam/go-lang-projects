package services

import (
	"log"
	"os"

	"github.com/shubhamnikam/go-lang-projects/crud-with-fiber-gorm-sqlite/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseContext struct {
	Db *gorm.DB
}

var DbContext DatabaseContext

func DbConnect() {
	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_NAME")), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to db:", os.Getenv("DB_NAME"), "\n", err.Error())
	}

	log.Println("Connected to db:", os.Getenv("DB_NAME"))
	db.Logger = logger.Default.LogMode(logger.Info)
	
	handleMigration(db)
	
	DbContext = DatabaseContext{Db: db}
}

func handleMigration(db *gorm.DB) {
	log.Println("Running migrations...")
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
}