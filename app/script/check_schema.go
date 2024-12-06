//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"polaris-api/domain/models"
	"reflect"

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
		modelType := reflect.TypeOf(model).Elem()
		tableName := db.NamingStrategy.TableName(modelType.Name()) // テーブル名を取得

		fmt.Printf("Checking schema for table: %s\n", tableName)

		// テーブルが存在するか確認
		if !db.Migrator().HasTable(model) {
			log.Printf("Table %s does not exist.\n", tableName)
			continue
		}

		// 各フィールドの確認
		for i := 0; i < modelType.NumField(); i++ {
			field := modelType.Field(i)
			columnName := db.NamingStrategy.ColumnName("", field.Name)

			if !db.Migrator().HasColumn(model, columnName) {
				log.Printf("Column %s in table %s does not exist.\n", columnName, tableName)
			} else {
				fmt.Printf("Column %s in table %s is up-to-date.\n", columnName, tableName)
			}
		}
	}
}
