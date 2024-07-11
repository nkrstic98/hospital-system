package models

import (
	"github.com/google/uuid"
	"time"
)

type Actor struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;"`
	CreatedAt time.Time
	UpdatedAt time.Time

	RoleID string  `gorm:"type:string;not null"`
	TeamID *string `gorm:"type:string"`

	Resources []Resource `gorm:"foreignKey:TeamLead;references:id;constraint:OnDelete:SET NULL;"`
}
