package infrastructure

import (
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var once sync.Once

// NewDatabase はデータベース接続を初期化します
func NewDatabase() {
	once.Do(func() {
		dsn := os.Getenv("DATABASE_URL")
		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) // ローカル変数を使用せず、グローバル変数に代入
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}
	})
}

// GetDatabaseConnection は初期化されたデータベース接続を返します
func GetDatabaseConnection() *gorm.DB {
	return db
}
