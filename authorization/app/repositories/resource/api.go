package resource

import (
	"github.com/google/uuid"
	"hospital-system/authorization/models"
)

type Repository interface {
	Insert(resource models.Resource) error
	GetByIDs(ids []string) ([]models.Resource, error)
	GetAll() ([]models.Resource, error)
	UpdateArchived(id uuid.UUID) error
}
