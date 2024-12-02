package models

import (
	"time"
)

type Follows struct {
	FollowerUserID string    `gorm:"type:uuid;not null"`
	FolloweeUserID string    `gorm:"type:uuid;not null"`
	FollowDate     time.Time `gorm:"type:timestamp;default:current_timestamp"`

	Follower User `gorm:"foreignKey:FollowerUserID"`
	Followee User `gorm:"foreignKey:FolloweeUserID"`
}

func (Follows) TableName() string {
	return "Follows"
}
