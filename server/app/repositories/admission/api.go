package admission

import (
	"github.com/google/uuid"
	"hospital-system/server/models"
)

type Repository interface {
	Insert(admission models.Admission) (*models.Admission, error)
	Delete(id uuid.UUID) error
	Get(id uuid.UUID) (*models.Admission, error)
	Update(admission *models.Admission) (*models.Admission, error)
	GetByPatientId(id uuid.UUID) ([]models.Admission, error)
	GetByStatuses(statuses []string) ([]models.Admission, error)
	GetByIDs(ids []uuid.UUID) ([]models.Admission, error)
	GetLabsByAdmissionID(admissionID uuid.UUID) ([]models.Lab, error)
}
