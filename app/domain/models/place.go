package models

type Place struct {
	ID           string   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name         string   `gorm:"size:255" json:"name"`
	Description  *string  `json:"description"`
	Country      string   `gorm:"size:50" json:"country"`
	State        string   `gorm:"size:50" json:"state"`
	City         string   `gorm:"size:150" json:"city"`
	ZipCode      *string  `gorm:"size:10" json:"zipcode"`
	AddressLine1 string   `gorm:"size:500" json:"addressline1"`
	AddressLine2 *string  `gorm:"size:500" json:"addressline2"`
	Latitude     *float64 `json:"latitude"`
	Longitude    *float64 `json:"longitude"`
	Geometry     float64  `gorm:"type:geometry(Point, 4326)" json:"geometry"`

	Posts []Post  `json:"posts"`
	Media []Media `json:"media"`
	// Events             []Event
	// PlaceCategories    []PlaceCategory
	// PlaceScenes        []PlaceScene
}

// TableName specifies the table name explicitly.
func (Place) TableName() string {
	return "Place" // 実際のテーブル名を記載
}
