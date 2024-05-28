package session

type Service interface {
	CreateSession(claims TokenClaims) (string, error)
	GetSession(token string) (*TokenClaims, error)
	RefreshSession(token string) error
	DeleteSession(token string) error
}
