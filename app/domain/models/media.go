package models

import "time"

type Media struct {
	ID         string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	PostID     string    `gorm:"type:uuid;not null"`
	PlaceID    string    `gorm:"type:uuid;not null"`
	MediaType  string    `gorm:"size:10;not null"`
	MediaURL   string    `gorm:"size:500;not null"`
	AltText    *string   `gorm:"size:255"`
	UploadedAt time.Time `gorm:"type:timestamp;default:current_timestamp"`

	Post  Post  `gorm:"foreignKey:PostID"`
	Place Place `gorm:"foreignKey:PlaceID"`
}

func (Media) TableName() string {
	return "Media"
}
