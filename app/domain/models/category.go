package models

type Category struct {
	ID              int             `gorm:"primaryKey;autoIncrement" json:"id"`
	Name            string          `gorm:"unique;size:100" json:"name"`
	PlaceCategories []PlaceCategory `json:"-"`
}

type PlaceCategory struct {
	PlaceID    string `gorm:"type:uuid"`
	CategoryID int
}

// TableName specifies the table name explicitly.
func (Category) TableName() string {
	return "Category" // 実際のテーブル名を記載
}

// TableName specifies the table name explicitly.
func (PlaceCategory) TableName() string {
	return "PlaceCategory" // 実際のテーブル名を記載
}
