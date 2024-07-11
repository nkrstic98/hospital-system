package actor

import "github.com/google/uuid"

type Actor struct {
	ActorID     uuid.UUID
	Role        string
	Team        *string
	Permissions map[string]string
}
