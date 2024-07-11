package user

import (
	"github.com/google/uuid"
)

type User struct {
	ID                           uuid.UUID `json:"id"`
	Firstname                    string    `json:"firstname"`
	Lastname                     string    `json:"lastname"`
	NationalIdentificationNumber string    `json:"national_identification_number"`
	Username                     string    `json:"username"`
	Email                        string    `json:"email"`

	Role        string            `json:"role"`
	Team        *string           `json:"team"`
	Permissions map[string]string `json:"permissions"`
}
