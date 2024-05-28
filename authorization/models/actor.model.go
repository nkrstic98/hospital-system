package models

import (
	"github.com/google/uuid"
	"time"
)

type Actor struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;"`
	CreatedAt time.Time
	UpdatedAt time.Time

	RoleID uint  `gorm:"type:uuid;not null"`
	TeamID *uint `gorm:"type:uuid"`

	Resources []Resource `gorm:"foreignKey:TeamLead;references:id;constraint:OnDelete:SET NULL;"`
}
