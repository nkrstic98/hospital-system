package role

import "hospital-system/authorization/models"

type Repository interface {
	Insert(role models.Role) error
	Get(id uint) (models.Role, error)
	GetByName(name string) (models.Role, error)
	GetAll() ([]models.Role, error)
}
