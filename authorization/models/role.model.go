package models

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Section string
type Permission string

const (
	Section_Employees  Section = "EMPLOYEES"
	Section_Intake     Section = "INTAKE"
	Section_Admissions Section = "ADMISSIONS"
	Section_Patients   Section = "PATIENTS"
	Section_Labs       Section = "LABS"
	Section_Imaging    Section = "IMAGING"

	Section_PatientsHistory   Section = "PATIENTS:HISTORY"
	Section_PatientsInfo      Section = "PATIENTS:INFO"
	Section_PatientsVitals    Section = "PATIENTS:VITALS"
	Section_PatientsDiagnosis Section = "PATIENTS:DIAGNOSIS"
	Section_PatientsConsults  Section = "PATIENTS:CONSULTS"
	Section_PatientsTransfer  Section = "PATIENTS:TRANSFER"
	Section_PatientsDischarge Section = "PATIENTS:DISCHARGE"

	Section_PatientsMedicationsPrescribe Section = "PATIENTS:MEDICATIONS:PRESCRIBE"
	Section_PatientsMedicationsGive      Section = "PATIENTS:MEDICATIONS:GIVE"

	Section_PatientsLabsOrder  Section = "PATIENTS:LABS:ORDER"
	Section_PatientsLabsResult Section = "PATIENTS:LABS:RESULT"

	Section_PatientsLogs Section = "PATIENTS:LOGS"
	Section_PatientTeam  Section = "PATIENTS:TEAM"

	Section_PatientsImagingOrder  Section = "PATIENTS:IMAGING:ORDER"
	Section_PatientsImagingResult Section = "PATIENTS:IMAGING:RESULT"
)

const (
	Permission_READ  Permission = "READ"
	Permission_WRITE Permission = "WRITE"
)

type Role struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	ID   string `gorm:"type:string;primaryKey"`
	Name string `gorm:"uniqueIndex;not null"`

	Permissions json.RawMessage `gorm:"type:jsonb"`

	Actors []Actor `gorm:"foreignKey:RoleID;constraint:OnDelete:SET NULL;"`
}
