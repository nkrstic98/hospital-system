package user

import (
	"github.com/google/uuid"
	"hospital-system/server/app/dto"
)

type Service interface {
	CreateUser(user dto.User) error
	GetUser(id uuid.UUID) (*dto.User, error)
	GetByUsername(username string) (*dto.User, error)
	ValidateUserPassword(userId uuid.UUID, password string) (bool, error)
	GetUsers() ([]dto.User, error)
}
