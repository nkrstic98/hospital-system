package models

import "gorm.io/gorm"

type PermissionType string

// TODO: Add permission types

type Permission struct {
	gorm.Model

	Name        string `gorm:"uniqueIndex;not null"`
	DisplayName string `gorm:"uniqueIndex;not null"`

	Roles []Role `gorm:"many2many:role_permissions;constraint:OnDelete:SET NULL;"`
}
