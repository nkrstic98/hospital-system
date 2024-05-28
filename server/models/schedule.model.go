package models

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Schedule struct {
	gorm.Model
	StartDate time.Time       `gorm:"not null"`
	EndDate   time.Time       `gorm:"not null"`
	Schedule  json.RawMessage `gorm:"type:jsonb;not null"`
}
