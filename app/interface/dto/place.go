package dto

type PlacesResponse struct {
	PlaceMedia []PlaceMedia `json:"places"`     // 場所とメディア情報
	NextCursor *NextCursor  `json:"nextCursor"` // 次のカーソル情報
}

// PlaceMedia は場所のメディア情報を表すレスポンスDTO
type PlaceMedia struct {
	PID      string  `gorm:"column:pid" json:"pid"` // Place ID
	MID      string  `gorm:"column:mid" json:"mid"` // Media ID
	Name     string  `json:"name"`                  // 場所の名前
	State    string  `json:"state"`                 // 都道府県
	City     string  `json:"city"`                  // 市区町村
	Src      string  `json:"src"`                   // メディアのソースURL
	Type     string  `json:"type"`                  // メディアの種類
	Alt      *string `json:"alt"`                   // 代替テキスト（省略可能）
	Distance float64 `json:"distance"`              // 距離（キロメートル単位）
}

type NextCursor struct {
	Distance float64 `json:"distance"` // カーソル用の距離
	PID      string  `json:"pid"`      // カーソル用のPlace ID
	MID      string  `json:"mid"`      // カーソル用のMedia ID
}

// SearchPlace は検索結果を表すレスポンスDTO
type SearchPlace struct {
	ID       string  `json:"id"`       // 場所のID
	Name     string  `json:"name"`     // 場所の名前
	State    string  `json:"state"`    // 州または県
	City     string  `json:"city"`     // 市または町
	Code     string  `json:"code"`     // 郵便番号
	Address1 string  `json:"address1"` // 住所1
	Address2 string  `json:"address2"` // 住所2
	Distance float64 `json:"distance"` // 距離（キロメートル単位）
}
