package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model

	Name        string `gorm:"uniqueIndex;not null"`
	DisplayName string `gorm:"uniqueIndex;not null"`

	Actors      []Actor      `gorm:"foreignKey:RoleID;constraint:OnDelete:SET NULL;"`
	Permissions []Permission `gorm:"many2many:role_permissions;constraint:OnDelete:SET NULL;"`
}
