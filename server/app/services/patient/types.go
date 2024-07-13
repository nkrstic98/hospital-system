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

	Admissions []AdmissionDetails `json:"admissions"`
}

type Anamnesis struct {
	ChiefComplaint          string   `json:"chief_complaint"`
	HistoryOfPresentIllness string   `json:"history_of_present_illness"`
	PastMedicalHistory      string   `json:"past_medical_history"`
	Medications             []string `json:"medications"`
	Allergies               []string `json:"allergies"`
	FamilyHistory           string   `json:"family_history"`
	SocialHistory           string   `json:"social_history"`
	PhysicalExamination     string   `json:"physical_examination"`
}

type Vitals struct {
	BodyTemperature   float64 `json:"body_temperature"`
	HeartRate         int     `json:"heart_rate"`
	BloodPressure     string  `json:"blood_pressure"`
	RespiratoryRate   int     `json:"respiratory_rate"`
	OxygenSaturation  int     `json:"oxygen_saturation"`
	PainLevel         int     `json:"pain_level"`
	BloodGlucoseLevel int     `json:"blood_glucose_level"`
}

type MedicationInfo struct {
	Medication string `json:"medication"`
	Dose       string `json:"dose"`
}

type Log struct {
	Timestamp   time.Time `json:"timestamp"`
	Message     string    `json:"message"`
	Action      string    `json:"action"`
	Details     string    `json:"details"`
	PerformedBy string    `json:"performed_by"`
}

type LabTest struct {
	Name           string   `json:"name"`
	Unit           string   `json:"unit"`
	MinValue       float64  `json:"min_value"`
	MaxValue       float64  `json:"max_value"`
	ReferenceRange string   `json:"reference_range"`
	Result         *float64 `json:"result"`
}

type Lab struct {
	ID          uuid.UUID `json:"id"`
	RequestedAt time.Time `json:"requested_at"`
	ProcessedAt time.Time `json:"processed_at"`

	TestType    string     `json:"test_type"`
	TestResults *[]LabTest `json:"test_results"`
}

type AdmissionDetails struct {
	ID uuid.UUID `json:"id"`

	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Status    string    `json:"status"`

	Anamnesis Anamnesis `json:"intake_info"`

	Vitals      Vitals            `json:"vitals"`
	Diagnosis   *string           `json:"diagnosis"`
	Medications *[]MedicationInfo `json:"medications"`
	Labs        *[]Lab            `json:"labs"`

	Logs []Log `json:"logs"`

	Patient    Patient   `json:"patient"`
	Department string    `json:"department"`
	Physician  uuid.UUID `json:"physician"`
}
