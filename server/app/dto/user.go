package dto

import (
	"github.com/google/uuid"
	"time"
)

type RegisterUserRequest struct {
	Firstname                    string    `json:"firstname"`
	Lastname                     string    `json:"lastname"`
	NationalIdentificationNumber string    `json:"national_identification_number"`
	Email                        string    `json:"email"`
	JoiningDate                  time.Time `json:"joining_date"`
	Role                         string    `json:"role"`
	Team                         *string   `json:"team"`
}

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
