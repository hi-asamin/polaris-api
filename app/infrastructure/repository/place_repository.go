package repository

import (
	"polaris-api/domain"
	"polaris-api/infrastructure"
	"polaris-api/interface/dto"
	"polaris-api/utils"
)

type PlaceRepository struct {}

func (r *PlaceRepository) FindByName(name string, lon, lat float64) ([]dto.SearchPlace, error) {
	db := infrastructure.GetDatabaseConnection()
	// SQLファイルを読み込み
	sqlQuery, err := utils.LoadSQLFile("infrastructure/repository/sql/search_places_by_name.sql")
	if err != nil {
		// エラーを呼び出し元に返却
		return nil, domain.Wrap(err, 500, "SQLファイルの読み込みに失敗")
	}

	// 検索結果を格納するスライスを初期化
	places := []dto.SearchPlace{}

	err = db.Raw(sqlQuery, lon, lat, name).Scan(&places).Error
	if err != nil {
		return nil, domain.Wrap(err, 500, "データベースアクセス時にエラー発生")
	}
	
	return places, nil
}
