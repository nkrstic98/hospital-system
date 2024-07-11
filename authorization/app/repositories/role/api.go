package role

import "hospital-system/authorization/models"

type Repository interface {
	Insert(role models.Role) error
	Get(id string) (models.Role, error)
	GetAll() ([]models.Role, error)
}
