package dto

// SearchPlace は検索結果を表すレスポンスDTO
type SearchPlace struct {
	ID       string  `json:"id"`        // 場所のID
	Name     string  `json:"name"`      // 場所の名前
	State    string  `json:"state"`     // 州または県
	City     string  `json:"city"`      // 市または町
	Code     string  `json:"code"`      // 郵便番号
	Address1 string  `json:"address1"`  // 住所1
	Address2 string  `json:"address2"`  // 住所2
	Distance float64 `json:"distance"`  // 距離（キロメートル単位）
}
