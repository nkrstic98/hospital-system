package models

import "github.com/google/uuid"

type Resource struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Team     string    `gorm:"not null"`
	TeamLead uuid.UUID `gorm:"type:uuid;not null"`
	Archived *bool     `gorm:"default:null"`
}
