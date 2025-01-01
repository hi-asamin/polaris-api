package repository

import (
	"fmt"
	"strings"

	"gorm.io/gorm"

	"polaris-api/constants"
	"polaris-api/domain"
	"polaris-api/domain/models"
	"polaris-api/infrastructure"
	"polaris-api/infrastructure/repository/sql"
	"polaris-api/interface/model"
	"polaris-api/utils"
)

type PlaceRepository struct{}

func (r *PlaceRepository) FindAll(
	cursorMID string,
	limit int,
	categoryIds []int,
) (*model.PlacesResponse, error) {
	db := infrastructure.GetDatabaseConnection()

	// 検索結果を格納するスライスを初期化
	places := []model.PlaceMedia{}

	// 空文字列を NULL に変換
	cursorMIDValue := utils.EmptyStringToNull(cursorMID)

	// SQLファイルを読み込み
	sqlQuery := sql.FindAllPlaces()
	// クエリで `LIMIT+1` のレコードを取得
	err := db.Raw(sqlQuery, cursorMIDValue, limit+1, categoryIds).Scan(&places).Error
	if err != nil {
		return nil, domain.Wrap(err, 500, "データベースアクセス時にエラー発生")
	}

	// `hasNextPage` フラグと `nextCursor` を初期化
	var nextCursor *model.NextCursor = nil

	// 検索結果が `limit+1` 件の場合、次のカーソルを設定
	if len(places) > limit {
		lastPlace := places[limit] // `limit+1` 番目の要素が次のカーソル情報
		nextCursor = &model.NextCursor{
			MID: lastPlace.MID,
		}

		// リストの最後の要素を削除して `limit` 件にする
		places = places[:limit]
	}

	// レスポンスDTOを構築
	response := &model.PlacesResponse{
		PlaceMedia: places,
		NextCursor: nextCursor,
	}

	return response, nil
}

func (r *PlaceRepository) FindPlacesByName(
	keywords []string,
	lon, lat float64,
) ([]model.SearchPlace, error) {
	db := infrastructure.GetDatabaseConnection()

	// 検索結果を格納するスライスを初期化
	places := []model.SearchPlace{}

	// 動的な ILIKE 条件を構築
	ilikeConditions := []string{}
	params := []interface{}{lon, lat}
	paramIndex := 3

	for _, word := range keywords {
		ilikeConditions = append(ilikeConditions, fmt.Sprintf(`(
			"Place".name ILIKE $%d::TEXT OR
			"Place".description ILIKE $%d::TEXT OR
			"Place".city ILIKE $%d::TEXT
		)`, paramIndex, paramIndex+1, paramIndex+2))
		params = append(params, "%"+word+"%", "%"+word+"%", "%"+word+"%")
		paramIndex += 3
	}

	// ILIKE条件を結合
	whereClause := strings.Join(ilikeConditions, " AND ")

	// SQLクエリを動的に構築
	query := fmt.Sprintf(sql.SearchPlacesBaseQuery(), whereClause)

	// クエリを実行して結果を取得
	err := db.Raw(query, params...).Scan(&places).Error
	if err != nil {
		return nil, domain.Wrap(err, 500, "データベースアクセス時にエラー発生")
	}

	return places, nil
}

func (r *PlaceRepository) FindByID(id string) (*models.Place, error) {
	db := infrastructure.GetDatabaseConnection()

	var place models.Place
	if err := db.Select(
		"id, name, description, country, state, city, zip_code, address_line1, address_line2, latitude, longitude",
	).First(&place, "id = ?", id).Error; err != nil {
		return nil, domain.Wrap(err, 500, "データベースアクセス時にエラー発生")
	}

	return &place, nil
}

// createGeometryは緯度経度からWKT形式のgeometryデータを生成します
func createGeometry(lat, lon *float64) string {
	if lat != nil && lon != nil {
		return fmt.Sprintf("POINT(%f %f)", *lon, *lat)
	}
	return ""
}

