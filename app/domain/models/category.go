package models

type Category struct {
	ID              int             `gorm:"primaryKey;autoIncrement" json:"id"`
	Name            string          `gorm:"unique;size:100;not null" json:"name"`
	PlaceCategories []PlaceCategory `gorm:"foreignKey:CategoryID" json:"-"`
}

// TableName specifies the table name explicitly.
func (Category) TableName() string {
	return "Category" // 実際のテーブル名を記載
}

type PlaceCategory struct {
	PlaceID    string `gorm:"type:uuid;not null"`
	CategoryID int    `gorm:"not null"`

	Place    Place    `gorm:"foreignKey:PlaceID"`
	Category Category `gorm:"foreignKey:CategoryID"`
}

// TableName specifies the table name explicitly.
func (PlaceCategory) TableName() string {
	return "PlaceCategory" // 実際のテーブル名を記載
}
