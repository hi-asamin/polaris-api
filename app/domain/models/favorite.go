package models

import (
	"time"
)

type Favorite struct {
	UserID      string    `gorm:"type:uuid;not null"`
	PostID      string    `gorm:"type:uuid;not null"`
	FolderID    string    `gorm:"type:uuid;not null"`
	FavoritedAt time.Time `gorm:"type:timestamp;default:current_timestamp"`

	User   User           `gorm:"foreignKey:UserID"`
	Post   Post           `gorm:"foreignKey:PostID"`
	Folder FavoriteFolder `gorm:"foreignKey:FolderID"`
}

func (Favorite) TableName() string {
	return "Favorite"
}

type FavoriteFolder struct {
	ID         string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID     string    `gorm:"type:uuid;not null"`
	FolderName string    `gorm:"size:100;not null"`
	CreatedAt  time.Time `gorm:"type:timestamp;default:current_timestamp"`

	User      User       `gorm:"foreignKey:UserID"`
	Favorites []Favorite `gorm:"foreignKey:FolderID"`
}

func (FavoriteFolder) TableName() string {
	return "FavoriteFolder"
}
