package models

import "time"

type Media struct {
	ID          string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	PostID      string    `gorm:"type:uuid"`
	PlaceID     string    `gorm:"type:uuid"`
	MediaType   string    `gorm:"size:10"`
	MediaURL    string    `gorm:"size:500"`
	AltText     *string   `gorm:"size:255"`
	UploadedAt  time.Time `gorm:"default:now()"`
}