package handler

import (
	"strconv"

	"polaris-api/domain"
	"polaris-api/interface/dto"
	"polaris-api/usecase"
)

type PlaceHandler struct {}

func (h *PlaceHandler) GetPlaceByName(q, lonStr, latStr string) ([]dto.SearchPlace, error) {
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
	places, err := u.GetPlaceByID(q, lon, lat)
	if err != nil {
		return nil, err
	}

	return places, nil
}
