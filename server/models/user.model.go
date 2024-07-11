package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Firstname                    string `gorm:"not null"`
	Lastname                     string `gorm:"not null"`
	NationalIdentificationNumber string `gorm:"not null;uniqueIndex"`
	Username                     string `gorm:"not null;uniqueIndex"`
	Email                        string `gorm:"not null;uniqueIndex"`
	Password                     string `gorm:"not null"`
}
