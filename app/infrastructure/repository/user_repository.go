package repository

import (
	"polaris-api/domain"
	"polaris-api/domain/models"
	"polaris-api/infrastructure"
)

type UserRepository struct{}

func (r *UserRepository) FindByUserName(userName string) (*models.User, error) {
	db := infrastructure.GetDatabaseConnection()

	var user models.User
	// user_name を条件に検索
	if err := db.Where("user_name = ?", userName).First(&user).Error; err != nil {
		return nil, domain.Wrap(err, 500, "ユーザー情報取得時にエラー発生")
	}

	return &user, nil
}
