package models

import "time"

// Gender Enum
type Gender string

const (
    Male   Gender = "MALE"
    Female Gender = "FEMALE"
    Other  Gender = "OTHER"
)

type User struct {
	ID                string          `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserName          string          `gorm:"unique;size:50"`
	DisplayName       string          `gorm:"size:255"`
	Email             string          `gorm:"unique;size:255"`
	SelfIntroduction  *string
	SNSLinks          *string         `gorm:"type:json"`
	DateOfBirth       *time.Time      `gorm:"type:date"`
	Gender            *Gender
	Country           *string         `gorm:"size:100"`
	Region            *string         `gorm:"size:100"`
	Language          *string         `gorm:"size:50"`
	IsPublic          bool            `gorm:"default:true"`
	OptIn             bool            `gorm:"default:true"`

	Posts             []Post          `gorm:"foreignKey:UserID"`
	// Favorites         []Favorite      `gorm:"foreignKey:UserID"`
	// FavoriteFolders   []FavoriteFolder `gorm:"foreignKey:UserID"`
	// Followings        []Follows       `gorm:"foreignKey:FollowerUserID"`
	// Followers         []Follows       `gorm:"foreignKey:FolloweeUserID"`
}