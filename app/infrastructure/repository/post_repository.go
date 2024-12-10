package repository

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"polaris-api/domain"
	"polaris-api/domain/models"
	"polaris-api/infrastructure"
)

type PostRepository struct{}

// UploadImageはS3にファイルをアップロードします
func (r *PostRepository) CreatePost(userID, placeID, body string, published bool, fileNames []string) error {
	db := infrastructure.GetDatabaseConnection()

	// 現在時刻を取得
	now := time.Now()

	// PublishedAtの値を設定
	var publishedAt *time.Time
	if published {
		publishedAt = &now
	}

	// Postモデルを作成
	post := &models.Post{
		ID:          uuid.New().String(), // UUIDを生成して設定
		UserID:      userID,
		PlaceID:     placeID,
		Body:        &body,
		Published:   true,
		PublishedAt: publishedAt,
		PostedAt:    time.Now(),
		CreatedAt:   time.Now(),
	}

	// Mediaモデルを作成
	var media []models.Media
	for _, fileName := range fileNames {
		media = append(media, models.Media{
			PostID:     post.ID,
			PlaceID:    placeID,
			MediaType:  "image",
			MediaURL:   fileName,
			UploadedAt: time.Now(),
		})
	}
	return db.Transaction(func(tx *gorm.DB) error {
		// Postを保存
		if err := tx.Create(post).Error; err != nil {
			return domain.Wrap(err, 500, "投稿本文の保存に失敗")
		}

		// Mediaを保存（存在する場合）
		if len(media) > 0 {
			for i := range media {
				media[i].PostID = post.ID
			}
			if err := tx.Create(&media).Error; err != nil {
				return domain.Wrap(err, 500, "メディア情報をデータベースに格納失敗")
			}
		}

		return nil
	})
}
