package handler

import (
	"strconv"

	"polaris-api/domain"
	"polaris-api/domain/models"
	"polaris-api/interface/dto"
	"polaris-api/usecase"
)

type PlaceHandler struct{}

func (h *PlaceHandler) GetPlaces(
	lonStr, latStr, cursorDistanceStr, cursorPID, cursorMID, limitStr string,
) (*dto.PlacesResponse, error) {
	// クエリパラメータの検証と変換
	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		return nil, domain.Wrap(err, 400, "緯度のパラメータエラー")
	}

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		return nil, domain.Wrap(err, 400, "経度のパラメータエラー")
	}

	cursorDistance := 0.0
	if cursorDistanceStr != "" {
		cursorDistance, err = strconv.ParseFloat(cursorDistanceStr, 64)
		if err != nil {
			return nil, domain.Wrap(err, 400, "カーソル距離のパラメータエラー")
		}
	}

	limit := 20 // デフォルト値
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			return nil, domain.Wrap(err, 400, "リミットのパラメータエラー")
		}
	}

	// Usecaseの呼び出し
	u := &usecase.PlaceUseCase{}
	response, err := u.GetPlaces(lon, lat, cursorDistance, cursorPID, cursorMID, limit)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (h *PlaceHandler) GetPlacesByName(q, lonStr, latStr string) ([]dto.SearchPlace, error) {
	// 検索ワードが空の場合は空配列を返却する
	if q == "" {
		return []dto.SearchPlace{}, nil
	}

	// クエリパラメータの検証と変換
	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		return nil, domain.Wrap(err, 400, "パラメーターエラー発生")
	}

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		return nil, domain.Wrap(err, 400, "パラメーターエラー発生")
	}

	u := &usecase.PlaceUseCase{}
	places, err := u.GetPlacesByName(q, lon, lat)
	if err != nil {
		return nil, err
	}

	return places, nil
}

func (h *PlaceHandler) GetPlacesByID(id string) (*models.Place, error) {
	u := &usecase.PlaceUseCase{}

	// IDの存在チェック
	if id == "" {
		return nil, domain.New(400, "IDが空です")
	}

	place, err := u.GetPlaceByID(id)
	if err != nil {
		return nil, err
	}

	return place, nil
}

func (h *PlaceHandler) GetPlacesNearBySpots(id, lonStr, latStr, limitStr string) (*dto.PlacesResponse, error) {
	u := &usecase.PlaceUseCase{}

	// IDの存在チェック
	if id == "" {
		return nil, domain.New(400, "IDが空です")
	}

	// クエリパラメータの検証と変換
	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		return nil, domain.Wrap(err, 400, "緯度のパラメータエラー")
	}

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		return nil, domain.Wrap(err, 400, "経度のパラメータエラー")
	}

	limit := 20 // デフォルト値
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			return nil, domain.Wrap(err, 400, "リミットのパラメータエラー")
		}
	}

	places, err := u.GetPlacesNearBySpots(id, lon, lat, limit)
	if err != nil {
		return nil, err
	}

	return places, nil
}
