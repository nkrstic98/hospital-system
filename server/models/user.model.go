package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`

	Firstname                    string `gorm:"not null"`
	Lastname                     string `gorm:"not null"`
	NationalIdentificationNumber string `gorm:"not null;uniqueIndex"`
	Username                     string `gorm:"not null;uniqueIndex"`
	Email                        string `gorm:"not null;uniqueIndex"`
	Password                     string `gorm:"not null"`

	OrderedLabs   []Lab `gorm:"foreignKey:RequestedBy;references:id;constraint:OnDelete:SET NULL;"`
	ProcessedLabs []Lab `gorm:"foreignKey:ProcessedBy;references:id;constraint:OnDelete:SET NULL;"`
}
