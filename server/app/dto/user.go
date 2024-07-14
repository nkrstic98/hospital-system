package dto

import (
	"github.com/google/uuid"
)

type RegisterUserRequest struct {
	Firstname                    string  `json:"firstname"`
	Lastname                     string  `json:"lastname"`
	NationalIdentificationNumber string  `json:"nationalIdentificationNumber"`
	Email                        string  `json:"email"`
	Role                         string  `json:"role"`
	Team                         *string `json:"team"`
}

type User struct {
	ID                           uuid.UUID         `json:"id"`
	Firstname                    string            `json:"firstname"`
	Lastname                     string            `json:"lastname"`
	NationalIdentificationNumber string            `json:"nationalIdentificationNumber"`
	Username                     string            `json:"username"`
	Email                        string            `json:"email"`
	Role                         string            `json:"role"`
	Team                         *string           `json:"team"`
	Permissions                  map[string]string `json:"permissions"`
}

type GetDepartmentsRequest struct {
	Team *string `json:"team"`
	Role *string `json:"role"`
}

type GetDepartmentsResponse struct {
	Departments map[string]Department `json:"departments"`
}

type Department struct {
	DisplayName string `json:"displayName"`
	Users       []User `json:"users"`
}
