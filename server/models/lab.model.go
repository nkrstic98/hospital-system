package models

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Lab struct {
	ID          uuid.UUID    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	RequestedAt time.Time    `gorm:"not null"`
	ProcessedAt sql.NullTime `gorm:"default:null"`

	TestType    string          `gorm:"not null"`
	TestResults json.RawMessage `gorm:"default:null"`

	AdmissionID uuid.UUID  `gorm:"type:uuid;not null"`
	RequestedBy uuid.UUID  `gorm:"type:uuid;not null"`
	ProcessedBy *uuid.UUID `gorm:"type:uuid;default:null"`
}
