package patient

import (
	"github.com/google/uuid"
	"hospital-system/server/app/dto"
)

type Service interface {
	RegisterPatient(patient dto.Patient) (*dto.Patient, error)
	GetPatient(id string) (*dto.Patient, error)
	GetPatientName(id uuid.UUID) (string, error)
	RegisterPatientAdmission(patientId uuid.UUID, admission dto.AdmissionDetails) error
	GetAdmissionsByStatuses(statuses []string) ([]dto.AdmissionDetails, error)
	GetActiveAdmissionsByPhysician(physicianId uuid.UUID) ([]dto.AdmissionDetails, error)
	GetAdmission(id uuid.UUID) (*dto.AdmissionDetails, error)
	GetAdmissions() ([]dto.Admission, error)
}
