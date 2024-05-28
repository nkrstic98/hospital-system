package actor

import (
	"github.com/google/uuid"
	"hospital-system/authorization/models"
)

type Repository interface {
	Insert(actor models.Actor) error
	Get(id uuid.UUID) (models.Actor, error)
	GetAll() ([]models.Actor, error)
	GetByTeamID(teamID uint) ([]models.Actor, error)
}
