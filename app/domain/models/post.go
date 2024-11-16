package models

import "time"

type Post struct {
	ID              string      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID          string      `gorm:"type:uuid"`
	PlaceID         string      `gorm:"type:uuid"`
	PostedAt        time.Time   `gorm:"default:now()"`
	Body            *string
	Published       bool        `gorm:"default:false"`
	PublishedAt     *time.Time
	CreatedAt       time.Time   `gorm:"default:now()"`
	UpdatedAt       time.Time   `gorm:"autoUpdateTime"`

	User            User        `gorm:"foreignKey:UserID"`
	Place           Place       `gorm:"foreignKey:PlaceID"`
	Media           []Media
	// Favorites       []Favorite
	// PostHashtags    []PostHashtag
}