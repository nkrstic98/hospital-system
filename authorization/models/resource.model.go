package models

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
)

type Resource struct {
	ID              uuid.UUID       `gorm:"type:uuid;primaryKey;"`
	Team            string          `gorm:"not null"`
	TeamLead        uuid.UUID       `gorm:"type:uuid;not null"`
	TeamAssignments json.RawMessage `gorm:"not null;type:jsonb"`
	Journey         json.RawMessage `gorm:"default:null;type:jsonb"`
	PendingTransfer json.RawMessage `gorm:"default:null;type:jsonb"`
	Archived        sql.NullBool    `gorm:"default:null"`
}
