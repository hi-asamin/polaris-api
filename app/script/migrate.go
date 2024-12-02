//go:build ignore
// +build ignore

package main

import (
	"log"
	"os"
	"polaris-api/domain/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// データベース接続情報を取得
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_DSN is not set in environment variables")
	}

	// データベース接続
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 必要に応じてログレベルを変更
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// マイグレーションを実行
	log.Println("Starting database migration...")
	err = db.AutoMigrate(
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
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration completed successfully!")
}
