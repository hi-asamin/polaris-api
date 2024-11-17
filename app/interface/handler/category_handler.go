package handler

import (
	"polaris-api/domain/models"
	"polaris-api/usecase"
)

type CategoryHandler struct{}

func (h *CategoryHandler) GetCategories() ([]models.Category, error) {
	u := &usecase.CategoryUseCase{}
	categories, err := u.GetCategories()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
