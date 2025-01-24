package models

import (
	"time"

	"gorm.io/datatypes"
)

type Place struct {
	ID           string         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name         string         `gorm:"size:255" json:"name"`
	Description  *string        `json:"description"`
	Country      string         `gorm:"size:50;not null" json:"country"`
	State        string         `gorm:"size:50;not null" json:"state"`
	City         string         `gorm:"size:150" json:"city"`
	ZipCode      *string        `gorm:"size:10" json:"zipcode"`
	AddressLine1 string         `gorm:"size:500;not null" json:"addressline1"`
	AddressLine2 *string        `gorm:"size:500" json:"addressline2"`
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
