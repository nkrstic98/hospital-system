package models

import (
	"github.com/google/uuid"
	"time"
)

type Patient struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`

	Firstname                    string `gorm:"not null"`
	Lastname                     string `gorm:"not null"`
	NationalIdentificationNumber string `gorm:"not null;uniqueIndex"`
	MedicalRecordNumber          string `gorm:"not null;uniqueIndex"`
	Email                        string `gorm:"not null;uniqueIndex"`
	PhoneNumber                  string `gorm:"not null"`

	Admissions []Admission `gorm:"foreignKey:PatientID;constraint:OnDelete:SET NULL;"`
}
