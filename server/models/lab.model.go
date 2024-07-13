package models

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type Lab struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	RequestedAt time.Time `gorm:"not null"`
	ProcessedAt time.Time `gorm:"default:null"`

	TestType    string          `gorm:"not null"`
	TestResults json.RawMessage `gorm:"default:null"`

	AdmissionID uuid.UUID `gorm:"type:uuid;not null"`
}
