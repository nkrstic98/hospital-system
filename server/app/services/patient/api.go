package patient

import "github.com/google/uuid"

type Service interface {
	CreatePatient(patient Patient) (*Patient, error)
	GetPatient(id string) (*Patient, error)
	GetPatientName(id uuid.UUID) (string, error)
	RegisterPatientAdmission(patientId uuid.UUID, admission AdmissionDetails) error
	GetAdmissionsByStatuses(statuses []string) ([]AdmissionDetails, error)
	GetActiveAdmissionsByPhysician(physicianId uuid.UUID) ([]AdmissionDetails, error)
	GetAdmission(id uuid.UUID) (*AdmissionDetails, error)
}