func (r *PlaceRepository) CreatePlace(req *model.CreatePlaceRequest) error {
	db := infrastructure.GetDatabaseConnection()

	// Placeモデルに変換
	newPlace := &models.Place{
		Name:         req.Name,
		Description:  req.Description,
		Country:      req.Country,
		State:        req.State,
		City:         req.City,
		ZipCode:      req.ZipCode,
		AddressLine1: req.AddressLine1,
		AddressLine2: req.AddressLine2,
		Latitude:     req.Latitude,
		Longitude:    req.Longitude,
		Geometry:     createGeometry(req.Latitude, req.Longitude),
	}

	return db.Transaction(func(tx *gorm.DB) error {
		// Placeを保存
		if err := tx.Create(newPlace).Error; err != nil {
			return domain.Wrap(err, 500, "場所の保存に失敗")
		}

		// CategoryIdsが1つ以上存在する場合のみPlaceCategoryを保存
		if len(req.CategoryIds) > 0 {
			for _, categoryID := range req.CategoryIds {
				placeCategory := &models.PlaceCategory{
					PlaceID:    newPlace.ID,
					CategoryID: categoryID,
				}
				if err := tx.Create(placeCategory).Error; err != nil {
					return domain.Wrap(err, 500, "カテゴリ情報の保存に失敗")
				}
			}
		}

		return nil
	})
}

func (r *PlaceRepository) FindNearBySpots(excludeID, cursorMID string, lon, lat float64, limit int) (*model.PlacesResponse, error) {
	db := infrastructure.GetDatabaseConnection()
	// 検索結果を格納するスライスを初期化
	places := []model.PlaceMedia{}
	nearPlaceDistance := constants.NearPlaceDistance
	// 空文字列を NULL に変換
	cursorMIDValue := utils.EmptyStringToNull(cursorMID)

	// SQLファイルを読み込み
	sqlQuery := sql.FindNearBySpots()
	err := db.Raw(sqlQuery, lon, lat, nearPlaceDistance, excludeID, cursorMIDValue, limit+1).Scan(&places).Error
	if err != nil {
		return nil, domain.Wrap(err, 500, "データベースアクセス時にエラー発生")
	}

	// `hasNextPage` フラグと `nextCursor` を初期化
	var nextCursor *model.NextCursor = nil

	// 検索結果が `limit+1` 件の場合、次のカーソルを設定
	if len(places) > limit {
		lastPlace := places[limit] // `limit+1` 番目の要素が次のカーソル情報
		nextCursor = &model.NextCursor{
			Distance: lastPlace.Distance,
			MID:      lastPlace.MID,
		}

		// リストの最後の要素を削除して `limit` 件にする
		places = places[:limit]
	}

	// レスポンスDTOを構築
	response := &model.PlacesResponse{
		PlaceMedia: places,
		NextCursor: nextCursor,
	}

	return response, nil
}

func (r *PlaceRepository) FindPlacesByNameWithMedia(
	keywords []string,
	cursorMID string,
	limit int,
) (*model.PlacesResponse, error) {
	db := infrastructure.GetDatabaseConnection()

	// 検索結果を格納するスライスを初期化
	places := []model.PlaceMedia{}

	// 空文字列を NULL に変換
	cursorMIDValue := utils.EmptyStringToNull(cursorMID)

	// キーワードの有無で処理を分岐
	var whereClause string
	params := []interface{}{limit + 1, cursorMIDValue}

	if len(keywords) > 0 {
		// 動的な ILIKE 条件を構築
		ilikeConditions := []string{}
		paramIndex := 3

		for _, word := range keywords {
			ilikeConditions = append(ilikeConditions, fmt.Sprintf(`(
				"Place".name ILIKE $%d::TEXT OR
				"Place".description ILIKE $%d::TEXT OR
				"Place".city ILIKE $%d::TEXT
			)`, paramIndex, paramIndex+1, paramIndex+2))
			params = append(params, "%"+word+"%", "%"+word+"%", "%"+word+"%")
			paramIndex += 3
		}

		// ILIKE条件を結合
		whereClause = strings.Join(ilikeConditions, " AND ")
	} else {
		// キーワードがない場合は常に真となる条件を設定
		whereClause = "TRUE"
	}

	// SQLクエリを動的に構築
	query := fmt.Sprintf(sql.FindPlacesBaseQuery(), whereClause)

	// クエリを実行して結果を取得
	err := db.Raw(query, params...).Scan(&places).Error
	if err != nil {
		return nil, domain.Wrap(err, 500, "データベースアクセス時にエラー発生")
	}

	// 次のカーソル情報を初期化
	var nextCursor *model.NextCursor = nil

	// 検索結果が `limit+1` 件の場合、次のカーソルを設定
	if len(places) > limit {
		lastPlace := places[limit] // `limit+1` 番目の要素が次のカーソル情報
		nextCursor = &model.NextCursor{
			MID: lastPlace.MID,
		}

		// リストの最後の要素を削除して `limit` 件にする
		places = places[:limit]
	}

	// レスポンスDTOを構築
	response := &model.PlacesResponse{
		PlaceMedia: places,
		NextCursor: nextCursor,
	}

	return response, nil
}
