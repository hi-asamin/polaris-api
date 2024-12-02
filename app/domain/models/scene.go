package models

type Scene struct {
	ID          int          `gorm:"primaryKey;autoIncrement"`
	Name        string       `gorm:"size:100;unique;not null"`
	PlaceScenes []PlaceScene `gorm:"foreignKey:SceneID"`
}

func (Scene) TableName() string {
	return "Scene"
}

type PlaceScene struct {
	PlaceID string `gorm:"type:uuid;not null"`
	SceneID int    `gorm:"not null"`

	Place Place `gorm:"foreignKey:PlaceID"`
	Scene Scene `gorm:"foreignKey:SceneID"`
}

func (PlaceScene) TableName() string {
	return "PlaceScene"
}
