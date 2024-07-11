package session

import "github.com/golang-jwt/jwt/v5"

type TokenClaims struct {
	jwt.RegisteredClaims
	Username                     string            `json:"username"`
	Firstname                    string            `json:"firstname"`
	Lastname                     string            `json:"lastname"`
	NationalIdentificationNumber string            `json:"national_identification_number"`
	Email                        string            `json:"email"`
	Role                         string            `json:"role"`
	Team                         *string           `json:"team"`
	Permissions                  map[string]string `json:"permissions"`
}
