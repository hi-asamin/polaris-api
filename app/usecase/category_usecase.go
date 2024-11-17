package usecase

import (
	"polaris-api/domain/models"
	"polaris-api/infrastructure/repository"
)

type CategoryUseCase struct{}

func (u *CategoryUseCase) GetCategories() ([]models.Category, error) {
	repo := &repository.CategoryRepository{}
	categories, err := repo.FindAll()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
