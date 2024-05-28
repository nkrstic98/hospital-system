package patient

import "github.com/google/uuid"

type Service interface {
	CreatePatient(patient Patient) (*Patient, error)
	GetPatient(id string) (*Patient, error)
	GetPatientName(id uuid.UUID) (string, error)
	RegisterPatientAdmission(patientId uuid.UUID, admission Admission) error
	GetAdmissionsByStatuses(statuses []string) ([]Admission, error)
	Discharge(admissionId uuid.UUID) error
}
