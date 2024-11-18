package repository

import (
	"polaris-api/domain"
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

func (r *PlaceRepository) FindByName(name string, lon, lat float64) ([]dto.SearchPlace, error) {
	db := infrastructure.GetDatabaseConnection()

	// 検索結果を格納するスライスを初期化
	places := []dto.SearchPlace{}

	// SQLファイルを読み込み
	sqlQuery := sql.SearchPlacesByNameQuery()
	err := db.Raw(sqlQuery, lon, lat, name).Scan(&places).Error
	if err != nil {
		return nil, domain.Wrap(err, 500, "データベースアクセス時にエラー発生")
	}

	return places, nil
}
