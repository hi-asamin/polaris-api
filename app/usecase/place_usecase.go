package usecase

import (
	"polaris-api/domain"
	"polaris-api/domain/models"
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

func (u *PlaceUseCase) GetPlaceByID(id string) (*models.Place, error) {
	repo := &repository.PlaceRepository{}
	place, err := repo.FindByID(id)
	if err != nil {
		return nil, domain.Wrap(err, 500, "場所情報の取得に失敗")
	}

	// place が nil の場合にエラーを返す
	if place == nil {
		return nil, domain.New(404, "該当する場所情報が見つかりません")
	}

	return place, nil
}

func (u *PlaceUseCase) GetPlacesNearBySpots(id string, lon, lat float64, limit int) (*dto.PlacesResponse, error) {
	repo := &repository.PlaceRepository{}

	places, err := repo.FindNearBySpots(id, lon, lat, limit)
	if err != nil {
		return nil, domain.Wrap(err, 500, "場所情報の取得に失敗")
	}

	return places, nil
}
