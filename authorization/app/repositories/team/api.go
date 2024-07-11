package team

import "hospital-system/authorization/models"

type Repository interface {
	Insert(team models.Team) error
	Get(id string) (models.Team, error)
	GetAll() ([]models.Team, error)
}
