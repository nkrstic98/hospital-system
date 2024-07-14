package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Lab struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	RequestedAt time.Time `gorm:"not null"`
	ProcessedAt time.Time `gorm:"default:null"`

	TestType    string          `gorm:"not null"`
	TestResults json.RawMessage `gorm:"default:null"`

	AdmissionID uuid.UUID `gorm:"type:uuid;not null"`
}
