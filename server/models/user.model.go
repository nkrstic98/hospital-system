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
	PhoneNumber                  string
	MailingAddress               string
	City                         string
	State                        string
	Zip                          string
	Gender                       string
	Birthday                     time.Time
	JoiningDate                  time.Time `gorm:"not null"`
	Verified                     bool      `gorm:"default:false"`
	Archived                     *bool
}
