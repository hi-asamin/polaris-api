package repository

import (
	"polaris-api/domain/models"
	"polaris-api/infrastructure"
)

type MediaRepository struct{}

func (r *MediaRepository) FindByPostID(postID string) ([]models.Media, error) {
	db := infrastructure.GetDatabaseConnection()

	medias := []models.Media{}
	if err := db.Where("post_id = ?", postID).Find(&medias).Error; err != nil {
		return nil, err
	}

	return medias, nil
}
