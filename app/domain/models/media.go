package models

import "time"

type Media struct {
	ID         string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	PostID     string    `gorm:"type:uuid;not null" json:"postId"`
	PlaceID    string    `gorm:"type:uuid;not null" json:"placeId"`
	MediaType  string    `gorm:"size:10;not null" json:"mediaType"`
	MediaURL   string    `gorm:"size:500;not null" json:"mediaUrl"`
	AltText    *string   `gorm:"size:255" json:"altText"`
	UploadedAt time.Time `gorm:"type:timestamp;default:current_timestamp" json:"uploadedAt"`

	Post  Post  `gorm:"foreignKey:PostID" json:"-"`
	Place Place `gorm:"foreignKey:PlaceID" json:"-"`
}

func (Media) TableName() string {
	return "Media"
}
