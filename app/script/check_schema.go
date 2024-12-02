//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"polaris-api/domain/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// データベース接続
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 差分確認
	modelsToCheck := []interface{}{
		&models.User{},
		&models.Follows{},
		&models.Place{},
		&models.Event{},
		&models.Favorite{},
		&models.FavoriteFolder{},
		&models.Post{},
		&models.Media{},
		&models.Category{},
		&models.PlaceCategory{},
		&models.Scene{},
		&models.PlaceScene{},
	}

	for _, model := range modelsToCheck {
		fmt.Printf("Checking schema for %T...\n", model)
		if err := checkSchema(db, model); err != nil {
			log.Printf("Schema mismatch for %T: %v\n", model, err)
		} else {
			fmt.Printf("Schema for %T is already up-to-date.\n", model)
		}
	}
}

func checkSchema(db *gorm.DB, model interface{}) error {
	tx := db.Session(&gorm.Session{DryRun: true})
	return tx.AutoMigrate(model)
}
