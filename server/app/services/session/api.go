package session

import "hospital-system/server/app/dto"

type Service interface {
	CreateSession(claims dto.TokenClaims) (string, error)
	GetSession(token string) (*dto.TokenClaims, error)
	RefreshSession(token string) error
	DeleteSession(token string) error
}
