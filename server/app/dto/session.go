package dto

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type LogoutRequest struct {
	Token string `json:"token"`
}

type ValidateSessionRequest struct {
	Token string `json:"token"`
}

type ValidateSessionResponse struct {
	User User `json:"user"`
}

type TokenClaims struct {
	jwt.RegisteredClaims
	UserID                       uuid.UUID         `json:"user_id"`
	Username                     string            `json:"username"`
	Firstname                    string            `json:"firstname"`
	Lastname                     string            `json:"lastname"`
	NationalIdentificationNumber string            `json:"national_identification_number"`
	Email                        string            `json:"email"`
	Role                         string            `json:"role"`
	Team                         *string           `json:"team"`
	Permissions                  map[string]string `json:"permissions"`
}
