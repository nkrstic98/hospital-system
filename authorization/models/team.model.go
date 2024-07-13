package models

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Team struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	ID   string `gorm:"type:string;primaryKey"`
	Name string `gorm:"uniqueIndex;not null"`

	Permissions json.RawMessage `gorm:"type:jsonb"`

	Actors    []Actor    `gorm:"foreignKey:TeamID;constraint:OnDelete:SET NULL;"`
	Resources []Resource `gorm:"foreignKey:Team;references:id;constraint:OnDelete:SET NULL;"`
}
