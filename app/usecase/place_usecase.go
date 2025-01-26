package usecase

import (
	"polaris-api/domain"
	"polaris-api/domain/models"
	"polaris-api/infrastructure/repository"
	"polaris-api/interface/types"
)

type PlaceUseCase struct{}

func (u *PlaceUseCase) GetPlaces(
	cursorMID string,
	limit int,
	categoryIds []int,
) (*types.PlacesResponse, error) {
	// リポジトリインスタンスの作成
	repo := &repository.PlaceRepository{}

	// リポジトリからデータ取得
	response, err := repo.FindAll(cursorMID, limit, categoryIds)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (u *PlaceUseCase) GetSuggestPlaces(keywords []string, lon, lat float64) ([]types.SearchPlace, error) {
	repo := &repository.PlaceRepository{}
	places, err := repo.FindPlacesByName(keywords, lon, lat)
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

func (u *PlaceUseCase) CreatePlace(req *types.CreatePlaceRequest) error {
	repo := &repository.PlaceRepository{}
	locationServiceRepo := &repository.LocationServiceRepository{}

	address := ""
	if req.State != "" {
		address += req.State
	}
	if req.City != "" {
		address += req.City
	}
	if req.AddressLine1 != "" {
		address += req.AddressLine1
	}
	if req.AddressLine2 != nil {
		address += *req.AddressLine2
	}

	geometry, err := locationServiceRepo.GeocodeAddress(address)
	if err != nil {
		return err
	}

	err = repo.CreatePlace(req, geometry)
	if err != nil {
		return err
	}

	return nil
}

func (u *PlaceUseCase) GetPlacesNearBySpots(id, cursorMID string, lon, lat float64, limit int) (*types.PlacesResponse, error) {
	repo := &repository.PlaceRepository{}

	places, err := repo.FindNearBySpots(id, cursorMID, lon, lat, limit)
	if err != nil {
		return nil, err
	}

	return places, nil
}

func (u *PlaceUseCase) GetPlacesBaseQuery(
	keywords []string,
	cursorMID string,
	limit int,
) (*types.PlacesResponse, error) {
	// リポジトリインスタンスの作成
	repo := &repository.PlaceRepository{}

	// リポジトリからデータ取得
	response, err := repo.FindPlacesByNameWithMedia(keywords, cursorMID, limit)
	if err != nil {
		return nil, err
	}

	return response, nil
}
