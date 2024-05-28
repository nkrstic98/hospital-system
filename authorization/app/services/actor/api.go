package actor

import "github.com/google/uuid"

type Service interface {
	AddActor(actor Actor) error
	GetActor(id uuid.UUID) (Actor, error)
	GetActors() ([]Actor, error)
	GetActorsByTeamID(teamID uint) ([]Actor, error)
}
