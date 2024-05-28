package models

import "gorm.io/gorm"

type Team struct {
	gorm.Model

	Name        string `gorm:"uniqueIndex;not null"`
	DisplayName string `gorm:"uniqueIndex;not null"`

	Actors    []Actor    `gorm:"foreignKey:TeamID;constraint:OnDelete:SET NULL;"`
	Resources []Resource `gorm:"foreignKey:Team;references:name;constraint:OnDelete:SET NULL;"`
}
