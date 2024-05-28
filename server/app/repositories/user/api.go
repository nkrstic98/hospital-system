package user

import (
	"github.com/google/uuid"
	"hospital-system/server/models"
)

type Repository interface {
	Insert(user models.User) error
	Get(id uuid.UUID) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetByIDs(ids []uuid.UUID) ([]models.User, error)
	GetAll() ([]models.User, error)
	Delete(id uuid.UUID) error
}
