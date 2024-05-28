package actors

import "github.com/google/uuid"

type AddActorRequest struct {
	ActorID uuid.UUID `json:"actor_id"`
	Role    string    `json:"role"`
	Team    *string   `json:"team"`
}
