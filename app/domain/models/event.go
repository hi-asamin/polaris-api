package models

import (
	"time"
)

type Event struct {
	ID               string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	PlaceID          string    `gorm:"type:uuid;not null"`
	EventDate        time.Time `gorm:"type:date;not null"`
	EventDescription string    `gorm:"type:text;not null"`

	Place Place `gorm:"foreignKey:PlaceID"`
}

func (Event) TableName() string {
	return "Event"
}
