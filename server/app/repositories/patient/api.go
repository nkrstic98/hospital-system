package patient

import (
	"github.com/google/uuid"
	"hospital-system/server/models"
)

type Repository interface {
	Insert(user models.Patient) (*models.Patient, error)
	Get(id uuid.UUID) (*models.Patient, error)
	GetByPersonalID(personalID string) (*models.Patient, error)
}
