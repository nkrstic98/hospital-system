package dto

import (
	"time"

	"github.com/google/uuid"
)

type RegisterPatientRequest struct {
	Firstname                    string `json:"firstname"`
	Lastname                     string `json:"lastname"`
	NationalIdentificationNumber string `json:"nationalIdentificationNumber"`
	MedicalRecordNumber          string `json:"medicalRecordNumber"`
	Email                        string `json:"email"`
	PhoneNumber                  string `json:"phoneNumber"`
}

type AdmitPatientRequest struct {
	PatientId               uuid.UUID `json:"patientId"`
	Department              string    `json:"department"`
	Physician               uuid.UUID `json:"physician"`
	ChiefComplaint          string    `json:"chiefComplaint"`
	HistoryOfPresentIllness string    `json:"historyOfPresentIllness"`
	PastMedicalHistory      string    `json:"pastMedicalHistory"`
	Medications             []string  `json:"medications"`
	Allergies               []string  `json:"allergies"`
	FamilyHistory           string    `json:"familyHistory"`
	SocialHistory           string    `json:"socialHistory"`
	PhysicalExamination     string    `json:"physicalExamination"`
	AdmittedBy              uuid.UUID `json:"admittedBy"`
	AdmittedByTeam          string    `json:"admittedByTeam"`
}

type OrderLabTestRequest struct {
	LabTest     string    `json:"labTest"`
	AdmissionId uuid.UUID `json:"admissionId"`
}

type AddTeamMemberRequest struct {
	AdmissionId uuid.UUID `json:"admissionId"`
	UserId      uuid.UUID `json:"userId"`
}

type RemoveTeamMemberRequest struct {
	AdmissionId uuid.UUID `json:"admissionId"`
	UserId      uuid.UUID `json:"userId"`
}

type AddTeamMemberPermissionsRequest struct {
	AdmissionId uuid.UUID `json:"admissionId"`
	UserId      uuid.UUID `json:"userId"`
	Section     string    `json:"section"`
	Permission  string    `json:"permission"`
}

type RemoveTeamMemberPermissionsRequest struct {
	AdmissionId uuid.UUID `json:"admissionId"`
	UserId      uuid.UUID `json:"userId"`
	Section     string    `json:"section"`
}

type TransferPatientRequest struct {
	AdmissionId uuid.UUID `json:"admissionId"`
	ToTeam      string    `json:"toTeam"`
	ToTeamLead  uuid.UUID `json:"toTeamLead"`
}

//============================================================

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
	ID         uuid.UUID  `json:"id"`
	StartTime  time.Time  `json:"startTime"`
	EndTime    *time.Time `json:"endTime"`
	Status     string     `json:"status"`
	Patient    string     `json:"patient"`
	Department string     `json:"department"`
	Physician  string     `json:"physician"`
}

type AdmissionDetails struct {
	ID uuid.UUID `json:"id"`

	StartTime time.Time  `json:"startTime"`
	EndTime   *time.Time `json:"endTime"`
	Status    string     `json:"status"`

	Anamnesis Anamnesis `json:"anamnesis"`

	Vitals      Vitals           `json:"vitals"`
	Diagnosis   *string          `json:"diagnosis"`
	Medications []MedicationInfo `json:"medications"`
	Labs        []Lab            `json:"labs"`

	Logs []Log `json:"logs"`

	Patient  Patient  `json:"patient"`
	CareTeam CareTeam `json:"careTeam"`

	PendingTransfer *JourneyStep `json:"pendingTransfer"`
}

type Anamnesis struct {
	ChiefComplaint          string   `json:"chiefComplaint"`
	HistoryOfPresentIllness string   `json:"historyOfPresentIllness"`
	PastMedicalHistory      string   `json:"pastMedicalHistory"`
	Medications             []string `json:"medications"`
	Allergies               []string `json:"allergies"`
	FamilyHistory           string   `json:"familyHistory"`
	SocialHistory           string   `json:"socialHistory"`
	PhysicalExamination     string   `json:"physicalExamination"`
}

type Vitals struct {
	BodyTemperature   string `json:"bodyTemperature"`
	HeartRate         string `json:"heartRate"`
	BloodPressure     string `json:"bloodPressure"`
	RespiratoryRate   string `json:"respiratoryRate"`
	OxygenSaturation  string `json:"oxygenSaturation"`
	PainLevel         string `json:"painLevel"`
	BloodGlucoseLevel string `json:"bloodGlucoseLevel"`
}

type MedicationInfo struct {
	Medication string `json:"medication"`
	Dose       string `json:"dose"`
}

type Log struct {
	Timestamp   time.Time `json:"timestamp"`
	Action      string    `json:"action"`
	Message     string    `json:"message"`
	Details     string    `json:"details"`
	PerformedBy string    `json:"performedBy"`
}

type LabTest struct {
	Name           string   `json:"name"`
	Unit           string   `json:"unit"`
	MinValue       float64  `json:"minValue"`
	MaxValue       float64  `json:"maxValue"`
	ReferenceRange string   `json:"referenceRange"`
	Result         *float64 `json:"result"`
}

type Lab struct {
	ID          uuid.UUID  `json:"id"`
	RequestedAt time.Time  `json:"requestedAt"`
	ProcessedAt *time.Time `json:"processedAt"`

	TestType    string     `json:"testType"`
	TestResults *[]LabTest `json:"testResults"`

	RequestedBy uuid.UUID  `json:"requestedBy"`
	ProcessedBy *uuid.UUID `json:"processedBy"`
}

type CareTeam struct {
	Team            string             `json:"team"`
	Department      string             `json:"department"`
	TeamLead        uuid.UUID          `json:"teamLead"`
	Assignments     map[uuid.UUID]User `json:"assignments"`
	Journey         []JourneyStep      `json:"journey"`
	PendingTransfer *JourneyStep       `json:"pendingTransfer"`
}

type JourneyStep struct {
	TransferTime time.Time `json:"transferTime"`
	FromTeam     string    `json:"fromTeam"`
	ToTeam       string    `json:"toTeam"`
	FromTeamLead uuid.UUID `json:"fromTeamLead"`
	ToTeamLead   uuid.UUID `json:"toTeamLead"`
}
