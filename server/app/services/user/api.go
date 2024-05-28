package user

import "github.com/google/uuid"

type Service interface {
	CreateUser(user User) (*uuid.UUID, error)
	GetUser(id uuid.UUID) (*User, error)
	GetByUsername(username string) (*User, error)
	ValidateUserPassword(userId uuid.UUID, password string) (bool, error)
	GetUsers() ([]User, error)
}
