package usecase

import (
	"polaris-api/domain"
	"polaris-api/infrastructure/repository"
	"polaris-api/interface/dto"
)

type PlaceUseCase struct {}

func (u *PlaceUseCase) GetPlaceByID(name string, lon, lat float64) ([]dto.SearchPlace, error) {
	repo := &repository.PlaceRepository{}
	places, err := repo.FindByName(name, lon, lat)
	if err != nil {
		return nil, domain.Wrap(err, 500, "場所情報の取得に失敗")
	}

	return places, nil
}
