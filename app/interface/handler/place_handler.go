package handler

import (
	"strconv"
	"strings"

	"polaris-api/domain"
	"polaris-api/domain/models"
	"polaris-api/interface/types"
	"polaris-api/usecase"
)

type PlaceHandler struct{}

func (h *PlaceHandler) GetPlaces(
	cursorMID, limitStr string,
	categoryIds []int,
) (*types.PlacesResponse, error) {
	var err error
	// クエリパラメータの検証と変換
	limit := 20 // デフォルト値
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			return nil, domain.Wrap(err, 400, "リミットのパラメータエラー")
		}
	}

	// Usecaseの呼び出し
	u := &usecase.PlaceUseCase{}
	response, err := u.GetPlaces(cursorMID, limit, categoryIds)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (h *PlaceHandler) NewPlace(req *types.CreatePlaceRequest) error {
	// Usecaseの呼び出し
	u := &usecase.PlaceUseCase{}

	err := u.CreatePlace(req)
	if err != nil {
		return err
	}

	return nil
}

func (h *PlaceHandler) GetSuggestPlaces(q, lonStr, latStr string) ([]types.SearchPlace, error) {
	// 検索ワードが空の場合は空配列を返却する
	if q == "" {
		return []types.SearchPlace{}, nil
	}
	// キーワードをスペースで分割
	keywords := strings.Fields(q)
	if len(keywords) == 0 {
		return []types.SearchPlace{}, nil
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
	places, err := u.GetSuggestPlaces(keywords, lon, lat)
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

func (h *PlaceHandler) UpdatePlace(id string, payload *types.PlaceUpdatePayload) error {
	u := &usecase.PlaceUseCase{}

	// IDの存在チェック
	if id == "" {
		return domain.New(400, "IDが空です")
	}

	// 更新したいカラムだけを map へ詰め込む
	updateFields := map[string]interface{}{}

	if payload.Name != nil {
		updateFields["name"] = payload.Name
	}
	if payload.Description != nil {
		updateFields["description"] = *payload.Description
	}
	if payload.Country != nil {
		updateFields["country"] = payload.Country
	}
	if payload.State != nil {
		updateFields["state"] = payload.State
	}
	if payload.City != nil {
		updateFields["city"] = payload.City
	}
	if payload.ZipCode != nil {
		updateFields["zip_code"] = payload.ZipCode
	}
	if payload.AddressLine1 != nil {
		updateFields["address_line1"] = payload.AddressLine1
	}
	if payload.AddressLine2 != nil {
		updateFields["address_line2"] = payload.AddressLine2
	}
	if payload.PhoneNumber != nil {
		updateFields["phone_number"] = payload.PhoneNumber
	}
	if payload.CategoryIds != nil {
		updateFields["category_ids"] = payload.CategoryIds
	}
	if payload.Links != nil {
		updateFields["links"] = payload.Links
	}

	err := u.UpdatePlace(id, updateFields)
	if err != nil {
		return err
	}

	return nil
}

func (h *PlaceHandler) GetPlacesNearBySpots(id, lonStr, latStr, cursorMID, limitStr string) (*types.PlacesResponse, error) {
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

	places, err := u.GetPlacesNearBySpots(id, cursorMID, lon, lat, limit)
	if err != nil {
		return nil, err
	}

	return places, nil
}

func (h *PlaceHandler) GetPlacesBaseQuery(
	keywords, cursorMID, limitStr string,
) (*types.PlacesResponse, error) {
	// キーワードをスペースで分割
	words := strings.Fields(keywords)

	var err error
	limit := 20 // デフォルト値
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			return nil, domain.Wrap(err, 400, "リミットのパラメータエラー")
		}
	}

	// Usecaseの呼び出し
	u := &usecase.PlaceUseCase{}
	response, err := u.GetPlacesBaseQuery(words, cursorMID, limit)
	if err != nil {
		return nil, err
	}

	return response, nil
}
