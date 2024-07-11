package patient

import (
	"github.com/google/uuid"
	"time"
)

type Patient struct {
	ID                           uuid.UUID `json:"id"`
	Firstname                    string    `json:"firstname"`
	Lastname                     string    `json:"lastname"`
	NationalIdentificationNumber string    `json:"nationalIdentificationNumber"`
	MedicalRecordNumber          string    `json:"medicalRecordNumber"`
	Email                        string    `json:"email"`
	PhoneNumber                  string    `json:"phoneNumber"`

	Admissions []Admission `json:"admissions"`
}

type Admission struct {
	ID          uuid.UUID `json:"id"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	Status      string    `json:"status"`
	Symptoms    string    `json:"symptoms"`
	Medications []string  `json:"medications"`
	Allergies   []string  `json:"allergies"`
	Diagnosis   string    `json:"diagnosis"`

	PatientID  uuid.UUID `json:"patientId"`
	Department string    `json:"department"`
	Physician  uuid.UUID `json:"physician"`
}
