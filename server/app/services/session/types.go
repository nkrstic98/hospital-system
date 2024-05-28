package session

import "github.com/golang-jwt/jwt/v5"

type TokenClaims struct {
	jwt.RegisteredClaims
	Username    string   `json:"username"`
	Firstname   string   `json:"firstname"`
	Lastname    string   `json:"lastname"`
	Role        string   `json:"role"`
	Team        *string  `json:"team"`
	Permissions []string `json:"permissions"`
}
