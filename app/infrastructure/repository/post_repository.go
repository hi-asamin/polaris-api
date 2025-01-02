package repository

import (
	"time"

	"gorm.io/gorm"

	"polaris-api/domain"
	"polaris-api/domain/models"
	"polaris-api/infrastructure"
	"polaris-api/interface/types"
)

type PostRepository struct{}

// UploadImageはS3にファイルをアップロードします
func (r *PostRepository) CreatePost(userID, placeID, placeName, body string, published bool, fileInfos []types.FileInfo) error {
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
	for _, fileInfo := range fileInfos {
		media = append(media, models.Media{
			PostID:     post.ID,
			PlaceID:    placeID,
			MediaType:  fileInfo.FileType,
			MediaURL:   fileInfo.FileName,
			AltText:    &placeName,
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

func (r *PostRepository) DeletePost(postID string) error {
	db := infrastructure.GetDatabaseConnection()

	// トランザクション開始
	return db.Transaction(func(tx *gorm.DB) error {
		// 関連するメディアを先に削除
		if err := tx.Where("post_id = ?", postID).Delete(&models.Media{}).Error; err != nil {
			return domain.Wrap(err, 500, "メディア情報の削除に失敗")
		}

		// 投稿を削除
		if err := tx.Where("id = ?", postID).Delete(&models.Post{}).Error; err != nil {
			return domain.Wrap(err, 500, "投稿の削除に失敗")
		}

		return nil
	})
}
