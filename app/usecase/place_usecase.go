package usecase

import (
	"polaris-api/domain"
	"polaris-api/infrastructure/repository"
	"polaris-api/interface/dto"
)

type PlaceUseCase struct{}

func (u *PlaceUseCase) GetPlaces(
	lon, lat float64,
	cursorDistance float64, cursorPID, cursorMID string,
	limit int,
) (*dto.PlacesResponse, error) {
	// リポジトリインスタンスの作成
	repo := &repository.PlaceRepository{}

	// リポジトリからデータ取得
	response, err := repo.FindAll(lon, lat, cursorDistance, cursorPID, cursorMID, limit)
	if err != nil {
		return nil, domain.Wrap(err, 500, "場所情報の取得に失敗")
	}

	return response, nil
}

func (u *PlaceUseCase) GetPlacesByName(name string, lon, lat float64) ([]dto.SearchPlace, error) {
	repo := &repository.PlaceRepository{}
	places, err := repo.FindByName(name, lon, lat)
	if err != nil {
		return nil, domain.Wrap(err, 500, "場所情報の取得に失敗")
	}

	return places, nil
}
