package repository

import (
	"fmt"
	"log"
	"strings"

	"polaris-api/constants"
	"polaris-api/domain"
	"polaris-api/domain/models"
	"polaris-api/infrastructure"
	"polaris-api/infrastructure/repository/sql"
	"polaris-api/interface/dto"
	"polaris-api/utils"
)

type PlaceRepository struct{}

func (r *PlaceRepository) FindAll(
	lon, lat, cursorDistance float64,
	cursorPID, cursorMID string,
	limit int,
) (*dto.PlacesResponse, error) {
	db := infrastructure.GetDatabaseConnection()

	// 検索結果を格納するスライスを初期化
	places := []dto.PlaceMedia{}

	// 空文字列を NULL に変換
	cursorPIDValue := utils.EmptyStringToNull(cursorPID)
	cursorMIDValue := utils.EmptyStringToNull(cursorMID)

	// SQLファイルを読み込み
	sqlQuery := sql.GetPlaceAndMediaQuery()
	// クエリで `LIMIT+1` のレコードを取得
	err := db.Raw(sqlQuery, lon, lat, cursorDistance, cursorPIDValue, cursorMIDValue, limit+1).Scan(&places).Error
	if err != nil {
		return nil, domain.Wrap(err, 500, "データベースアクセス時にエラー発生")
	}

	// `hasNextPage` フラグと `nextCursor` を初期化
	var nextCursor *dto.NextCursor = nil

	// 検索結果が `limit+1` 件の場合、次のカーソルを設定
	if len(places) > limit {
		lastPlace := places[limit] // `limit+1` 番目の要素が次のカーソル情報
		nextCursor = &dto.NextCursor{
			Distance: lastPlace.Distance,
			PID:      lastPlace.PID,
			MID:      lastPlace.MID,
		}

		// リストの最後の要素を削除して `limit` 件にする
		places = places[:limit]
	}

	// レスポンスDTOを構築
	response := &dto.PlacesResponse{
		PlaceMedia: places,
		NextCursor: nextCursor,
	}

	return response, nil
}

func (r *PlaceRepository) SearchPlacesBaseQuery(
	keywords []string,
	lon, lat float64,
) ([]dto.SearchPlace, error) {
	db := infrastructure.GetDatabaseConnection()

	// 検索結果を格納するスライスを初期化
	places := []dto.SearchPlace{}

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
	log.Printf("%s\n", whereClause)

	// SQLクエリを動的に構築
	query := fmt.Sprintf(sql.SearchPlacesBaseQuery(), whereClause)
	log.Printf("%s\n", query)

	log.Printf("%s\n", params)

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

func (r *PlaceRepository) CreatePlace(req *dto.CreatePlaceRequest) error {
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

	// データベースに登録
	if err := db.Create(newPlace).Error; err != nil {
		return domain.Wrap(err, 500, "データベースアクセス時にエラー発生")
	}

	return nil
}

func (r *PlaceRepository) FindNearBySpots(excludeID string, lon, lat float64, limit int) (*dto.PlacesResponse, error) {
	// 検索結果を格納するスライスを初期化
	places := []dto.PlaceMedia{}

	nearPlaceDistance := constants.NearPlaceDistance

	db := infrastructure.GetDatabaseConnection()

	// SQLファイルを読み込み
	sqlQuery := sql.FindNearBySpots()
	err := db.Raw(sqlQuery, lat, lon, nearPlaceDistance, excludeID, limit+1).Scan(&places).Error
	if err != nil {
		return nil, domain.Wrap(err, 500, "データベースアクセス時にエラー発生")
	}

	// `hasNextPage` フラグと `nextCursor` を初期化
	var nextCursor *dto.NextCursor = nil

	// 検索結果が `limit+1` 件の場合、次のカーソルを設定
	if len(places) > limit {
		lastPlace := places[limit] // `limit+1` 番目の要素が次のカーソル情報
		nextCursor = &dto.NextCursor{
			Distance: lastPlace.Distance,
			PID:      lastPlace.PID,
			MID:      lastPlace.MID,
		}

		// リストの最後の要素を削除して `limit` 件にする
		places = places[:limit]
	}

	// レスポンスDTOを構築
	response := &dto.PlacesResponse{
		PlaceMedia: places,
		NextCursor: nextCursor,
	}

	return response, nil
}

func (r *PlaceRepository) FindPlacesBaseQuery(
	keywords []string,
	cursorPID, cursorMID string,
	limit int,
) (*dto.PlacesResponse, error) {
	db := infrastructure.GetDatabaseConnection()

	// 検索結果を格納するスライスを初期化
	places := []dto.PlaceMedia{}

	// 空文字列を NULL に変換
	cursorPIDValue := utils.EmptyStringToNull(cursorPID)
	cursorMIDValue := utils.EmptyStringToNull(cursorMID)

	// 動的な ILIKE 条件を構築
	ilikeConditions := []string{}
	params := []interface{}{limit, cursorPIDValue, cursorMIDValue}
	paramIndex := 4

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
	log.Printf("%s\n", whereClause)

	// SQLクエリを動的に構築
	query := fmt.Sprintf(sql.FindPlacesBaseQuery(), whereClause)
	log.Printf("%s\n", query)

	log.Printf("%s\n", params)

	// クエリを実行して結果を取得
	err := db.Raw(query, params...).Scan(&places).Error
	if err != nil {
		return nil, domain.Wrap(err, 500, "データベースアクセス時にエラー発生")
	}

	// 次のカーソル情報を初期化
	var nextCursor *dto.NextCursor = nil

	// 検索結果が `limit+1` 件の場合、次のカーソルを設定
	if len(places) > limit {
		lastPlace := places[limit] // `limit+1` 番目の要素が次のカーソル情報
		nextCursor = &dto.NextCursor{
			PID: lastPlace.PID,
			MID: lastPlace.MID,
		}

		// リストの最後の要素を削除して `limit` 件にする
		places = places[:limit]
	}

	// レスポンスDTOを構築
	response := &dto.PlacesResponse{
		PlaceMedia: places,
		NextCursor: nextCursor,
	}

	return response, nil
}
