package user

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID                           uuid.UUID `json:"id"`
	Firstname                    string    `json:"firstname"`
	Lastname                     string    `json:"lastname"`
	NationalIdentificationNumber string    `json:"national_identification_number"`
	Username                     string    `json:"username"`
	Email                        string    `json:"email"`
	PhoneNumber                  string    `json:"phone_number"`
	MailingAddress               string    `json:"mailing_address"`
	City                         string    `json:"city"`
	State                        string    `json:"state"`
	Zip                          string    `json:"zip"`
	Gender                       string    `json:"gender"`
	Birthday                     time.Time `json:"birthday"`
	JoiningDate                  time.Time `json:"joining_date"`
	Verified                     bool      `json:"verified"`
	Archived                     *bool     `json:"archived"`

	Role        string   `json:"role"`
	Team        *string  `json:"team"`
	Permissions []string `json:"permissions"`
}
