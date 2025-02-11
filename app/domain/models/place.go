package models

import (
	"time"

	"gorm.io/datatypes"
)

type Place struct {
	ID           string         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name         string         `gorm:"size:255;uniqueIndex:idx_place_name_state_city_line1" json:"name"`
	Description  *string        `json:"description"`
	Country      string         `gorm:"size:50;not null" json:"country"`
	ZipCode      *string        `gorm:"size:10" json:"zipCode"`
	State        string         `gorm:"size:50;not null;uniqueIndex:idx_place_name_state_city_line1" json:"state"`
	City         string         `gorm:"size:150;uniqueIndex:idx_place_name_state_city_line1" json:"city"`
	AddressLine1 string         `gorm:"size:500;not null;uniqueIndex:idx_place_name_state_city_line1" json:"addressLine1"`
	AddressLine2 *string        `gorm:"size:500" json:"addressLine2"`
	PhoneNumber  *string        `gorm:"size:20" json:"phoneNumber"`
	Latitude     *float64       `json:"latitude"`
	Longitude    *float64       `json:"longitude"`
	Geometry     string         `gorm:"type:geometry(Point, 4326)" json:"geometry"`
	Links        datatypes.JSON `gorm:"type:jsonb" json:"links"`
	CreatedAt    time.Time      `gorm:"type:timestamp;default:current_timestamp" json:"createdAt"`
	UpdatedAt    time.Time      `gorm:"type:timestamp;autoUpdateTime" json:"updatedAt"`

	Posts           []Post  `json:"posts"`
	Media           []Media `json:"medias"`
	Events          []Event `json:"events"`
	PlaceCategories []PlaceCategory
	PlaceScenes     []PlaceScene
}

// TableName specifies the table name explicitly.
func (Place) TableName() string {
	return "Place" // 実際のテーブル名を記載
}
