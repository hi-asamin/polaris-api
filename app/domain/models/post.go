package models

import "time"

type Post struct {
	ID          string     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID      string     `gorm:"type:uuid;not null" json:"userId"`
	PlaceID     string     `gorm:"type:uuid;not null" json:"placeId"`
	PostedAt    time.Time  `gorm:"type:timestamp;default:current_timestamp" json:"postedAt"`
	Body        *string    `json:"body"`
	Published   bool       `gorm:"default:false" json:"published"`
	PublishedAt *time.Time `json:"publishedAt"`
	CreatedAt   time.Time  `gorm:"type:timestamp;default:current_timestamp" json:"createdAt"`
	UpdatedAt   time.Time  `gorm:"type:timestamp;autoUpdateTime" json:"updatedAt"`

	User  User    `gorm:"foreignKey:UserID" json:"-"`
	Place Place   `gorm:"foreignKey:PlaceID" json:"-"`
	Media []Media `gorm:"foreignKey:PostID" json:"-"`
}

func (Post) TableName() string {
	return "Post"
}
