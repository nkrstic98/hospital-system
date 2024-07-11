package models

import (
	"github.com/google/uuid"
	"time"
)

type Patient struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Firstname                    string `gorm:"not null"`
	Lastname                     string `gorm:"not null"`
	NationalIdentificationNumber string `gorm:"not null;uniqueIndex"`
	MedicalRecordNumber          string `gorm:"not null;uniqueIndex"`
	Email                        string `gorm:"not null;uniqueIndex"`
	PhoneNumber                  string `gorm:"not null"`

	Admissions []Admission `gorm:"foreignKey:PatientID;constraint:OnDelete:SET NULL;"`
}
