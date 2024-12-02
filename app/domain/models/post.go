package models

import "time"

type Post struct {
	ID          string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID      string    `gorm:"type:uuid;not null"`
	PlaceID     string    `gorm:"type:uuid;not null"`
	PostedAt    time.Time `gorm:"type:timestamp;default:current_timestamp"`
	Body        *string
	Published   bool `gorm:"default:false"`
	PublishedAt *time.Time
	CreatedAt   time.Time `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt   time.Time `gorm:"type:timestamp;autoUpdateTime"`

	User  User    `gorm:"foreignKey:UserID"`
	Place Place   `gorm:"foreignKey:PlaceID"`
	Media []Media `gorm:"foreignKey:PostID"`
}

func (Post) TableName() string {
	return "Post"
}
