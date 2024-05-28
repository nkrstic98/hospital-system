package models

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

const (
	StatusPending    = "pending"
	StatusAdmitted   = "admitted"
	StatusDischarged = "discharged"
)

type Admission struct {
	ID          uuid.UUID       `gorm:"type:uuid;primaryKey"`
	StartTime   time.Time       `gorm:"not null"`
	EndTime     time.Time       `gorm:"default:null"`
	Status      string          `gorm:"not null;default:'pending'"`
	Symptoms    string          `gorm:"not null"`
	Medications json.RawMessage `gorm:"type:json"`
	Allergies   json.RawMessage `gorm:"type:json"`
	Diagnosis   string          `gorm:"default:null"`

	PatientID uuid.UUID `gorm:"type:uuid;not null"`
}
