package usecase

import (
	"polaris-api/domain"
	"polaris-api/domain/models"
	"polaris-api/infrastructure/repository"
	"polaris-api/interface/model"
)

type PlaceUseCase struct{}

func (u *PlaceUseCase) GetPlaces(
	cursorPID, cursorMID string,
	limit int,
	categoryIds []int,
) (*model.PlacesResponse, error) {
	// リポジトリインスタンスの作成
	repo := &repository.PlaceRepository{}

	// リポジトリからデータ取得
	response, err := repo.FindAll(cursorPID, cursorMID, limit, categoryIds)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (u *PlaceUseCase) GetPlacesByName(keywords []string, lon, lat float64) ([]model.SearchPlace, error) {
	repo := &repository.PlaceRepository{}
	places, err := repo.SearchPlacesBaseQuery(keywords, lon, lat)
	if err != nil {
		return nil, err
	}

	return places, nil
}

func (u *PlaceUseCase) GetPlaceByID(id string) (*models.Place, error) {
	repo := &repository.PlaceRepository{}
	place, err := repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// place が nil の場合にエラーを返す
	if place == nil {
		return nil, domain.New(404, "該当する場所情報が見つかりません")
	}

	return place, nil
}

func (u *PlaceUseCase) CreatePlace(req *model.CreatePlaceRequest) error {
	repo := &repository.PlaceRepository{}

	err := repo.CreatePlace(req)
	if err != nil {
		return err
	}

	return nil
}

func (u *PlaceUseCase) GetPlacesNearBySpots(id string, lon, lat float64, limit int) (*model.PlacesResponse, error) {
	repo := &repository.PlaceRepository{}

	places, err := repo.FindNearBySpots(id, lon, lat, limit)
	if err != nil {
		return nil, err
	}

	return places, nil
}

func (u *PlaceUseCase) GetPlacesBaseQuery(
	keywords []string,
	cursorPID, cursorMID string,
	limit int,
) (*model.PlacesResponse, error) {
	// リポジトリインスタンスの作成
	repo := &repository.PlaceRepository{}

	// リポジトリからデータ取得
	response, err := repo.FindPlacesBaseQuery(keywords, cursorPID, cursorMID, limit)
	if err != nil {
		return nil, err
	}

	return response, nil
}
