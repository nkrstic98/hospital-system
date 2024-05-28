package admin

import (
	"github.com/google/uuid"
	"time"
)

type AddUserRequest struct {
	Firstname                    string    `json:"firstname"`
	Lastname                     string    `json:"lastname"`
	NationalIdentificationNumber string    `json:"national_identification_number"`
	Email                        string    `json:"email"`
	JoiningDate                  time.Time `json:"joining_date"`
	Role                         string    `json:"role"`
	Team                         *string   `json:"team"`
}

type AddUserResponse struct {
	Id uuid.UUID `json:"id"`
}

type AddPatientRequest struct {
	Firstname                    string    `json:"firstname"`
	Lastname                     string    `json:"lastname"`
	NationalIdentificationNumber string    `json:"nationalIdentificationNumber"`
	MedicalRecordNumber          string    `json:"medicalRecordNumber"`
	Birthday                     time.Time `json:"birthday"`
	Gender                       string    `json:"gender"`
	Email                        string    `json:"email"`
	PhoneNumber                  string    `json:"phoneNumber"`
}

type AdmitPatientRequest struct {
	PatientId   uuid.UUID `json:"patientId"`
	Department  string    `json:"department"`
	Physician   uuid.UUID `json:"physician"`
	Symptoms    string    `json:"symptoms"`
	Medications []string  `json:"medications"`
	Allergies   []string  `json:"allergies"`
}

type GetAdmissionsRequest struct {
	Statuses []string `json:"statuses"`
}

type GetAdmissionsResponse struct {
	Admissions []AdmissionResponse `json:"admissions"`
}

type AdmissionResponse struct {
	Id            uuid.UUID `json:"id"`
	Patient       string    `json:"patient"`
	Department    string    `json:"department"`
	Physician     string    `json:"physician"`
	AdmissionTime time.Time `json:"admissionTime"`
	Status        string    `json:"status"`
}
