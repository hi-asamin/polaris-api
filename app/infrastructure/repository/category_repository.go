package repository

import (
	"polaris-api/domain"
	"polaris-api/domain/models"
	"polaris-api/infrastructure"
)

type CategoryRepository struct{}

func (r *CategoryRepository) FindAll() ([]models.Category, error) {
	db := infrastructure.GetDatabaseConnection()

	// 検索結果を格納するスライスを初期化
	categories := []models.Category{}

	if err := db.Find(&categories).Error; err != nil {
		return nil, domain.Wrap(err, 500, "データベースアクセス時にエラー発生")
	}

	return categories, nil
}
