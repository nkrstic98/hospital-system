package session

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ValidateSessionRequest struct {
	Token string `json:"token"`
}

type LogoutRequest struct {
	Token string `json:"token"`
}
