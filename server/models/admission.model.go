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
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`

	StartTime time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	EndTime   time.Time `gorm:"default:null"`
	Status    string    `gorm:"not null;default:'pending'"`

	Anamnesis json.RawMessage `gorm:"not null;type:jsonb"`

	Vitals      json.RawMessage `gorm:"default:null;type:jsonb"`
	Diagnosis   *string         `gorm:"default:null"`
	Medications json.RawMessage `gorm:"default:null;type:jsonb"`

	Logs json.RawMessage `gorm:"type:jsonb"`

	PatientID uuid.UUID `gorm:"type:uuid;not null"`
	Labs      []Lab     `gorm:"foreignKey:AdmissionID;constraint:OnDelete:SET NULL;"`
}
